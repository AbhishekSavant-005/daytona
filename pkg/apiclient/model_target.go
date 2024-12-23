/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Target type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Target{}

// Target struct for Target
type Target struct {
	Default          bool              `json:"default"`
	EnvVars          map[string]string `json:"envVars"`
	Id               string            `json:"id"`
	LastJob          *Job              `json:"lastJob,omitempty"`
	LastJobId        *string           `json:"lastJobId,omitempty"`
	Metadata         *TargetMetadata   `json:"metadata,omitempty"`
	Name             string            `json:"name"`
	ProviderMetadata *string           `json:"providerMetadata,omitempty"`
	TargetConfig     TargetConfig      `json:"targetConfig"`
	TargetConfigId   string            `json:"targetConfigId"`
	Workspaces       []Workspace       `json:"workspaces"`
}

type _Target Target

// NewTarget instantiates a new Target object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTarget(default_ bool, envVars map[string]string, id string, name string, targetConfig TargetConfig, targetConfigId string, workspaces []Workspace) *Target {
	this := Target{}
	this.Default = default_
	this.EnvVars = envVars
	this.Id = id
	this.Name = name
	this.TargetConfig = targetConfig
	this.TargetConfigId = targetConfigId
	this.Workspaces = workspaces
	return &this
}

// NewTargetWithDefaults instantiates a new Target object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetWithDefaults() *Target {
	this := Target{}
	return &this
}

// GetDefault returns the Default field value
func (o *Target) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *Target) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *Target) SetDefault(v bool) {
	o.Default = v
}

// GetEnvVars returns the EnvVars field value
func (o *Target) GetEnvVars() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.EnvVars
}

// GetEnvVarsOk returns a tuple with the EnvVars field value
// and a boolean to check if the value has been set.
func (o *Target) GetEnvVarsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvVars, true
}

// SetEnvVars sets field value
func (o *Target) SetEnvVars(v map[string]string) {
	o.EnvVars = v
}

// GetId returns the Id field value
func (o *Target) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Target) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Target) SetId(v string) {
	o.Id = v
}

// GetLastJob returns the LastJob field value if set, zero value otherwise.
func (o *Target) GetLastJob() Job {
	if o == nil || IsNil(o.LastJob) {
		var ret Job
		return ret
	}
	return *o.LastJob
}

// GetLastJobOk returns a tuple with the LastJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetLastJobOk() (*Job, bool) {
	if o == nil || IsNil(o.LastJob) {
		return nil, false
	}
	return o.LastJob, true
}

// HasLastJob returns a boolean if a field has been set.
func (o *Target) HasLastJob() bool {
	if o != nil && !IsNil(o.LastJob) {
		return true
	}

	return false
}

// SetLastJob gets a reference to the given Job and assigns it to the LastJob field.
func (o *Target) SetLastJob(v Job) {
	o.LastJob = &v
}

// GetLastJobId returns the LastJobId field value if set, zero value otherwise.
func (o *Target) GetLastJobId() string {
	if o == nil || IsNil(o.LastJobId) {
		var ret string
		return ret
	}
	return *o.LastJobId
}

// GetLastJobIdOk returns a tuple with the LastJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetLastJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastJobId) {
		return nil, false
	}
	return o.LastJobId, true
}

// HasLastJobId returns a boolean if a field has been set.
func (o *Target) HasLastJobId() bool {
	if o != nil && !IsNil(o.LastJobId) {
		return true
	}

	return false
}

// SetLastJobId gets a reference to the given string and assigns it to the LastJobId field.
func (o *Target) SetLastJobId(v string) {
	o.LastJobId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Target) GetMetadata() TargetMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret TargetMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetMetadataOk() (*TargetMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Target) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given TargetMetadata and assigns it to the Metadata field.
func (o *Target) SetMetadata(v TargetMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *Target) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Target) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Target) SetName(v string) {
	o.Name = v
}

// GetProviderMetadata returns the ProviderMetadata field value if set, zero value otherwise.
func (o *Target) GetProviderMetadata() string {
	if o == nil || IsNil(o.ProviderMetadata) {
		var ret string
		return ret
	}
	return *o.ProviderMetadata
}

// GetProviderMetadataOk returns a tuple with the ProviderMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Target) GetProviderMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderMetadata) {
		return nil, false
	}
	return o.ProviderMetadata, true
}

// HasProviderMetadata returns a boolean if a field has been set.
func (o *Target) HasProviderMetadata() bool {
	if o != nil && !IsNil(o.ProviderMetadata) {
		return true
	}

	return false
}

// SetProviderMetadata gets a reference to the given string and assigns it to the ProviderMetadata field.
func (o *Target) SetProviderMetadata(v string) {
	o.ProviderMetadata = &v
}

// GetTargetConfig returns the TargetConfig field value
func (o *Target) GetTargetConfig() TargetConfig {
	if o == nil {
		var ret TargetConfig
		return ret
	}

	return o.TargetConfig
}

// GetTargetConfigOk returns a tuple with the TargetConfig field value
// and a boolean to check if the value has been set.
func (o *Target) GetTargetConfigOk() (*TargetConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetConfig, true
}

// SetTargetConfig sets field value
func (o *Target) SetTargetConfig(v TargetConfig) {
	o.TargetConfig = v
}

// GetTargetConfigId returns the TargetConfigId field value
func (o *Target) GetTargetConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetConfigId
}

// GetTargetConfigIdOk returns a tuple with the TargetConfigId field value
// and a boolean to check if the value has been set.
func (o *Target) GetTargetConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetConfigId, true
}

// SetTargetConfigId sets field value
func (o *Target) SetTargetConfigId(v string) {
	o.TargetConfigId = v
}

// GetWorkspaces returns the Workspaces field value
func (o *Target) GetWorkspaces() []Workspace {
	if o == nil {
		var ret []Workspace
		return ret
	}

	return o.Workspaces
}

// GetWorkspacesOk returns a tuple with the Workspaces field value
// and a boolean to check if the value has been set.
func (o *Target) GetWorkspacesOk() ([]Workspace, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workspaces, true
}

// SetWorkspaces sets field value
func (o *Target) SetWorkspaces(v []Workspace) {
	o.Workspaces = v
}

func (o Target) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Target) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["default"] = o.Default
	toSerialize["envVars"] = o.EnvVars
	toSerialize["id"] = o.Id
	if !IsNil(o.LastJob) {
		toSerialize["lastJob"] = o.LastJob
	}
	if !IsNil(o.LastJobId) {
		toSerialize["lastJobId"] = o.LastJobId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ProviderMetadata) {
		toSerialize["providerMetadata"] = o.ProviderMetadata
	}
	toSerialize["targetConfig"] = o.TargetConfig
	toSerialize["targetConfigId"] = o.TargetConfigId
	toSerialize["workspaces"] = o.Workspaces
	return toSerialize, nil
}

func (o *Target) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"default",
		"envVars",
		"id",
		"name",
		"targetConfig",
		"targetConfigId",
		"workspaces",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTarget := _Target{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTarget)

	if err != nil {
		return err
	}

	*o = Target(varTarget)

	return err
}

type NullableTarget struct {
	value *Target
	isSet bool
}

func (v NullableTarget) Get() *Target {
	return v.value
}

func (v *NullableTarget) Set(val *Target) {
	v.value = val
	v.isSet = true
}

func (v NullableTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTarget(val *Target) *NullableTarget {
	return &NullableTarget{value: val, isSet: true}
}

func (v NullableTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
