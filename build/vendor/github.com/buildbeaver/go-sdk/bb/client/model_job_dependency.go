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

// checks if the JobDependency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobDependency{}

// JobDependency struct for JobDependency
type JobDependency struct {
	// The name of the workflow containing the job that must complete, or an empty string for the default workflow.
	Workflow string `json:"workflow"`
	// The name of the job that must complete.
	JobName string `json:"job_name"`
	// The set of artifacts required from the job that must complete.
	ArtifactDependencies []ArtifactDependency `json:"artifact_dependencies"`
	AdditionalProperties map[string]interface{}
}

type _JobDependency JobDependency

// NewJobDependency instantiates a new JobDependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDependency(workflow string, jobName string, artifactDependencies []ArtifactDependency) *JobDependency {
	this := JobDependency{}
	this.Workflow = workflow
	this.JobName = jobName
	this.ArtifactDependencies = artifactDependencies
	return &this
}

// NewJobDependencyWithDefaults instantiates a new JobDependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDependencyWithDefaults() *JobDependency {
	this := JobDependency{}
	return &this
}

// GetWorkflow returns the Workflow field value
func (o *JobDependency) GetWorkflow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
func (o *JobDependency) GetWorkflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workflow, true
}

// SetWorkflow sets field value
func (o *JobDependency) SetWorkflow(v string) {
	o.Workflow = v
}

// GetJobName returns the JobName field value
func (o *JobDependency) GetJobName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobName
}

// GetJobNameOk returns a tuple with the JobName field value
// and a boolean to check if the value has been set.
func (o *JobDependency) GetJobNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobName, true
}

// SetJobName sets field value
func (o *JobDependency) SetJobName(v string) {
	o.JobName = v
}

// GetArtifactDependencies returns the ArtifactDependencies field value
func (o *JobDependency) GetArtifactDependencies() []ArtifactDependency {
	if o == nil {
		var ret []ArtifactDependency
		return ret
	}

	return o.ArtifactDependencies
}

// GetArtifactDependenciesOk returns a tuple with the ArtifactDependencies field value
// and a boolean to check if the value has been set.
func (o *JobDependency) GetArtifactDependenciesOk() ([]ArtifactDependency, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArtifactDependencies, true
}

// SetArtifactDependencies sets field value
func (o *JobDependency) SetArtifactDependencies(v []ArtifactDependency) {
	o.ArtifactDependencies = v
}

func (o JobDependency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobDependency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workflow"] = o.Workflow
	toSerialize["job_name"] = o.JobName
	toSerialize["artifact_dependencies"] = o.ArtifactDependencies

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobDependency) UnmarshalJSON(bytes []byte) (err error) {
	varJobDependency := _JobDependency{}

	if err = json.Unmarshal(bytes, &varJobDependency); err == nil {
		*o = JobDependency(varJobDependency)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "workflow")
		delete(additionalProperties, "job_name")
		delete(additionalProperties, "artifact_dependencies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobDependency struct {
	value *JobDependency
	isSet bool
}

func (v NullableJobDependency) Get() *JobDependency {
	return v.value
}

func (v *NullableJobDependency) Set(val *JobDependency) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDependency(val *JobDependency) *NullableJobDependency {
	return &NullableJobDependency{value: val, isSet: true}
}

func (v NullableJobDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


