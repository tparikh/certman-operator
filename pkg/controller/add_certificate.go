package controller

import (
	"github.com/certman-operator/pkg/controller/certificate"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, certificate.Add)
}
