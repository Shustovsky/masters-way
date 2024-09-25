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

// checks if the SchemasAICommentIssueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasAICommentIssueResponse{}

// SchemasAICommentIssueResponse struct for SchemasAICommentIssueResponse
type SchemasAICommentIssueResponse struct {
	Goal string
}

type _SchemasAICommentIssueResponse SchemasAICommentIssueResponse

// NewSchemasAICommentIssueResponse instantiates a new SchemasAICommentIssueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasAICommentIssueResponse(goal string) *SchemasAICommentIssueResponse {
	this := SchemasAICommentIssueResponse{}
	this.Goal = goal
	return &this
}

// NewSchemasAICommentIssueResponseWithDefaults instantiates a new SchemasAICommentIssueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasAICommentIssueResponseWithDefaults() *SchemasAICommentIssueResponse {
	this := SchemasAICommentIssueResponse{}
	return &this
}

// GetGoal returns the Goal field value
func (o *SchemasAICommentIssueResponse) GetGoal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Goal
}

// GetGoalOk returns a tuple with the Goal field value
// and a boolean to check if the value has been set.
func (o *SchemasAICommentIssueResponse) GetGoalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Goal, true
}

// SetGoal sets field value
func (o *SchemasAICommentIssueResponse) SetGoal(v string) {
	o.Goal = v
}

func (o SchemasAICommentIssueResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasAICommentIssueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["goal"] = o.Goal
	return toSerialize, nil
}

func (o *SchemasAICommentIssueResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"goal",
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

	varSchemasAICommentIssueResponse := _SchemasAICommentIssueResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasAICommentIssueResponse)

	if err != nil {
		return err
	}

	*o = SchemasAICommentIssueResponse(varSchemasAICommentIssueResponse)

	return err
}

type NullableSchemasAICommentIssueResponse struct {
	value *SchemasAICommentIssueResponse
	isSet bool
}

func (v NullableSchemasAICommentIssueResponse) Get() *SchemasAICommentIssueResponse {
	return v.value
}

func (v *NullableSchemasAICommentIssueResponse) Set(val *SchemasAICommentIssueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasAICommentIssueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasAICommentIssueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasAICommentIssueResponse(val *SchemasAICommentIssueResponse) *NullableSchemasAICommentIssueResponse {
	return &NullableSchemasAICommentIssueResponse{value: val, isSet: true}
}

func (v NullableSchemasAICommentIssueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasAICommentIssueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


