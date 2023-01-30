/*
Asset Components

An Asset Component is an internal part of an Asset which is defined by the Hierarchy API.  Currently the following components are supported, * Bearings * Shafts

API version: 0.1
Contact: team.gob@enlight.skf.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// FixedSpeed struct for FixedSpeed
type FixedSpeed struct {
	Type string `json:"type"`
	Unit string `json:"unit"`
	Value float64 `json:"value"`
}

// NewFixedSpeed instantiates a new FixedSpeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedSpeed(type_ string, unit string, value float64) *FixedSpeed {
	this := FixedSpeed{}
	this.Type = type_
	this.Unit = unit
	this.Value = value
	return &this
}

// NewFixedSpeedWithDefaults instantiates a new FixedSpeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedSpeedWithDefaults() *FixedSpeed {
	this := FixedSpeed{}
	return &this
}

// GetType returns the Type field value
func (o *FixedSpeed) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FixedSpeed) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FixedSpeed) SetType(v string) {
	o.Type = v
}

// GetUnit returns the Unit field value
func (o *FixedSpeed) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *FixedSpeed) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *FixedSpeed) SetUnit(v string) {
	o.Unit = v
}

// GetValue returns the Value field value
func (o *FixedSpeed) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FixedSpeed) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FixedSpeed) SetValue(v float64) {
	o.Value = v
}

func (o FixedSpeed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFixedSpeed struct {
	value *FixedSpeed
	isSet bool
}

func (v NullableFixedSpeed) Get() *FixedSpeed {
	return v.value
}

func (v *NullableFixedSpeed) Set(val *FixedSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedSpeed(val *FixedSpeed) *NullableFixedSpeed {
	return &NullableFixedSpeed{value: val, isSet: true}
}

func (v NullableFixedSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


