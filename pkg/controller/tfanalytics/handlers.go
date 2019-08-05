package tfanalytics

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
func (r *ReconcileTFAnalytics) handleAPIDeployment() (bool, error) {
	// Define a new API deployment object
	apiDeployment := newDeploymentForAPI(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, apiDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this API deployment already exists
	foundAPIDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: apiDeployment.Name, Namespace: apiDeployment.Namespace}, foundAPIDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Analytics API deployment", "Deploy.Name", apiDeployment.Name)
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
	r.reqLogger.Info("Skip reconcile: Analytics API Deployment already exists", "Deploy.Name", foundAPIDeploy.Name)
	return false, nil
}

// API service handler
// create/update(if exists)
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleAPIService() (bool, error) {
	// Define a new API service object
	apiService := newServicesForAPI(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, apiService, r.scheme); err != nil {
		return false, err
	}
	// Check if this API Service already exists
	foundAPIService := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: apiService.Name, Namespace: apiService.Namespace}, foundAPIService)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Analytics API Service", "Service.Name", apiService.Name)
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
	r.reqLogger.Info("Skip reconcile: API Service already exists", "Service.Name", foundAPIService.Name)
	return false, nil
}

// AlarmGen deployment handler
// create/update(if exists) config AlarmGen deployment
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleAlarmGenDeployment() (bool, error) {
	// Define a new AlarmGen deployment object
	alarmGenDeployment := newDeploymentForAlarmGen(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, alarmGenDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this AlarmGen deployment already exists
	foundAlarmGenDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: alarmGenDeployment.Name, Namespace: alarmGenDeployment.Namespace}, foundAlarmGenDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new AlarmGen deployment", "Deploy.Name", alarmGenDeployment.Name)
		err = r.client.Create(context.TODO(), alarmGenDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: AlarmGen Deployment already exists", "Deploy.Name", foundAlarmGenDeploy.Name)
	return false, nil
}

// AlarmGen service handler
// create/update(if exists)
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleAlarmGenService() (bool, error) {
	// Define a new AlarmGen service object
	svcmService := newServicesForAlarmGen(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, svcmService, r.scheme); err != nil {
		return false, err
	}
	// Check if this AlarmGen Service already exists
	foundSvcmService := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: svcmService.Name, Namespace: svcmService.Namespace}, foundSvcmService)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new AlarmGen Service", "Service.Name", svcmService.Name)
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
	r.reqLogger.Info("Skip reconcile: AlarmGen Service already exists", "Service.Name", foundSvcmService.Name)
	return false, nil
}

// Collector deployment handler
// create/update(if exists) config Collector deployment
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleCollectorDeployment() (bool, error) {
	// Define a new Collector deployment object
	collectorDeployment := newDeploymentForCollector(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, collectorDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this Collector deployment already exists
	foundCollectorDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: collectorDeployment.Name, Namespace: collectorDeployment.Namespace}, foundCollectorDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Collector deployment", "Deploy.Name", collectorDeployment.Name)
		err = r.client.Create(context.TODO(), collectorDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: Collector Deployment already exists", "Deploy.Name", foundCollectorDeploy.Name)
	return false, nil
}

// Collector service handler
// create/update(if exists)
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleCollectorService() (bool, error) {
	// Define a new Collector service object
	collectorService := newServicesForCollector(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, collectorService, r.scheme); err != nil {
		return false, err
	}
	// Check if this Collector Service already exists
	foundCollectorService := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: collectorService.Name, Namespace: collectorService.Namespace}, foundCollectorService)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Collector Service", "Service.Name", collectorService.Name)
		err = r.client.Create(context.TODO(), collectorService)
		if err != nil {
			return false, err
		}
		// Service has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Service already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: Collector Service already exists", "Service.Name", foundCollectorService.Name)
	return false, nil
}

