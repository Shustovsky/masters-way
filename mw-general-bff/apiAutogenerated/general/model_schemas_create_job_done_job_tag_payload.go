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

// checks if the SchemasCreateJobDoneJobTagPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasCreateJobDoneJobTagPayload{}

// SchemasCreateJobDoneJobTagPayload struct for SchemasCreateJobDoneJobTagPayload
type SchemasCreateJobDoneJobTagPayload struct {
	JobDoneUuid string
	JobTagUuid string
}

type _SchemasCreateJobDoneJobTagPayload SchemasCreateJobDoneJobTagPayload

// NewSchemasCreateJobDoneJobTagPayload instantiates a new SchemasCreateJobDoneJobTagPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasCreateJobDoneJobTagPayload(jobDoneUuid string, jobTagUuid string) *SchemasCreateJobDoneJobTagPayload {
	this := SchemasCreateJobDoneJobTagPayload{}
	this.JobDoneUuid = jobDoneUuid
	this.JobTagUuid = jobTagUuid
	return &this
}

// NewSchemasCreateJobDoneJobTagPayloadWithDefaults instantiates a new SchemasCreateJobDoneJobTagPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasCreateJobDoneJobTagPayloadWithDefaults() *SchemasCreateJobDoneJobTagPayload {
	this := SchemasCreateJobDoneJobTagPayload{}
	return &this
}

// GetJobDoneUuid returns the JobDoneUuid field value
func (o *SchemasCreateJobDoneJobTagPayload) GetJobDoneUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobDoneUuid
}

// GetJobDoneUuidOk returns a tuple with the JobDoneUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobDoneJobTagPayload) GetJobDoneUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobDoneUuid, true
}

// SetJobDoneUuid sets field value
func (o *SchemasCreateJobDoneJobTagPayload) SetJobDoneUuid(v string) {
	o.JobDoneUuid = v
}

// GetJobTagUuid returns the JobTagUuid field value
func (o *SchemasCreateJobDoneJobTagPayload) GetJobTagUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobTagUuid
}

// GetJobTagUuidOk returns a tuple with the JobTagUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobDoneJobTagPayload) GetJobTagUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobTagUuid, true
}

// SetJobTagUuid sets field value
func (o *SchemasCreateJobDoneJobTagPayload) SetJobTagUuid(v string) {
	o.JobTagUuid = v
}

func (o SchemasCreateJobDoneJobTagPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasCreateJobDoneJobTagPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jobDoneUuid"] = o.JobDoneUuid
	toSerialize["jobTagUuid"] = o.JobTagUuid
	return toSerialize, nil
}

func (o *SchemasCreateJobDoneJobTagPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jobDoneUuid",
		"jobTagUuid",
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

	varSchemasCreateJobDoneJobTagPayload := _SchemasCreateJobDoneJobTagPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasCreateJobDoneJobTagPayload)

	if err != nil {
		return err
	}

	*o = SchemasCreateJobDoneJobTagPayload(varSchemasCreateJobDoneJobTagPayload)

	return err
}

type NullableSchemasCreateJobDoneJobTagPayload struct {
	value *SchemasCreateJobDoneJobTagPayload
	isSet bool
}

func (v NullableSchemasCreateJobDoneJobTagPayload) Get() *SchemasCreateJobDoneJobTagPayload {
	return v.value
}

func (v *NullableSchemasCreateJobDoneJobTagPayload) Set(val *SchemasCreateJobDoneJobTagPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasCreateJobDoneJobTagPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasCreateJobDoneJobTagPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasCreateJobDoneJobTagPayload(val *SchemasCreateJobDoneJobTagPayload) *NullableSchemasCreateJobDoneJobTagPayload {
	return &NullableSchemasCreateJobDoneJobTagPayload{value: val, isSet: true}
}

func (v NullableSchemasCreateJobDoneJobTagPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasCreateJobDoneJobTagPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


