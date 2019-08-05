package controller

import (
	"github.com/Svimba/tungstenfabric-operator/pkg/controller/tfanalytics"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, tfanalytics.Add)
}
