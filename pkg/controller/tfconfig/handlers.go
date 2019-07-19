package tfconfig

import (
	"context"
	"fmt"

	betav1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// API deployment handler
// create/update(if exists) config API deployment
// return true/false(Requeue), error
func (r *ReconcileTFConfig) handleAPIDeployment() (bool, error) {
	// Define a new API deployment object
	apiDeployment := newDeploymentForAPI(r.instance)
	// Set TFConfig instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, apiDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this API deployment already exists
	foundAPIDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: apiDeployment.Name, Namespace: apiDeployment.Namespace}, foundAPIDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new API deployment", "Deploy.Namespace", apiDeployment.Namespace, "Deploy.Name", apiDeployment.Name)
		err = r.client.Create(context.TODO(), apiDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}

	// Check replicas of the deployment, update if needed
	if *foundAPIDeploy.Spec.Replicas != *r.instance.Spec.APISpec.Replicas {
		r.reqLogger.Info(fmt.Sprintf("Current replicas: %d  Desired: %d reconfiguring...", int(*foundAPIDeploy.Spec.Replicas), int(*r.instance.Spec.APISpec.Replicas)))
		foundAPIDeploy.Spec.Replicas = r.instance.Spec.APISpec.Replicas
		err = r.client.Update(context.TODO(), foundAPIDeploy)
		if err != nil {
			r.reqLogger.Error(err, "Cannot update replicas for deployment:", foundAPIDeploy.Name)
			return false, err
		}
		r.reqLogger.Info("Replicas have been changed")
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: API Deployment already exists", "Deploy.Namespace", foundAPIDeploy.Namespace, "Deploy.Name", foundAPIDeploy.Name)
	return false, nil
}

// API service handler
// create/update(if exists)
// return true/false(Requeue), error
func (r *ReconcileTFConfig) handleAPIService() (bool, error) {
	// Define a new API service object
	apiService := newServicesForAPI(r.instance)
	// Set TFConfig instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, apiService, r.scheme); err != nil {
		return false, err
	}
	// Check if this API Service already exists
	foundAPIService := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: apiService.Name, Namespace: apiService.Namespace}, foundAPIService)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new API Service", "Service.Namespace", apiService.Namespace, "Service.Name", apiService.Name)
		err = r.client.Create(context.TODO(), apiService)
		if err != nil {
			return false, err
		}
		// Service has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Service already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: API Service already exists", "Service.Namespace", foundAPIService.Namespace, "Service.Name", foundAPIService.Name)
	return false, nil
}

// SVCMonitor deployment handler
// create/update(if exists) config SVCMonitor deployment
// return true/false(Requeue), error
func (r *ReconcileTFConfig) handleSVCMonitorDeployment() (bool, error) {
	// Define a new SVCMonitor deployment object
	svcmDeployment := newDeploymentForSVCMonitor(r.instance)
	// Set TFConfig instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, svcmDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this SVCMonitor deployment already exists
	foundSVCMonitorDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: svcmDeployment.Name, Namespace: svcmDeployment.Namespace}, foundSVCMonitorDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new SVCMonitor deployment", "Deploy.Namespace", svcmDeployment.Namespace, "Deploy.Name", svcmDeployment.Name)
		err = r.client.Create(context.TODO(), svcmDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: SVCMonitor Deployment already exists", "Deploy.Namespace", foundSVCMonitorDeploy.Namespace, "Deploy.Name", foundSVCMonitorDeploy.Name)
	return false, nil
}

// SVCMonitor service handler
// create/update(if exists)
// return true/false(Requeue), error
func (r *ReconcileTFConfig) handleSVCMonitorService() (bool, error) {
	// Define a new SVCMonitor service object
	svcmService := newServicesForSVCMonitor(r.instance)
	// Set TFConfig instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, svcmService, r.scheme); err != nil {
		return false, err
	}
	// Check if this SVCMonitor Service already exists
	foundSvcmService := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: svcmService.Name, Namespace: svcmService.Namespace}, foundSvcmService)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new SVCMonitor Service", "Service.Namespace", svcmService.Namespace, "Service.Name", svcmService.Name)
		err = r.client.Create(context.TODO(), svcmService)
		if err != nil {
			return false, err
		}
		// Service has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Service already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: SVCMonitor Service already exists", "Service.Namespace", foundSvcmService.Namespace, "Service.Name", foundSvcmService.Name)
	return false, nil
}

// Schema deployment handler
// create/update(if exists) config Schema deployment
// return true/false(Requeue), error
func (r *ReconcileTFConfig) handleSchemaDeployment() (bool, error) {
	// Define a new Schema deployment object
	schemaDeployment := newDeploymentForSchema(r.instance)
	// Set TFConfig instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, schemaDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this Schema deployment already exists
	foundSchemaDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: schemaDeployment.Name, Namespace: schemaDeployment.Namespace}, foundSchemaDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Schema deployment", "Deploy.Namespace", schemaDeployment.Namespace, "Deploy.Name", schemaDeployment.Name)
		err = r.client.Create(context.TODO(), schemaDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: Schema Deployment already exists", "Deploy.Namespace", foundSchemaDeploy.Namespace, "Deploy.Name", foundSchemaDeploy.Name)
	return false, nil
}

// DeviceMgr deployment handler
// create/update(if exists) config DeviceMgr deployment
// return true/false(Requeue), error
func (r *ReconcileTFConfig) handleDeviceMgrDeployment() (bool, error) {
	// Define a new DeviceMgr deployment object
	devicemgrDeployment := newDeploymentForDeviceMgr(r.instance)
	// Set TFConfig instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, devicemgrDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this DeviceMgr deployment already exists
	foundDeviceMgrDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: devicemgrDeployment.Name, Namespace: devicemgrDeployment.Namespace}, foundDeviceMgrDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new DeviceMgr deployment", "Deploy.Namespace", devicemgrDeployment.Namespace, "Deploy.Name", devicemgrDeployment.Name)
		err = r.client.Create(context.TODO(), devicemgrDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: DeviceMgr Deployment already exists", "Deploy.Namespace", foundDeviceMgrDeploy.Namespace, "Deploy.Name", foundDeviceMgrDeploy.Name)
	return false, nil
}

// DeviceMgr service handler
// create/update(if exists)
// return true/false(Requeue), error
func (r *ReconcileTFConfig) handleDeviceMgrService() (bool, error) {
	// Define a new DeviceMgr service object
	devicemgrService := newServicesForDeviceMgr(r.instance)
	// Set TFConfig instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, devicemgrService, r.scheme); err != nil {
		return false, err
	}
	// Check if this DeviceMgr Service already exists
	foundDeviceMgrService := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: devicemgrService.Name, Namespace: devicemgrService.Namespace}, foundDeviceMgrService)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new DeviceMgr Service", "Service.Namespace", devicemgrService.Namespace, "Service.Name", devicemgrService.Name)
		err = r.client.Create(context.TODO(), devicemgrService)
		if err != nil {
			return false, err
		}
		// Service has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Service already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: DeviceMgr Service already exists", "Service.Namespace", foundDeviceMgrService.Namespace, "Service.Name", foundDeviceMgrService.Name)
	return false, nil
}
