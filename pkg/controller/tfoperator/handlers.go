package tfoperator

import (
	"context"

	configv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/config/v1alpha1"
	betav1 "k8s.io/api/apps/v1beta1"
	extbetav1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// Config Operator handler
// return true/false(Requeue), error
func (r *ReconcileTFOperator) handleConfigOperator() (bool, error) {
	// Define a new Config CRD object
	crdConfig := newCRDForConfig(r.instance)
	// Set TFOperator instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, crdConfig, r.scheme); err != nil {
		return false, err
	}
	// Check if this Config CRD already exists
	foundCRDConfig := &extbetav1.CustomResourceDefinition{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: crdConfig.Name, Namespace: crdConfig.Namespace}, foundCRDConfig)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Config CRD", "Deploy.Namespace", crdConfig.Namespace, "Deploy.Name", crdConfig.Name)
		err = r.client.Create(context.TODO(), crdConfig)
		if err != nil {
			return false, err
		}
		// CRD has been created successfully - don't requeue
	} else if err != nil {
		return false, err
	}
	// Config CRD already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: Config CRD already exists", "Deploy.Namespace", foundCRDConfig.Namespace, "Deploy.Name", foundCRDConfig.Name)

	// Define a new Config Operator object
	configOperator := newOperatorForConfig(r.instance)
	// Set TFOperator instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, configOperator, r.scheme); err != nil {
		return false, err
	}
	// Check if this Config Operator already exists
	foundConfigOperator := &betav1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: configOperator.Name, Namespace: configOperator.Namespace}, foundConfigOperator)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Config Operator", "Deploy.Namespace", configOperator.Namespace, "Deploy.Name", configOperator.Name)
		err = r.client.Create(context.TODO(), configOperator)
		if err != nil {
			return false, err
		}
		// Operator has been created successfully - don't requeue
	} else if err != nil {
		return false, err
	}
	// Config Operator already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: Config Operator already exists", "Deploy.Namespace", foundConfigOperator.Namespace, "Deploy.Name", foundConfigOperator.Name)

	// Define a new CR for Config Operator object
	crConfig := newCRForConfig(r.instance)
	// Set TFOperator instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, crConfig, r.scheme); err != nil {
		return false, err
	}
	// Check if this CR for Config Operator already exists
	foundCRConfig := &configv1alpha1.TFConfig{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: crConfig.Name, Namespace: crConfig.Namespace}, foundCRConfig)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new CR for Config Operator", "Deploy.Namespace", crConfig.Namespace, "Deploy.Name", crConfig.Name)
		err = r.client.Create(context.TODO(), crConfig)
		if err != nil {
			return false, err
		}
		// CR has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// CR for Config Operator already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: CR for Config Operator already exists", "Deploy.Namespace", foundCRConfig.Namespace, "Deploy.Name", foundCRConfig.Name)

	return false, nil

}