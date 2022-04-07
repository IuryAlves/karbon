package multiclusterservice

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
)

func SelectRegion(region string) {
	ctx   := context.Background()
	config := ctrl.GetConfigOrDie()
	dynamic := dynamic.NewForConfigOrDie(config)
	resourceId := schema.GroupVersionResource{
		Group:    "networking.gke.io",
		Version:  "v1",
		Resource: "multiclusterservices",
	}
	patchContent := fmt.Sprintf(`{"spec": {"clusters": [{"link": "%s"}]}}`, region)
	Patch(dynamic, ctx, resourceId, "whereami-mcs", "whereami", patchContent)
	fmt.Println("Using region: " + region)
}

func Patch(dynamic dynamic.Interface, ctx context.Context, resourceId schema.GroupVersionResource,
	name, namespace, patchContent string) *unstructured.Unstructured {
	resource, err := dynamic.Resource(resourceId).Namespace(namespace).
		Patch(ctx, name, types.MergePatchType, []byte(patchContent), v1.PatchOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return resource
}