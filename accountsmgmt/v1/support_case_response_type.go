/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// SupportCaseResponseKind is the name of the type used to represent objects
// of type 'support_case_response'.
const SupportCaseResponseKind = "SupportCaseResponse"

// SupportCaseResponseLinkKind is the name of the type used to represent links
// to objects of type 'support_case_response'.
const SupportCaseResponseLinkKind = "SupportCaseResponseLink"

// SupportCaseResponseNilKind is the name of the type used to nil references
// to objects of type 'support_case_response'.
const SupportCaseResponseNilKind = "SupportCaseResponseNil"

// SupportCaseResponse represents the values of the 'support_case_response' type.
//
//
type SupportCaseResponse struct {
	id             *string
	href           *string
	link           bool
	uri            *string
	caseNumber     *string
	clusterId      *string
	clusterUuid    *string
	description    *string
	severity       *string
	status         *string
	subscriptionId *string
	summary        *string
}

// Kind returns the name of the type of the object.
func (o *SupportCaseResponse) Kind() string {
	if o == nil {
		return SupportCaseResponseNilKind
	}
	if o.link {
		return SupportCaseResponseLinkKind
	}
	return SupportCaseResponseKind
}

// ID returns the identifier of the object.
func (o *SupportCaseResponse) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *SupportCaseResponse) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *SupportCaseResponse) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *SupportCaseResponse) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *SupportCaseResponse) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *SupportCaseResponse) Empty() bool {
	return o == nil || (o.id == nil &&
		o.uri == nil &&
		o.caseNumber == nil &&
		o.clusterId == nil &&
		o.clusterUuid == nil &&
		o.description == nil &&
		o.severity == nil &&
		o.status == nil &&
		o.subscriptionId == nil &&
		o.summary == nil &&
		true)
}

// URI returns the value of the 'URI' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Support case uri.
func (o *SupportCaseResponse) URI() string {
	if o != nil && o.uri != nil {
		return *o.uri
	}
	return ""
}

// GetURI returns the value of the 'URI' attribute and
// a flag indicating if the attribute has a value.
//
// Support case uri.
func (o *SupportCaseResponse) GetURI() (value string, ok bool) {
	ok = o != nil && o.uri != nil
	if ok {
		value = *o.uri
	}
	return
}

// CaseNumber returns the value of the 'case_number' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Support case number.
func (o *SupportCaseResponse) CaseNumber() string {
	if o != nil && o.caseNumber != nil {
		return *o.caseNumber
	}
	return ""
}

// GetCaseNumber returns the value of the 'case_number' attribute and
// a flag indicating if the attribute has a value.
//
// Support case number.
func (o *SupportCaseResponse) GetCaseNumber() (value string, ok bool) {
	ok = o != nil && o.caseNumber != nil
	if ok {
		value = *o.caseNumber
	}
	return
}

// ClusterId returns the value of the 'cluster_id' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// (optional) cluster id of the cluster on which we created the support case for.
func (o *SupportCaseResponse) ClusterId() string {
	if o != nil && o.clusterId != nil {
		return *o.clusterId
	}
	return ""
}

// GetClusterId returns the value of the 'cluster_id' attribute and
// a flag indicating if the attribute has a value.
//
// (optional) cluster id of the cluster on which we created the support case for.
func (o *SupportCaseResponse) GetClusterId() (value string, ok bool) {
	ok = o != nil && o.clusterId != nil
	if ok {
		value = *o.clusterId
	}
	return
}

// ClusterUuid returns the value of the 'cluster_uuid' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// (optional) cluster uuid of the cluster on which we created the support case for.
func (o *SupportCaseResponse) ClusterUuid() string {
	if o != nil && o.clusterUuid != nil {
		return *o.clusterUuid
	}
	return ""
}

// GetClusterUuid returns the value of the 'cluster_uuid' attribute and
// a flag indicating if the attribute has a value.
//
// (optional) cluster uuid of the cluster on which we created the support case for.
func (o *SupportCaseResponse) GetClusterUuid() (value string, ok bool) {
	ok = o != nil && o.clusterUuid != nil
	if ok {
		value = *o.clusterUuid
	}
	return
}

