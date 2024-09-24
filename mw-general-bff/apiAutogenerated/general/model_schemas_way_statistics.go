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

// checks if the SchemasWayStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasWayStatistics{}

// SchemasWayStatistics struct for SchemasWayStatistics
type SchemasWayStatistics struct {
	LabelStatistics SchemasLabelStatistics
	OverallInformation SchemasOverallInformation
	TimeSpentByDayChart []SchemasTimeSpentByDayPoint
}

type _SchemasWayStatistics SchemasWayStatistics

// NewSchemasWayStatistics instantiates a new SchemasWayStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasWayStatistics(labelStatistics SchemasLabelStatistics, overallInformation SchemasOverallInformation, timeSpentByDayChart []SchemasTimeSpentByDayPoint) *SchemasWayStatistics {
	this := SchemasWayStatistics{}
	this.LabelStatistics = labelStatistics
	this.OverallInformation = overallInformation
	this.TimeSpentByDayChart = timeSpentByDayChart
	return &this
}

// NewSchemasWayStatisticsWithDefaults instantiates a new SchemasWayStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasWayStatisticsWithDefaults() *SchemasWayStatistics {
	this := SchemasWayStatistics{}
	return &this
}

// GetLabelStatistics returns the LabelStatistics field value
func (o *SchemasWayStatistics) GetLabelStatistics() SchemasLabelStatistics {
	if o == nil {
		var ret SchemasLabelStatistics
		return ret
	}

	return o.LabelStatistics
}

// GetLabelStatisticsOk returns a tuple with the LabelStatistics field value
// and a boolean to check if the value has been set.
func (o *SchemasWayStatistics) GetLabelStatisticsOk() (*SchemasLabelStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelStatistics, true
}

// SetLabelStatistics sets field value
func (o *SchemasWayStatistics) SetLabelStatistics(v SchemasLabelStatistics) {
	o.LabelStatistics = v
}

// GetOverallInformation returns the OverallInformation field value
func (o *SchemasWayStatistics) GetOverallInformation() SchemasOverallInformation {
	if o == nil {
		var ret SchemasOverallInformation
		return ret
	}

	return o.OverallInformation
}

// GetOverallInformationOk returns a tuple with the OverallInformation field value
// and a boolean to check if the value has been set.
func (o *SchemasWayStatistics) GetOverallInformationOk() (*SchemasOverallInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverallInformation, true
}

// SetOverallInformation sets field value
func (o *SchemasWayStatistics) SetOverallInformation(v SchemasOverallInformation) {
	o.OverallInformation = v
}

// GetTimeSpentByDayChart returns the TimeSpentByDayChart field value
func (o *SchemasWayStatistics) GetTimeSpentByDayChart() []SchemasTimeSpentByDayPoint {
	if o == nil {
		var ret []SchemasTimeSpentByDayPoint
		return ret
	}

	return o.TimeSpentByDayChart
}

// GetTimeSpentByDayChartOk returns a tuple with the TimeSpentByDayChart field value
// and a boolean to check if the value has been set.
func (o *SchemasWayStatistics) GetTimeSpentByDayChartOk() ([]SchemasTimeSpentByDayPoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeSpentByDayChart, true
}

// SetTimeSpentByDayChart sets field value
func (o *SchemasWayStatistics) SetTimeSpentByDayChart(v []SchemasTimeSpentByDayPoint) {
	o.TimeSpentByDayChart = v
}

func (o SchemasWayStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasWayStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labelStatistics"] = o.LabelStatistics
	toSerialize["overallInformation"] = o.OverallInformation
	toSerialize["timeSpentByDayChart"] = o.TimeSpentByDayChart
	return toSerialize, nil
}

func (o *SchemasWayStatistics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"labelStatistics",
		"overallInformation",
		"timeSpentByDayChart",
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

	varSchemasWayStatistics := _SchemasWayStatistics{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasWayStatistics)

	if err != nil {
		return err
	}

	*o = SchemasWayStatistics(varSchemasWayStatistics)

	return err
}

type NullableSchemasWayStatistics struct {
	value *SchemasWayStatistics
	isSet bool
}

func (v NullableSchemasWayStatistics) Get() *SchemasWayStatistics {
	return v.value
}

func (v *NullableSchemasWayStatistics) Set(val *SchemasWayStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasWayStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasWayStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasWayStatistics(val *SchemasWayStatistics) *NullableSchemasWayStatistics {
	return &NullableSchemasWayStatistics{value: val, isSet: true}
}

func (v NullableSchemasWayStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasWayStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


