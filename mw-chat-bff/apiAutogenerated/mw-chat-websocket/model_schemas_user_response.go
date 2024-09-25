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

// checks if the SchemasUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUserResponse{}

// SchemasUserResponse struct for SchemasUserResponse
type SchemasUserResponse struct {
	ImageUrl string
	Name string
	Role string
	UserId string
}

type _SchemasUserResponse SchemasUserResponse

// NewSchemasUserResponse instantiates a new SchemasUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUserResponse(imageUrl string, name string, role string, userId string) *SchemasUserResponse {
	this := SchemasUserResponse{}
	this.ImageUrl = imageUrl
	this.Name = name
	this.Role = role
	this.UserId = userId
	return &this
}

// NewSchemasUserResponseWithDefaults instantiates a new SchemasUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUserResponseWithDefaults() *SchemasUserResponse {
	this := SchemasUserResponse{}
	return &this
}

// GetImageUrl returns the ImageUrl field value
func (o *SchemasUserResponse) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *SchemasUserResponse) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetName returns the Name field value
func (o *SchemasUserResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasUserResponse) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *SchemasUserResponse) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *SchemasUserResponse) SetRole(v string) {
	o.Role = v
}

// GetUserId returns the UserId field value
func (o *SchemasUserResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SchemasUserResponse) SetUserId(v string) {
	o.UserId = v
}

func (o SchemasUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["name"] = o.Name
	toSerialize["role"] = o.Role
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *SchemasUserResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"imageUrl",
		"name",
		"role",
		"userId",
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

	varSchemasUserResponse := _SchemasUserResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasUserResponse)

	if err != nil {
		return err
	}

	*o = SchemasUserResponse(varSchemasUserResponse)

	return err
}

type NullableSchemasUserResponse struct {
	value *SchemasUserResponse
	isSet bool
}

func (v NullableSchemasUserResponse) Get() *SchemasUserResponse {
	return v.value
}

func (v *NullableSchemasUserResponse) Set(val *SchemasUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUserResponse(val *SchemasUserResponse) *NullableSchemasUserResponse {
	return &NullableSchemasUserResponse{value: val, isSet: true}
}

func (v NullableSchemasUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


