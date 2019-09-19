package emergencyresponsedemo

import (
	"context"
	erdemov1alpha1 "github.com/integr8ly/erd-operator/pkg/apis/erdemo/v1alpha1"
	erdHandlers "github.com/integr8ly/erd-operator/pkg/controller/emergencyresponsedemo/handlers"
	"github.com/integr8ly/erd-operator/pkg/controller/emergencyresponsedemo/handlers/helpers/status"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_emergencyresponsedemo")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new EmergencyResponseDemo Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileEmergencyResponseDemo{
		client:  mgr.GetClient(),
		config:  mgr.GetConfig(),
		builder: erdHandlers.NewBuilder(mgr.GetClient()),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("emergencyresponsedemo-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource EmergencyResponseDemo
	err = c.Watch(&source.Kind{Type: &erdemov1alpha1.EmergencyResponseDemo{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type : &corev1.Secret{}}, &handler.EnqueueRequestsFromMapFunc{
		ToRequests:handler.ToRequestsFunc(func(a handler.MapObject) []reconcile.Request {
			requests := make([]reconcile.Request, 0)

			labels := a.Meta.GetLabels()
			value, ok := labels["erd"]

			if ok {
				requests = append(requests, reconcile.Request{NamespacedName: types.NamespacedName{
					Name: value,
					Namespace: a.Meta.GetNamespace(),
				}})
			}

			return requests
		}),
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileEmergencyResponseDemo implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileEmergencyResponseDemo{}

// ReconcileEmergencyResponseDemo reconciles a EmergencyResponseDemo object
type ReconcileEmergencyResponseDemo struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	config *rest.Config
	builder erdHandlers.Builder
}

// Reconcile reads that state of the cluster for a EmergencyResponseDemo object and makes changes based on the state read
// and what is in the EmergencyResponseDemo.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileEmergencyResponseDemo) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	var err error

	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling EmergencyResponseDemo")

	// Fetch the EmergencyResponseDemo instance
	instance := &erdemov1alpha1.EmergencyResponseDemo{}
	err = r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if instance.DeletionTimestamp != nil && instance.Status.Type != erdemov1alpha1.EmergencyResponseDemoDelete {
		err = r.setDeleteStatus(instance)
		if err != nil {
			return reconcile.Result{}, err
		}

		reqLogger.Info("Deleting ERD instance", "ERD.Namespace", instance.Namespace, "ERD.Name", instance.Name)

		return reconcile.Result{}, nil
	}

	//build handler based on instance status
	requestHandler, err := r.builder.Build(instance.Status.Type)
	if err != nil {
		reqLogger.Error(err, "Failed to build handler")
		return reconcile.Result{}, err
	}

	//handle request
	result, err := requestHandler.Handle(instance)
	if err != nil {
		reqLogger.Error(err, "Error handling request", "Status", instance.Status.Type)
		return result, err
	}

	//end of reconcile request
	reqLogger.Info("End of reconcile request", "Namespace", instance.Namespace, "Name", instance.Name)
	return  result, nil
}

func (r *ReconcileEmergencyResponseDemo) setDeleteStatus(instance *erdemov1alpha1.EmergencyResponseDemo) error {
	statusHelper := status.Helper{}

	instance.Status = statusHelper.DeleteStatus()

	return r.client.Status().Update(context.TODO(), instance)
}
