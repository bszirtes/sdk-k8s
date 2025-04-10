// Copyright (c) 2020-2022 Doc.ai and/or its affiliates.
//
// Copyright (c) 2020-2022 Cisco and/or its affiliates.
//
// Copyright (c) 2020-2022 VMware, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	networkservicemeshiov1 "github.com/bszirtes/sdk-k8s/pkg/tools/k8s/apis/networkservicemesh.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkServices implements NetworkServiceInterface
type FakeNetworkServices struct {
	Fake *FakeNetworkservicemeshV1
	ns   string
}

var networkservicesResource = schema.GroupVersionResource{Group: "networkservicemesh.io", Version: "v1", Resource: "networkservices"}

var networkservicesKind = schema.GroupVersionKind{Group: "networkservicemesh.io", Version: "v1", Kind: "NetworkService"}

// Get takes name of the networkService, and returns the corresponding networkService object, and an error if there is any.
func (c *FakeNetworkServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkservicemeshiov1.NetworkService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkservicesResource, c.ns, name), &networkservicemeshiov1.NetworkService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkservicemeshiov1.NetworkService), err
}

// List takes label and field selectors, and returns the list of NetworkServices that match those selectors.
func (c *FakeNetworkServices) List(ctx context.Context, opts v1.ListOptions) (result *networkservicemeshiov1.NetworkServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkservicesResource, networkservicesKind, c.ns, opts), &networkservicemeshiov1.NetworkServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkservicemeshiov1.NetworkServiceList{ListMeta: obj.(*networkservicemeshiov1.NetworkServiceList).ListMeta}
	for _, item := range obj.(*networkservicemeshiov1.NetworkServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkServices.
func (c *FakeNetworkServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkservicesResource, c.ns, opts))

}

// Create takes the representation of a networkService and creates it.  Returns the server's representation of the networkService, and an error, if there is any.
func (c *FakeNetworkServices) Create(ctx context.Context, networkService *networkservicemeshiov1.NetworkService, opts v1.CreateOptions) (result *networkservicemeshiov1.NetworkService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkservicesResource, c.ns, networkService), &networkservicemeshiov1.NetworkService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkservicemeshiov1.NetworkService), err
}

// Update takes the representation of a networkService and updates it. Returns the server's representation of the networkService, and an error, if there is any.
func (c *FakeNetworkServices) Update(ctx context.Context, networkService *networkservicemeshiov1.NetworkService, opts v1.UpdateOptions) (result *networkservicemeshiov1.NetworkService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkservicesResource, c.ns, networkService), &networkservicemeshiov1.NetworkService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkservicemeshiov1.NetworkService), err
}

// Delete takes name of the networkService and deletes it. Returns an error if one occurs.
func (c *FakeNetworkServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkservicesResource, c.ns, name, opts), &networkservicemeshiov1.NetworkService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &networkservicemeshiov1.NetworkServiceList{})
	return err
}

// Patch applies the patch and returns the patched networkService.
func (c *FakeNetworkServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkservicemeshiov1.NetworkService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkservicesResource, c.ns, name, pt, data, subresources...), &networkservicemeshiov1.NetworkService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkservicemeshiov1.NetworkService), err
}
