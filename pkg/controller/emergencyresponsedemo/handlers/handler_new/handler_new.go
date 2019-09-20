package handler_new

import (
	"context"
	"github.com/integr8ly/erd-operator/pkg/apis/erdemo/v1alpha1"
	"github.com/integr8ly/erd-operator/pkg/controller/emergencyresponsedemo/handlers/helpers/status"
	"github.com/integr8ly/erd-operator/pkg/controller/emergencyresponsedemo/handlers/meta"
	"github.com/integr8ly/erd-operator/pkg/lib/kube/metahelper"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type handler struct {
	client       client.Client
	statusHelper meta.StatusBluePrint
}

func New(c client.Client) *handler {
	return &handler{
		client:       c,
		statusHelper: &status.Helper{},
	}
}

func (h *handler) Handle(instance *v1alpha1.EmergencyResponseDemo) (reconcile.Result, error) {
	var err error

	err = h.bootstrap(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	err = h.setInitStatus(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (h *handler) bootstrap(instance *v1alpha1.EmergencyResponseDemo) error {
	metaHelper, err := metahelper.New(instance)
	if err != nil {
		return err
	}

	metaHelper.AddFinalizer(v1alpha1.SchemeGroupVersion.Group)
	instance.SetFinalizers(metaHelper.Finalizers())

	err = h.client.Update(context.TODO(), instance)
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) setInitStatus(instance *v1alpha1.EmergencyResponseDemo) error {
	instance.Status = h.statusHelper.InitStatus()

	return h.client.Status().Update(context.TODO(), instance)
}
