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

// checks if the JobDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobDefinition{}

// JobDefinition struct for JobDefinition
type JobDefinition struct {
	// Job name. This can optionally include a workflow as a prefix (dot-separated), as an alternative to specifying an explicit 'workflow' element.
	Name string `json:"name"`
	// Workflow the job is a part of. If not specified then the job is part of the default workflow.
	Workflow *string `json:"workflow,omitempty"`
	// Optional human-readable description of the job.
	Description *string `json:"description,omitempty"`
	// Type of the job (e.g. docker, exec etc.)
	Type *string `json:"type,omitempty"`
	// RunsOn contains a set of labels that this job requires runners to have.
	RunsOn []string `json:"runs_on,omitempty"`
	Docker *DockerConfigDefinition `json:"docker,omitempty"`
	// Determines how the runner will execute steps within this job
	StepExecution string `json:"step_execution"`
	// Dependencies on other jobs and their artifacts (see dependency syntax)
	Depends []string `json:"depends,omitempty"`
	// Services to run in the background for the duration of the job; services are started before the first step is run, and stopped after the last step completes
	Services []ServiceDefinition `json:"services,omitempty"`
	// Shell commands to execute to generate a unique fingerprint for the jobs; two jobs in the same repo with the same name and fingerprint are considered identical
	Fingerprint []string `json:"fingerprint,omitempty"`
	// A list of all artifacts the job is expected to produce that will be saved to the artifact store at the end of the job's execution
	Artifacts []ArtifactDefinition `json:"artifacts,omitempty"`
	// A list of environment variables to export prior to executing the job
	Environment map[string]SecretStringDefinition `json:"environment"`
	// The set of steps within the job
	Steps []StepDefinition `json:"steps"`
	AdditionalProperties map[string]interface{}
}

type _JobDefinition JobDefinition

// NewJobDefinition instantiates a new JobDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDefinition(name string, stepExecution string, environment map[string]SecretStringDefinition, steps []StepDefinition) *JobDefinition {
	this := JobDefinition{}
	this.Name = name
	this.StepExecution = stepExecution
	this.Environment = environment
	this.Steps = steps
	return &this
}

// NewJobDefinitionWithDefaults instantiates a new JobDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDefinitionWithDefaults() *JobDefinition {
	this := JobDefinition{}
	return &this
}

// GetName returns the Name field value
func (o *JobDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobDefinition) SetName(v string) {
	o.Name = v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *JobDefinition) GetWorkflow() string {
	if o == nil || IsNil(o.Workflow) {
		var ret string
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetWorkflowOk() (*string, bool) {
	if o == nil || IsNil(o.Workflow) {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *JobDefinition) HasWorkflow() bool {
	if o != nil && !IsNil(o.Workflow) {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given string and assigns it to the Workflow field.
func (o *JobDefinition) SetWorkflow(v string) {
	o.Workflow = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JobDefinition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JobDefinition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JobDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *JobDefinition) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *JobDefinition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *JobDefinition) SetType(v string) {
	o.Type = &v
}

// GetRunsOn returns the RunsOn field value if set, zero value otherwise.
func (o *JobDefinition) GetRunsOn() []string {
	if o == nil || IsNil(o.RunsOn) {
		var ret []string
		return ret
	}
	return o.RunsOn
}

// GetRunsOnOk returns a tuple with the RunsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetRunsOnOk() ([]string, bool) {
	if o == nil || IsNil(o.RunsOn) {
		return nil, false
	}
	return o.RunsOn, true
}

// HasRunsOn returns a boolean if a field has been set.
func (o *JobDefinition) HasRunsOn() bool {
	if o != nil && !IsNil(o.RunsOn) {
		return true
	}

	return false
}

// SetRunsOn gets a reference to the given []string and assigns it to the RunsOn field.
func (o *JobDefinition) SetRunsOn(v []string) {
	o.RunsOn = v
}

// GetDocker returns the Docker field value if set, zero value otherwise.
func (o *JobDefinition) GetDocker() DockerConfigDefinition {
	if o == nil || IsNil(o.Docker) {
		var ret DockerConfigDefinition
		return ret
	}
	return *o.Docker
}

// GetDockerOk returns a tuple with the Docker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetDockerOk() (*DockerConfigDefinition, bool) {
	if o == nil || IsNil(o.Docker) {
		return nil, false
	}
	return o.Docker, true
}

// HasDocker returns a boolean if a field has been set.
func (o *JobDefinition) HasDocker() bool {
	if o != nil && !IsNil(o.Docker) {
		return true
	}

	return false
}

// SetDocker gets a reference to the given DockerConfigDefinition and assigns it to the Docker field.
func (o *JobDefinition) SetDocker(v DockerConfigDefinition) {
	o.Docker = &v
}

// GetStepExecution returns the StepExecution field value
func (o *JobDefinition) GetStepExecution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StepExecution
}

// GetStepExecutionOk returns a tuple with the StepExecution field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetStepExecutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepExecution, true
}

