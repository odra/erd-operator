package handler_init

import (
	"context"
	"github.com/integr8ly/erd-operator/pkg/apis/erdemo/v1alpha1"
	"github.com/integr8ly/erd-operator/pkg/controller/emergencyresponsedemo/handlers/helpers/service"
	"github.com/integr8ly/erd-operator/pkg/controller/emergencyresponsedemo/handlers/helpers/status"
	"github.com/integr8ly/erd-operator/pkg/controller/emergencyresponsedemo/handlers/meta"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

type handler struct {
	client        client.Client
	statusHelper  meta.StatusBluePrint
	serviceHelper meta.ServiceBluePrint
}

func New(c client.Client) *handler {
	return &handler{
		client:        c,
		statusHelper:  &status.Helper{},
		serviceHelper: &service.Helper{},
	}
}

func (h *handler) Handle(instance *v1alpha1.EmergencyResponseDemo) (reconcile.Result, error) {
	secret, err := h.getSecret(instance)
	if err != nil {
		if innerErr := h.setSecretErrorStatus(instance); innerErr != nil {
			return reconcile.Result{}, innerErr
		}
		return reconcile.Result{}, err
	}

	err = h.checkServices(secret)
	if err != nil {
		if innerErr := h.setServiceErrorStatus(instance, err); innerErr != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, err
	}

	err = h.setReadyStatus(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (h *handler) getSecret(instance *v1alpha1.EmergencyResponseDemo) (*corev1.Secret, error) {
	secret := &corev1.Secret{}

	err := h.client.Get(context.TODO(), instance.SecretNamespacedName(), secret)

	return secret, err
}

func (h *handler) setSecretErrorStatus(instance *v1alpha1.EmergencyResponseDemo) error {
	instance.Status = h.statusHelper.SecretError()

	return h.client.Status().Update(context.TODO(), instance)
}

func (h *handler) checkServices(secret *corev1.Secret) error {
	return h.serviceHelper.CheckServices(secret)
}

func (h *handler) setServiceErrorStatus(instance *v1alpha1.EmergencyResponseDemo, err error) error {
	instance.Status = h.statusHelper.ServiceError(err)

	return h.client.Status().Update(context.TODO(), instance)
}

func (h *handler) setReadyStatus(instance *v1alpha1.EmergencyResponseDemo) error {
	reason := new(string)
	*reason = "Ready"

	message := new(string)
	*message = "Emergency Response Demo is ready"

	lastTime := metav1.NewTime(time.Now())

	instance.Status = v1alpha1.EmergencyResponseDemoStatus{
		Type:               v1alpha1.EmergencyResponseDemoReady,
		Status:             corev1.ConditionTrue,
		Reason:             reason,
		Message:            message,
		LastHeartbeatTime:  &lastTime,
		LastTransitionTime: &lastTime,
	}

	return h.client.Status().Update(context.TODO(), instance)
}
