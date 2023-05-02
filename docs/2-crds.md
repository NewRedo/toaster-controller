Creating an API
===============

We will now create an API for our application. This consists of resources and controllers.

# Steps

1. Create an API for the Toast:
  ```shell
  operator-sdk create api --group breakfast --kind Toast --version v1 --resource --controller
  ```

2. Open `api/v1/toast_types.go`; notice the TODO comments.

  Each API resource has a Spec, where you describe the parameters, and a Status.

3. Update the `Toast` API;

  To the `ToastSpec`, add a `Toastiness` property:
  ```go
  // ToastSpec defines the desired state of Toast
  type ToastSpec struct {
    // Toastiness is how long to cook the toast, from 0 (raw) to 10 (blackend)
    Toastiness int `json:"toastiness,omitempty"`
  }
  ```
  And to the `ToastStatus` add `Toastiness` and `Ready` properties:
  ```go
  // ToastStatus defines the observed state of Toast
  type ToastStatus struct {
    // Toastiness is the current level of toastiness
    Toastiness int  `json:"toastiness,omitempty"`
    // Ready is whether the Toast has been fully cooked
    Ready bool `json:"ready,omitempty"`
    // LastUpdated is the last time the Toast was updated
    LastUpdated metav1.Time `json:"lastUpdated,omitempty"`
  }
  ```

5. Update the API manifests
  ```shell
  make manifests
  ```

# Using the API

We can install the new Custom Resource Definition (CRD) for Toast into our cluster:
```shell
kubectl apply -f breakfast/config/crd/bases/breakfast.newredo.com_toasts.yaml
```

And we can create objects with kubectl. The operator SDK has created some sample YAML files for us; open the Toast sample in `breakfast/config/samples/breakfast_v1_toast.yaml` and add a `toastiness` property:
```diff
 spec:
-  # TODO(user): Add fields here
+  toastiness: 5
```

Let's create a Toast and view it with kubectl:
```shell
kubectl apply -f toast/config/samples/toast_v1_toast.yaml
kubectl get toast
kubectl describe toasts
```

Look at the `2-crds` branch for the output from this step.

# Next step

[Creating a controller](./3-controller.md)
