package controller

import (
	"github.com/Emergency-Response-Demo/erd-operator/pkg/controller/emergencyresponsedemo"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, emergencyresponsedemo.Add)
}
