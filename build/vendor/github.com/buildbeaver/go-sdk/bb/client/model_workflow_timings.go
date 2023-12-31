/*
BuildBeaver Dynamic Build API - OpenAPI 3.0

This is the BuildBeaver Dynamic Build API.

API version: 0.3.00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the WorkflowTimings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowTimings{}

// WorkflowTimings struct for WorkflowTimings
type WorkflowTimings struct {
	QueuedAt *time.Time `json:"queued_at,omitempty"`
	SubmittedAt *time.Time `json:"submitted_at,omitempty"`
	RunningAt *time.Time `json:"running_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	CanceledAt *time.Time `json:"canceled_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTimings WorkflowTimings

// NewWorkflowTimings instantiates a new WorkflowTimings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTimings() *WorkflowTimings {
	this := WorkflowTimings{}
	return &this
}

// NewWorkflowTimingsWithDefaults instantiates a new WorkflowTimings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTimingsWithDefaults() *WorkflowTimings {
	this := WorkflowTimings{}
	return &this
}

// GetQueuedAt returns the QueuedAt field value if set, zero value otherwise.
func (o *WorkflowTimings) GetQueuedAt() time.Time {
	if o == nil || IsNil(o.QueuedAt) {
		var ret time.Time
		return ret
	}
	return *o.QueuedAt
}

// GetQueuedAtOk returns a tuple with the QueuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTimings) GetQueuedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.QueuedAt) {
		return nil, false
	}
	return o.QueuedAt, true
}

// HasQueuedAt returns a boolean if a field has been set.
func (o *WorkflowTimings) HasQueuedAt() bool {
	if o != nil && !IsNil(o.QueuedAt) {
		return true
	}

	return false
}

// SetQueuedAt gets a reference to the given time.Time and assigns it to the QueuedAt field.
func (o *WorkflowTimings) SetQueuedAt(v time.Time) {
	o.QueuedAt = &v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *WorkflowTimings) GetSubmittedAt() time.Time {
	if o == nil || IsNil(o.SubmittedAt) {
		var ret time.Time
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTimings) GetSubmittedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmittedAt) {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *WorkflowTimings) HasSubmittedAt() bool {
	if o != nil && !IsNil(o.SubmittedAt) {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given time.Time and assigns it to the SubmittedAt field.
func (o *WorkflowTimings) SetSubmittedAt(v time.Time) {
	o.SubmittedAt = &v
}

// GetRunningAt returns the RunningAt field value if set, zero value otherwise.
func (o *WorkflowTimings) GetRunningAt() time.Time {
	if o == nil || IsNil(o.RunningAt) {
		var ret time.Time
		return ret
	}
	return *o.RunningAt
}

// GetRunningAtOk returns a tuple with the RunningAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTimings) GetRunningAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RunningAt) {
		return nil, false
	}
	return o.RunningAt, true
}

// HasRunningAt returns a boolean if a field has been set.
func (o *WorkflowTimings) HasRunningAt() bool {
	if o != nil && !IsNil(o.RunningAt) {
		return true
	}

	return false
}

// SetRunningAt gets a reference to the given time.Time and assigns it to the RunningAt field.
func (o *WorkflowTimings) SetRunningAt(v time.Time) {
	o.RunningAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *WorkflowTimings) GetFinishedAt() time.Time {
	if o == nil || IsNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTimings) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *WorkflowTimings) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *WorkflowTimings) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise.
func (o *WorkflowTimings) GetCanceledAt() time.Time {
	if o == nil || IsNil(o.CanceledAt) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTimings) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CanceledAt) {
		return nil, false
	}
	return o.CanceledAt, true
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *WorkflowTimings) HasCanceledAt() bool {
	if o != nil && !IsNil(o.CanceledAt) {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given time.Time and assigns it to the CanceledAt field.
func (o *WorkflowTimings) SetCanceledAt(v time.Time) {
	o.CanceledAt = &v
}

func (o WorkflowTimings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowTimings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QueuedAt) {
		toSerialize["queued_at"] = o.QueuedAt
	}
	if !IsNil(o.SubmittedAt) {
		toSerialize["submitted_at"] = o.SubmittedAt
	}
	if !IsNil(o.RunningAt) {
		toSerialize["running_at"] = o.RunningAt
	}
	if !IsNil(o.FinishedAt) {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if !IsNil(o.CanceledAt) {
		toSerialize["canceled_at"] = o.CanceledAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowTimings) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTimings := _WorkflowTimings{}

	if err = json.Unmarshal(bytes, &varWorkflowTimings); err == nil {
		*o = WorkflowTimings(varWorkflowTimings)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "queued_at")
		delete(additionalProperties, "submitted_at")
		delete(additionalProperties, "running_at")
		delete(additionalProperties, "finished_at")
		delete(additionalProperties, "canceled_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTimings struct {
	value *WorkflowTimings
	isSet bool
}

func (v NullableWorkflowTimings) Get() *WorkflowTimings {
	return v.value
}

func (v *NullableWorkflowTimings) Set(val *WorkflowTimings) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTimings) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTimings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTimings(val *WorkflowTimings) *NullableWorkflowTimings {
	return &NullableWorkflowTimings{value: val, isSet: true}
}

func (v NullableWorkflowTimings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTimings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


