package tas

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
)

var _ = Describe("NetworkPolicy Reconciliation", func() {

	BeforeEach(func() {
		recorder = record.NewFakeRecorder(10)
		reconciler = &TrustyAIServiceReconciler{
			Client:        k8sClient,
			Scheme:        scheme.Scheme,
			EventRecorder: recorder,
			Namespace:     operatorNamespace,
		}
		ctx = context.Background()
	})

	Context("When ensuring a NetworkPolicy for a TrustyAIService", func() {
		It("Should create a NetworkPolicy with the correct spec", func() {
			namespace := "np-test-namespace-1"
			instance := createDefaultPVCCustomResource(namespace)

			WaitFor(func() error {
				return createNamespace(ctx, k8sClient, namespace)
			}, "failed to create namespace")

			WaitFor(func() error {
				return reconciler.ensureNetworkPolicy(ctx, instance)
			}, "failed to create NetworkPolicy")

			np := &networkingv1.NetworkPolicy{}
			npName := networkPolicyName(instance.Name)
			err := k8sClient.Get(ctx, types.NamespacedName{Name: npName, Namespace: namespace}, np)
			Expect(err).NotTo(HaveOccurred())

			Expect(np.Name).To(Equal(npName))
			Expect(np.Namespace).To(Equal(namespace))

			// Pod selector targets TrustyAI pods
			Expect(np.Spec.PodSelector.MatchLabels).To(HaveKeyWithValue("app", instance.Name))
			Expect(np.Spec.PodSelector.MatchLabels).To(HaveKeyWithValue("app.kubernetes.io/instance", instance.Name))

			// PolicyTypes is Ingress only
			Expect(np.Spec.PolicyTypes).To(HaveLen(1))
			Expect(np.Spec.PolicyTypes[0]).To(Equal(networkingv1.PolicyTypeIngress))

			// Three ingress rules
			Expect(np.Spec.Ingress).To(HaveLen(3))
		})

		It("Should allow KServe predictor pods on inference ports", func() {
			namespace := "np-test-namespace-2"
			instance := createDefaultPVCCustomResource(namespace)

			WaitFor(func() error {
				return createNamespace(ctx, k8sClient, namespace)
			}, "failed to create namespace")

			WaitFor(func() error {
				return reconciler.ensureNetworkPolicy(ctx, instance)
			}, "failed to create NetworkPolicy")

			np := &networkingv1.NetworkPolicy{}
			err := k8sClient.Get(ctx, types.NamespacedName{Name: networkPolicyName(instance.Name), Namespace: namespace}, np)
			Expect(err).NotTo(HaveOccurred())

			kserveRule := np.Spec.Ingress[0]

			// From: KServe pods with label existence selector
			Expect(kserveRule.From).To(HaveLen(1))
			Expect(kserveRule.From[0].PodSelector).NotTo(BeNil())
			Expect(kserveRule.From[0].PodSelector.MatchExpressions).To(HaveLen(1))
			Expect(kserveRule.From[0].PodSelector.MatchExpressions[0].Key).To(Equal(kserveInferenceServiceLabel))
			Expect(kserveRule.From[0].PodSelector.MatchExpressions[0].Operator).To(Equal(metav1.LabelSelectorOpExists))

			// Ports: 8080 and 4443
			Expect(kserveRule.Ports).To(HaveLen(2))
			port8080 := intstr.FromInt32(8080)
			port4443 := intstr.FromInt32(4443)
			Expect(kserveRule.Ports[0].Port).To(Equal(&port8080))
			Expect(kserveRule.Ports[1].Port).To(Equal(&port4443))
		})

		It("Should allow Prometheus scraping on port 8080", func() {
			namespace := "np-test-namespace-3"
			instance := createDefaultPVCCustomResource(namespace)

			WaitFor(func() error {
				return createNamespace(ctx, k8sClient, namespace)
			}, "failed to create namespace")

			WaitFor(func() error {
				return reconciler.ensureNetworkPolicy(ctx, instance)
			}, "failed to create NetworkPolicy")

			np := &networkingv1.NetworkPolicy{}
			err := k8sClient.Get(ctx, types.NamespacedName{Name: networkPolicyName(instance.Name), Namespace: namespace}, np)
			Expect(err).NotTo(HaveOccurred())

			prometheusRule := np.Spec.Ingress[1]

			// From: monitoring namespaces
			Expect(prometheusRule.From).To(HaveLen(1))
			Expect(prometheusRule.From[0].NamespaceSelector).NotTo(BeNil())
			Expect(prometheusRule.From[0].NamespaceSelector.MatchExpressions).To(HaveLen(1))
			nsExpr := prometheusRule.From[0].NamespaceSelector.MatchExpressions[0]
			Expect(nsExpr.Key).To(Equal("kubernetes.io/metadata.name"))
			Expect(nsExpr.Operator).To(Equal(metav1.LabelSelectorOpIn))
			Expect(nsExpr.Values).To(ConsistOf("openshift-monitoring", "openshift-user-workload-monitoring"))

			// Port: 8080 only
			Expect(prometheusRule.Ports).To(HaveLen(1))
			port8080 := intstr.FromInt32(8080)
			Expect(prometheusRule.Ports[0].Port).To(Equal(&port8080))
		})

		It("Should allow all traffic on oauth-proxy port 8443", func() {
			namespace := "np-test-namespace-4"
			instance := createDefaultPVCCustomResource(namespace)

			WaitFor(func() error {
				return createNamespace(ctx, k8sClient, namespace)
			}, "failed to create namespace")

			WaitFor(func() error {
				return reconciler.ensureNetworkPolicy(ctx, instance)
			}, "failed to create NetworkPolicy")

			np := &networkingv1.NetworkPolicy{}
			err := k8sClient.Get(ctx, types.NamespacedName{Name: networkPolicyName(instance.Name), Namespace: namespace}, np)
			Expect(err).NotTo(HaveOccurred())

			oauthRule := np.Spec.Ingress[2]

			// No From restriction — open to all
			Expect(oauthRule.From).To(BeEmpty())

			// Port: 8443 only
			Expect(oauthRule.Ports).To(HaveLen(1))
			port8443 := intstr.FromInt32(8443)
			Expect(oauthRule.Ports[0].Port).To(Equal(&port8443))
		})

		It("Should be idempotent on repeated calls", func() {
			namespace := "np-test-namespace-5"
			instance := createDefaultPVCCustomResource(namespace)

			WaitFor(func() error {
				return createNamespace(ctx, k8sClient, namespace)
			}, "failed to create namespace")

			WaitFor(func() error {
				return reconciler.ensureNetworkPolicy(ctx, instance)
			}, "failed to create NetworkPolicy on first call")

			// Second call should succeed without error
			err := reconciler.ensureNetworkPolicy(ctx, instance)
			Expect(err).NotTo(HaveOccurred())

			// Should still be exactly one NetworkPolicy
			np := &networkingv1.NetworkPolicy{}
			err = k8sClient.Get(ctx, types.NamespacedName{Name: networkPolicyName(instance.Name), Namespace: namespace}, np)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
