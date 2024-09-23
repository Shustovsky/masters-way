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

// checks if the SchemasAIEstimateIssuePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasAIEstimateIssuePayload{}

// SchemasAIEstimateIssuePayload struct for SchemasAIEstimateIssuePayload
type SchemasAIEstimateIssuePayload struct {
	Goal string `json:"goal"`
	Issue string `json:"issue"`
}

type _SchemasAIEstimateIssuePayload SchemasAIEstimateIssuePayload

// NewSchemasAIEstimateIssuePayload instantiates a new SchemasAIEstimateIssuePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasAIEstimateIssuePayload(goal string, issue string) *SchemasAIEstimateIssuePayload {
	this := SchemasAIEstimateIssuePayload{}
	this.Goal = goal
	this.Issue = issue
	return &this
}

// NewSchemasAIEstimateIssuePayloadWithDefaults instantiates a new SchemasAIEstimateIssuePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasAIEstimateIssuePayloadWithDefaults() *SchemasAIEstimateIssuePayload {
	this := SchemasAIEstimateIssuePayload{}
	return &this
}

// GetGoal returns the Goal field value
func (o *SchemasAIEstimateIssuePayload) GetGoal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Goal
}

// GetGoalOk returns a tuple with the Goal field value
// and a boolean to check if the value has been set.
func (o *SchemasAIEstimateIssuePayload) GetGoalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Goal, true
}

// SetGoal sets field value
func (o *SchemasAIEstimateIssuePayload) SetGoal(v string) {
	o.Goal = v
}

// GetIssue returns the Issue field value
func (o *SchemasAIEstimateIssuePayload) GetIssue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issue
}

// GetIssueOk returns a tuple with the Issue field value
// and a boolean to check if the value has been set.
func (o *SchemasAIEstimateIssuePayload) GetIssueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issue, true
}

// SetIssue sets field value
func (o *SchemasAIEstimateIssuePayload) SetIssue(v string) {
	o.Issue = v
}

func (o SchemasAIEstimateIssuePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasAIEstimateIssuePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["goal"] = o.Goal
	toSerialize["issue"] = o.Issue
	return toSerialize, nil
}

func (o *SchemasAIEstimateIssuePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"goal",
		"issue",
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

	varSchemasAIEstimateIssuePayload := _SchemasAIEstimateIssuePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasAIEstimateIssuePayload)

	if err != nil {
		return err
	}

	*o = SchemasAIEstimateIssuePayload(varSchemasAIEstimateIssuePayload)

	return err
}

type NullableSchemasAIEstimateIssuePayload struct {
	value *SchemasAIEstimateIssuePayload
	isSet bool
}

func (v NullableSchemasAIEstimateIssuePayload) Get() *SchemasAIEstimateIssuePayload {
	return v.value
}

func (v *NullableSchemasAIEstimateIssuePayload) Set(val *SchemasAIEstimateIssuePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasAIEstimateIssuePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasAIEstimateIssuePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasAIEstimateIssuePayload(val *SchemasAIEstimateIssuePayload) *NullableSchemasAIEstimateIssuePayload {
	return &NullableSchemasAIEstimateIssuePayload{value: val, isSet: true}
}

func (v NullableSchemasAIEstimateIssuePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasAIEstimateIssuePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

