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

// checks if the SchemasAIDecomposeIssuePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasAIDecomposeIssuePayload{}

// SchemasAIDecomposeIssuePayload struct for SchemasAIDecomposeIssuePayload
type SchemasAIDecomposeIssuePayload struct {
	Goal string `json:"goal"`
	Message string `json:"message"`
}

type _SchemasAIDecomposeIssuePayload SchemasAIDecomposeIssuePayload

// NewSchemasAIDecomposeIssuePayload instantiates a new SchemasAIDecomposeIssuePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasAIDecomposeIssuePayload(goal string, message string) *SchemasAIDecomposeIssuePayload {
	this := SchemasAIDecomposeIssuePayload{}
	this.Goal = goal
	this.Message = message
	return &this
}

// NewSchemasAIDecomposeIssuePayloadWithDefaults instantiates a new SchemasAIDecomposeIssuePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasAIDecomposeIssuePayloadWithDefaults() *SchemasAIDecomposeIssuePayload {
	this := SchemasAIDecomposeIssuePayload{}
	return &this
}

// GetGoal returns the Goal field value
func (o *SchemasAIDecomposeIssuePayload) GetGoal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Goal
}

// GetGoalOk returns a tuple with the Goal field value
// and a boolean to check if the value has been set.
func (o *SchemasAIDecomposeIssuePayload) GetGoalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Goal, true
}

// SetGoal sets field value
func (o *SchemasAIDecomposeIssuePayload) SetGoal(v string) {
	o.Goal = v
}

// GetMessage returns the Message field value
func (o *SchemasAIDecomposeIssuePayload) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SchemasAIDecomposeIssuePayload) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SchemasAIDecomposeIssuePayload) SetMessage(v string) {
	o.Message = v
}

func (o SchemasAIDecomposeIssuePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasAIDecomposeIssuePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["goal"] = o.Goal
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *SchemasAIDecomposeIssuePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"goal",
		"message",
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

	varSchemasAIDecomposeIssuePayload := _SchemasAIDecomposeIssuePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasAIDecomposeIssuePayload)

	if err != nil {
		return err
	}

	*o = SchemasAIDecomposeIssuePayload(varSchemasAIDecomposeIssuePayload)

	return err
}

type NullableSchemasAIDecomposeIssuePayload struct {
	value *SchemasAIDecomposeIssuePayload
	isSet bool
}

func (v NullableSchemasAIDecomposeIssuePayload) Get() *SchemasAIDecomposeIssuePayload {
	return v.value
}

func (v *NullableSchemasAIDecomposeIssuePayload) Set(val *SchemasAIDecomposeIssuePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasAIDecomposeIssuePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasAIDecomposeIssuePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasAIDecomposeIssuePayload(val *SchemasAIDecomposeIssuePayload) *NullableSchemasAIDecomposeIssuePayload {
	return &NullableSchemasAIDecomposeIssuePayload{value: val, isSet: true}
}

func (v NullableSchemasAIDecomposeIssuePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasAIDecomposeIssuePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


