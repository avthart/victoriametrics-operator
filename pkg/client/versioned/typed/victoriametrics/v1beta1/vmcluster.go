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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/VictoriaMetrics/operator/pkg/apis/victoriametrics/v1beta1"
	scheme "github.com/VictoriaMetrics/operator/pkg/client/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VMClustersGetter has a method to return a VMClusterInterface.
// A group's client should implement this interface.
type VMClustersGetter interface {
	VMClusters(namespace string) VMClusterInterface
}

// VMClusterInterface has methods to work with VMCluster resources.
type VMClusterInterface interface {
	Create(ctx context.Context, vMCluster *v1beta1.VMCluster, opts v1.CreateOptions) (*v1beta1.VMCluster, error)
	Update(ctx context.Context, vMCluster *v1beta1.VMCluster, opts v1.UpdateOptions) (*v1beta1.VMCluster, error)
	UpdateStatus(ctx context.Context, vMCluster *v1beta1.VMCluster, opts v1.UpdateOptions) (*v1beta1.VMCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VMCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VMClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMCluster, err error)
	VMClusterExpansion
}

// vMClusters implements VMClusterInterface
type vMClusters struct {
	client rest.Interface
	ns     string
}

// newVMClusters returns a VMClusters
func newVMClusters(c *VictoriametricsV1beta1Client, namespace string) *vMClusters {
	return &vMClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vMCluster, and returns the corresponding vMCluster object, and an error if there is any.
func (c *vMClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VMCluster, err error) {
	result = &v1beta1.VMCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vmclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VMClusters that match those selectors.
func (c *vMClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VMClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VMClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vmclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vMClusters.
func (c *vMClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vmclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vMCluster and creates it.  Returns the server's representation of the vMCluster, and an error, if there is any.
func (c *vMClusters) Create(ctx context.Context, vMCluster *v1beta1.VMCluster, opts v1.CreateOptions) (result *v1beta1.VMCluster, err error) {
	result = &v1beta1.VMCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vmclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vMCluster and updates it. Returns the server's representation of the vMCluster, and an error, if there is any.
func (c *vMClusters) Update(ctx context.Context, vMCluster *v1beta1.VMCluster, opts v1.UpdateOptions) (result *v1beta1.VMCluster, err error) {
	result = &v1beta1.VMCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vmclusters").
		Name(vMCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vMClusters) UpdateStatus(ctx context.Context, vMCluster *v1beta1.VMCluster, opts v1.UpdateOptions) (result *v1beta1.VMCluster, err error) {
	result = &v1beta1.VMCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vmclusters").
		Name(vMCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vMCluster and deletes it. Returns an error if one occurs.
func (c *vMClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vmclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vMClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vmclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vMCluster.
func (c *vMClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMCluster, err error) {
	result = &v1beta1.VMCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vmclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
