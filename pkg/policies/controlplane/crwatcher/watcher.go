package crwatcher

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/fluxninja/aperture/operator/api"
	"github.com/fluxninja/aperture/pkg/panichandler"
	"github.com/hashicorp/go-multierror"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	policyv1alpha1 "github.com/fluxninja/aperture/operator/api/policy/v1alpha1"
	"github.com/fluxninja/aperture/pkg/log"
	"github.com/fluxninja/aperture/pkg/notifiers"
)

// watcher holds the state of the watcher.
type watcher struct {
	waitGroup sync.WaitGroup
	notifiers.Trackers
	dynamicConfigTrackers notifiers.Trackers
	ctx                   context.Context
	cancel                context.CancelFunc
	client.Client
	scheme           *runtime.Scheme
	recorder         record.EventRecorder
	resourcesDeleted map[types.NamespacedName]bool
}

// Make sure Watcher implements notifiers.Watcher interface.
var _ notifiers.Watcher = &watcher{}

// NewWatcher prepares watcher instance for the Kuberneter Policy.
func NewWatcher() (*watcher, error) {
	ctx, cancel := context.WithCancel(context.Background())

	watcher := &watcher{
		Trackers:              notifiers.NewDefaultTrackers(),
		dynamicConfigTrackers: notifiers.NewDefaultTrackers(),
		ctx:                   ctx,
		cancel:                cancel,
	}

	return watcher, nil
}

// Start starts the watcher go routines and handles Policy Custom resource events from Kubernetes.
func (w *watcher) Start() error {
	err := w.Trackers.Start()
	if err != nil {
		return err
	}
	err = w.dynamicConfigTrackers.Start()
	if err != nil {
		return err
	}

	w.waitGroup.Add(1)

	panichandler.Go(func() {
		defer w.waitGroup.Done()
		operation := func() error {
			scheme := runtime.NewScheme()

			utilruntime.Must(clientgoscheme.AddToScheme(scheme))

			utilruntime.Must(api.AddToScheme(scheme))

			mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
				Scheme:             scheme,
				MetricsBindAddress: "0",
				LeaderElection:     false,
				Namespace:          os.Getenv("APERTURE_CONTROLLER_NAMESPACE"),
			})
			if err != nil {
				log.Error().Err(err).Msg("Failed to create Kubernetes Reconciler for Policy")
				return nil
			}
			w.Client = mgr.GetClient()
			w.scheme = scheme
			w.recorder = mgr.GetEventRecorderFor("aperture-policy")

			if err = w.SetupWithManager(mgr); err != nil {
				log.Error().Err(err).Msg("Failed to create Kubernetes controller for policy")
				return nil
			}
			return mgr.Start(w.ctx)
		}

		boff := backoff.NewConstantBackOff(5 * time.Second)

		_ = backoff.Retry(operation, backoff.WithContext(boff, w.ctx))
		log.Info().Msg("Stopping kubernetes watcher for Policy")
	})

	return nil
}

// Stop stops the watcher go routines.
func (w *watcher) Stop() error {
	w.cancel()
	w.waitGroup.Wait()
	var err, merr error
	err = w.Trackers.Stop()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	err = w.dynamicConfigTrackers.Stop()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	return merr
}

