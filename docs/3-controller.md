Implementing a controller
=========================

We will now create our ToastController to manage the Toast resources.

# Steps

1. Update the `ToastReconciler.Reconcile()` function to implement the Toaster update loop.
  Find and open the `breakfast/internal/controller/toast_controller.go` file.

  First, we define a constant to represent how long to wait between updates:
  ```go
  const TOAST_INTERVAL = 10 * time.Second
  ```

  Then we can fill in the Reconcile function as follows:
  ```go
  func (r *ToastReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    logger := log.FromContext(ctx)

    // TODO(user): your logic here
    toast := &breakfastv1.Toast{}
    err := r.Get(ctx, req.NamespacedName, toast)

    if err != nil && client.IgnoreNotFound(err) == nil {
      logger.Info("Toast not found, assumed deleted")
      return ctrl.Result{}, nil
    }
    if client.IgnoreNotFound(err) != nil {
      logger.Error(err, "unable to fetch Toast")
      return ctrl.Result{}, err
    }

    if time.Now().Before(toast.Status.LastUpdated.Add(TOAST_INTERVAL)) {
      return ctrl.Result{}, nil
    }

    if toast.Status.Toastiness < toast.Spec.Toastiness {
      logger.Info("status toastiness < spec toastiness")
      toast.Status.Toastiness++
      toast.Status.Ready = false
      toast.Status.LastUpdated = metav1.Now()

      if err = r.Status().Update(ctx, toast); err != nil {
        logger.Error(err, "unable to update Toast")
        return ctrl.Result{}, err
      }

      return ctrl.Result{
        RequeueAfter: TOAST_INTERVAL,
      }, nil
    }

    // Update the status after the toastiness reaches the desired level
    toast.Status.Ready = true
    if err = r.Status().Update(ctx, toast); err != nil {
      logger.Error(err, "unable to update Toast")
      return ctrl.Result{}, err
    }
    // Don't requeue if the toast is ready
    return ctrl.Result{}, nil
  }
  ```

  This requires the packages `time` and `metav1` to be added to the import list:
  ```diff
    import (
      "context"

  +   "time"

  +   metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
      "k8s.io/apimachinery/pkg/runtime"
      ctrl "sigs.k8s.io/controller-runtime"
      "sigs.k8s.io/controller-runtime/pkg/client"
      "sigs.k8s.io/controller-runtime/pkg/log"

      breakfastv1 "github.com/NewRedo/toaster-controller/breakfast/api/v1"
    )
  ```

2. Build with `make docker-build TAG=breakfast:latest`

3. If using a local kubernetes deployment, prevent kubernetes from pulling the image from a remote repository:

  In `breakfast/config/manager/manager.yaml`:
  ```diff
            args:
              - --leader-elect
            image: controller:latest
  +         imagePullPolicy: IfNotPresent
            name: manager
            securityContext:
  ```

4. Deploy with `make deploy TAG=breakfast:latest`

# Testing the controller

Add the sample toast resource, that was created in the previous section, to the cluster.
```shell
kubectl apply -f config/samples/breakfast_v1_toast.yaml
```

Repeated invocations of `kubectl describe toast` will show the toastiness increasing until it reaches the desired
level.

# Next Steps

[Summary](./4-summary.md)
