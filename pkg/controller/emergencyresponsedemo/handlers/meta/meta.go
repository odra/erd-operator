package meta

import (
	"github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type ServiceBluePrint interface {
	CheckServices(secret *corev1.Secret) error
}

type StatusBluePrint interface {
	InitStatus() v1alpha1.EmergencyResponseDemoStatus
	DeleteStatus() v1alpha1.EmergencyResponseDemoStatus
	ServiceError(err error) v1alpha1.EmergencyResponseDemoStatus
	SecretError() v1alpha1.EmergencyResponseDemoStatus
}
