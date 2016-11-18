/*
Copyright 2016 The Kubernetes Authors.

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

package registry

import (
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/storage"
	"k8s.io/kubernetes/pkg/watch"

	"golang.org/x/net/context"
)

type proxyStore struct {
	store     storage.Interface
}

func NewProxyStore(c storage.Interface) *proxyStore {
	return &proxyStore{
		store:     c,
	}
}

// Versioner implements storage.Interface.Versioner.
func (s *proxyStore) Versioner() storage.Versioner {
	return s.store.Versioner()
}

// Get implements storage.Interface.Get.
func (s *proxyStore) Get(ctx context.Context, key string, out runtime.Object, ignoreNotFound bool) error {
	return s.store.Get(ctx, key, out, ignoreNotFound)
}

// Create implements storage.Interface.Create.
func (s *proxyStore) Create(ctx context.Context, key string, obj, out runtime.Object, ttl uint64) error {
	return s.store.Create(ctx, key, obj, out, ttl)
}

// Delete implements storage.Interface.Delete.
func (s *proxyStore) Delete(ctx context.Context, key string, out runtime.Object, precondtions *storage.Preconditions) error {
	return s.store.Delete(ctx, key, out, precondtions)
}


// GuaranteedUpdate implements storage.Interface.GuaranteedUpdate.
func (s *proxyStore) GuaranteedUpdate(ctx context.Context, key string, out runtime.Object, ignoreNotFound bool, precondtions *storage.Preconditions, tryUpdate storage.UpdateFunc) error {
	return s.store.GuaranteedUpdate(ctx, key, out, ignoreNotFound, precondtions, tryUpdate)
}

// GetToList implements storage.Interface.GetToList.
func (s *proxyStore) GetToList(ctx context.Context, key string, pred storage.SelectionPredicate, listObj runtime.Object) error {
	return s.store.GetToList(ctx, key, pred, listObj)
}

// List implements storage.Interface.List.
func (s *proxyStore) List(ctx context.Context, key, resourceVersion string, pred storage.SelectionPredicate, listObj runtime.Object) error {
	return s.store.List(ctx, key, resourceVersion, pred, listObj)
}

// Watch implements storage.Interface.Watch.
func (s *proxyStore) Watch(ctx context.Context, key string, resourceVersion string, pred storage.SelectionPredicate) (watch.Interface, error) {
	watchInterface, err := s.store.Watch(ctx, key, resourceVersion, pred)
	return watchInterface, err
}

// WatchList implements storage.Interface.WatchList.
func (s *proxyStore) WatchList(ctx context.Context, key string, resourceVersion string, pred storage.SelectionPredicate) (watch.Interface, error) {
	return s.store.WatchList(ctx, key, resourceVersion, pred)
}

