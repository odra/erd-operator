package status

import (
	"github.com/integr8ly/erd-operator/pkg/apis/erdemo/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

func (sh *Helper) ServiceError(err error) v1alpha1.EmergencyResponseDemoStatus {
	reason := new(string)
	*reason = "ERDServiceError"

	message := new(string)
	*message = err.Error()

	lastTime := metav1.NewTime(time.Now())

	return v1alpha1.EmergencyResponseDemoStatus{
		Type:               v1alpha1.EmergencyResponseDemoError,
		Status:             corev1.ConditionTrue,
		Reason:             reason,
		Message:            message,
		LastHeartbeatTime:  &lastTime,
		LastTransitionTime: &lastTime,
	}
}

func (sh *Helper) SecretError() v1alpha1.EmergencyResponseDemoStatus {
	reason := new(string)
	*reason = "ERDSecretNotFound"

	message := new(string)
	*message = "Could not find an ERD secret in this namespace"

	lastTime := metav1.NewTime(time.Now())

	return v1alpha1.EmergencyResponseDemoStatus{
		Type:               v1alpha1.EmergencyResponseDemoError,
		Status:             corev1.ConditionTrue,
		Reason:             reason,
		Message:            message,
		LastHeartbeatTime:  &lastTime,
		LastTransitionTime: &lastTime,
	}
}
