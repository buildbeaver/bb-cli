/*
BuildBeaver Dynamic Build API - OpenAPI 3.0

This is the BuildBeaver Dynamic Build API.

API version: 0.3.00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ExternalResourceID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalResourceID{}

// ExternalResourceID struct for ExternalResourceID
type ExternalResourceID struct {
	// The name of the external system, e.g. GitHub
	ExternalSystem string `json:"external_system"`
	// The resource within the external system, e.g. github_repo.id
	ResourceId string `json:"resource_id"`
	AdditionalProperties map[string]interface{}
}

type _ExternalResourceID ExternalResourceID

// NewExternalResourceID instantiates a new ExternalResourceID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalResourceID(externalSystem string, resourceId string) *ExternalResourceID {
	this := ExternalResourceID{}
	this.ExternalSystem = externalSystem
	this.ResourceId = resourceId
	return &this
}

// NewExternalResourceIDWithDefaults instantiates a new ExternalResourceID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalResourceIDWithDefaults() *ExternalResourceID {
	this := ExternalResourceID{}
	return &this
}

// GetExternalSystem returns the ExternalSystem field value
func (o *ExternalResourceID) GetExternalSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalSystem
}

// GetExternalSystemOk returns a tuple with the ExternalSystem field value
// and a boolean to check if the value has been set.
func (o *ExternalResourceID) GetExternalSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalSystem, true
}

// SetExternalSystem sets field value
func (o *ExternalResourceID) SetExternalSystem(v string) {
	o.ExternalSystem = v
}

// GetResourceId returns the ResourceId field value
func (o *ExternalResourceID) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ExternalResourceID) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *ExternalResourceID) SetResourceId(v string) {
	o.ResourceId = v
}

func (o ExternalResourceID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalResourceID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["external_system"] = o.ExternalSystem
	toSerialize["resource_id"] = o.ResourceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExternalResourceID) UnmarshalJSON(bytes []byte) (err error) {
	varExternalResourceID := _ExternalResourceID{}

	if err = json.Unmarshal(bytes, &varExternalResourceID); err == nil {
		*o = ExternalResourceID(varExternalResourceID)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "external_system")
		delete(additionalProperties, "resource_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalResourceID struct {
	value *ExternalResourceID
	isSet bool
}

func (v NullableExternalResourceID) Get() *ExternalResourceID {
	return v.value
}

func (v *NullableExternalResourceID) Set(val *ExternalResourceID) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalResourceID) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalResourceID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalResourceID(val *ExternalResourceID) *NullableExternalResourceID {
	return &NullableExternalResourceID{value: val, isSet: true}
}

func (v NullableExternalResourceID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalResourceID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


