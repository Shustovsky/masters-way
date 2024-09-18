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

// checks if the SchemasMessageReader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasMessageReader{}

// SchemasMessageReader struct for SchemasMessageReader
type SchemasMessageReader struct {
	ImageUrl string
	Name string
	ReadDate string
	UserId string
}

type _SchemasMessageReader SchemasMessageReader

// NewSchemasMessageReader instantiates a new SchemasMessageReader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasMessageReader(imageUrl string, name string, readDate string, userId string) *SchemasMessageReader {
	this := SchemasMessageReader{}
	this.ImageUrl = imageUrl
	this.Name = name
	this.ReadDate = readDate
	this.UserId = userId
	return &this
}

// NewSchemasMessageReaderWithDefaults instantiates a new SchemasMessageReader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasMessageReaderWithDefaults() *SchemasMessageReader {
	this := SchemasMessageReader{}
	return &this
}

// GetImageUrl returns the ImageUrl field value
func (o *SchemasMessageReader) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageReader) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *SchemasMessageReader) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetName returns the Name field value
func (o *SchemasMessageReader) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageReader) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasMessageReader) SetName(v string) {
	o.Name = v
}

// GetReadDate returns the ReadDate field value
func (o *SchemasMessageReader) GetReadDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReadDate
}

// GetReadDateOk returns a tuple with the ReadDate field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageReader) GetReadDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadDate, true
}

// SetReadDate sets field value
func (o *SchemasMessageReader) SetReadDate(v string) {
	o.ReadDate = v
}

// GetUserId returns the UserId field value
func (o *SchemasMessageReader) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SchemasMessageReader) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SchemasMessageReader) SetUserId(v string) {
	o.UserId = v
}

func (o SchemasMessageReader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasMessageReader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["name"] = o.Name
	toSerialize["readDate"] = o.ReadDate
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *SchemasMessageReader) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"imageUrl",
		"name",
		"readDate",
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

	varSchemasMessageReader := _SchemasMessageReader{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasMessageReader)

	if err != nil {
		return err
	}

	*o = SchemasMessageReader(varSchemasMessageReader)

	return err
}

type NullableSchemasMessageReader struct {
	value *SchemasMessageReader
	isSet bool
}

func (v NullableSchemasMessageReader) Get() *SchemasMessageReader {
	return v.value
}

func (v *NullableSchemasMessageReader) Set(val *SchemasMessageReader) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasMessageReader) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasMessageReader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasMessageReader(val *SchemasMessageReader) *NullableSchemasMessageReader {
	return &NullableSchemasMessageReader{value: val, isSet: true}
}

func (v NullableSchemasMessageReader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasMessageReader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


