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

// GetRelatedComponentsResponseContext struct for GetRelatedComponentsResponseContext
type GetRelatedComponentsResponseContext struct {
	Source RelationSource `json:"source"`
	Type RelationType `json:"type"`
	Id string `json:"id"`
}

// NewGetRelatedComponentsResponseContext instantiates a new GetRelatedComponentsResponseContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRelatedComponentsResponseContext(source RelationSource, type_ RelationType, id string) *GetRelatedComponentsResponseContext {
	this := GetRelatedComponentsResponseContext{}
	this.Source = source
	this.Type = type_
	this.Id = id
	return &this
}

// NewGetRelatedComponentsResponseContextWithDefaults instantiates a new GetRelatedComponentsResponseContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRelatedComponentsResponseContextWithDefaults() *GetRelatedComponentsResponseContext {
	this := GetRelatedComponentsResponseContext{}
	return &this
}

// GetSource returns the Source field value
func (o *GetRelatedComponentsResponseContext) GetSource() RelationSource {
	if o == nil {
		var ret RelationSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GetRelatedComponentsResponseContext) GetSourceOk() (*RelationSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GetRelatedComponentsResponseContext) SetSource(v RelationSource) {
	o.Source = v
}

// GetType returns the Type field value
func (o *GetRelatedComponentsResponseContext) GetType() RelationType {
	if o == nil {
		var ret RelationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetRelatedComponentsResponseContext) GetTypeOk() (*RelationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetRelatedComponentsResponseContext) SetType(v RelationType) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetRelatedComponentsResponseContext) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetRelatedComponentsResponseContext) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetRelatedComponentsResponseContext) SetId(v string) {
	o.Id = v
}

func (o GetRelatedComponentsResponseContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGetRelatedComponentsResponseContext struct {
	value *GetRelatedComponentsResponseContext
	isSet bool
}

func (v NullableGetRelatedComponentsResponseContext) Get() *GetRelatedComponentsResponseContext {
	return v.value
}

func (v *NullableGetRelatedComponentsResponseContext) Set(val *GetRelatedComponentsResponseContext) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRelatedComponentsResponseContext) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRelatedComponentsResponseContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRelatedComponentsResponseContext(val *GetRelatedComponentsResponseContext) *NullableGetRelatedComponentsResponseContext {
	return &NullableGetRelatedComponentsResponseContext{value: val, isSet: true}
}

func (v NullableGetRelatedComponentsResponseContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRelatedComponentsResponseContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

