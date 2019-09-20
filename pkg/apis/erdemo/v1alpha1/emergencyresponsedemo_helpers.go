package v1alpha1

import "k8s.io/apimachinery/pkg/types"

func (in *EmergencyResponseDemo) SecretNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: in.Namespace,
		Name:      in.Spec.SecretName,
	}
}
