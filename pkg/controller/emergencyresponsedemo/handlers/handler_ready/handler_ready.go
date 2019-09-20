package handler_ready

import (
	"context"
	"github.com/integr8ly/erd-operator/pkg/apis/erdemo/v1alpha1"
	"github.com/integr8ly/erd-operator/pkg/controller/emergencyresponsedemo/handlers/helpers/service"
	"github.com/integr8ly/erd-operator/pkg/controller/emergencyresponsedemo/handlers/helpers/status"
	"github.com/integr8ly/erd-operator/pkg/controller/emergencyresponsedemo/handlers/meta"
	"github.com/integr8ly/erd-operator/pkg/lib/kube/metahelper"
	corev1 "k8s.io/api/core/v1"
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
	hasChanges, err := h.ensure(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	if hasChanges {
		//TODO
		//err = h.setInitStatus(instance)
		//if err != nil {
		//	return reconcile.Result{}, err
		//}
	}

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

	return reconcile.Result{RequeueAfter: time.Minute * 3}, nil
}

func (h *handler) ensure(instance *v1alpha1.EmergencyResponseDemo) (bool, error) {
	hasChanges := false

	metaHelper, err := metahelper.New(instance)
	if err != nil {
		return false, err
	}

	if !metaHelper.HasFinalizer(v1alpha1.SchemeGroupVersion.Group) {
		metaHelper.AddFinalizer(v1alpha1.SchemeGroupVersion.Group)
		instance.SetFinalizers(metaHelper.Finalizers())
		hasChanges = true
	}

	if hasChanges {
		err = h.client.Update(context.TODO(), instance)
		if err != nil {
			return hasChanges, err
		}
	}

	return hasChanges, nil
}

func (h *handler) checkServices(secret *corev1.Secret) error {
	return h.serviceHelper.CheckServices(secret)
}

func (h *handler) getSecret(instance *v1alpha1.EmergencyResponseDemo) (*corev1.Secret, error) {
	secret := &corev1.Secret{}

	err := h.client.Get(context.TODO(), instance.SecretNamespacedName(), secret)

	return secret, err
}

func (h *handler) setInitStatus(instance *v1alpha1.EmergencyResponseDemo) error {
	instance.Status = h.statusHelper.InitStatus()

	return h.client.Status().Update(context.TODO(), instance)
}

func (h *handler) setServiceErrorStatus(instance *v1alpha1.EmergencyResponseDemo, err error) error {
	instance.Status = h.statusHelper.ServiceError(err)

	return h.client.Status().Update(context.TODO(), instance)
}

func (h *handler) setSecretErrorStatus(instance *v1alpha1.EmergencyResponseDemo) error {
	instance.Status = h.statusHelper.SecretError()

	return h.client.Status().Update(context.TODO(), instance)
}
