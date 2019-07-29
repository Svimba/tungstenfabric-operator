package tfoperator

import (
	"context"
	"fmt"

	configv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/config/v1alpha1"
	corev1 "k8s.io/api/core/v1"
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

	// Define a new CR for Config Operator object
	crConfig := newCRForConfig(r.instance, r.defaults)
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

// ConfigMaps handler, prepare global configmaps for TF
// return true/false(Requeue), error
func (r *ReconcileTFOperator) handleConfigMaps() (bool, error) {

	r.reqLogger.Info(fmt.Sprintf("ZOOKEEPER PORTS: %v", r.defaults.Get("zookeeper").Services[0].Ports[0].Port))
	// Define a new ConfigMap object
	configMap := getConfigMapForRabbitMQ(r.instance, r.defaults)
	// Set TFOperator instance as the owner and controller
	r.reqLogger.Info(fmt.Sprintf("RabbitMQ CFGMAP: %s", configMap))
	if err := controllerutil.SetControllerReference(r.instance, configMap, r.scheme); err != nil {
		return false, err
	}
	// Check if this Config CRD already exists
	foundConfigMap := &corev1.ConfigMap{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: configMap.Name, Namespace: configMap.Namespace}, foundConfigMap)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Config Map for RabbitMQ", "Namespace", configMap.Namespace, "Name", configMap.Name)
		err = r.client.Create(context.TODO(), configMap)
		if err != nil {
			return false, err
		}
		// ConfigMap has been created successfully - don't requeue
	} else if err != nil {
		return false, err
	}
	// ConfigMap already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: ConfigMap for RabbitMQ already exists", "Namespace", foundConfigMap.Namespace, "Name", foundConfigMap.Name)
	return false, nil
}
