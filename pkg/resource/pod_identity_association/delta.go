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

package pod_identity_association

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

	if ackcompare.HasNilDifference(a.ko.Spec.ClientRequestToken, b.ko.Spec.ClientRequestToken) {
		delta.Add("Spec.ClientRequestToken", a.ko.Spec.ClientRequestToken, b.ko.Spec.ClientRequestToken)
	} else if a.ko.Spec.ClientRequestToken != nil && b.ko.Spec.ClientRequestToken != nil {
		if *a.ko.Spec.ClientRequestToken != *b.ko.Spec.ClientRequestToken {
			delta.Add("Spec.ClientRequestToken", a.ko.Spec.ClientRequestToken, b.ko.Spec.ClientRequestToken)
		}
	}
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
	if ackcompare.HasNilDifference(a.ko.Spec.Namespace, b.ko.Spec.Namespace) {
		delta.Add("Spec.Namespace", a.ko.Spec.Namespace, b.ko.Spec.Namespace)
	} else if a.ko.Spec.Namespace != nil && b.ko.Spec.Namespace != nil {
		if *a.ko.Spec.Namespace != *b.ko.Spec.Namespace {
			delta.Add("Spec.Namespace", a.ko.Spec.Namespace, b.ko.Spec.Namespace)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RoleARN, b.ko.Spec.RoleARN) {
		delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
	} else if a.ko.Spec.RoleARN != nil && b.ko.Spec.RoleARN != nil {
		if *a.ko.Spec.RoleARN != *b.ko.Spec.RoleARN {
			delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.RoleRef, b.ko.Spec.RoleRef) {
		delta.Add("Spec.RoleRef", a.ko.Spec.RoleRef, b.ko.Spec.RoleRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ServiceAccount, b.ko.Spec.ServiceAccount) {
		delta.Add("Spec.ServiceAccount", a.ko.Spec.ServiceAccount, b.ko.Spec.ServiceAccount)
	} else if a.ko.Spec.ServiceAccount != nil && b.ko.Spec.ServiceAccount != nil {
		if *a.ko.Spec.ServiceAccount != *b.ko.Spec.ServiceAccount {
			delta.Add("Spec.ServiceAccount", a.ko.Spec.ServiceAccount, b.ko.Spec.ServiceAccount)
		}
	}
	desiredACKTags, _ := convertToOrderedACKTags(a.ko.Spec.Tags)
	latestACKTags, _ := convertToOrderedACKTags(b.ko.Spec.Tags)
	if !ackcompare.MapStringStringEqual(desiredACKTags, latestACKTags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}