// SetStepExecution sets field value
func (o *JobDefinition) SetStepExecution(v string) {
	o.StepExecution = v
}

// GetDepends returns the Depends field value if set, zero value otherwise.
func (o *JobDefinition) GetDepends() []string {
	if o == nil || IsNil(o.Depends) {
		var ret []string
		return ret
	}
	return o.Depends
}

// GetDependsOk returns a tuple with the Depends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetDependsOk() ([]string, bool) {
	if o == nil || IsNil(o.Depends) {
		return nil, false
	}
	return o.Depends, true
}

// HasDepends returns a boolean if a field has been set.
func (o *JobDefinition) HasDepends() bool {
	if o != nil && !IsNil(o.Depends) {
		return true
	}

	return false
}

// SetDepends gets a reference to the given []string and assigns it to the Depends field.
func (o *JobDefinition) SetDepends(v []string) {
	o.Depends = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *JobDefinition) GetServices() []ServiceDefinition {
	if o == nil || IsNil(o.Services) {
		var ret []ServiceDefinition
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetServicesOk() ([]ServiceDefinition, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *JobDefinition) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []ServiceDefinition and assigns it to the Services field.
func (o *JobDefinition) SetServices(v []ServiceDefinition) {
	o.Services = v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *JobDefinition) GetFingerprint() []string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret []string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetFingerprintOk() ([]string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *JobDefinition) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given []string and assigns it to the Fingerprint field.
func (o *JobDefinition) SetFingerprint(v []string) {
	o.Fingerprint = v
}

// GetArtifacts returns the Artifacts field value if set, zero value otherwise.
func (o *JobDefinition) GetArtifacts() []ArtifactDefinition {
	if o == nil || IsNil(o.Artifacts) {
		var ret []ArtifactDefinition
		return ret
	}
	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetArtifactsOk() ([]ArtifactDefinition, bool) {
	if o == nil || IsNil(o.Artifacts) {
		return nil, false
	}
	return o.Artifacts, true
}

// HasArtifacts returns a boolean if a field has been set.
func (o *JobDefinition) HasArtifacts() bool {
	if o != nil && !IsNil(o.Artifacts) {
		return true
	}

	return false
}

// SetArtifacts gets a reference to the given []ArtifactDefinition and assigns it to the Artifacts field.
func (o *JobDefinition) SetArtifacts(v []ArtifactDefinition) {
	o.Artifacts = v
}

// GetEnvironment returns the Environment field value
func (o *JobDefinition) GetEnvironment() map[string]SecretStringDefinition {
	if o == nil {
		var ret map[string]SecretStringDefinition
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetEnvironmentOk() (*map[string]SecretStringDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *JobDefinition) SetEnvironment(v map[string]SecretStringDefinition) {
	o.Environment = v
}

// GetSteps returns the Steps field value
func (o *JobDefinition) GetSteps() []StepDefinition {
	if o == nil {
		var ret []StepDefinition
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetStepsOk() ([]StepDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *JobDefinition) SetSteps(v []StepDefinition) {
	o.Steps = v
}

func (o JobDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Workflow) {
		toSerialize["workflow"] = o.Workflow
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RunsOn) {
		toSerialize["runs_on"] = o.RunsOn
	}
	if !IsNil(o.Docker) {
		toSerialize["docker"] = o.Docker
	}
	toSerialize["step_execution"] = o.StepExecution
	if !IsNil(o.Depends) {
		toSerialize["depends"] = o.Depends
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if !IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if !IsNil(o.Artifacts) {
		toSerialize["artifacts"] = o.Artifacts
	}
	toSerialize["environment"] = o.Environment
	toSerialize["steps"] = o.Steps

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobDefinition) UnmarshalJSON(bytes []byte) (err error) {
	varJobDefinition := _JobDefinition{}

	if err = json.Unmarshal(bytes, &varJobDefinition); err == nil {
		*o = JobDefinition(varJobDefinition)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "workflow")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "runs_on")
		delete(additionalProperties, "docker")
		delete(additionalProperties, "step_execution")
		delete(additionalProperties, "depends")
		delete(additionalProperties, "services")
		delete(additionalProperties, "fingerprint")
		delete(additionalProperties, "artifacts")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "steps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobDefinition struct {
	value *JobDefinition
	isSet bool
}

func (v NullableJobDefinition) Get() *JobDefinition {
	return v.value
}

func (v *NullableJobDefinition) Set(val *JobDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDefinition(val *JobDefinition) *NullableJobDefinition {
	return &NullableJobDefinition{value: val, isSet: true}
}

func (v NullableJobDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

