package handlers

import (
	"github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type HandlerBluePrint interface {
	Handle(instance *v1alpha1.EmergencyResponseDemo) (reconcile.Result, error)
}
