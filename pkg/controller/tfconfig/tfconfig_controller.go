package tfconfig

import (
	"context"
	"fmt"

	configv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/config/v1alpha1"
	"github.com/go-logr/logr"
	betav1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_tfconfig")

// Add creates a new TFConfig Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileTFConfig{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("tfconfig-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource TFConfig
	err = c.Watch(&source.Kind{Type: &configv1alpha1.TFConfig{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner TFConfig
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &configv1alpha1.TFConfig{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Deployments and requeue the owner TFConfig
	err = c.Watch(&source.Kind{Type: &betav1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &configv1alpha1.TFConfig{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Services and requeue the owner TFConfig
	err = c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &configv1alpha1.TFConfig{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileTFConfig{}

// ReconcileTFConfig reconciles a TFConfig object
type ReconcileTFConfig struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client    client.Client
	scheme    *runtime.Scheme
	instance  *configv1alpha1.TFConfig
	reqLogger logr.Logger
}

// Hndl is Structure for handler for function for extend params
type Hndl struct {
	Name    string
	Func    func() (bool, error)
	Enabled bool
}

// Reconcile reads that state of the cluster for a TFConfig object and makes changes based on the state read
// and what is in the TFConfig.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileTFConfig) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	r.reqLogger = log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	r.reqLogger.Info("Reconciling TFConfig")

	// Fetch the TFConfig instance
	r.instance = &configv1alpha1.TFConfig{}
	err := r.client.Get(context.TODO(), request.NamespacedName, r.instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Set list of available ConfigMaps
	err = r.setAvailableConfigMaps()
	if err != nil {
		r.reqLogger.Error(err, "In setAvailableConfigMaps")
		return reconcile.Result{}, err
	}
	r.reqLogger.Info(fmt.Sprintf("List of available ConfigMaps: %s", r.instance.Status.ConfigMapList))

	var handlerList []Hndl
	handlerList = append(handlerList, Hndl{Name: "ApiDeployment", Func: r.handleAPIDeployment, Enabled: r.instance.Spec.APISpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "ApiService", Func: r.handleAPIService, Enabled: r.instance.Spec.APISpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "SVCMonitorDeployment", Func: r.handleSVCMonitorDeployment, Enabled: r.instance.Spec.SVCMonitorSpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "SVCMonitorService", Func: r.handleSVCMonitorService, Enabled: r.instance.Spec.SVCMonitorSpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "SchemaDeployment", Func: r.handleSchemaDeployment, Enabled: r.instance.Spec.SchemaSpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "DeviceMgrDeployment", Func: r.handleDeviceMgrDeployment, Enabled: r.instance.Spec.DeviceMgrSpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "DeviceMgrService", Func: r.handleDeviceMgrService, Enabled: r.instance.Spec.DeviceMgrSpec.Enabled})

	for _, handler := range handlerList {
		if !handler.Enabled {
			continue
		}
		requeue, err := handler.Func()
		if err != nil {
			return reconcile.Result{}, err
		}
		if requeue {
			return reconcile.Result{Requeue: true}, nil
		}
	}
	return reconcile.Result{}, nil
}

// setAvailableConfigMaps Prepares and sets list of available config maps to Status.CofigMapList
func (r *ReconcileTFConfig) setAvailableConfigMaps() error {

	var listAvailableCfgMaps []string
	for _, cfm := range r.instance.Spec.CofigMapList {
		exists, err := r.checkConfigMapExists(cfm)
		if err != nil {
			return err
		}
		if exists {
			listAvailableCfgMaps = append(listAvailableCfgMaps, cfm)
		}
	}
	r.instance.Status.ConfigMapList = listAvailableCfgMaps
	if err := r.client.Status().Update(context.TODO(), r.instance); err != nil {
		return err
	}
	return nil
}

// checkConfigMapExists check if config map exists in K8s API
func (r *ReconcileTFConfig) checkConfigMapExists(name string) (bool, error) {
	foundCfgMap := &corev1.ConfigMap{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: r.instance.Namespace}, foundCfgMap)
	if err != nil && errors.IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
