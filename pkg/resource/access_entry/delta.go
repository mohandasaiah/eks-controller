// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package access_entry

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customPreCompare(delta, a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.ClusterName, b.ko.Spec.ClusterName) {
		delta.Add("Spec.ClusterName", a.ko.Spec.ClusterName, b.ko.Spec.ClusterName)
	} else if a.ko.Spec.ClusterName != nil && b.ko.Spec.ClusterName != nil {
		if *a.ko.Spec.ClusterName != *b.ko.Spec.ClusterName {
			delta.Add("Spec.ClusterName", a.ko.Spec.ClusterName, b.ko.Spec.ClusterName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ClusterRef, b.ko.Spec.ClusterRef) {
		delta.Add("Spec.ClusterRef", a.ko.Spec.ClusterRef, b.ko.Spec.ClusterRef)
	}
	if len(a.ko.Spec.KubernetesGroups) != len(b.ko.Spec.KubernetesGroups) {
		delta.Add("Spec.KubernetesGroups", a.ko.Spec.KubernetesGroups, b.ko.Spec.KubernetesGroups)
	} else if len(a.ko.Spec.KubernetesGroups) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.KubernetesGroups, b.ko.Spec.KubernetesGroups) {
			delta.Add("Spec.KubernetesGroups", a.ko.Spec.KubernetesGroups, b.ko.Spec.KubernetesGroups)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PrincipalARN, b.ko.Spec.PrincipalARN) {
		delta.Add("Spec.PrincipalARN", a.ko.Spec.PrincipalARN, b.ko.Spec.PrincipalARN)
	} else if a.ko.Spec.PrincipalARN != nil && b.ko.Spec.PrincipalARN != nil {
		if *a.ko.Spec.PrincipalARN != *b.ko.Spec.PrincipalARN {
			delta.Add("Spec.PrincipalARN", a.ko.Spec.PrincipalARN, b.ko.Spec.PrincipalARN)
		}
	}
	desiredACKTags, _ := convertToOrderedACKTags(a.ko.Spec.Tags)
	latestACKTags, _ := convertToOrderedACKTags(b.ko.Spec.Tags)
	if !ackcompare.MapStringStringEqual(desiredACKTags, latestACKTags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Type, b.ko.Spec.Type) {
		delta.Add("Spec.Type", a.ko.Spec.Type, b.ko.Spec.Type)
	} else if a.ko.Spec.Type != nil && b.ko.Spec.Type != nil {
		if *a.ko.Spec.Type != *b.ko.Spec.Type {
			delta.Add("Spec.Type", a.ko.Spec.Type, b.ko.Spec.Type)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Username, b.ko.Spec.Username) {
		delta.Add("Spec.Username", a.ko.Spec.Username, b.ko.Spec.Username)
	} else if a.ko.Spec.Username != nil && b.ko.Spec.Username != nil {
		if *a.ko.Spec.Username != *b.ko.Spec.Username {
			delta.Add("Spec.Username", a.ko.Spec.Username, b.ko.Spec.Username)
		}
	}

	return delta
}
