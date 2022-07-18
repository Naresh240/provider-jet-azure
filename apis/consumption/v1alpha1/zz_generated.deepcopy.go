//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroup) DeepCopyInto(out *BudgetResourceGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroup.
func (in *BudgetResourceGroup) DeepCopy() *BudgetResourceGroup {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetResourceGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupList) DeepCopyInto(out *BudgetResourceGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BudgetResourceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupList.
func (in *BudgetResourceGroupList) DeepCopy() *BudgetResourceGroupList {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetResourceGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupObservation) DeepCopyInto(out *BudgetResourceGroupObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupObservation.
func (in *BudgetResourceGroupObservation) DeepCopy() *BudgetResourceGroupObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupParameters) DeepCopyInto(out *BudgetResourceGroupParameters) {
	*out = *in
	if in.Amount != nil {
		in, out := &in.Amount, &out.Amount
		*out = new(float64)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]FilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]NotificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupIDRef != nil {
		in, out := &in.ResourceGroupIDRef, &out.ResourceGroupIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupIDSelector != nil {
		in, out := &in.ResourceGroupIDSelector, &out.ResourceGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeGrain != nil {
		in, out := &in.TimeGrain, &out.TimeGrain
		*out = new(string)
		**out = **in
	}
	if in.TimePeriod != nil {
		in, out := &in.TimePeriod, &out.TimePeriod
		*out = make([]TimePeriodParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupParameters.
func (in *BudgetResourceGroupParameters) DeepCopy() *BudgetResourceGroupParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupSpec) DeepCopyInto(out *BudgetResourceGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupSpec.
func (in *BudgetResourceGroupSpec) DeepCopy() *BudgetResourceGroupSpec {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetResourceGroupStatus) DeepCopyInto(out *BudgetResourceGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetResourceGroupStatus.
func (in *BudgetResourceGroupStatus) DeepCopy() *BudgetResourceGroupStatus {
	if in == nil {
		return nil
	}
	out := new(BudgetResourceGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscription) DeepCopyInto(out *BudgetSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscription.
func (in *BudgetSubscription) DeepCopy() *BudgetSubscription {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionFilterObservation) DeepCopyInto(out *BudgetSubscriptionFilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionFilterObservation.
func (in *BudgetSubscriptionFilterObservation) DeepCopy() *BudgetSubscriptionFilterObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionFilterParameters) DeepCopyInto(out *BudgetSubscriptionFilterParameters) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]FilterDimensionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = make([]FilterNotParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]BudgetSubscriptionFilterTagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionFilterParameters.
func (in *BudgetSubscriptionFilterParameters) DeepCopy() *BudgetSubscriptionFilterParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionFilterTagObservation) DeepCopyInto(out *BudgetSubscriptionFilterTagObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionFilterTagObservation.
func (in *BudgetSubscriptionFilterTagObservation) DeepCopy() *BudgetSubscriptionFilterTagObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionFilterTagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionFilterTagParameters) DeepCopyInto(out *BudgetSubscriptionFilterTagParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionFilterTagParameters.
func (in *BudgetSubscriptionFilterTagParameters) DeepCopy() *BudgetSubscriptionFilterTagParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionFilterTagParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionList) DeepCopyInto(out *BudgetSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BudgetSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionList.
func (in *BudgetSubscriptionList) DeepCopy() *BudgetSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionNotificationObservation) DeepCopyInto(out *BudgetSubscriptionNotificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionNotificationObservation.
func (in *BudgetSubscriptionNotificationObservation) DeepCopy() *BudgetSubscriptionNotificationObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionNotificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionNotificationParameters) DeepCopyInto(out *BudgetSubscriptionNotificationParameters) {
	*out = *in
	if in.ContactEmails != nil {
		in, out := &in.ContactEmails, &out.ContactEmails
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ContactGroups != nil {
		in, out := &in.ContactGroups, &out.ContactGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ContactRoles != nil {
		in, out := &in.ContactRoles, &out.ContactRoles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdType != nil {
		in, out := &in.ThresholdType, &out.ThresholdType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionNotificationParameters.
func (in *BudgetSubscriptionNotificationParameters) DeepCopy() *BudgetSubscriptionNotificationParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionNotificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionObservation) DeepCopyInto(out *BudgetSubscriptionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionObservation.
func (in *BudgetSubscriptionObservation) DeepCopy() *BudgetSubscriptionObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionParameters) DeepCopyInto(out *BudgetSubscriptionParameters) {
	*out = *in
	if in.Amount != nil {
		in, out := &in.Amount, &out.Amount
		*out = new(float64)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = make([]BudgetSubscriptionFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]BudgetSubscriptionNotificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubscriptionID != nil {
		in, out := &in.SubscriptionID, &out.SubscriptionID
		*out = new(string)
		**out = **in
	}
	if in.TimeGrain != nil {
		in, out := &in.TimeGrain, &out.TimeGrain
		*out = new(string)
		**out = **in
	}
	if in.TimePeriod != nil {
		in, out := &in.TimePeriod, &out.TimePeriod
		*out = make([]BudgetSubscriptionTimePeriodParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionParameters.
func (in *BudgetSubscriptionParameters) DeepCopy() *BudgetSubscriptionParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionSpec) DeepCopyInto(out *BudgetSubscriptionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionSpec.
func (in *BudgetSubscriptionSpec) DeepCopy() *BudgetSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionStatus) DeepCopyInto(out *BudgetSubscriptionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionStatus.
func (in *BudgetSubscriptionStatus) DeepCopy() *BudgetSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionTimePeriodObservation) DeepCopyInto(out *BudgetSubscriptionTimePeriodObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionTimePeriodObservation.
func (in *BudgetSubscriptionTimePeriodObservation) DeepCopy() *BudgetSubscriptionTimePeriodObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionTimePeriodObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSubscriptionTimePeriodParameters) DeepCopyInto(out *BudgetSubscriptionTimePeriodParameters) {
	*out = *in
	if in.EndDate != nil {
		in, out := &in.EndDate, &out.EndDate
		*out = new(string)
		**out = **in
	}
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSubscriptionTimePeriodParameters.
func (in *BudgetSubscriptionTimePeriodParameters) DeepCopy() *BudgetSubscriptionTimePeriodParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetSubscriptionTimePeriodParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionObservation) DeepCopyInto(out *DimensionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionObservation.
func (in *DimensionObservation) DeepCopy() *DimensionObservation {
	if in == nil {
		return nil
	}
	out := new(DimensionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionParameters) DeepCopyInto(out *DimensionParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionParameters.
func (in *DimensionParameters) DeepCopy() *DimensionParameters {
	if in == nil {
		return nil
	}
	out := new(DimensionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterDimensionObservation) DeepCopyInto(out *FilterDimensionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterDimensionObservation.
func (in *FilterDimensionObservation) DeepCopy() *FilterDimensionObservation {
	if in == nil {
		return nil
	}
	out := new(FilterDimensionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterDimensionParameters) DeepCopyInto(out *FilterDimensionParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterDimensionParameters.
func (in *FilterDimensionParameters) DeepCopy() *FilterDimensionParameters {
	if in == nil {
		return nil
	}
	out := new(FilterDimensionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterNotDimensionObservation) DeepCopyInto(out *FilterNotDimensionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterNotDimensionObservation.
func (in *FilterNotDimensionObservation) DeepCopy() *FilterNotDimensionObservation {
	if in == nil {
		return nil
	}
	out := new(FilterNotDimensionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterNotDimensionParameters) DeepCopyInto(out *FilterNotDimensionParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterNotDimensionParameters.
func (in *FilterNotDimensionParameters) DeepCopy() *FilterNotDimensionParameters {
	if in == nil {
		return nil
	}
	out := new(FilterNotDimensionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterNotObservation) DeepCopyInto(out *FilterNotObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterNotObservation.
func (in *FilterNotObservation) DeepCopy() *FilterNotObservation {
	if in == nil {
		return nil
	}
	out := new(FilterNotObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterNotParameters) DeepCopyInto(out *FilterNotParameters) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]FilterNotDimensionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]NotTagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterNotParameters.
func (in *FilterNotParameters) DeepCopy() *FilterNotParameters {
	if in == nil {
		return nil
	}
	out := new(FilterNotParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterObservation) DeepCopyInto(out *FilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterObservation.
func (in *FilterObservation) DeepCopy() *FilterObservation {
	if in == nil {
		return nil
	}
	out := new(FilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterParameters) DeepCopyInto(out *FilterParameters) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]DimensionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = make([]NotParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]FilterTagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterParameters.
func (in *FilterParameters) DeepCopy() *FilterParameters {
	if in == nil {
		return nil
	}
	out := new(FilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterTagObservation) DeepCopyInto(out *FilterTagObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterTagObservation.
func (in *FilterTagObservation) DeepCopy() *FilterTagObservation {
	if in == nil {
		return nil
	}
	out := new(FilterTagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterTagParameters) DeepCopyInto(out *FilterTagParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterTagParameters.
func (in *FilterTagParameters) DeepCopy() *FilterTagParameters {
	if in == nil {
		return nil
	}
	out := new(FilterTagParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotDimensionObservation) DeepCopyInto(out *NotDimensionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotDimensionObservation.
func (in *NotDimensionObservation) DeepCopy() *NotDimensionObservation {
	if in == nil {
		return nil
	}
	out := new(NotDimensionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotDimensionParameters) DeepCopyInto(out *NotDimensionParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotDimensionParameters.
func (in *NotDimensionParameters) DeepCopy() *NotDimensionParameters {
	if in == nil {
		return nil
	}
	out := new(NotDimensionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotObservation) DeepCopyInto(out *NotObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotObservation.
func (in *NotObservation) DeepCopy() *NotObservation {
	if in == nil {
		return nil
	}
	out := new(NotObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotParameters) DeepCopyInto(out *NotParameters) {
	*out = *in
	if in.Dimension != nil {
		in, out := &in.Dimension, &out.Dimension
		*out = make([]NotDimensionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotParameters.
func (in *NotParameters) DeepCopy() *NotParameters {
	if in == nil {
		return nil
	}
	out := new(NotParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotTagObservation) DeepCopyInto(out *NotTagObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotTagObservation.
func (in *NotTagObservation) DeepCopy() *NotTagObservation {
	if in == nil {
		return nil
	}
	out := new(NotTagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotTagParameters) DeepCopyInto(out *NotTagParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotTagParameters.
func (in *NotTagParameters) DeepCopy() *NotTagParameters {
	if in == nil {
		return nil
	}
	out := new(NotTagParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationObservation) DeepCopyInto(out *NotificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationObservation.
func (in *NotificationObservation) DeepCopy() *NotificationObservation {
	if in == nil {
		return nil
	}
	out := new(NotificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationParameters) DeepCopyInto(out *NotificationParameters) {
	*out = *in
	if in.ContactEmails != nil {
		in, out := &in.ContactEmails, &out.ContactEmails
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ContactGroups != nil {
		in, out := &in.ContactGroups, &out.ContactGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ContactRoles != nil {
		in, out := &in.ContactRoles, &out.ContactRoles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdType != nil {
		in, out := &in.ThresholdType, &out.ThresholdType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationParameters.
func (in *NotificationParameters) DeepCopy() *NotificationParameters {
	if in == nil {
		return nil
	}
	out := new(NotificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagObservation) DeepCopyInto(out *TagObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagObservation.
func (in *TagObservation) DeepCopy() *TagObservation {
	if in == nil {
		return nil
	}
	out := new(TagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagParameters) DeepCopyInto(out *TagParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagParameters.
func (in *TagParameters) DeepCopy() *TagParameters {
	if in == nil {
		return nil
	}
	out := new(TagParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimePeriodObservation) DeepCopyInto(out *TimePeriodObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimePeriodObservation.
func (in *TimePeriodObservation) DeepCopy() *TimePeriodObservation {
	if in == nil {
		return nil
	}
	out := new(TimePeriodObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimePeriodParameters) DeepCopyInto(out *TimePeriodParameters) {
	*out = *in
	if in.EndDate != nil {
		in, out := &in.EndDate, &out.EndDate
		*out = new(string)
		**out = **in
	}
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimePeriodParameters.
func (in *TimePeriodParameters) DeepCopy() *TimePeriodParameters {
	if in == nil {
		return nil
	}
	out := new(TimePeriodParameters)
	in.DeepCopyInto(out)
	return out
}