// Description returns the value of the 'description' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Support case desciption.
func (o *SupportCaseResponse) Description() string {
	if o != nil && o.description != nil {
		return *o.description
	}
	return ""
}

// GetDescription returns the value of the 'description' attribute and
// a flag indicating if the attribute has a value.
//
// Support case desciption.
func (o *SupportCaseResponse) GetDescription() (value string, ok bool) {
	ok = o != nil && o.description != nil
	if ok {
		value = *o.description
	}
	return
}

// Severity returns the value of the 'severity' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Support case severity.
func (o *SupportCaseResponse) Severity() string {
	if o != nil && o.severity != nil {
		return *o.severity
	}
	return ""
}

// GetSeverity returns the value of the 'severity' attribute and
// a flag indicating if the attribute has a value.
//
// Support case severity.
func (o *SupportCaseResponse) GetSeverity() (value string, ok bool) {
	ok = o != nil && o.severity != nil
	if ok {
		value = *o.severity
	}
	return
}

// Status returns the value of the 'status' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Support case status.
func (o *SupportCaseResponse) Status() string {
	if o != nil && o.status != nil {
		return *o.status
	}
	return ""
}

// GetStatus returns the value of the 'status' attribute and
// a flag indicating if the attribute has a value.
//
// Support case status.
func (o *SupportCaseResponse) GetStatus() (value string, ok bool) {
	ok = o != nil && o.status != nil
	if ok {
		value = *o.status
	}
	return
}

// SubscriptionId returns the value of the 'subscription_id' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// (optional) subscription id of the subscription on which we created the support case for.
func (o *SupportCaseResponse) SubscriptionId() string {
	if o != nil && o.subscriptionId != nil {
		return *o.subscriptionId
	}
	return ""
}

// GetSubscriptionId returns the value of the 'subscription_id' attribute and
// a flag indicating if the attribute has a value.
//
// (optional) subscription id of the subscription on which we created the support case for.
func (o *SupportCaseResponse) GetSubscriptionId() (value string, ok bool) {
	ok = o != nil && o.subscriptionId != nil
	if ok {
		value = *o.subscriptionId
	}
	return
}

// Summary returns the value of the 'summary' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Support case title.
func (o *SupportCaseResponse) Summary() string {
	if o != nil && o.summary != nil {
		return *o.summary
	}
	return ""
}

// GetSummary returns the value of the 'summary' attribute and
// a flag indicating if the attribute has a value.
//
// Support case title.
func (o *SupportCaseResponse) GetSummary() (value string, ok bool) {
	ok = o != nil && o.summary != nil
	if ok {
		value = *o.summary
	}
	return
}

// SupportCaseResponseListKind is the name of the type used to represent list of objects of
// type 'support_case_response'.
const SupportCaseResponseListKind = "SupportCaseResponseList"

// SupportCaseResponseListLinkKind is the name of the type used to represent links to list
// of objects of type 'support_case_response'.
const SupportCaseResponseListLinkKind = "SupportCaseResponseListLink"

// SupportCaseResponseNilKind is the name of the type used to nil lists of objects of
// type 'support_case_response'.
const SupportCaseResponseListNilKind = "SupportCaseResponseListNil"

// SupportCaseResponseList is a list of values of the 'support_case_response' type.
type SupportCaseResponseList struct {
	href  *string
	link  bool
	items []*SupportCaseResponse
}

// Kind returns the name of the type of the object.
func (l *SupportCaseResponseList) Kind() string {
	if l == nil {
		return SupportCaseResponseListNilKind
	}
	if l.link {
		return SupportCaseResponseListLinkKind
	}
	return SupportCaseResponseListKind
}

// Link returns true iif this is a link.
func (l *SupportCaseResponseList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *SupportCaseResponseList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *SupportCaseResponseList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *SupportCaseResponseList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *SupportCaseResponseList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *SupportCaseResponseList) Get(i int) *SupportCaseResponse {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *SupportCaseResponseList) Slice() []*SupportCaseResponse {
	var slice []*SupportCaseResponse
	if l == nil {
		slice = make([]*SupportCaseResponse, 0)
	} else {
		slice = make([]*SupportCaseResponse, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *SupportCaseResponseList) Each(f func(item *SupportCaseResponse) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *SupportCaseResponseList) Range(f func(index int, item *SupportCaseResponse) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
