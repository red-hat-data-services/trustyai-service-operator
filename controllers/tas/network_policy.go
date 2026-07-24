package tas

import (
	"context"
	"fmt"
	"reflect"

	trustyaiopendatahubiov1 "github.com/trustyai-explainability/trustyai-service-operator/api/tas/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	kserveInferenceServiceLabel = "serving.kserve.io/inferenceservice"
)

func networkPolicyName(instanceName string) string {
	return instanceName + "-kserve-ingress"
}

func (r *TrustyAIServiceReconciler) ensureNetworkPolicy(ctx context.Context, instance *trustyaiopendatahubiov1.TrustyAIService) error {
	logger := log.FromContext(ctx)
	npName := networkPolicyName(instance.Name)

	existing := &networkingv1.NetworkPolicy{}
	err := r.Get(ctx, types.NamespacedName{Name: npName, Namespace: instance.Namespace}, existing)
	if err == nil {
		if !r.networkPolicyNeedsUpdate(existing) {
			return nil
		}
		existing.Spec = r.desiredNetworkPolicySpec(instance)
		if err := r.Update(ctx, existing); err != nil {
			return fmt.Errorf("failed to update NetworkPolicy %s: %v", npName, err)
		}
		logger.Info("Updated NetworkPolicy", "name", npName)
		return nil
	}
	if !errors.IsNotFound(err) {
		return fmt.Errorf("failed to check for existing NetworkPolicy %s: %v", npName, err)
	}

	np := &networkingv1.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      npName,
			Namespace: instance.Namespace,
		},
		Spec: r.desiredNetworkPolicySpec(instance),
	}

	if err := ctrl.SetControllerReference(instance, np, r.Scheme); err != nil {
		return fmt.Errorf("failed to set controller reference on NetworkPolicy %s: %v", npName, err)
	}

	if err := r.Create(ctx, np); err != nil {
		return fmt.Errorf("failed to create NetworkPolicy %s: %v", npName, err)
	}

	logger.Info("Created NetworkPolicy", "name", npName)
	return nil
}

func (r *TrustyAIServiceReconciler) desiredNetworkPolicySpec(instance *trustyaiopendatahubiov1.TrustyAIService) networkingv1.NetworkPolicySpec {
	protocolTCP := corev1.ProtocolTCP
	port8080 := intstr.FromInt32(8080)
	port4443 := intstr.FromInt32(4443)
	port8443 := intstr.FromInt32(8443)

	return networkingv1.NetworkPolicySpec{
		PodSelector: metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app":                        instance.Name,
				"app.kubernetes.io/instance": instance.Name,
			},
		},
		PolicyTypes: []networkingv1.PolicyType{
			networkingv1.PolicyTypeIngress,
		},
		Ingress: []networkingv1.NetworkPolicyIngressRule{
			{
				From: []networkingv1.NetworkPolicyPeer{
					{
						PodSelector: &metav1.LabelSelector{
							MatchExpressions: []metav1.LabelSelectorRequirement{
								{
									Key:      kserveInferenceServiceLabel,
									Operator: metav1.LabelSelectorOpExists,
								},
							},
						},
					},
				},
				Ports: []networkingv1.NetworkPolicyPort{
					{Protocol: &protocolTCP, Port: &port8080},
					{Protocol: &protocolTCP, Port: &port4443},
				},
			},
			{
				From: []networkingv1.NetworkPolicyPeer{
					{
						NamespaceSelector: &metav1.LabelSelector{
							MatchExpressions: []metav1.LabelSelectorRequirement{
								{
									Key:      "kubernetes.io/metadata.name",
									Operator: metav1.LabelSelectorOpIn,
									Values: []string{
										"openshift-monitoring",
										"openshift-user-workload-monitoring",
									},
								},
							},
						},
					},
				},
				Ports: []networkingv1.NetworkPolicyPort{
					{Protocol: &protocolTCP, Port: &port8080},
				},
			},
			{
				Ports: []networkingv1.NetworkPolicyPort{
					{Protocol: &protocolTCP, Port: &port8443},
				},
			},
		},
	}
}

func (r *TrustyAIServiceReconciler) networkPolicyNeedsUpdate(existing *networkingv1.NetworkPolicy) bool {
	if len(existing.Spec.Ingress) != 3 {
		return true
	}
	if len(existing.Spec.PolicyTypes) != 1 || existing.Spec.PolicyTypes[0] != networkingv1.PolicyTypeIngress {
		return true
	}
	kserveRule := existing.Spec.Ingress[0]
	if len(kserveRule.From) != 1 || kserveRule.From[0].PodSelector == nil {
		return true
	}
	exprs := kserveRule.From[0].PodSelector.MatchExpressions
	if len(exprs) != 1 || exprs[0].Key != kserveInferenceServiceLabel || exprs[0].Operator != metav1.LabelSelectorOpExists {
		return true
	}
	if len(kserveRule.Ports) != 2 {
		return true
	}
	oauthRule := existing.Spec.Ingress[2]
	if len(oauthRule.From) != 0 {
		return true
	}
	if len(oauthRule.Ports) != 1 || !reflect.DeepEqual(oauthRule.Ports[0].Port, &intstr.IntOrString{Type: intstr.Int, IntVal: 8443}) {
		return true
	}
	return false
}
