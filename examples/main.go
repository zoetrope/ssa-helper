package main

import (
	"context"
	"fmt"

	"github.com/zoetrope/ssa-helper/applyconfigurations"
	appsv1 "github.com/zoetrope/ssa-helper/applyconfigurations/apps/v1"
	corev1 "github.com/zoetrope/ssa-helper/applyconfigurations/core/v1"
	metav1 "github.com/zoetrope/ssa-helper/applyconfigurations/meta/v1"
	apiappsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	// Prepare
	ctx := context.Background()
	scm := runtime.NewScheme()
	err := scheme.AddToScheme(scm)
	if err != nil {
		fmt.Println(err)
		return
	}
	k8sClient, err := getClient(scm)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Server-Side Apply
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
	err = applyconfigurations.Apply(ctx, k8sClient, dep1, "field-manager-name")
	if err != nil {
		fmt.Println(err)
		return
	}

	// OwnerReference
	owner := &apiappsv1.Deployment{}
	key, err := dep1.ObjectKey()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = k8sClient.Get(ctx, key, owner)
	if err != nil {
		fmt.Println(err)
		return
	}
	ownerRef, err := applyconfigurations.OwnerReference(owner, scm)
	if err != nil {
		fmt.Println(err)
		return
	}

	// DeepCopy
	dep2 := dep1.DeepCopy()
	dep2.WithName("nginx2").
		WithOwnerReferences(ownerRef)
	err = applyconfigurations.Apply(ctx, k8sClient, dep2, "field-manager-name")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getClient(scm *runtime.Scheme) (client.Client, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	cli, err := client.New(cfg, client.Options{Scheme: scm})
	if err != nil {
		return nil, err
	}
	return cli, nil
}
