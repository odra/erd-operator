package handler_error

import (
	"context"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/controller/emergencyresponsedemo/handlers/helpers/status"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/controller/emergencyresponsedemo/handlers/meta"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
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
	now := time.Now()
	latest := instance.Status.LastHeartbeatTime.Time
	diff := now.Sub(latest).Minutes()

	if diff < 0.1 {
		return reconcile.Result{}, nil
	}

	err := h.setInitStatus(instance)
	if err != nil {
		return reconcile.Result{RequeueAfter: time.Minute * 3}, err
	}

	return reconcile.Result{Requeue: true}, nil
}

func (h *handler) setInitStatus(instance *v1alpha1.EmergencyResponseDemo) error {
	instance.Status = h.statusHelper.InitStatus()

	return h.client.Status().Update(context.TODO(), instance)
}
