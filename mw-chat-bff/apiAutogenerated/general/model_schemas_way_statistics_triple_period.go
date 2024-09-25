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

// checks if the SchemasWayStatisticsTriplePeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasWayStatisticsTriplePeriod{}

// SchemasWayStatisticsTriplePeriod struct for SchemasWayStatisticsTriplePeriod
type SchemasWayStatisticsTriplePeriod struct {
	LastMonth SchemasWayStatistics
	LastWeek SchemasWayStatistics
	TotalTime SchemasWayStatistics
}

type _SchemasWayStatisticsTriplePeriod SchemasWayStatisticsTriplePeriod

// NewSchemasWayStatisticsTriplePeriod instantiates a new SchemasWayStatisticsTriplePeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasWayStatisticsTriplePeriod(lastMonth SchemasWayStatistics, lastWeek SchemasWayStatistics, totalTime SchemasWayStatistics) *SchemasWayStatisticsTriplePeriod {
	this := SchemasWayStatisticsTriplePeriod{}
	this.LastMonth = lastMonth
	this.LastWeek = lastWeek
	this.TotalTime = totalTime
	return &this
}

// NewSchemasWayStatisticsTriplePeriodWithDefaults instantiates a new SchemasWayStatisticsTriplePeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasWayStatisticsTriplePeriodWithDefaults() *SchemasWayStatisticsTriplePeriod {
	this := SchemasWayStatisticsTriplePeriod{}
	return &this
}

// GetLastMonth returns the LastMonth field value
func (o *SchemasWayStatisticsTriplePeriod) GetLastMonth() SchemasWayStatistics {
	if o == nil {
		var ret SchemasWayStatistics
		return ret
	}

	return o.LastMonth
}

// GetLastMonthOk returns a tuple with the LastMonth field value
// and a boolean to check if the value has been set.
func (o *SchemasWayStatisticsTriplePeriod) GetLastMonthOk() (*SchemasWayStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastMonth, true
}

// SetLastMonth sets field value
func (o *SchemasWayStatisticsTriplePeriod) SetLastMonth(v SchemasWayStatistics) {
	o.LastMonth = v
}

// GetLastWeek returns the LastWeek field value
func (o *SchemasWayStatisticsTriplePeriod) GetLastWeek() SchemasWayStatistics {
	if o == nil {
		var ret SchemasWayStatistics
		return ret
	}

	return o.LastWeek
}

// GetLastWeekOk returns a tuple with the LastWeek field value
// and a boolean to check if the value has been set.
func (o *SchemasWayStatisticsTriplePeriod) GetLastWeekOk() (*SchemasWayStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastWeek, true
}

// SetLastWeek sets field value
func (o *SchemasWayStatisticsTriplePeriod) SetLastWeek(v SchemasWayStatistics) {
	o.LastWeek = v
}

// GetTotalTime returns the TotalTime field value
func (o *SchemasWayStatisticsTriplePeriod) GetTotalTime() SchemasWayStatistics {
	if o == nil {
		var ret SchemasWayStatistics
		return ret
	}

	return o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value
// and a boolean to check if the value has been set.
func (o *SchemasWayStatisticsTriplePeriod) GetTotalTimeOk() (*SchemasWayStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTime, true
}

// SetTotalTime sets field value
func (o *SchemasWayStatisticsTriplePeriod) SetTotalTime(v SchemasWayStatistics) {
	o.TotalTime = v
}

func (o SchemasWayStatisticsTriplePeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasWayStatisticsTriplePeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lastMonth"] = o.LastMonth
	toSerialize["lastWeek"] = o.LastWeek
	toSerialize["totalTime"] = o.TotalTime
	return toSerialize, nil
}

func (o *SchemasWayStatisticsTriplePeriod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lastMonth",
		"lastWeek",
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

	varSchemasWayStatisticsTriplePeriod := _SchemasWayStatisticsTriplePeriod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasWayStatisticsTriplePeriod)

	if err != nil {
		return err
	}

	*o = SchemasWayStatisticsTriplePeriod(varSchemasWayStatisticsTriplePeriod)

	return err
}

type NullableSchemasWayStatisticsTriplePeriod struct {
	value *SchemasWayStatisticsTriplePeriod
	isSet bool
}

func (v NullableSchemasWayStatisticsTriplePeriod) Get() *SchemasWayStatisticsTriplePeriod {
	return v.value
}

func (v *NullableSchemasWayStatisticsTriplePeriod) Set(val *SchemasWayStatisticsTriplePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasWayStatisticsTriplePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasWayStatisticsTriplePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasWayStatisticsTriplePeriod(val *SchemasWayStatisticsTriplePeriod) *NullableSchemasWayStatisticsTriplePeriod {
	return &NullableSchemasWayStatisticsTriplePeriod{value: val, isSet: true}
}

func (v NullableSchemasWayStatisticsTriplePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasWayStatisticsTriplePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


