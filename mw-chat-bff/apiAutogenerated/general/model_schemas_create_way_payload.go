/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SchemasCreateWayPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasCreateWayPayload{}

// SchemasCreateWayPayload struct for SchemasCreateWayPayload
type SchemasCreateWayPayload struct {
	CopiedFromWayUuid NullableString
	EstimationTime int32
	GoalDescription string
	IsCompleted bool
	IsPrivate bool
	Name string
	OwnerUuid string
}

type _SchemasCreateWayPayload SchemasCreateWayPayload

// NewSchemasCreateWayPayload instantiates a new SchemasCreateWayPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasCreateWayPayload(copiedFromWayUuid NullableString, estimationTime int32, goalDescription string, isCompleted bool, isPrivate bool, name string, ownerUuid string) *SchemasCreateWayPayload {
	this := SchemasCreateWayPayload{}
	this.CopiedFromWayUuid = copiedFromWayUuid
	this.EstimationTime = estimationTime
	this.GoalDescription = goalDescription
	this.IsCompleted = isCompleted
	this.IsPrivate = isPrivate
	this.Name = name
	this.OwnerUuid = ownerUuid
	return &this
}

// NewSchemasCreateWayPayloadWithDefaults instantiates a new SchemasCreateWayPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasCreateWayPayloadWithDefaults() *SchemasCreateWayPayload {
	this := SchemasCreateWayPayload{}
	return &this
}

// GetCopiedFromWayUuid returns the CopiedFromWayUuid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SchemasCreateWayPayload) GetCopiedFromWayUuid() string {
	if o == nil || o.CopiedFromWayUuid.Get() == nil {
		var ret string
		return ret
	}

	return *o.CopiedFromWayUuid.Get()
}

// GetCopiedFromWayUuidOk returns a tuple with the CopiedFromWayUuid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasCreateWayPayload) GetCopiedFromWayUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CopiedFromWayUuid.Get(), o.CopiedFromWayUuid.IsSet()
}

// SetCopiedFromWayUuid sets field value
func (o *SchemasCreateWayPayload) SetCopiedFromWayUuid(v string) {
	o.CopiedFromWayUuid.Set(&v)
}

// GetEstimationTime returns the EstimationTime field value
func (o *SchemasCreateWayPayload) GetEstimationTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EstimationTime
}

// GetEstimationTimeOk returns a tuple with the EstimationTime field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateWayPayload) GetEstimationTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimationTime, true
}

// SetEstimationTime sets field value
func (o *SchemasCreateWayPayload) SetEstimationTime(v int32) {
	o.EstimationTime = v
}

// GetGoalDescription returns the GoalDescription field value
func (o *SchemasCreateWayPayload) GetGoalDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoalDescription
}

// GetGoalDescriptionOk returns a tuple with the GoalDescription field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateWayPayload) GetGoalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoalDescription, true
}

// SetGoalDescription sets field value
func (o *SchemasCreateWayPayload) SetGoalDescription(v string) {
	o.GoalDescription = v
}

// GetIsCompleted returns the IsCompleted field value
func (o *SchemasCreateWayPayload) GetIsCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateWayPayload) GetIsCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCompleted, true
}

// SetIsCompleted sets field value
func (o *SchemasCreateWayPayload) SetIsCompleted(v bool) {
	o.IsCompleted = v
}

// GetIsPrivate returns the IsPrivate field value
func (o *SchemasCreateWayPayload) GetIsPrivate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateWayPayload) GetIsPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrivate, true
}

// SetIsPrivate sets field value
func (o *SchemasCreateWayPayload) SetIsPrivate(v bool) {
	o.IsPrivate = v
}

// GetName returns the Name field value
func (o *SchemasCreateWayPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateWayPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasCreateWayPayload) SetName(v string) {
	o.Name = v
}

// GetOwnerUuid returns the OwnerUuid field value
func (o *SchemasCreateWayPayload) GetOwnerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerUuid
}

// GetOwnerUuidOk returns a tuple with the OwnerUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateWayPayload) GetOwnerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerUuid, true
}

// SetOwnerUuid sets field value
func (o *SchemasCreateWayPayload) SetOwnerUuid(v string) {
	o.OwnerUuid = v
}

func (o SchemasCreateWayPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasCreateWayPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["copiedFromWayUuid"] = o.CopiedFromWayUuid.Get()
	toSerialize["estimationTime"] = o.EstimationTime
	toSerialize["goalDescription"] = o.GoalDescription
	toSerialize["isCompleted"] = o.IsCompleted
	toSerialize["isPrivate"] = o.IsPrivate
	toSerialize["name"] = o.Name
	toSerialize["ownerUuid"] = o.OwnerUuid
	return toSerialize, nil
}

func (o *SchemasCreateWayPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"copiedFromWayUuid",
		"estimationTime",
		"goalDescription",
		"isCompleted",
		"isPrivate",
		"name",
		"ownerUuid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSchemasCreateWayPayload := _SchemasCreateWayPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasCreateWayPayload)

	if err != nil {
		return err
	}

	*o = SchemasCreateWayPayload(varSchemasCreateWayPayload)

	return err
}

type NullableSchemasCreateWayPayload struct {
	value *SchemasCreateWayPayload
	isSet bool
}

func (v NullableSchemasCreateWayPayload) Get() *SchemasCreateWayPayload {
	return v.value
}

func (v *NullableSchemasCreateWayPayload) Set(val *SchemasCreateWayPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasCreateWayPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasCreateWayPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasCreateWayPayload(val *SchemasCreateWayPayload) *NullableSchemasCreateWayPayload {
	return &NullableSchemasCreateWayPayload{value: val, isSet: true}
}

func (v NullableSchemasCreateWayPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasCreateWayPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


