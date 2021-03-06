# Server-Side Apply Helper for Kubernetes Custom Controller

client-go provides [ApplyConfiguration types](https://pkg.go.dev/k8s.io/client-go/applyconfigurations) for typesafe Server-Side Apply.
This repository provides helper functions for convenient use of server-side apply in your custom controller.

## DeepCopy

ApplyConfiguration types provided by client-go don't have a DeepCopy function.
Especially it would be inconvenient when you embed an ApplyConfiguration type into your custom resource.
This repository provides the DeepCopy functions for ApplyConfiguration types generated by [deepcopy-gen](https://github.com/kubernetes/code-generator/tree/master/cmd/deepcopy-gen).

You can call DeepCopy of an ApplyConfiguration type as follows:

```go
dep1 := appsv1.Deployment("nginx", "default").
	WithLabels(map[string]string{"component": "nginx"}).
	WithSpec(appsv1.DeploymentSpec().
		WithReplicas(1).
		WithSelector(metav1.LabelSelector().WithMatchLabels(map[string]string{"component": "nginx"})).
		WithTemplate(corev1.PodTemplateSpec().
			WithLabels(map[string]string{"component": "nginx"}).
			WithSpec(corev1.PodSpec().
				WithContainers(corev1.Container().
					WithName("nginx").
					WithImage("nginx:latest"),
				),
			),
		),
	)

dep2 := dep1.DeepCopy()
```

## Apply

The implementation of Server-Side Apply using Extract function (e.g. [ExtractDeployment](https://pkg.go.dev/k8s.io/client-go/applyconfigurations/apps/v1#ExtractDeployment)) is a bit redundant.
This repository provides Apply function for ApplyConfiguration types.

```go
dep := appsv1.Deployment("nginx", "default").
	WithLabels(map[string]string{"component": "nginx"}).
	WithSpec(appsv1.DeploymentSpec().
		WithReplicas(1).
		WithSelector(metav1.LabelSelector().WithMatchLabels(map[string]string{"component": "nginx"})).
		WithTemplate(corev1.PodTemplateSpec().
			WithLabels(map[string]string{"component": "nginx"}).
			WithSpec(corev1.PodSpec().
				WithContainers(corev1.Container().
					WithName("nginx").
					WithImage("nginx:latest"),
				),
			),
		),
	)
err := applyconfigurations.Apply(ctx, k8sClient, dep, "field-manager-name")
if err != nil {
	return err
}
```
## OwnerReference

controller-runtime provides [controllerutil.SetControllerReference](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/controller/controllerutil#SetControllerReference).
This function cannot be used for ApplyConfiguration types.

This repository provides OwnerReference function to get the owner reference of other resource for ApplyConfiguration types.
You can use WithOwnerReferences function to set the owner reference to ApplyConfiguration types.

```go
ownerRef, err := applyconfigurations.OwnerReference(owner, scheme)
if err != nil {
	return err
}

dep := appsv1.Deployment("nginx", "default").
	WithOwnerReferences(ownerRef).
	WithLabels(map[string]string{"component": "nginx"}).
	WithSpec(appsv1.DeploymentSpec().
		WithReplicas(1).
		WithSelector(metav1.LabelSelector().WithMatchLabels(map[string]string{"component": "nginx"})).
		WithTemplate(corev1.PodTemplateSpec().
			WithLabels(map[string]string{"component": "nginx"}).
			WithSpec(corev1.PodSpec().
				WithContainers(corev1.Container().
					WithName("nginx").
					WithImage("nginx:latest"),
				),
			),
		),
	)
```
