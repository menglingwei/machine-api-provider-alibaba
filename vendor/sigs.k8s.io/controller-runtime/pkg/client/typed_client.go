/*
Copyright 2018 The Kubernetes Authors.

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

package client

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
)

var _ Reader = &typedClient{}
var _ Writer = &typedClient{}
var _ StatusWriter = &typedClient{}

// client is a client.Client that reads and writes directly from/to an API server.  It lazily initializes
// new clients at the time they are used, and caches the client.
type typedClient struct {
	cache      *clientCache
	paramCodec runtime.ParameterCodec
}

// Create implements client.Client
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *typedClient) Create(ctx context.Context, obj Object, opts ...CreateOption) error {
=======
func (c *typedClient) Create(ctx context.Context, obj runtime.Object, opts ...CreateOption) error {
>>>>>>> 79bfea2d (update vendor)
=======
func (c *typedClient) Create(ctx context.Context, obj Object, opts ...CreateOption) error {
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (c *typedClient) Create(ctx context.Context, obj Object, opts ...CreateOption) error {
>>>>>>> 03397665 (update api)
	o, err := c.cache.getObjMeta(obj)
	if err != nil {
		return err
	}

	createOpts := &CreateOptions{}
	createOpts.ApplyOptions(opts)
	return o.Post().
		NamespaceIfScoped(o.GetNamespace(), o.isNamespaced()).
		Resource(o.resource()).
		Body(obj).
		VersionedParams(createOpts.AsCreateOptions(), c.paramCodec).
		Do(ctx).
		Into(obj)
}

// Update implements client.Client
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *typedClient) Update(ctx context.Context, obj Object, opts ...UpdateOption) error {
=======
func (c *typedClient) Update(ctx context.Context, obj runtime.Object, opts ...UpdateOption) error {
>>>>>>> 79bfea2d (update vendor)
=======
func (c *typedClient) Update(ctx context.Context, obj Object, opts ...UpdateOption) error {
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (c *typedClient) Update(ctx context.Context, obj Object, opts ...UpdateOption) error {
>>>>>>> 03397665 (update api)
	o, err := c.cache.getObjMeta(obj)
	if err != nil {
		return err
	}

	updateOpts := &UpdateOptions{}
	updateOpts.ApplyOptions(opts)
	return o.Put().
		NamespaceIfScoped(o.GetNamespace(), o.isNamespaced()).
		Resource(o.resource()).
		Name(o.GetName()).
		Body(obj).
		VersionedParams(updateOpts.AsUpdateOptions(), c.paramCodec).
		Do(ctx).
		Into(obj)
}

// Delete implements client.Client
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *typedClient) Delete(ctx context.Context, obj Object, opts ...DeleteOption) error {
=======
func (c *typedClient) Delete(ctx context.Context, obj runtime.Object, opts ...DeleteOption) error {
>>>>>>> 79bfea2d (update vendor)
=======
func (c *typedClient) Delete(ctx context.Context, obj Object, opts ...DeleteOption) error {
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (c *typedClient) Delete(ctx context.Context, obj Object, opts ...DeleteOption) error {
>>>>>>> 03397665 (update api)
	o, err := c.cache.getObjMeta(obj)
	if err != nil {
		return err
	}

	deleteOpts := DeleteOptions{}
	deleteOpts.ApplyOptions(opts)

	return o.Delete().
		NamespaceIfScoped(o.GetNamespace(), o.isNamespaced()).
		Resource(o.resource()).
		Name(o.GetName()).
		Body(deleteOpts.AsDeleteOptions()).
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		Do(ctx).
=======
		Context(ctx).
		Do().
>>>>>>> 79bfea2d (update vendor)
=======
		Do(ctx).
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
		Do(ctx).
>>>>>>> 03397665 (update api)
		Error()
}

// DeleteAllOf implements client.Client
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *typedClient) DeleteAllOf(ctx context.Context, obj Object, opts ...DeleteAllOfOption) error {
=======
func (c *typedClient) DeleteAllOf(ctx context.Context, obj runtime.Object, opts ...DeleteAllOfOption) error {
>>>>>>> 79bfea2d (update vendor)
=======
func (c *typedClient) DeleteAllOf(ctx context.Context, obj Object, opts ...DeleteAllOfOption) error {
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (c *typedClient) DeleteAllOf(ctx context.Context, obj Object, opts ...DeleteAllOfOption) error {
>>>>>>> 03397665 (update api)
	o, err := c.cache.getObjMeta(obj)
	if err != nil {
		return err
	}

	deleteAllOfOpts := DeleteAllOfOptions{}
	deleteAllOfOpts.ApplyOptions(opts)

	return o.Delete().
		NamespaceIfScoped(deleteAllOfOpts.ListOptions.Namespace, o.isNamespaced()).
		Resource(o.resource()).
		VersionedParams(deleteAllOfOpts.AsListOptions(), c.paramCodec).
		Body(deleteAllOfOpts.AsDeleteOptions()).
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		Do(ctx).
=======
		Context(ctx).
		Do().
>>>>>>> 79bfea2d (update vendor)
=======
		Do(ctx).
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
		Do(ctx).
>>>>>>> 03397665 (update api)
		Error()
}

// Patch implements client.Client
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *typedClient) Patch(ctx context.Context, obj Object, patch Patch, opts ...PatchOption) error {
=======
func (c *typedClient) Patch(ctx context.Context, obj runtime.Object, patch Patch, opts ...PatchOption) error {
>>>>>>> 79bfea2d (update vendor)
=======
func (c *typedClient) Patch(ctx context.Context, obj Object, patch Patch, opts ...PatchOption) error {
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (c *typedClient) Patch(ctx context.Context, obj Object, patch Patch, opts ...PatchOption) error {
>>>>>>> 03397665 (update api)
	o, err := c.cache.getObjMeta(obj)
	if err != nil {
		return err
	}

	data, err := patch.Data(obj)
	if err != nil {
		return err
	}

	patchOpts := &PatchOptions{}
	return o.Patch(patch.Type()).
		NamespaceIfScoped(o.GetNamespace(), o.isNamespaced()).
		Resource(o.resource()).
		Name(o.GetName()).
		VersionedParams(patchOpts.ApplyOptions(opts).AsPatchOptions(), c.paramCodec).
		Body(data).
		Do(ctx).
		Into(obj)
}

// Get implements client.Client
func (c *typedClient) Get(ctx context.Context, key ObjectKey, obj Object) error {
	r, err := c.cache.getResource(obj)
	if err != nil {
		return err
	}
	return r.Get().
		NamespaceIfScoped(key.Namespace, r.isNamespaced()).
		Resource(r.resource()).
		Name(key.Name).Do(ctx).Into(obj)
}

// List implements client.Client
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *typedClient) List(ctx context.Context, obj ObjectList, opts ...ListOption) error {
=======
func (c *typedClient) List(ctx context.Context, obj runtime.Object, opts ...ListOption) error {
>>>>>>> 79bfea2d (update vendor)
=======
func (c *typedClient) List(ctx context.Context, obj ObjectList, opts ...ListOption) error {
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (c *typedClient) List(ctx context.Context, obj ObjectList, opts ...ListOption) error {
>>>>>>> 03397665 (update api)
	r, err := c.cache.getResource(obj)
	if err != nil {
		return err
	}
	listOpts := ListOptions{}
	listOpts.ApplyOptions(opts)
	return r.Get().
		NamespaceIfScoped(listOpts.Namespace, r.isNamespaced()).
		Resource(r.resource()).
		VersionedParams(listOpts.AsListOptions(), c.paramCodec).
		Do(ctx).
		Into(obj)
}

// UpdateStatus used by StatusWriter to write status.
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *typedClient) UpdateStatus(ctx context.Context, obj Object, opts ...UpdateOption) error {
=======
func (c *typedClient) UpdateStatus(ctx context.Context, obj runtime.Object, opts ...UpdateOption) error {
>>>>>>> 79bfea2d (update vendor)
=======
func (c *typedClient) UpdateStatus(ctx context.Context, obj Object, opts ...UpdateOption) error {
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (c *typedClient) UpdateStatus(ctx context.Context, obj Object, opts ...UpdateOption) error {
>>>>>>> 03397665 (update api)
	o, err := c.cache.getObjMeta(obj)
	if err != nil {
		return err
	}
	// TODO(droot): examine the returned error and check if it error needs to be
	// wrapped to improve the UX ?
	// It will be nice to receive an error saying the object doesn't implement
	// status subresource and check CRD definition
	return o.Put().
		NamespaceIfScoped(o.GetNamespace(), o.isNamespaced()).
		Resource(o.resource()).
		Name(o.GetName()).
		SubResource("status").
		Body(obj).
		VersionedParams((&UpdateOptions{}).ApplyOptions(opts).AsUpdateOptions(), c.paramCodec).
		Do(ctx).
		Into(obj)
}

// PatchStatus used by StatusWriter to write status.
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (c *typedClient) PatchStatus(ctx context.Context, obj Object, patch Patch, opts ...PatchOption) error {
=======
func (c *typedClient) PatchStatus(ctx context.Context, obj runtime.Object, patch Patch, opts ...PatchOption) error {
>>>>>>> 79bfea2d (update vendor)
=======
func (c *typedClient) PatchStatus(ctx context.Context, obj Object, patch Patch, opts ...PatchOption) error {
>>>>>>> e879a141 (alibabacloud machine-api provider)
=======
func (c *typedClient) PatchStatus(ctx context.Context, obj Object, patch Patch, opts ...PatchOption) error {
>>>>>>> 03397665 (update api)
	o, err := c.cache.getObjMeta(obj)
	if err != nil {
		return err
	}

	data, err := patch.Data(obj)
	if err != nil {
		return err
	}

	patchOpts := &PatchOptions{}
	return o.Patch(patch.Type()).
		NamespaceIfScoped(o.GetNamespace(), o.isNamespaced()).
		Resource(o.resource()).
		Name(o.GetName()).
		SubResource("status").
		Body(data).
		VersionedParams(patchOpts.ApplyOptions(opts).AsPatchOptions(), c.paramCodec).
		Do(ctx).
		Into(obj)
}
