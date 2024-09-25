/*
Masters way chat-websocket API

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

// checks if the SchemasSendMessagePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasSendMessagePayload{}

// SchemasSendMessagePayload struct for SchemasSendMessagePayload
type SchemasSendMessagePayload struct {
	Message SchemasMessageResponse
	Users []string
}

type _SchemasSendMessagePayload SchemasSendMessagePayload

// NewSchemasSendMessagePayload instantiates a new SchemasSendMessagePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasSendMessagePayload(message SchemasMessageResponse, users []string) *SchemasSendMessagePayload {
	this := SchemasSendMessagePayload{}
	this.Message = message
	this.Users = users
	return &this
}

// NewSchemasSendMessagePayloadWithDefaults instantiates a new SchemasSendMessagePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasSendMessagePayloadWithDefaults() *SchemasSendMessagePayload {
	this := SchemasSendMessagePayload{}
	return &this
}

// GetMessage returns the Message field value
func (o *SchemasSendMessagePayload) GetMessage() SchemasMessageResponse {
	if o == nil {
		var ret SchemasMessageResponse
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SchemasSendMessagePayload) GetMessageOk() (*SchemasMessageResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SchemasSendMessagePayload) SetMessage(v SchemasMessageResponse) {
	o.Message = v
}

// GetUsers returns the Users field value
func (o *SchemasSendMessagePayload) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *SchemasSendMessagePayload) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *SchemasSendMessagePayload) SetUsers(v []string) {
	o.Users = v
}

func (o SchemasSendMessagePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasSendMessagePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *SchemasSendMessagePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"users",
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

	varSchemasSendMessagePayload := _SchemasSendMessagePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasSendMessagePayload)

	if err != nil {
		return err
	}

	*o = SchemasSendMessagePayload(varSchemasSendMessagePayload)

	return err
}

type NullableSchemasSendMessagePayload struct {
	value *SchemasSendMessagePayload
	isSet bool
}

func (v NullableSchemasSendMessagePayload) Get() *SchemasSendMessagePayload {
	return v.value
}

func (v *NullableSchemasSendMessagePayload) Set(val *SchemasSendMessagePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasSendMessagePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasSendMessagePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasSendMessagePayload(val *SchemasSendMessagePayload) *NullableSchemasSendMessagePayload {
	return &NullableSchemasSendMessagePayload{value: val, isSet: true}
}

func (v NullableSchemasSendMessagePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasSendMessagePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


