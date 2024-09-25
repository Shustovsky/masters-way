/*
Masters way chat API

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

// checks if the SchemasUploadFileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUploadFileResponse{}

// SchemasUploadFileResponse struct for SchemasUploadFileResponse
type SchemasUploadFileResponse struct {
	Id string
	Name string
	OwnerId string
	PreviewUrl string
	SrcUrl string
}

type _SchemasUploadFileResponse SchemasUploadFileResponse

// NewSchemasUploadFileResponse instantiates a new SchemasUploadFileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUploadFileResponse(id string, name string, ownerId string, previewUrl string, srcUrl string) *SchemasUploadFileResponse {
	this := SchemasUploadFileResponse{}
	this.Id = id
	this.Name = name
	this.OwnerId = ownerId
	this.PreviewUrl = previewUrl
	this.SrcUrl = srcUrl
	return &this
}

// NewSchemasUploadFileResponseWithDefaults instantiates a new SchemasUploadFileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUploadFileResponseWithDefaults() *SchemasUploadFileResponse {
	this := SchemasUploadFileResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SchemasUploadFileResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SchemasUploadFileResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SchemasUploadFileResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SchemasUploadFileResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasUploadFileResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasUploadFileResponse) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *SchemasUploadFileResponse) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *SchemasUploadFileResponse) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *SchemasUploadFileResponse) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetPreviewUrl returns the PreviewUrl field value
func (o *SchemasUploadFileResponse) GetPreviewUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviewUrl
}

// GetPreviewUrlOk returns a tuple with the PreviewUrl field value
// and a boolean to check if the value has been set.
func (o *SchemasUploadFileResponse) GetPreviewUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviewUrl, true
}

// SetPreviewUrl sets field value
func (o *SchemasUploadFileResponse) SetPreviewUrl(v string) {
	o.PreviewUrl = v
}

// GetSrcUrl returns the SrcUrl field value
func (o *SchemasUploadFileResponse) GetSrcUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SrcUrl
}

// GetSrcUrlOk returns a tuple with the SrcUrl field value
// and a boolean to check if the value has been set.
func (o *SchemasUploadFileResponse) GetSrcUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SrcUrl, true
}

// SetSrcUrl sets field value
func (o *SchemasUploadFileResponse) SetSrcUrl(v string) {
	o.SrcUrl = v
}

func (o SchemasUploadFileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUploadFileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["ownerId"] = o.OwnerId
	toSerialize["previewUrl"] = o.PreviewUrl
	toSerialize["srcUrl"] = o.SrcUrl
	return toSerialize, nil
}

func (o *SchemasUploadFileResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"ownerId",
		"previewUrl",
		"srcUrl",
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

	varSchemasUploadFileResponse := _SchemasUploadFileResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasUploadFileResponse)

	if err != nil {
		return err
	}

	*o = SchemasUploadFileResponse(varSchemasUploadFileResponse)

	return err
}

type NullableSchemasUploadFileResponse struct {
	value *SchemasUploadFileResponse
	isSet bool
}

func (v NullableSchemasUploadFileResponse) Get() *SchemasUploadFileResponse {
	return v.value
}

func (v *NullableSchemasUploadFileResponse) Set(val *SchemasUploadFileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUploadFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUploadFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUploadFileResponse(val *SchemasUploadFileResponse) *NullableSchemasUploadFileResponse {
	return &NullableSchemasUploadFileResponse{value: val, isSet: true}
}

func (v NullableSchemasUploadFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUploadFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


