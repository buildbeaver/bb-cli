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

// checks if the Step type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Step{}

// Step struct for Step
type Step struct {
	// A link to the Step resource on the BuildBeaver server
	Url string `json:"url"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Etag string `json:"etag"`
	// Step name, in URL format
	Name string `json:"name"`
	// Optional human-readable description of the job.
	Description string `json:"description"`
	// A list of one or more shell commands to execute during the step.
	Commands []string `json:"commands"`
	// Dependencies this step has on other steps within the job (see dependency syntax)
	Depends []StepDependency `json:"depends"`
	// ID of the job this step forms a part of
	JobId string `json:"job_id"`
	// RepoID that the step is building from.
	RepoId string `json:"repo_id"`
	// RunnerID that ran the step (or empty if the step has not run yet).
	RunnerId string `json:"runner_id"`
	// LogDescriptorID points to the log for this step.
	LogDescriptorId string `json:"log_descriptor_id"`
	// Status reflects where the step is in processing.
	Status string `json:"status"`
	// Error is set if the step finished with an error (or empty if the step succeeded).
	Error *string `json:"error,omitempty"`
	Timings WorkflowTimings `json:"timings"`
	// URL of the runner that ran the step (or empty if the step has not run yet).
	RunnerUrl string `json:"runner_url"`
	// URL of the log for this step.
	LogDescriptorUrl string `json:"log_descriptor_url"`
	AdditionalProperties map[string]interface{}
}

type _Step Step

// NewStep instantiates a new Step object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStep(url string, id string, createdAt time.Time, updatedAt time.Time, etag string, name string, description string, commands []string, depends []StepDependency, jobId string, repoId string, runnerId string, logDescriptorId string, status string, timings WorkflowTimings, runnerUrl string, logDescriptorUrl string) *Step {
	this := Step{}
	this.Url = url
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Etag = etag
	this.Name = name
	this.Description = description
	this.Commands = commands
	this.Depends = depends
	this.JobId = jobId
	this.RepoId = repoId
	this.RunnerId = runnerId
	this.LogDescriptorId = logDescriptorId
	this.Status = status
	this.Timings = timings
	this.RunnerUrl = runnerUrl
	this.LogDescriptorUrl = logDescriptorUrl
	return &this
}

// NewStepWithDefaults instantiates a new Step object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStepWithDefaults() *Step {
	this := Step{}
	return &this
}

// GetUrl returns the Url field value
func (o *Step) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Step) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Step) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *Step) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Step) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Step) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Step) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Step) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Step) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Step) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Step) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Step) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Step) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Step) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Step) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Step) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetEtag returns the Etag field value
func (o *Step) GetEtag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Etag
}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
func (o *Step) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etag, true
}

// SetEtag sets field value
func (o *Step) SetEtag(v string) {
	o.Etag = v
}

// GetName returns the Name field value
func (o *Step) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Step) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Step) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Step) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Step) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Step) SetDescription(v string) {
	o.Description = v
}

// GetCommands returns the Commands field value
func (o *Step) GetCommands() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value
// and a boolean to check if the value has been set.
func (o *Step) GetCommandsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Commands, true
}

// SetCommands sets field value
func (o *Step) SetCommands(v []string) {
	o.Commands = v
}

// GetDepends returns the Depends field value
func (o *Step) GetDepends() []StepDependency {
	if o == nil {
		var ret []StepDependency
		return ret
	}

	return o.Depends
}

// GetDependsOk returns a tuple with the Depends field value
// and a boolean to check if the value has been set.
func (o *Step) GetDependsOk() ([]StepDependency, bool) {
	if o == nil {
		return nil, false
	}
	return o.Depends, true
}

// SetDepends sets field value
func (o *Step) SetDepends(v []StepDependency) {
	o.Depends = v
}

// GetJobId returns the JobId field value
func (o *Step) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *Step) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *Step) SetJobId(v string) {
	o.JobId = v
}

// GetRepoId returns the RepoId field value
func (o *Step) GetRepoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value
// and a boolean to check if the value has been set.
func (o *Step) GetRepoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepoId, true
}

// SetRepoId sets field value
func (o *Step) SetRepoId(v string) {
	o.RepoId = v
}

// GetRunnerId returns the RunnerId field value
func (o *Step) GetRunnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunnerId
}

// GetRunnerIdOk returns a tuple with the RunnerId field value
// and a boolean to check if the value has been set.
func (o *Step) GetRunnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunnerId, true
}

// SetRunnerId sets field value
func (o *Step) SetRunnerId(v string) {
	o.RunnerId = v
}

// GetLogDescriptorId returns the LogDescriptorId field value
func (o *Step) GetLogDescriptorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogDescriptorId
}

// GetLogDescriptorIdOk returns a tuple with the LogDescriptorId field value
// and a boolean to check if the value has been set.
func (o *Step) GetLogDescriptorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogDescriptorId, true
}

// SetLogDescriptorId sets field value
func (o *Step) SetLogDescriptorId(v string) {
	o.LogDescriptorId = v
}

// GetStatus returns the Status field value
func (o *Step) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Step) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Step) SetStatus(v string) {
	o.Status = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *Step) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Step) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *Step) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *Step) SetError(v string) {
	o.Error = &v
}

// GetTimings returns the Timings field value
func (o *Step) GetTimings() WorkflowTimings {
	if o == nil {
		var ret WorkflowTimings
		return ret
	}

	return o.Timings
}

// GetTimingsOk returns a tuple with the Timings field value
// and a boolean to check if the value has been set.
func (o *Step) GetTimingsOk() (*WorkflowTimings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timings, true
}

// SetTimings sets field value
func (o *Step) SetTimings(v WorkflowTimings) {
	o.Timings = v
}

// GetRunnerUrl returns the RunnerUrl field value
func (o *Step) GetRunnerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunnerUrl
}

// GetRunnerUrlOk returns a tuple with the RunnerUrl field value
// and a boolean to check if the value has been set.
func (o *Step) GetRunnerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunnerUrl, true
}

// SetRunnerUrl sets field value
func (o *Step) SetRunnerUrl(v string) {
	o.RunnerUrl = v
}

// GetLogDescriptorUrl returns the LogDescriptorUrl field value
func (o *Step) GetLogDescriptorUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogDescriptorUrl
}

// GetLogDescriptorUrlOk returns a tuple with the LogDescriptorUrl field value
// and a boolean to check if the value has been set.
func (o *Step) GetLogDescriptorUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogDescriptorUrl, true
}

// SetLogDescriptorUrl sets field value
func (o *Step) SetLogDescriptorUrl(v string) {
	o.LogDescriptorUrl = v
}

func (o Step) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Step) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	toSerialize["etag"] = o.Etag
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["commands"] = o.Commands
	toSerialize["depends"] = o.Depends
	toSerialize["job_id"] = o.JobId
	toSerialize["repo_id"] = o.RepoId
	toSerialize["runner_id"] = o.RunnerId
	toSerialize["log_descriptor_id"] = o.LogDescriptorId
	toSerialize["status"] = o.Status
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	toSerialize["timings"] = o.Timings
	toSerialize["runner_url"] = o.RunnerUrl
	toSerialize["log_descriptor_url"] = o.LogDescriptorUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Step) UnmarshalJSON(bytes []byte) (err error) {
	varStep := _Step{}

	if err = json.Unmarshal(bytes, &varStep); err == nil {
		*o = Step(varStep)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "etag")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "commands")
		delete(additionalProperties, "depends")
		delete(additionalProperties, "job_id")
		delete(additionalProperties, "repo_id")
		delete(additionalProperties, "runner_id")
		delete(additionalProperties, "log_descriptor_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "error")
		delete(additionalProperties, "timings")
		delete(additionalProperties, "runner_url")
		delete(additionalProperties, "log_descriptor_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStep struct {
	value *Step
	isSet bool
}

func (v NullableStep) Get() *Step {
	return v.value
}

func (v *NullableStep) Set(val *Step) {
	v.value = val
	v.isSet = true
}

func (v NullableStep) IsSet() bool {
	return v.isSet
}

func (v *NullableStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStep(val *Step) *NullableStep {
	return &NullableStep{value: val, isSet: true}
}

func (v NullableStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


