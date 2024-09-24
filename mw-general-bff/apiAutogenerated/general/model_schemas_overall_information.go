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

// checks if the SchemasOverallInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasOverallInformation{}

// SchemasOverallInformation struct for SchemasOverallInformation
type SchemasOverallInformation struct {
	AverageJobTime int32
	AverageTimePerCalendarDay int32
	AverageTimePerWorkingDay int32
	FinishedJobs int32
	TotalReports int32
	TotalTime int32
}

type _SchemasOverallInformation SchemasOverallInformation

// NewSchemasOverallInformation instantiates a new SchemasOverallInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasOverallInformation(averageJobTime int32, averageTimePerCalendarDay int32, averageTimePerWorkingDay int32, finishedJobs int32, totalReports int32, totalTime int32) *SchemasOverallInformation {
	this := SchemasOverallInformation{}
	this.AverageJobTime = averageJobTime
	this.AverageTimePerCalendarDay = averageTimePerCalendarDay
	this.AverageTimePerWorkingDay = averageTimePerWorkingDay
	this.FinishedJobs = finishedJobs
	this.TotalReports = totalReports
	this.TotalTime = totalTime
	return &this
}

// NewSchemasOverallInformationWithDefaults instantiates a new SchemasOverallInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasOverallInformationWithDefaults() *SchemasOverallInformation {
	this := SchemasOverallInformation{}
	return &this
}

// GetAverageJobTime returns the AverageJobTime field value
func (o *SchemasOverallInformation) GetAverageJobTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AverageJobTime
}

// GetAverageJobTimeOk returns a tuple with the AverageJobTime field value
// and a boolean to check if the value has been set.
func (o *SchemasOverallInformation) GetAverageJobTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageJobTime, true
}

// SetAverageJobTime sets field value
func (o *SchemasOverallInformation) SetAverageJobTime(v int32) {
	o.AverageJobTime = v
}

// GetAverageTimePerCalendarDay returns the AverageTimePerCalendarDay field value
func (o *SchemasOverallInformation) GetAverageTimePerCalendarDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AverageTimePerCalendarDay
}

// GetAverageTimePerCalendarDayOk returns a tuple with the AverageTimePerCalendarDay field value
// and a boolean to check if the value has been set.
func (o *SchemasOverallInformation) GetAverageTimePerCalendarDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageTimePerCalendarDay, true
}

// SetAverageTimePerCalendarDay sets field value
func (o *SchemasOverallInformation) SetAverageTimePerCalendarDay(v int32) {
	o.AverageTimePerCalendarDay = v
}

// GetAverageTimePerWorkingDay returns the AverageTimePerWorkingDay field value
func (o *SchemasOverallInformation) GetAverageTimePerWorkingDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AverageTimePerWorkingDay
}

// GetAverageTimePerWorkingDayOk returns a tuple with the AverageTimePerWorkingDay field value
// and a boolean to check if the value has been set.
func (o *SchemasOverallInformation) GetAverageTimePerWorkingDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageTimePerWorkingDay, true
}

// SetAverageTimePerWorkingDay sets field value
func (o *SchemasOverallInformation) SetAverageTimePerWorkingDay(v int32) {
	o.AverageTimePerWorkingDay = v
}

// GetFinishedJobs returns the FinishedJobs field value
func (o *SchemasOverallInformation) GetFinishedJobs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FinishedJobs
}

// GetFinishedJobsOk returns a tuple with the FinishedJobs field value
// and a boolean to check if the value has been set.
func (o *SchemasOverallInformation) GetFinishedJobsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedJobs, true
}

// SetFinishedJobs sets field value
func (o *SchemasOverallInformation) SetFinishedJobs(v int32) {
	o.FinishedJobs = v
}

// GetTotalReports returns the TotalReports field value
func (o *SchemasOverallInformation) GetTotalReports() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalReports
}

// GetTotalReportsOk returns a tuple with the TotalReports field value
// and a boolean to check if the value has been set.
func (o *SchemasOverallInformation) GetTotalReportsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalReports, true
}

// SetTotalReports sets field value
func (o *SchemasOverallInformation) SetTotalReports(v int32) {
	o.TotalReports = v
}

// GetTotalTime returns the TotalTime field value
func (o *SchemasOverallInformation) GetTotalTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value
// and a boolean to check if the value has been set.
func (o *SchemasOverallInformation) GetTotalTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTime, true
}

// SetTotalTime sets field value
func (o *SchemasOverallInformation) SetTotalTime(v int32) {
	o.TotalTime = v
}

func (o SchemasOverallInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasOverallInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["averageJobTime"] = o.AverageJobTime
	toSerialize["averageTimePerCalendarDay"] = o.AverageTimePerCalendarDay
	toSerialize["averageTimePerWorkingDay"] = o.AverageTimePerWorkingDay
	toSerialize["finishedJobs"] = o.FinishedJobs
	toSerialize["totalReports"] = o.TotalReports
	toSerialize["totalTime"] = o.TotalTime
	return toSerialize, nil
}

func (o *SchemasOverallInformation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"averageJobTime",
		"averageTimePerCalendarDay",
		"averageTimePerWorkingDay",
		"finishedJobs",
		"totalReports",
		"totalTime",
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

	varSchemasOverallInformation := _SchemasOverallInformation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasOverallInformation)

	if err != nil {
		return err
	}

	*o = SchemasOverallInformation(varSchemasOverallInformation)

	return err
}

type NullableSchemasOverallInformation struct {
	value *SchemasOverallInformation
	isSet bool
}

func (v NullableSchemasOverallInformation) Get() *SchemasOverallInformation {
	return v.value
}

func (v *NullableSchemasOverallInformation) Set(val *SchemasOverallInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasOverallInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasOverallInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasOverallInformation(val *SchemasOverallInformation) *NullableSchemasOverallInformation {
	return &NullableSchemasOverallInformation{value: val, isSet: true}
}

func (v NullableSchemasOverallInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasOverallInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


