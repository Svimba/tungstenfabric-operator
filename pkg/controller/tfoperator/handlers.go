package tfoperator

import (
	"context"
	"fmt"
	"reflect"

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

// CfgMapHandler is structure of handlers function
type CfgMapHandler struct {
	Name   string
	CfgMap *corev1.ConfigMap
}

func (r *ReconcileTFOperator) compareConfigMaps(newcm *corev1.ConfigMap, curcm *corev1.ConfigMap) bool {
	r.reqLogger.Info(fmt.Sprintf("Comparing: %v <<<>>> %v", newcm.Data, curcm.Data))
	return reflect.DeepEqual(newcm.Data, curcm.Data)
}

// ConfigMaps handler, prepare global configmaps for TF
// return true/false(Requeue), error
func (r *ReconcileTFOperator) handleConfigMaps() (bool, error) {
	var handlerList []CfgMapHandler
	handlerList = append(handlerList, CfgMapHandler{Name: "tf-zookeeper-cfgmap", CfgMap: getConfigMapForZookeeper(r.instance, r.defaults)})
	handlerList = append(handlerList, CfgMapHandler{Name: "tf-rabbitmq-cfgmap", CfgMap: getConfigMapForRabbitMQ(r.instance, r.defaults)})
	handlerList = append(handlerList, CfgMapHandler{Name: "tf-cassandra-cfgmap", CfgMap: getConfigMapForCassandra(r.instance, r.defaults)})

	for _, hdl := range handlerList {

		r.reqLogger.Info("Checking: ", hdl.Name, "cfgmap")
		if err := controllerutil.SetControllerReference(r.instance, hdl.CfgMap, r.scheme); err != nil {
			return false, err
		}
		// Check if this Config CRD already exists
		foundConfigMap := &corev1.ConfigMap{}
		err := r.client.Get(context.TODO(), types.NamespacedName{Name: hdl.CfgMap.Name, Namespace: hdl.CfgMap.Namespace}, foundConfigMap)
		if err != nil && errors.IsNotFound(err) {
			r.reqLogger.Info(fmt.Sprintf("Creating a new Config Map for %s Name %s", hdl.Name, hdl.CfgMap.Name))
			err = r.client.Create(context.TODO(), hdl.CfgMap)
			if err != nil {
				return false, err
			}
			// ConfigMap has been created successfully - don't requeue
		} else if err != nil {
			return false, err
		} else {
			// ConfigMap already exists - check if is updated
			r.reqLogger.Info(fmt.Sprintf("Skip reconcile: ConfigMap for %s already exists Name %s", hdl.Name, foundConfigMap.Name))
			updated := r.compareConfigMaps(hdl.CfgMap, foundConfigMap)
			r.reqLogger.Info(fmt.Sprintf("Checking ConfigMap is updated: %t", updated))
			if !updated {
				r.reqLogger.Info(fmt.Sprintf("Updating ConfigMap %s", hdl.Name))

				foundConfigMap.Data = hdl.CfgMap.Data
				err = r.client.Update(context.TODO(), foundConfigMap)
				if err != nil {
					r.reqLogger.Error(err, fmt.Sprintf("Cannot update ConfigMap %s", hdl.Name))
					return false, err
				}
			}
		}
	}
	return false, nil
}
