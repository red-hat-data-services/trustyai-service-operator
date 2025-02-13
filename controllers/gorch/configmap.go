package gorch

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

func (r *GuardrailsOrchestratorReconciler) getImageFromConfigMap(ctx context.Context, configMapKey string, configMapName string, namespace string) (string, error) {
	configMap := &corev1.ConfigMap{}
	err := r.Get(ctx, types.NamespacedName{Name: configMapName, Namespace: namespace}, configMap)
	if err != nil {
		if errors.IsNotFound(err) {
			return "", nil
		}
		return "", fmt.Errorf("error reading configmap %s", configMapName)
	}

	containerImage, ok := configMap.Data[configMapKey]

	if !ok {
		return "", fmt.Errorf("configmap %s does not contain necessary keys", configMapName)
	}
	return containerImage, nil
}