// GetDynamicConfigWatcher returns the config watcher.
func (w *watcher) GetDynamicConfigWatcher() notifiers.Watcher {
	return w.dynamicConfigTrackers
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Policy object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (w *watcher) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	if w.resourcesDeleted == nil {
		w.resourcesDeleted = make(map[types.NamespacedName]bool)
	}

	instance := &policyv1alpha1.Policy{}
	err := w.Get(ctx, req.NamespacedName, instance)
	if err != nil && errors.IsNotFound(err) {
		if !w.resourcesDeleted[req.NamespacedName] {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Debug().Msg(fmt.Sprintf("Handling deletion of resources for Instance '%s' in Namespace '%s'", req.Name, req.Namespace))
			instance.Name = req.Name
			instance.Namespace = req.Namespace
			w.deleteResources(ctx, instance.DeepCopy())
			w.resourcesDeleted[req.NamespacedName] = true
			log.Debug().Msg("Policy resource not found. Ignoring since object must be deleted")
		}
		return ctrl.Result{}, nil
	} else if err != nil {
		// Error reading the object - requeue the request.
		log.Error().Err(err).Msg("failed to get Policy")
		return ctrl.Result{}, err
	}

	// Handing delete operation
	if instance.GetDeletionTimestamp() != nil {
		log.Debug().Msg(fmt.Sprintf("Handling deletion of resources for Instance '%s' in Namespace '%s'", instance.GetName(), instance.GetNamespace()))
		w.deleteResources(ctx, instance.DeepCopy())
		w.resourcesDeleted[req.NamespacedName] = true
		return ctrl.Result{}, nil
	}

	instance.Status.Status = "uploading"
	if err := w.updateStatus(ctx, instance.DeepCopy()); err != nil {
		return ctrl.Result{}, err
	}
	w.resourcesDeleted[req.NamespacedName] = false

	if err := w.reconcilePolicy(ctx, instance); err != nil {
		return ctrl.Result{}, err
	}

	instance.Status.Status = "uploaded"
	if err := w.updateStatus(ctx, instance.DeepCopy()); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (w *watcher) deleteResources(ctx context.Context, instance *policyv1alpha1.Policy) {
	w.RemoveEvent(notifiers.Key(instance.GetName()))
	w.dynamicConfigTrackers.RemoveEvent(notifiers.Key(instance.GetName()))
}

// updateResource updates the Aperture resource in Kubernetes.
func (w *watcher) updateStatus(ctx context.Context, instance *policyv1alpha1.Policy) error {
	attempt := 5
	status := instance.DeepCopy().Status
	for attempt > 0 {
		attempt -= 1
		if err := w.Status().Update(ctx, instance); err != nil {
			if errors.IsConflict(err) {
				namespacesName := types.NamespacedName{
					Namespace: instance.GetNamespace(),
					Name:      instance.GetName(),
				}
				if err = w.Get(ctx, namespacesName, instance); err != nil {
					return err
				}
				instance.Status = status
				continue
			}
			return err
		}
	}
	return nil
}

// reconcilePolicy sends a write event to notifier to get it uploaded on the Etcd.
func (w *watcher) reconcilePolicy(ctx context.Context, instance *policyv1alpha1.Policy) error {
	w.WriteEvent(notifiers.Key(instance.GetName()), instance.Spec.Raw)
	w.dynamicConfigTrackers.WriteEvent(notifiers.Key(instance.GetName()), instance.DynamicConfig.Raw)

	w.recorder.Eventf(instance, corev1.EventTypeWarning, "UploadSuccessful", "Uploaded policy to trackers.")
	return nil
}

// eventFilters sets up a Predicate filter for the received events.
func eventFiltersForPolicy() predicate.Predicate {
	return predicate.Funcs{
		CreateFunc: func(create event.CreateEvent) bool {
			_, ok := create.Object.(*policyv1alpha1.Policy)
			return ok
		},
		UpdateFunc: func(update event.UpdateEvent) bool {
			new, ok1 := update.ObjectNew.(*policyv1alpha1.Policy)
			old, ok2 := update.ObjectOld.(*policyv1alpha1.Policy)
			if !ok1 || !ok2 {
				return false
			}

			return !bytes.Equal(old.Spec.Raw, new.Spec.Raw) || !bytes.Equal(old.DynamicConfig.Raw, new.DynamicConfig.Raw) || new.GetDeletionTimestamp() != nil
		},
		DeleteFunc: func(delete event.DeleteEvent) bool {
			return !delete.DeleteStateUnknown
		},
	}
}

// SetupWithManager sets up the controller with the Manager.
func (w *watcher) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&policyv1alpha1.Policy{}).
		WithEventFilter(eventFiltersForPolicy()).
		Complete(w)
}
