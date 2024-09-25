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

// checks if the SchemasUpdateCommentPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUpdateCommentPayload{}

// SchemasUpdateCommentPayload struct for SchemasUpdateCommentPayload
type SchemasUpdateCommentPayload struct {
	Description *string
}

// NewSchemasUpdateCommentPayload instantiates a new SchemasUpdateCommentPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUpdateCommentPayload() *SchemasUpdateCommentPayload {
	this := SchemasUpdateCommentPayload{}
	return &this
}

// NewSchemasUpdateCommentPayloadWithDefaults instantiates a new SchemasUpdateCommentPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUpdateCommentPayloadWithDefaults() *SchemasUpdateCommentPayload {
	this := SchemasUpdateCommentPayload{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemasUpdateCommentPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateCommentPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemasUpdateCommentPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemasUpdateCommentPayload) SetDescription(v string) {
	o.Description = &v
}

func (o SchemasUpdateCommentPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUpdateCommentPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableSchemasUpdateCommentPayload struct {
	value *SchemasUpdateCommentPayload
	isSet bool
}

func (v NullableSchemasUpdateCommentPayload) Get() *SchemasUpdateCommentPayload {
	return v.value
}

func (v *NullableSchemasUpdateCommentPayload) Set(val *SchemasUpdateCommentPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUpdateCommentPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUpdateCommentPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUpdateCommentPayload(val *SchemasUpdateCommentPayload) *NullableSchemasUpdateCommentPayload {
	return &NullableSchemasUpdateCommentPayload{value: val, isSet: true}
}

func (v NullableSchemasUpdateCommentPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUpdateCommentPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


