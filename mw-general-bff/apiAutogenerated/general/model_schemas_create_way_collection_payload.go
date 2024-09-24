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

// checks if the SchemasCreateWayCollectionPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasCreateWayCollectionPayload{}

// SchemasCreateWayCollectionPayload struct for SchemasCreateWayCollectionPayload
type SchemasCreateWayCollectionPayload struct {
	Name string
	OwnerUuid string
}

type _SchemasCreateWayCollectionPayload SchemasCreateWayCollectionPayload

// NewSchemasCreateWayCollectionPayload instantiates a new SchemasCreateWayCollectionPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasCreateWayCollectionPayload(name string, ownerUuid string) *SchemasCreateWayCollectionPayload {
	this := SchemasCreateWayCollectionPayload{}
	this.Name = name
	this.OwnerUuid = ownerUuid
	return &this
}

// NewSchemasCreateWayCollectionPayloadWithDefaults instantiates a new SchemasCreateWayCollectionPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasCreateWayCollectionPayloadWithDefaults() *SchemasCreateWayCollectionPayload {
	this := SchemasCreateWayCollectionPayload{}
	return &this
}

// GetName returns the Name field value
func (o *SchemasCreateWayCollectionPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateWayCollectionPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasCreateWayCollectionPayload) SetName(v string) {
	o.Name = v
}

// GetOwnerUuid returns the OwnerUuid field value
func (o *SchemasCreateWayCollectionPayload) GetOwnerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerUuid
}

// GetOwnerUuidOk returns a tuple with the OwnerUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateWayCollectionPayload) GetOwnerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerUuid, true
}

// SetOwnerUuid sets field value
func (o *SchemasCreateWayCollectionPayload) SetOwnerUuid(v string) {
	o.OwnerUuid = v
}

func (o SchemasCreateWayCollectionPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasCreateWayCollectionPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["ownerUuid"] = o.OwnerUuid
	return toSerialize, nil
}

func (o *SchemasCreateWayCollectionPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varSchemasCreateWayCollectionPayload := _SchemasCreateWayCollectionPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasCreateWayCollectionPayload)

	if err != nil {
		return err
	}

	*o = SchemasCreateWayCollectionPayload(varSchemasCreateWayCollectionPayload)

	return err
}

type NullableSchemasCreateWayCollectionPayload struct {
	value *SchemasCreateWayCollectionPayload
	isSet bool
}

func (v NullableSchemasCreateWayCollectionPayload) Get() *SchemasCreateWayCollectionPayload {
	return v.value
}

func (v *NullableSchemasCreateWayCollectionPayload) Set(val *SchemasCreateWayCollectionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasCreateWayCollectionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasCreateWayCollectionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasCreateWayCollectionPayload(val *SchemasCreateWayCollectionPayload) *NullableSchemasCreateWayCollectionPayload {
	return &NullableSchemasCreateWayCollectionPayload{value: val, isSet: true}
}

func (v NullableSchemasCreateWayCollectionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasCreateWayCollectionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


