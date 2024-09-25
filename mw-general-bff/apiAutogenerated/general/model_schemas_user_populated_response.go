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

// checks if the SchemasUserPopulatedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUserPopulatedResponse{}

// SchemasUserPopulatedResponse struct for SchemasUserPopulatedResponse
type SchemasUserPopulatedResponse struct {
	CreatedAt string
	CustomWayCollections []SchemasWayCollectionPopulatedResponse
	DefaultWayCollections SchemasDefaultWayCollections
	Description string
	Email string
	FavoriteForUsers []string
	FavoriteUsers []SchemasUserPlainResponse
	ImageUrl string
	IsMentor bool
	Name string
	Tags []SchemasUserTagResponse
	Uuid string
	WayRequests []SchemasWayPlainResponse
}

type _SchemasUserPopulatedResponse SchemasUserPopulatedResponse

// NewSchemasUserPopulatedResponse instantiates a new SchemasUserPopulatedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUserPopulatedResponse(createdAt string, customWayCollections []SchemasWayCollectionPopulatedResponse, defaultWayCollections SchemasDefaultWayCollections, description string, email string, favoriteForUsers []string, favoriteUsers []SchemasUserPlainResponse, imageUrl string, isMentor bool, name string, tags []SchemasUserTagResponse, uuid string, wayRequests []SchemasWayPlainResponse) *SchemasUserPopulatedResponse {
	this := SchemasUserPopulatedResponse{}
	this.CreatedAt = createdAt
	this.CustomWayCollections = customWayCollections
	this.DefaultWayCollections = defaultWayCollections
	this.Description = description
	this.Email = email
	this.FavoriteForUsers = favoriteForUsers
	this.FavoriteUsers = favoriteUsers
	this.ImageUrl = imageUrl
	this.IsMentor = isMentor
	this.Name = name
	this.Tags = tags
	this.Uuid = uuid
	this.WayRequests = wayRequests
	return &this
}

// NewSchemasUserPopulatedResponseWithDefaults instantiates a new SchemasUserPopulatedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUserPopulatedResponseWithDefaults() *SchemasUserPopulatedResponse {
	this := SchemasUserPopulatedResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *SchemasUserPopulatedResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SchemasUserPopulatedResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCustomWayCollections returns the CustomWayCollections field value
func (o *SchemasUserPopulatedResponse) GetCustomWayCollections() []SchemasWayCollectionPopulatedResponse {
	if o == nil {
		var ret []SchemasWayCollectionPopulatedResponse
		return ret
	}

	return o.CustomWayCollections
}

// GetCustomWayCollectionsOk returns a tuple with the CustomWayCollections field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetCustomWayCollectionsOk() ([]SchemasWayCollectionPopulatedResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomWayCollections, true
}

// SetCustomWayCollections sets field value
func (o *SchemasUserPopulatedResponse) SetCustomWayCollections(v []SchemasWayCollectionPopulatedResponse) {
	o.CustomWayCollections = v
}

// GetDefaultWayCollections returns the DefaultWayCollections field value
func (o *SchemasUserPopulatedResponse) GetDefaultWayCollections() SchemasDefaultWayCollections {
	if o == nil {
		var ret SchemasDefaultWayCollections
		return ret
	}

	return o.DefaultWayCollections
}

// GetDefaultWayCollectionsOk returns a tuple with the DefaultWayCollections field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetDefaultWayCollectionsOk() (*SchemasDefaultWayCollections, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultWayCollections, true
}

// SetDefaultWayCollections sets field value
func (o *SchemasUserPopulatedResponse) SetDefaultWayCollections(v SchemasDefaultWayCollections) {
	o.DefaultWayCollections = v
}

// GetDescription returns the Description field value
func (o *SchemasUserPopulatedResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SchemasUserPopulatedResponse) SetDescription(v string) {
	o.Description = v
}

// GetEmail returns the Email field value
func (o *SchemasUserPopulatedResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SchemasUserPopulatedResponse) SetEmail(v string) {
	o.Email = v
}

// GetFavoriteForUsers returns the FavoriteForUsers field value
func (o *SchemasUserPopulatedResponse) GetFavoriteForUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FavoriteForUsers
}

// GetFavoriteForUsersOk returns a tuple with the FavoriteForUsers field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetFavoriteForUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FavoriteForUsers, true
}

