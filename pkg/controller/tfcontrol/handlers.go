package tfcontrol

import (
	"context"
	"fmt"

	betav1 "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// Control deployment handler
// create/update(if exists) control process deployment
// return true/false(Requeue), error
func (r *ReconcileTFControl) handleControlDeployment() (bool, error) {
	// Define a new control deployment object
	controlDeployment := newDeploymentForControl(r.instance)
	// Set TFConfig instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, controlDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this API deployment already exists
	foundControlDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: controlDeployment.Name, Namespace: controlDeployment.Namespace}, foundControlDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Control service deployment", "Deploy.Namespace", controlDeployment.Namespace, "Deploy.Name", controlDeployment.Name)
		err = r.client.Create(context.TODO(), controlDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}

	// Check replicas of the deployment, update if needed
	if *foundControlDeploy.Spec.Replicas != *r.instance.Spec.ControlSpec.Replicas {
		r.reqLogger.Info(fmt.Sprintf("Current replicas: %d  Desired: %d reconfiguring...", int(*foundControlDeploy.Spec.Replicas), int(*r.instance.Spec.ControlSpec.Replicas)))
		foundControlDeploy.Spec.Replicas = r.instance.Spec.ControlSpec.Replicas
		err = r.client.Update(context.TODO(), foundControlDeploy)
		if err != nil {
			r.reqLogger.Error(err, "Cannot update replicas for deployment:", foundControlDeploy.Name)
			return false, err
		}
		r.reqLogger.Info("Replicas have been changed")
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: Control service Deployment already exists", "Deploy.Namespace", foundControlDeploy.Namespace, "Deploy.Name", foundControlDeploy.Name)
	return false, nil
}

// Named deployment handler
// create/update(if exists) control Named deployment
// return true/false(Requeue), error
func (r *ReconcileTFControl) handleNamedDeployment() (bool, error) {
	// Define a new Named deployment object
	namedDeployment := newDeploymentForNamed(r.instance)
	// Set TFConfig instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, namedDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this SVCMonitor deployment already exists
	foundNamedDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: namedDeployment.Name, Namespace: namedDeployment.Namespace}, foundNamedDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Named deployment", "Deploy.Namespace", namedDeployment.Namespace, "Deploy.Name", namedDeployment.Name)
		err = r.client.Create(context.TODO(), namedDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: Named Deployment already exists", "Deploy.Namespace", foundNamedDeploy.Namespace, "Deploy.Name", foundNamedDeploy.Name)
	return false, nil
}

// Dns deployment handler
// create/update(if exists) control dns deployment
// return true/false(Requeue), error
func (r *ReconcileTFControl) handleDnsDeployment() (bool, error) {
	// Define a new Dns deployment object
	dnsDeployment := newDeploymentForDns(r.instance)
	// Set TFControl instance as the owner and controller
	if err := controllerutil.SetControllerReference(r.instance, dnsDeployment, r.scheme); err != nil {
		return false, err
	}
	// Check if this Schema deployment already exists
	foundDnsDeploy := &betav1.Deployment{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: dnsDeployment.Name, Namespace: dnsDeployment.Namespace}, foundDnsDeploy)
	if err != nil && errors.IsNotFound(err) {
		r.reqLogger.Info("Creating a new Dns deployment", "Deploy.Namespace", dnsDeployment.Namespace, "Deploy.Name", dnsDeployment.Name)
		err = r.client.Create(context.TODO(), dnsDeployment)
		if err != nil {
			return false, err
		}
		// Deployment has been created successfully - don't requeue
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Deployment already exists - don't requeue
	r.reqLogger.Info("Skip reconcile: Dns Deployment already exists", "Deploy.Namespace", foundDnsDeploy.Namespace, "Deploy.Name", foundDnsDeploy.Name)
	return false, nil
}
