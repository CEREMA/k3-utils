/*
Copyright 2022 The Kubernetes Authors.

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

package fake

import (
	"context"

	v1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVolumeSnapshots implements VolumeSnapshotInterface
type FakeVolumeSnapshots struct {
	Fake *FakeSnapshotV1beta1
	ns   string
}

var volumesnapshotsResource = schema.GroupVersionResource{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Resource: "volumesnapshots"}

var volumesnapshotsKind = schema.GroupVersionKind{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Kind: "VolumeSnapshot"}

// Get takes name of the volumeSnapshot, and returns the corresponding volumeSnapshot object, and an error if there is any.
func (c *FakeVolumeSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumesnapshotsResource, c.ns, name), &v1beta1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshot), err
}

// List takes label and field selectors, and returns the list of VolumeSnapshots that match those selectors.
func (c *FakeVolumeSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumesnapshotsResource, volumesnapshotsKind, c.ns, opts), &v1beta1.VolumeSnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VolumeSnapshotList{ListMeta: obj.(*v1beta1.VolumeSnapshotList).ListMeta}
	for _, item := range obj.(*v1beta1.VolumeSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeSnapshots.
func (c *FakeVolumeSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumesnapshotsResource, c.ns, opts))

}

// Create takes the representation of a volumeSnapshot and creates it.  Returns the server's representation of the volumeSnapshot, and an error, if there is any.
func (c *FakeVolumeSnapshots) Create(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.CreateOptions) (result *v1beta1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumesnapshotsResource, c.ns, volumeSnapshot), &v1beta1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshot), err
}

// Update takes the representation of a volumeSnapshot and updates it. Returns the server's representation of the volumeSnapshot, and an error, if there is any.
func (c *FakeVolumeSnapshots) Update(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumesnapshotsResource, c.ns, volumeSnapshot), &v1beta1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeSnapshots) UpdateStatus(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.UpdateOptions) (*v1beta1.VolumeSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumesnapshotsResource, "status", c.ns, volumeSnapshot), &v1beta1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshot), err
}

// Delete takes name of the volumeSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeVolumeSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(volumesnapshotsResource, c.ns, name, opts), &v1beta1.VolumeSnapshot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(volumesnapshotsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VolumeSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched volumeSnapshot.
func (c *FakeVolumeSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumesnapshotsResource, c.ns, name, pt, data, subresources...), &v1beta1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshot), err
}
