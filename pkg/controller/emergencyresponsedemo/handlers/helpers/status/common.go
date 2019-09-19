package status

import (
	"github.com/integr8ly/erd-operator/pkg/apis/erdemo/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

func (h *Helper) InitStatus() v1alpha1.EmergencyResponseDemoStatus {
	reason := new(string)
	*reason = "Initialize"

	message := new(string)
	*message = "Emergency Response Demo is being initialized"

	lastTime := metav1.NewTime(time.Now())

	return v1alpha1.EmergencyResponseDemoStatus{
		Type:               v1alpha1.EmergencyResponseDemoInit,
		Status:             corev1.ConditionTrue,
		Reason:             reason,
		Message:            message,
		LastHeartbeatTime:  &lastTime,
		LastTransitionTime: &lastTime,
	}
}

func (h *Helper) DeleteStatus() v1alpha1.EmergencyResponseDemoStatus {
	reason := new(string)
	*reason = "Delete"

	message := new(string)
	*message = "Emergency Response Demo is being deleted"

	lastTime := metav1.NewTime(time.Now())

	return v1alpha1.EmergencyResponseDemoStatus{
		Type:               v1alpha1.EmergencyResponseDemoDelete,
		Status:             corev1.ConditionTrue,
		Reason:             reason,
		Message:            message,
		LastHeartbeatTime:  &lastTime,
		LastTransitionTime: &lastTime,
	}
}
