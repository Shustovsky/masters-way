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

// checks if the SchemasRoomPopulatedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasRoomPopulatedResponse{}

// SchemasRoomPopulatedResponse struct for SchemasRoomPopulatedResponse
type SchemasRoomPopulatedResponse struct {
	ImageUrl string
	Messages []SchemasMessageResponse
	Name string
	RoomId string
	RoomType string
	Users []SchemasUserResponse
}

type _SchemasRoomPopulatedResponse SchemasRoomPopulatedResponse

// NewSchemasRoomPopulatedResponse instantiates a new SchemasRoomPopulatedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasRoomPopulatedResponse(imageUrl string, messages []SchemasMessageResponse, name string, roomId string, roomType string, users []SchemasUserResponse) *SchemasRoomPopulatedResponse {
	this := SchemasRoomPopulatedResponse{}
	this.ImageUrl = imageUrl
	this.Messages = messages
	this.Name = name
	this.RoomId = roomId
	this.RoomType = roomType
	this.Users = users
	return &this
}

// NewSchemasRoomPopulatedResponseWithDefaults instantiates a new SchemasRoomPopulatedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasRoomPopulatedResponseWithDefaults() *SchemasRoomPopulatedResponse {
	this := SchemasRoomPopulatedResponse{}
	return &this
}

// GetImageUrl returns the ImageUrl field value
func (o *SchemasRoomPopulatedResponse) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *SchemasRoomPopulatedResponse) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *SchemasRoomPopulatedResponse) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetMessages returns the Messages field value
func (o *SchemasRoomPopulatedResponse) GetMessages() []SchemasMessageResponse {
	if o == nil {
		var ret []SchemasMessageResponse
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *SchemasRoomPopulatedResponse) GetMessagesOk() ([]SchemasMessageResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *SchemasRoomPopulatedResponse) SetMessages(v []SchemasMessageResponse) {
	o.Messages = v
}

// GetName returns the Name field value
func (o *SchemasRoomPopulatedResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasRoomPopulatedResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasRoomPopulatedResponse) SetName(v string) {
	o.Name = v
}

// GetRoomId returns the RoomId field value
func (o *SchemasRoomPopulatedResponse) GetRoomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value
// and a boolean to check if the value has been set.
func (o *SchemasRoomPopulatedResponse) GetRoomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomId, true
}

// SetRoomId sets field value
func (o *SchemasRoomPopulatedResponse) SetRoomId(v string) {
	o.RoomId = v
}

// GetRoomType returns the RoomType field value
func (o *SchemasRoomPopulatedResponse) GetRoomType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value
// and a boolean to check if the value has been set.
func (o *SchemasRoomPopulatedResponse) GetRoomTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomType, true
}

// SetRoomType sets field value
func (o *SchemasRoomPopulatedResponse) SetRoomType(v string) {
	o.RoomType = v
}

// GetUsers returns the Users field value
func (o *SchemasRoomPopulatedResponse) GetUsers() []SchemasUserResponse {
	if o == nil {
		var ret []SchemasUserResponse
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *SchemasRoomPopulatedResponse) GetUsersOk() ([]SchemasUserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *SchemasRoomPopulatedResponse) SetUsers(v []SchemasUserResponse) {
	o.Users = v
}

func (o SchemasRoomPopulatedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasRoomPopulatedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["messages"] = o.Messages
	toSerialize["name"] = o.Name
	toSerialize["roomId"] = o.RoomId
	toSerialize["roomType"] = o.RoomType
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *SchemasRoomPopulatedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"imageUrl",
		"messages",
		"name",
		"roomId",
		"roomType",
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

	varSchemasRoomPopulatedResponse := _SchemasRoomPopulatedResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasRoomPopulatedResponse)

	if err != nil {
		return err
	}

	*o = SchemasRoomPopulatedResponse(varSchemasRoomPopulatedResponse)

	return err
}

type NullableSchemasRoomPopulatedResponse struct {
	value *SchemasRoomPopulatedResponse
	isSet bool
}

func (v NullableSchemasRoomPopulatedResponse) Get() *SchemasRoomPopulatedResponse {
	return v.value
}

func (v *NullableSchemasRoomPopulatedResponse) Set(val *SchemasRoomPopulatedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasRoomPopulatedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasRoomPopulatedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasRoomPopulatedResponse(val *SchemasRoomPopulatedResponse) *NullableSchemasRoomPopulatedResponse {
	return &NullableSchemasRoomPopulatedResponse{value: val, isSet: true}
}

func (v NullableSchemasRoomPopulatedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasRoomPopulatedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


