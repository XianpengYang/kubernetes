/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	consistencydetector "k8s.io/client-go/util/consistencydetector"
	v1 "k8s.io/code-generator/examples/crd/apis/example/v1"
	examplev1 "k8s.io/code-generator/examples/crd/applyconfiguration/example/v1"
	scheme "k8s.io/code-generator/examples/crd/clientset/versioned/scheme"
)

// ClusterTestTypesGetter has a method to return a ClusterTestTypeInterface.
// A group's client should implement this interface.
type ClusterTestTypesGetter interface {
	ClusterTestTypes() ClusterTestTypeInterface
}

// ClusterTestTypeInterface has methods to work with ClusterTestType resources.
type ClusterTestTypeInterface interface {
	Create(ctx context.Context, clusterTestType *v1.ClusterTestType, opts metav1.CreateOptions) (*v1.ClusterTestType, error)
	Update(ctx context.Context, clusterTestType *v1.ClusterTestType, opts metav1.UpdateOptions) (*v1.ClusterTestType, error)
	UpdateStatus(ctx context.Context, clusterTestType *v1.ClusterTestType, opts metav1.UpdateOptions) (*v1.ClusterTestType, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterTestType, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterTestTypeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterTestType, err error)
	Apply(ctx context.Context, clusterTestType *examplev1.ClusterTestTypeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterTestType, err error)
	ApplyStatus(ctx context.Context, clusterTestType *examplev1.ClusterTestTypeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterTestType, err error)
	GetScale(ctx context.Context, clusterTestTypeName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(ctx context.Context, clusterTestTypeName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error)

	ClusterTestTypeExpansion
}

// clusterTestTypes implements ClusterTestTypeInterface
type clusterTestTypes struct {
	client rest.Interface
}

// newClusterTestTypes returns a ClusterTestTypes
func newClusterTestTypes(c *ExampleV1Client) *clusterTestTypes {
	return &clusterTestTypes{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterTestType, and returns the corresponding clusterTestType object, and an error if there is any.
func (c *clusterTestTypes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Get().
		Resource("clustertesttypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterTestTypes that match those selectors.
func (c *clusterTestTypes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterTestTypeList, err error) {
	defer func() {
		if err == nil {
			consistencydetector.CheckListFromCacheDataConsistencyIfRequested(ctx, "list request for clustertesttypes", c.list, opts, result)
		}
	}()
	return c.list(ctx, opts)
}

// list takes label and field selectors, and returns the list of ClusterTestTypes that match those selectors.
func (c *clusterTestTypes) list(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterTestTypeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterTestTypeList{}
	err = c.client.Get().
		Resource("clustertesttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterTestTypes.
func (c *clusterTestTypes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustertesttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterTestType and creates it.  Returns the server's representation of the clusterTestType, and an error, if there is any.
func (c *clusterTestTypes) Create(ctx context.Context, clusterTestType *v1.ClusterTestType, opts metav1.CreateOptions) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Post().
		Resource("clustertesttypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTestType).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterTestType and updates it. Returns the server's representation of the clusterTestType, and an error, if there is any.
func (c *clusterTestTypes) Update(ctx context.Context, clusterTestType *v1.ClusterTestType, opts metav1.UpdateOptions) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Put().
		Resource("clustertesttypes").
		Name(clusterTestType.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTestType).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterTestTypes) UpdateStatus(ctx context.Context, clusterTestType *v1.ClusterTestType, opts metav1.UpdateOptions) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Put().
		Resource("clustertesttypes").
		Name(clusterTestType.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTestType).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterTestType and deletes it. Returns an error if one occurs.
func (c *clusterTestTypes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustertesttypes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterTestTypes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustertesttypes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterTestType.
func (c *clusterTestTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterTestType, err error) {
	result = &v1.ClusterTestType{}
	err = c.client.Patch(pt).
		Resource("clustertesttypes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterTestType.
func (c *clusterTestTypes) Apply(ctx context.Context, clusterTestType *examplev1.ClusterTestTypeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterTestType, err error) {
	if clusterTestType == nil {
		return nil, fmt.Errorf("clusterTestType provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(clusterTestType)
	if err != nil {
		return nil, err
	}
	name := clusterTestType.Name
	if name == nil {
		return nil, fmt.Errorf("clusterTestType.Name must be provided to Apply")
	}
	result = &v1.ClusterTestType{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("clustertesttypes").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *clusterTestTypes) ApplyStatus(ctx context.Context, clusterTestType *examplev1.ClusterTestTypeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterTestType, err error) {
	if clusterTestType == nil {
		return nil, fmt.Errorf("clusterTestType provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(clusterTestType)
	if err != nil {
		return nil, err
	}

	name := clusterTestType.Name
	if name == nil {
		return nil, fmt.Errorf("clusterTestType.Name must be provided to Apply")
	}

	result = &v1.ClusterTestType{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("clustertesttypes").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// GetScale takes name of the clusterTestType, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *clusterTestTypes) GetScale(ctx context.Context, clusterTestTypeName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Get().
		Resource("clustertesttypes").
		Name(clusterTestTypeName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *clusterTestTypes) UpdateScale(ctx context.Context, clusterTestTypeName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Put().
		Resource("clustertesttypes").
		Name(clusterTestTypeName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}
