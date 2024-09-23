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

// checks if the SchemasAIDecomposeIssueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasAIDecomposeIssueResponse{}

// SchemasAIDecomposeIssueResponse struct for SchemasAIDecomposeIssueResponse
type SchemasAIDecomposeIssueResponse struct {
	Plans []string
}

type _SchemasAIDecomposeIssueResponse SchemasAIDecomposeIssueResponse

// NewSchemasAIDecomposeIssueResponse instantiates a new SchemasAIDecomposeIssueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasAIDecomposeIssueResponse(plans []string) *SchemasAIDecomposeIssueResponse {
	this := SchemasAIDecomposeIssueResponse{}
	this.Plans = plans
	return &this
}

// NewSchemasAIDecomposeIssueResponseWithDefaults instantiates a new SchemasAIDecomposeIssueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasAIDecomposeIssueResponseWithDefaults() *SchemasAIDecomposeIssueResponse {
	this := SchemasAIDecomposeIssueResponse{}
	return &this
}

// GetPlans returns the Plans field value
func (o *SchemasAIDecomposeIssueResponse) GetPlans() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value
// and a boolean to check if the value has been set.
func (o *SchemasAIDecomposeIssueResponse) GetPlansOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plans, true
}

// SetPlans sets field value
func (o *SchemasAIDecomposeIssueResponse) SetPlans(v []string) {
	o.Plans = v
}

func (o SchemasAIDecomposeIssueResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasAIDecomposeIssueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plans"] = o.Plans
	return toSerialize, nil
}

func (o *SchemasAIDecomposeIssueResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plans",
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

	varSchemasAIDecomposeIssueResponse := _SchemasAIDecomposeIssueResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasAIDecomposeIssueResponse)

	if err != nil {
		return err
	}

	*o = SchemasAIDecomposeIssueResponse(varSchemasAIDecomposeIssueResponse)

	return err
}

type NullableSchemasAIDecomposeIssueResponse struct {
	value *SchemasAIDecomposeIssueResponse
	isSet bool
}

func (v NullableSchemasAIDecomposeIssueResponse) Get() *SchemasAIDecomposeIssueResponse {
	return v.value
}

func (v *NullableSchemasAIDecomposeIssueResponse) Set(val *SchemasAIDecomposeIssueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasAIDecomposeIssueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasAIDecomposeIssueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasAIDecomposeIssueResponse(val *SchemasAIDecomposeIssueResponse) *NullableSchemasAIDecomposeIssueResponse {
	return &NullableSchemasAIDecomposeIssueResponse{value: val, isSet: true}
}

func (v NullableSchemasAIDecomposeIssueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasAIDecomposeIssueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