// SetFavoriteForUsers sets field value
func (o *SchemasUserPopulatedResponse) SetFavoriteForUsers(v []string) {
	o.FavoriteForUsers = v
}

// GetFavoriteUsers returns the FavoriteUsers field value
func (o *SchemasUserPopulatedResponse) GetFavoriteUsers() []SchemasUserPlainResponse {
	if o == nil {
		var ret []SchemasUserPlainResponse
		return ret
	}

	return o.FavoriteUsers
}

// GetFavoriteUsersOk returns a tuple with the FavoriteUsers field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetFavoriteUsersOk() ([]SchemasUserPlainResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.FavoriteUsers, true
}

// SetFavoriteUsers sets field value
func (o *SchemasUserPopulatedResponse) SetFavoriteUsers(v []SchemasUserPlainResponse) {
	o.FavoriteUsers = v
}

// GetImageUrl returns the ImageUrl field value
func (o *SchemasUserPopulatedResponse) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *SchemasUserPopulatedResponse) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetIsMentor returns the IsMentor field value
func (o *SchemasUserPopulatedResponse) GetIsMentor() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMentor
}

// GetIsMentorOk returns a tuple with the IsMentor field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetIsMentorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMentor, true
}

// SetIsMentor sets field value
func (o *SchemasUserPopulatedResponse) SetIsMentor(v bool) {
	o.IsMentor = v
}

// GetName returns the Name field value
func (o *SchemasUserPopulatedResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasUserPopulatedResponse) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value
func (o *SchemasUserPopulatedResponse) GetTags() []SchemasUserTagResponse {
	if o == nil {
		var ret []SchemasUserTagResponse
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetTagsOk() ([]SchemasUserTagResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *SchemasUserPopulatedResponse) SetTags(v []SchemasUserTagResponse) {
	o.Tags = v
}

// GetUuid returns the Uuid field value
func (o *SchemasUserPopulatedResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SchemasUserPopulatedResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetWayRequests returns the WayRequests field value
func (o *SchemasUserPopulatedResponse) GetWayRequests() []SchemasWayPlainResponse {
	if o == nil {
		var ret []SchemasWayPlainResponse
		return ret
	}

	return o.WayRequests
}

// GetWayRequestsOk returns a tuple with the WayRequests field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPopulatedResponse) GetWayRequestsOk() ([]SchemasWayPlainResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.WayRequests, true
}

// SetWayRequests sets field value
func (o *SchemasUserPopulatedResponse) SetWayRequests(v []SchemasWayPlainResponse) {
	o.WayRequests = v
}

func (o SchemasUserPopulatedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUserPopulatedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["customWayCollections"] = o.CustomWayCollections
	toSerialize["defaultWayCollections"] = o.DefaultWayCollections
	toSerialize["description"] = o.Description
	toSerialize["email"] = o.Email
	toSerialize["favoriteForUsers"] = o.FavoriteForUsers
	toSerialize["favoriteUsers"] = o.FavoriteUsers
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["isMentor"] = o.IsMentor
	toSerialize["name"] = o.Name
	toSerialize["tags"] = o.Tags
	toSerialize["uuid"] = o.Uuid
	toSerialize["wayRequests"] = o.WayRequests
	return toSerialize, nil
}

func (o *SchemasUserPopulatedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"customWayCollections",
		"defaultWayCollections",
		"description",
		"email",
		"favoriteForUsers",
		"favoriteUsers",
		"imageUrl",
		"isMentor",
		"name",
		"tags",
		"uuid",
		"wayRequests",
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

	varSchemasUserPopulatedResponse := _SchemasUserPopulatedResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasUserPopulatedResponse)

	if err != nil {
		return err
	}

	*o = SchemasUserPopulatedResponse(varSchemasUserPopulatedResponse)

	return err
}

type NullableSchemasUserPopulatedResponse struct {
	value *SchemasUserPopulatedResponse
	isSet bool
}

func (v NullableSchemasUserPopulatedResponse) Get() *SchemasUserPopulatedResponse {
	return v.value
}

func (v *NullableSchemasUserPopulatedResponse) Set(val *SchemasUserPopulatedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUserPopulatedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUserPopulatedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUserPopulatedResponse(val *SchemasUserPopulatedResponse) *NullableSchemasUserPopulatedResponse {
	return &NullableSchemasUserPopulatedResponse{value: val, isSet: true}
}

func (v NullableSchemasUserPopulatedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUserPopulatedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

