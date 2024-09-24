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

// checks if the SchemasCreateWayTagPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasCreateWayTagPayload{}

// SchemasCreateWayTagPayload struct for SchemasCreateWayTagPayload
type SchemasCreateWayTagPayload struct {
	Name string
	WayUuid string
}

type _SchemasCreateWayTagPayload SchemasCreateWayTagPayload

// NewSchemasCreateWayTagPayload instantiates a new SchemasCreateWayTagPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasCreateWayTagPayload(name string, wayUuid string) *SchemasCreateWayTagPayload {
	this := SchemasCreateWayTagPayload{}
	this.Name = name
	this.WayUuid = wayUuid
	return &this
}

// NewSchemasCreateWayTagPayloadWithDefaults instantiates a new SchemasCreateWayTagPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasCreateWayTagPayloadWithDefaults() *SchemasCreateWayTagPayload {
	this := SchemasCreateWayTagPayload{}
	return &this
}

// GetName returns the Name field value
func (o *SchemasCreateWayTagPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateWayTagPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasCreateWayTagPayload) SetName(v string) {
	o.Name = v
}

// GetWayUuid returns the WayUuid field value
func (o *SchemasCreateWayTagPayload) GetWayUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WayUuid
}

// GetWayUuidOk returns a tuple with the WayUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateWayTagPayload) GetWayUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WayUuid, true
}

// SetWayUuid sets field value
func (o *SchemasCreateWayTagPayload) SetWayUuid(v string) {
	o.WayUuid = v
}

func (o SchemasCreateWayTagPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasCreateWayTagPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["wayUuid"] = o.WayUuid
	return toSerialize, nil
}

func (o *SchemasCreateWayTagPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"wayUuid",
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

	varSchemasCreateWayTagPayload := _SchemasCreateWayTagPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasCreateWayTagPayload)

	if err != nil {
		return err
	}

	*o = SchemasCreateWayTagPayload(varSchemasCreateWayTagPayload)

	return err
}

type NullableSchemasCreateWayTagPayload struct {
	value *SchemasCreateWayTagPayload
	isSet bool
}

func (v NullableSchemasCreateWayTagPayload) Get() *SchemasCreateWayTagPayload {
	return v.value
}

func (v *NullableSchemasCreateWayTagPayload) Set(val *SchemasCreateWayTagPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasCreateWayTagPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasCreateWayTagPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasCreateWayTagPayload(val *SchemasCreateWayTagPayload) *NullableSchemasCreateWayTagPayload {
	return &NullableSchemasCreateWayTagPayload{value: val, isSet: true}
}

func (v NullableSchemasCreateWayTagPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasCreateWayTagPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


