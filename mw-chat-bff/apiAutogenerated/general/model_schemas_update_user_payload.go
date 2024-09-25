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

// checks if the SchemasUpdateUserPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUpdateUserPayload{}

// SchemasUpdateUserPayload struct for SchemasUpdateUserPayload
type SchemasUpdateUserPayload struct {
	Description *string
	ImageUrl *string
	IsMentor *bool
	Name *string
}

// NewSchemasUpdateUserPayload instantiates a new SchemasUpdateUserPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUpdateUserPayload() *SchemasUpdateUserPayload {
	this := SchemasUpdateUserPayload{}
	return &this
}

// NewSchemasUpdateUserPayloadWithDefaults instantiates a new SchemasUpdateUserPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUpdateUserPayloadWithDefaults() *SchemasUpdateUserPayload {
	this := SchemasUpdateUserPayload{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemasUpdateUserPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateUserPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemasUpdateUserPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemasUpdateUserPayload) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *SchemasUpdateUserPayload) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateUserPayload) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *SchemasUpdateUserPayload) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *SchemasUpdateUserPayload) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetIsMentor returns the IsMentor field value if set, zero value otherwise.
func (o *SchemasUpdateUserPayload) GetIsMentor() bool {
	if o == nil || IsNil(o.IsMentor) {
		var ret bool
		return ret
	}
	return *o.IsMentor
}

// GetIsMentorOk returns a tuple with the IsMentor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateUserPayload) GetIsMentorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMentor) {
		return nil, false
	}
	return o.IsMentor, true
}

// HasIsMentor returns a boolean if a field has been set.
func (o *SchemasUpdateUserPayload) HasIsMentor() bool {
	if o != nil && !IsNil(o.IsMentor) {
		return true
	}

	return false
}

// SetIsMentor gets a reference to the given bool and assigns it to the IsMentor field.
func (o *SchemasUpdateUserPayload) SetIsMentor(v bool) {
	o.IsMentor = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchemasUpdateUserPayload) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateUserPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchemasUpdateUserPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchemasUpdateUserPayload) SetName(v string) {
	o.Name = &v
}

func (o SchemasUpdateUserPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUpdateUserPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.IsMentor) {
		toSerialize["isMentor"] = o.IsMentor
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSchemasUpdateUserPayload struct {
	value *SchemasUpdateUserPayload
	isSet bool
}

func (v NullableSchemasUpdateUserPayload) Get() *SchemasUpdateUserPayload {
	return v.value
}

func (v *NullableSchemasUpdateUserPayload) Set(val *SchemasUpdateUserPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUpdateUserPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUpdateUserPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUpdateUserPayload(val *SchemasUpdateUserPayload) *NullableSchemasUpdateUserPayload {
	return &NullableSchemasUpdateUserPayload{value: val, isSet: true}
}

func (v NullableSchemasUpdateUserPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUpdateUserPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


