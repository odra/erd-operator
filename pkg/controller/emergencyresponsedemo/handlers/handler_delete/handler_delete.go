package handler_delete

import (
	"context"
	"errors"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/lib/kube/metahelper"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type handler struct {
	client client.Client
}

func New(c client.Client) *handler {
	return &handler{
		client: c,
	}
}

func (h *handler) Handle(instance *v1alpha1.EmergencyResponseDemo) (reconcile.Result, error) {
	err := h.removeFinalizer(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (h *handler) removeFinalizer(instance *v1alpha1.EmergencyResponseDemo) error {
	if instance.Status.Status != corev1.ConditionTrue {
		return errors.New("waiting for error status condition to be ready")
	}

	metaHelper, err := metahelper.New(instance)
	if err != nil {
		return err
	}

	metaHelper.RemoveFinalizer(v1alpha1.SchemeGroupVersion.Group)
	instance.SetFinalizers(metaHelper.Finalizers())

	return h.client.Update(context.TODO(), instance)
}