// QueryEngine deployment handler
// create/update(if exists) config QueryEngine deployment
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleQueryEngineDeployment() (bool, error) {
	// Define a new QueryEngine deployment object
	queryEngineDeployment := newDeploymentForQueryEngine(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, queryEngineDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this QueryEngine deployment already exists
	foundQueryEngineDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: queryEngineDeployment.Name, Namespace: queryEngineDeployment.Namespace}, foundQueryEngineDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new QueryEngine deployment", "Deploy.Name", queryEngineDeployment.Name)
		err = r.client.Create(context.TODO(), queryEngineDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: QueryEngine Deployment already exists", "Deploy.Name", foundQueryEngineDeploy.Name)
	return false, nil
}

// QueryEngine service handler
// create/update(if exists)
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleQueryEngineService() (bool, error) {
	// Define a new QueryEngine service object
	queryEngineService := newServicesForQueryEngine(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, queryEngineService, r.scheme); err != nil {
		return false, err
	}
	// Check if this QueryEngine Service already exists
	foundQueryEngineService := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: queryEngineService.Name, Namespace: queryEngineService.Namespace}, foundQueryEngineService)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new QueryEngine Service", "Service.Name", queryEngineService.Name)
		err = r.client.Create(context.TODO(), queryEngineService)
		if err != nil {
			return false, err
		}
		// Service has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Service already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: QueryEngine Service already exists", "Service.Name", foundQueryEngineService.Name)
	return false, nil
}

// SNMP deployment handler
// create/update(if exists) config SNMP deployment
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleSNMPDeployment() (bool, error) {
	// Define a new SNMP deployment object
	snmpDeployment := newDeploymentForSNMP(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, snmpDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this SNMP deployment already exists
	foundSNMPDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: snmpDeployment.Name, Namespace: snmpDeployment.Namespace}, foundSNMPDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new SNMP deployment", "Deploy.Name", snmpDeployment.Name)
		err = r.client.Create(context.TODO(), snmpDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: SNMP Deployment already exists", "Deploy.Name", foundSNMPDeploy.Name)
	return false, nil
}

// SNMP service handler
// create/update(if exists)
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleSNMPService() (bool, error) {
	// Define a new SNMP service object
	snmpService := newServicesForSNMP(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, snmpService, r.scheme); err != nil {
		return false, err
	}
	// Check if this SNMP Service already exists
	foundSNMPService := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: snmpService.Name, Namespace: snmpService.Namespace}, foundSNMPService)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new SNMP Service", "Service.Name", snmpService.Name)
		err = r.client.Create(context.TODO(), snmpService)
		if err != nil {
			return false, err
		}
		// Service has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Service already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: SNMP Service already exists", "Service.Name", foundSNMPService.Name)
	return false, nil
}

// Topology deployment handler
// create/update(if exists) config Topology deployment
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleTopologyDeployment() (bool, error) {
	// Define a new Topology deployment object
	topologyDeployment := newDeploymentForTopology(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, topologyDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this Topology deployment already exists
	foundTopologyDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: topologyDeployment.Name, Namespace: topologyDeployment.Namespace}, foundTopologyDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Topology deployment", "Deploy.Name", topologyDeployment.Name)
		err = r.client.Create(context.TODO(), topologyDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: Topology Deployment already exists", "Deploy.Name", foundTopologyDeploy.Name)
	return false, nil
}

// Topology service handler
// create/update(if exists)
// return true/false(Requeue), error
func (r *ReconcileTFAnalytics) handleTopologyService() (bool, error) {
	// Define a new Topology service object
	topologyService := newServicesForTopology(r.instance)
	// Set TFAnalytics instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, topologyService, r.scheme); err != nil {
		return false, err
	}
	// Check if this Topology Service already exists
	foundTopologyService := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: topologyService.Name, Namespace: topologyService.Namespace}, foundTopologyService)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Topology Service", "Service.Name", topologyService.Name)
		err = r.client.Create(context.TODO(), topologyService)
		if err != nil {
			return false, err
		}
		// Service has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Service already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: Topology Service already exists", "Service.Name", foundTopologyService.Name)
	return false, nil
}
