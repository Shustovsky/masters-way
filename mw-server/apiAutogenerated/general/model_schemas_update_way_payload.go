/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SchemasUpdateWayPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUpdateWayPayload{}

// SchemasUpdateWayPayload struct for SchemasUpdateWayPayload
type SchemasUpdateWayPayload struct {
	EstimationTime *int32 `json:"estimationTime,omitempty"`
	GoalDescription *string `json:"goalDescription,omitempty"`
	IsCompleted *bool `json:"isCompleted,omitempty"`
	IsPrivate *bool `json:"isPrivate,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewSchemasUpdateWayPayload instantiates a new SchemasUpdateWayPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUpdateWayPayload() *SchemasUpdateWayPayload {
	this := SchemasUpdateWayPayload{}
	return &this
}

// NewSchemasUpdateWayPayloadWithDefaults instantiates a new SchemasUpdateWayPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUpdateWayPayloadWithDefaults() *SchemasUpdateWayPayload {
	this := SchemasUpdateWayPayload{}
	return &this
}

// GetEstimationTime returns the EstimationTime field value if set, zero value otherwise.
func (o *SchemasUpdateWayPayload) GetEstimationTime() int32 {
	if o == nil || IsNil(o.EstimationTime) {
		var ret int32
		return ret
	}
	return *o.EstimationTime
}

// GetEstimationTimeOk returns a tuple with the EstimationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateWayPayload) GetEstimationTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimationTime) {
		return nil, false
	}
	return o.EstimationTime, true
}

// HasEstimationTime returns a boolean if a field has been set.
func (o *SchemasUpdateWayPayload) HasEstimationTime() bool {
	if o != nil && !IsNil(o.EstimationTime) {
		return true
	}

	return false
}

// SetEstimationTime gets a reference to the given int32 and assigns it to the EstimationTime field.
func (o *SchemasUpdateWayPayload) SetEstimationTime(v int32) {
	o.EstimationTime = &v
}

// GetGoalDescription returns the GoalDescription field value if set, zero value otherwise.
func (o *SchemasUpdateWayPayload) GetGoalDescription() string {
	if o == nil || IsNil(o.GoalDescription) {
		var ret string
		return ret
	}
	return *o.GoalDescription
}

// GetGoalDescriptionOk returns a tuple with the GoalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateWayPayload) GetGoalDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.GoalDescription) {
		return nil, false
	}
	return o.GoalDescription, true
}

// HasGoalDescription returns a boolean if a field has been set.
func (o *SchemasUpdateWayPayload) HasGoalDescription() bool {
	if o != nil && !IsNil(o.GoalDescription) {
		return true
	}

	return false
}

// SetGoalDescription gets a reference to the given string and assigns it to the GoalDescription field.
func (o *SchemasUpdateWayPayload) SetGoalDescription(v string) {
	o.GoalDescription = &v
}

// GetIsCompleted returns the IsCompleted field value if set, zero value otherwise.
func (o *SchemasUpdateWayPayload) GetIsCompleted() bool {
	if o == nil || IsNil(o.IsCompleted) {
		var ret bool
		return ret
	}
	return *o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateWayPayload) GetIsCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompleted) {
		return nil, false
	}
	return o.IsCompleted, true
}

// HasIsCompleted returns a boolean if a field has been set.
func (o *SchemasUpdateWayPayload) HasIsCompleted() bool {
	if o != nil && !IsNil(o.IsCompleted) {
		return true
	}

	return false
}

// SetIsCompleted gets a reference to the given bool and assigns it to the IsCompleted field.
func (o *SchemasUpdateWayPayload) SetIsCompleted(v bool) {
	o.IsCompleted = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *SchemasUpdateWayPayload) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate) {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateWayPayload) GetIsPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivate) {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *SchemasUpdateWayPayload) HasIsPrivate() bool {
	if o != nil && !IsNil(o.IsPrivate) {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *SchemasUpdateWayPayload) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchemasUpdateWayPayload) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateWayPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchemasUpdateWayPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchemasUpdateWayPayload) SetName(v string) {
	o.Name = &v
}

func (o SchemasUpdateWayPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUpdateWayPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EstimationTime) {
		toSerialize["estimationTime"] = o.EstimationTime
	}
	if !IsNil(o.GoalDescription) {
		toSerialize["goalDescription"] = o.GoalDescription
	}
	if !IsNil(o.IsCompleted) {
		toSerialize["isCompleted"] = o.IsCompleted
	}
	if !IsNil(o.IsPrivate) {
		toSerialize["isPrivate"] = o.IsPrivate
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSchemasUpdateWayPayload struct {
	value *SchemasUpdateWayPayload
	isSet bool
}

func (v NullableSchemasUpdateWayPayload) Get() *SchemasUpdateWayPayload {
	return v.value
}

func (v *NullableSchemasUpdateWayPayload) Set(val *SchemasUpdateWayPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUpdateWayPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUpdateWayPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUpdateWayPayload(val *SchemasUpdateWayPayload) *NullableSchemasUpdateWayPayload {
	return &NullableSchemasUpdateWayPayload{value: val, isSet: true}
}

func (v NullableSchemasUpdateWayPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUpdateWayPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

