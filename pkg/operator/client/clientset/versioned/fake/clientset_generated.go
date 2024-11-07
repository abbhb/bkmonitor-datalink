// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/client/clientset/versioned"
	bkv1alpha1 "github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/client/clientset/versioned/typed/logging/v1alpha1"
	fakebkv1alpha1 "github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/client/clientset/versioned/typed/logging/v1alpha1/fake"
	monitoringv1beta1 "github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/client/clientset/versioned/typed/monitoring/v1beta1"
	fakemonitoringv1beta1 "github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/client/clientset/versioned/typed/monitoring/v1beta1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// BkV1alpha1 retrieves the BkV1alpha1Client
func (c *Clientset) BkV1alpha1() bkv1alpha1.BkV1alpha1Interface {
	return &fakebkv1alpha1.FakeBkV1alpha1{Fake: &c.Fake}
}

// MonitoringV1beta1 retrieves the MonitoringV1beta1Client
func (c *Clientset) MonitoringV1beta1() monitoringv1beta1.MonitoringV1beta1Interface {
	return &fakemonitoringv1beta1.FakeMonitoringV1beta1{Fake: &c.Fake}
}
