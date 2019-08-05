package tfanalytics

import (
	"context"
	"fmt"

	analyticsv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/analytics/v1alpha1"
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

var log = logf.Log.WithName("controller_tfanalytics")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new TFAnalytics Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileTFAnalytics{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("tfanalytics-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource TFAnalytics
	err = c.Watch(&source.Kind{Type: &analyticsv1alpha1.TFAnalytics{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner TFAnalytics
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &analyticsv1alpha1.TFAnalytics{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Deployments and requeue the owner TFConfig
	err = c.Watch(&source.Kind{Type: &betav1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &analyticsv1alpha1.TFAnalytics{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Services and requeue the owner TFConfig
	err = c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &analyticsv1alpha1.TFAnalytics{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileTFAnalytics implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileTFAnalytics{}

// ReconcileTFAnalytics reconciles a TFAnalytics object
type ReconcileTFAnalytics struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client    client.Client
	scheme    *runtime.Scheme
	instance  *analyticsv1alpha1.TFAnalytics
	reqLogger logr.Logger
}

// Hndl is Structure for handler for function for extend params
type Hndl struct {
	Name    string
	Func    func() (bool, error)
	Enabled bool
}

// Reconcile reads that state of the cluster for a TFAnalytics object and makes changes based on the state read
// and what is in the TFAnalytics.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileTFAnalytics) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	r.reqLogger = log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	r.reqLogger.Info("Reconciling TFAnalytics")

	// Fetch the TFAnalytics instance
	r.instance = &analyticsv1alpha1.TFAnalytics{}
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
	handlerList = append(handlerList, Hndl{Name: "AlarmGenDeployment", Func: r.handleAlarmGenDeployment, Enabled: r.instance.Spec.AlarmGenSpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "AlarmGenService", Func: r.handleAlarmGenService, Enabled: r.instance.Spec.AlarmGenSpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "CollectorDeployment", Func: r.handleCollectorDeployment, Enabled: r.instance.Spec.CollectorSpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "CollectorService", Func: r.handleCollectorService, Enabled: r.instance.Spec.CollectorSpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "QueryEngineDeployment", Func: r.handleQueryEngineDeployment, Enabled: r.instance.Spec.QueryEngine.Enabled})
	handlerList = append(handlerList, Hndl{Name: "QueryEngineService", Func: r.handleQueryEngineService, Enabled: r.instance.Spec.QueryEngine.Enabled})
	handlerList = append(handlerList, Hndl{Name: "SNMPDeployment", Func: r.handleSNMPDeployment, Enabled: r.instance.Spec.SNMPSpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "SNMPService", Func: r.handleSNMPService, Enabled: r.instance.Spec.SNMPSpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "TopologyDeployment", Func: r.handleTopologyDeployment, Enabled: r.instance.Spec.TopologySpec.Enabled})
	handlerList = append(handlerList, Hndl{Name: "TopologyService", Func: r.handleTopologyService, Enabled: r.instance.Spec.TopologySpec.Enabled})

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

// setAvailableConfigMaps Prepares and sets list of available config maps to Status.ConfigMapList
func (r *ReconcileTFAnalytics) setAvailableConfigMaps() error {

	var listAvailableCfgMaps []string
	for _, cfm := range r.instance.Spec.ConfigMapList {
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
func (r *ReconcileTFAnalytics) checkConfigMapExists(name string) (bool, error) {
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
