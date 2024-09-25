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

// checks if the SchemasGoogleToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasGoogleToken{}

// SchemasGoogleToken struct for SchemasGoogleToken
type SchemasGoogleToken struct {
	AccessToken string
}

type _SchemasGoogleToken SchemasGoogleToken

// NewSchemasGoogleToken instantiates a new SchemasGoogleToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasGoogleToken(accessToken string) *SchemasGoogleToken {
	this := SchemasGoogleToken{}
	this.AccessToken = accessToken
	return &this
}

// NewSchemasGoogleTokenWithDefaults instantiates a new SchemasGoogleToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasGoogleTokenWithDefaults() *SchemasGoogleToken {
	this := SchemasGoogleToken{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *SchemasGoogleToken) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *SchemasGoogleToken) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *SchemasGoogleToken) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o SchemasGoogleToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasGoogleToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessToken"] = o.AccessToken
	return toSerialize, nil
}

func (o *SchemasGoogleToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accessToken",
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

	varSchemasGoogleToken := _SchemasGoogleToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasGoogleToken)

	if err != nil {
		return err
	}

	*o = SchemasGoogleToken(varSchemasGoogleToken)

	return err
}

type NullableSchemasGoogleToken struct {
	value *SchemasGoogleToken
	isSet bool
}

func (v NullableSchemasGoogleToken) Get() *SchemasGoogleToken {
	return v.value
}

func (v *NullableSchemasGoogleToken) Set(val *SchemasGoogleToken) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasGoogleToken) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasGoogleToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasGoogleToken(val *SchemasGoogleToken) *NullableSchemasGoogleToken {
	return &NullableSchemasGoogleToken{value: val, isSet: true}
}

func (v NullableSchemasGoogleToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasGoogleToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


