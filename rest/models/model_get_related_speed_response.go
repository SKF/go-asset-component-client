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

// GetRelatedSpeedResponse struct for GetRelatedSpeedResponse
type GetRelatedSpeedResponse struct {
	Fixedspeed GetRelatedSpeedResponseFixedspeed `json:"fixedspeed"`
	Context GetRelatedSpeedResponseContext `json:"context"`
}

// NewGetRelatedSpeedResponse instantiates a new GetRelatedSpeedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRelatedSpeedResponse(fixedspeed GetRelatedSpeedResponseFixedspeed, context GetRelatedSpeedResponseContext) *GetRelatedSpeedResponse {
	this := GetRelatedSpeedResponse{}
	this.Fixedspeed = fixedspeed
	this.Context = context
	return &this
}

// NewGetRelatedSpeedResponseWithDefaults instantiates a new GetRelatedSpeedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRelatedSpeedResponseWithDefaults() *GetRelatedSpeedResponse {
	this := GetRelatedSpeedResponse{}
	return &this
}

// GetFixedspeed returns the Fixedspeed field value
func (o *GetRelatedSpeedResponse) GetFixedspeed() GetRelatedSpeedResponseFixedspeed {
	if o == nil {
		var ret GetRelatedSpeedResponseFixedspeed
		return ret
	}

	return o.Fixedspeed
}

// GetFixedspeedOk returns a tuple with the Fixedspeed field value
// and a boolean to check if the value has been set.
func (o *GetRelatedSpeedResponse) GetFixedspeedOk() (*GetRelatedSpeedResponseFixedspeed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fixedspeed, true
}

// SetFixedspeed sets field value
func (o *GetRelatedSpeedResponse) SetFixedspeed(v GetRelatedSpeedResponseFixedspeed) {
	o.Fixedspeed = v
}

// GetContext returns the Context field value
func (o *GetRelatedSpeedResponse) GetContext() GetRelatedSpeedResponseContext {
	if o == nil {
		var ret GetRelatedSpeedResponseContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *GetRelatedSpeedResponse) GetContextOk() (*GetRelatedSpeedResponseContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *GetRelatedSpeedResponse) SetContext(v GetRelatedSpeedResponseContext) {
	o.Context = v
}

func (o GetRelatedSpeedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fixedspeed"] = o.Fixedspeed
	}
	if true {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableGetRelatedSpeedResponse struct {
	value *GetRelatedSpeedResponse
	isSet bool
}

func (v NullableGetRelatedSpeedResponse) Get() *GetRelatedSpeedResponse {
	return v.value
}

func (v *NullableGetRelatedSpeedResponse) Set(val *GetRelatedSpeedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRelatedSpeedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRelatedSpeedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRelatedSpeedResponse(val *GetRelatedSpeedResponse) *NullableGetRelatedSpeedResponse {
	return &NullableGetRelatedSpeedResponse{value: val, isSet: true}
}

func (v NullableGetRelatedSpeedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRelatedSpeedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


