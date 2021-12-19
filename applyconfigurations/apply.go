package applyconfigurations

import (
	"context"

	metav1 "github.com/zoetrope/ssa-helper/applyconfigurations/meta/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

type Extractor[T any] interface {
	Extract(obj client.Object, fieldManager string, subresource string) (T, error)
	Original() client.Object
	ObjectKey() (client.ObjectKey, error)
}

func Apply[T Extractor[T]](ctx context.Context, k8sClient client.Client, expected T, fieldManager string) error {
	obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(expected)
	if err != nil {
		return err
	}
	patch := &unstructured.Unstructured{
		Object: obj,
	}

	res := expected.Original()
	key, err := expected.ObjectKey()
	if err != nil {
		return err
	}
	err = k8sClient.Get(ctx, key, res)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	current, err := expected.Extract(res, fieldManager, "")
	if err != nil {
		return err
	}

	if equality.Semantic.DeepEqual(expected, current) {
		return nil
	}

	return k8sClient.Patch(ctx, patch, client.Apply, &client.PatchOptions{
		FieldManager: fieldManager,
		Force:        pointer.Bool(true),
	})
}

func OwnerReference(owner client.Object, scheme *runtime.Scheme) (*metav1.OwnerReferenceApplyConfiguration, error) {
	gvk, err := apiutil.GVKForObject(owner, scheme)
	if err != nil {
		return nil, err
	}
	return metav1.OwnerReference().
		WithAPIVersion(gvk.GroupVersion().String()).
		WithKind(gvk.Kind).
		WithName(owner.GetName()).
		WithUID(owner.GetUID()).
		WithBlockOwnerDeletion(true).
		WithController(true), nil
}
