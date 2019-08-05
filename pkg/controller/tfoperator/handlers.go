package tfoperator

import (
	"context"
	"fmt"
	"reflect"

	analyticsv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/analytics/v1alpha1"
	configv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/config/v1alpha1"
	controlv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/control/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcileTFOperator) updateSpec(newSpec *configv1alpha1.TFConfigSpec, curSpec *configv1alpha1.TFConfigSpec) bool {
	changed := false
	if !reflect.DeepEqual(newSpec.APISpec.EnvList, curSpec.APISpec.EnvList) {
		curSpec.APISpec.EnvList = newSpec.APISpec.EnvList
		changed = true
	}
	if !reflect.DeepEqual(newSpec.SVCMonitorSpec.EnvList, curSpec.SVCMonitorSpec.EnvList) {
		curSpec.SVCMonitorSpec.EnvList = newSpec.SVCMonitorSpec.EnvList
		changed = true
	}
	if !reflect.DeepEqual(newSpec.SchemaSpec.EnvList, curSpec.SchemaSpec.EnvList) {
		curSpec.SchemaSpec.EnvList = newSpec.SchemaSpec.EnvList
		changed = true
	}
	if !reflect.DeepEqual(newSpec.DeviceMgrSpec.EnvList, curSpec.DeviceMgrSpec.EnvList) {
		curSpec.DeviceMgrSpec.EnvList = newSpec.DeviceMgrSpec.EnvList
		changed = true
	}
	return changed
}

// Config Operator handler
// return true/false(Requeue), error
func (r *ReconcileTFOperator) handleConfigOperator() (bool, error) {
	// Define a new CR for Config Operator object
	crConfig := newCRForConfig(r.instance, r.defaults)
	// Set TFOperator instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, crConfig, r.scheme); err != nil {
		return false, err
	}
	// Check if this CR for Config Operator already exists
	foundCRConfig := &configv1alpha1.TFConfig{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: crConfig.Name, Namespace: crConfig.Namespace}, foundCRConfig)
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
	} else {
		// Update envs if needed
		if r.updateSpec(&crConfig.Spec, &foundCRConfig.Spec) {
			r.reqLogger.Info(fmt.Sprintf("Updating CR for TFConfig"))
			err = r.client.Update(context.TODO(), foundCRConfig)
			if err != nil {
				r.reqLogger.Error(err, fmt.Sprintf("Cannot update CR for TFConfig"))
				return false, err
			}
		}
	}
	// CR for Config Operator already exists - don't requeue
	r.reqLogger.Info("CR for Config Operator already exists and looks updated", "Deploy.Name", foundCRConfig.Name)
	return false, nil
}

// Control Operator handler
// return true/false(Requeue), error
func (r *ReconcileTFOperator) handleControlOperator() (bool, error) {
	// Define a new CR for Control Operator object
	crControl := newCRForControl(r.instance, r.defaults)
	// Set TFOperator instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, crControl, r.scheme); err != nil {
		return false, err
	}
	// Check if this CR for Control Operator already exists
	foundCRControl := &controlv1alpha1.TFControl{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: crControl.Name, Namespace: crControl.Namespace}, foundCRControl)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new CR for Control Operator", "Deploy.Namespace", crControl.Namespace, "Deploy.Name", crControl.Name)
		err = r.client.Create(context.TODO(), crControl)
		if err != nil {
			return false, err
		}
		// CR has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// CR for Control Operator already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: CR for Control Operator already exists", "Deploy.Namespace", foundCRControl.Namespace, "Deploy.Name", foundCRControl.Name)

	return false, nil

}

// Analytics Operator handler
// return true/false(Requeue), error
func (r *ReconcileTFOperator) handleAnalyticsOperator() (bool, error) {
	// Define a new CR for Analytics Operator object
	crAnalytics := newCRForAnalytics(r.instance, r.defaults)
	// Set TFOperator instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, crAnalytics, r.scheme); err != nil {
		return false, err
	}
	// Check if this CR for Analytics Operator already exists
	foundCRAnalytics := &analyticsv1alpha1.TFAnalytics{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: crAnalytics.Name, Namespace: crAnalytics.Namespace}, foundCRAnalytics)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new CR for Analytics Operator", "Deploy.Namespace", crAnalytics.Namespace, "Deploy.Name", crAnalytics.Name)
		err = r.client.Create(context.TODO(), crAnalytics)
		if err != nil {
			return false, err
		}
		// CR has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// CR for Analytics Operator already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: CR for Analytics Operator already exists", "Deploy.Namespace", foundCRAnalytics.Namespace, "Deploy.Name", foundCRAnalytics.Name)

	return false, nil

}

// CfgMapHandler is structure of handlers function
type CfgMapHandler struct {
	Name   string
	CfgMap *corev1.ConfigMap
}

func (r *ReconcileTFOperator) updateConfigMaps(newcm *corev1.ConfigMap, curcm *corev1.ConfigMap) bool {
	if !reflect.DeepEqual(newcm.Data, curcm.Data) {
		curcm.Data = newcm.Data
		return true
	}
	return false
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
			if r.updateConfigMaps(hdl.CfgMap, foundConfigMap) {
				r.reqLogger.Info(fmt.Sprintf("Updating ConfigMap %s", hdl.Name))

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
