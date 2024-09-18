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

// checks if the SchemasMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasMessageResponse{}

// SchemasMessageResponse struct for SchemasMessageResponse
type SchemasMessageResponse struct {
	Message string
	MessageId string
	MessageReaders []SchemasMessageReader
	OwnerId string
	OwnerImageUrl string
	OwnerName string
	RoomId string
}

type _SchemasMessageResponse SchemasMessageResponse

// NewSchemasMessageResponse instantiates a new SchemasMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasMessageResponse(message string, messageId string, messageReaders []SchemasMessageReader, ownerId string, ownerImageUrl string, ownerName string, roomId string) *SchemasMessageResponse {
	this := SchemasMessageResponse{}
	this.Message = message
	this.MessageId = messageId
	this.MessageReaders = messageReaders
	this.OwnerId = ownerId
	this.OwnerImageUrl = ownerImageUrl
	this.OwnerName = ownerName
	this.RoomId = roomId
	return &this
}

// NewSchemasMessageResponseWithDefaults instantiates a new SchemasMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasMessageResponseWithDefaults() *SchemasMessageResponse {
	this := SchemasMessageResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *SchemasMessageResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SchemasMessageResponse) SetMessage(v string) {
	o.Message = v
}

// GetMessageId returns the MessageId field value
func (o *SchemasMessageResponse) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageResponse) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value
func (o *SchemasMessageResponse) SetMessageId(v string) {
	o.MessageId = v
}

// GetMessageReaders returns the MessageReaders field value
func (o *SchemasMessageResponse) GetMessageReaders() []SchemasMessageReader {
	if o == nil {
		var ret []SchemasMessageReader
		return ret
	}

	return o.MessageReaders
}

// GetMessageReadersOk returns a tuple with the MessageReaders field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageResponse) GetMessageReadersOk() ([]SchemasMessageReader, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageReaders, true
}

// SetMessageReaders sets field value
func (o *SchemasMessageResponse) SetMessageReaders(v []SchemasMessageReader) {
	o.MessageReaders = v
}

// GetOwnerId returns the OwnerId field value
func (o *SchemasMessageResponse) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageResponse) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *SchemasMessageResponse) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetOwnerImageUrl returns the OwnerImageUrl field value
func (o *SchemasMessageResponse) GetOwnerImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerImageUrl
}

// GetOwnerImageUrlOk returns a tuple with the OwnerImageUrl field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageResponse) GetOwnerImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerImageUrl, true
}

// SetOwnerImageUrl sets field value
func (o *SchemasMessageResponse) SetOwnerImageUrl(v string) {
	o.OwnerImageUrl = v
}

// GetOwnerName returns the OwnerName field value
func (o *SchemasMessageResponse) GetOwnerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageResponse) GetOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerName, true
}

// SetOwnerName sets field value
func (o *SchemasMessageResponse) SetOwnerName(v string) {
	o.OwnerName = v
}

// GetRoomId returns the RoomId field value
func (o *SchemasMessageResponse) GetRoomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageResponse) GetRoomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomId, true
}

// SetRoomId sets field value
func (o *SchemasMessageResponse) SetRoomId(v string) {
	o.RoomId = v
}

func (o SchemasMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["messageId"] = o.MessageId
	toSerialize["messageReaders"] = o.MessageReaders
	toSerialize["ownerId"] = o.OwnerId
	toSerialize["ownerImageUrl"] = o.OwnerImageUrl
	toSerialize["ownerName"] = o.OwnerName
	toSerialize["roomId"] = o.RoomId
	return toSerialize, nil
}

func (o *SchemasMessageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"messageId",
		"messageReaders",
		"ownerId",
		"ownerImageUrl",
		"ownerName",
		"roomId",
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

	varSchemasMessageResponse := _SchemasMessageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasMessageResponse)

	if err != nil {
		return err
	}

	*o = SchemasMessageResponse(varSchemasMessageResponse)

	return err
}

type NullableSchemasMessageResponse struct {
	value *SchemasMessageResponse
	isSet bool
}

func (v NullableSchemasMessageResponse) Get() *SchemasMessageResponse {
	return v.value
}

func (v *NullableSchemasMessageResponse) Set(val *SchemasMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasMessageResponse(val *SchemasMessageResponse) *NullableSchemasMessageResponse {
	return &NullableSchemasMessageResponse{value: val, isSet: true}
}

func (v NullableSchemasMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


