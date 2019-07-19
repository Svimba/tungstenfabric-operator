package controller

import (
	"github.com/Svimba/tungstenfabric-operator/pkg/controller/tfconfig"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, tfconfig.Add)
}
