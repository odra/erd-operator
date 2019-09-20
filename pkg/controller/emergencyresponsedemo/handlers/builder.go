package handlers

import (
	"errors"
	"fmt"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/apis/erdemo/v1alpha1"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/controller/emergencyresponsedemo/handlers/handler_delete"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/controller/emergencyresponsedemo/handlers/handler_error"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/controller/emergencyresponsedemo/handlers/handler_init"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/controller/emergencyresponsedemo/handlers/handler_new"
	"github.com/Emergency-Response-Demo/erd-operator/pkg/controller/emergencyresponsedemo/handlers/handler_ready"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Builder struct {
	client client.Client
}

func NewBuilder(client client.Client) Builder {
	return Builder{client: client}
}

func (b *Builder) Build(status v1alpha1.EmergencyResponseDemoConditionType) (HandlerBluePrint, error) {
	if b.client == nil {
		return nil, errors.New("builder client is nil")
	}

	switch status {
	case v1alpha1.EmergencyResponseDemoNew:
		return handler_new.New(b.client), nil
	case v1alpha1.EmergencyResponseDemoInit:
		return handler_init.New(b.client), nil
	case v1alpha1.EmergencyResponseDemoDelete:
		return handler_delete.New(b.client), nil
	case v1alpha1.EmergencyResponseDemoError:
		return handler_error.New(b.client), nil
	case v1alpha1.EmergencyResponseDemoReady:
		return handler_ready.New(b.client), nil
	default:
		return nil, fmt.Errorf("invalid status: %s", status)
	}
}
