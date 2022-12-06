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

// Component struct for Component
type Component struct {
	// Unique identifier for a component
	Id *string `json:"id,omitempty"`
	// Hierarchy Asset ID where this component is located within
	Asset string `json:"asset"`
	// Identifier to the component that this one is attached to, required for bearings
	AttachedTo *string `json:"attachedTo,omitempty"`
	// What kind of component it is. Will decide what other attributes can be used.
	Type string `json:"type"`
	// A non unique identifier within the asset to help locate the component
	Position int32 `json:"position"`
	Origin *Origin `json:"origin,omitempty"`
	// The fixed rotation speed of the shaft, expressed in RPM
	FixedSpeed *float64 `json:"fixedSpeed,omitempty"`
	// A human-readable description of where the bearing is located
	PositionDescription *string `json:"positionDescription,omitempty"`
	// Bearing model number
	Designation *string `json:"designation,omitempty"`
	// Who is producing the bearing, usually a company name
	Manufacturer *string `json:"manufacturer,omitempty"`
	// A manufacturer designated identifier for a specific bearing
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Indicating what ring or rings that is rotating under operating conditions
	RotatingRing *string `json:"rotatingRing,omitempty"`
	// What side of the shaft the bearing is mounted
	ShaftSide *string `json:"shaftSide,omitempty"`
}

// NewComponent instantiates a new Component object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponent(asset string, type_ string, position int32) *Component {
	this := Component{}
	this.Asset = asset
	this.Type = type_
	this.Position = position
	return &this
}

// NewComponentWithDefaults instantiates a new Component object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentWithDefaults() *Component {
	this := Component{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Component) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Component) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Component) SetId(v string) {
	o.Id = &v
}

// GetAsset returns the Asset field value
func (o *Component) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *Component) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *Component) SetAsset(v string) {
	o.Asset = v
}

// GetAttachedTo returns the AttachedTo field value if set, zero value otherwise.
func (o *Component) GetAttachedTo() string {
	if o == nil || o.AttachedTo == nil {
		var ret string
		return ret
	}
	return *o.AttachedTo
}

// GetAttachedToOk returns a tuple with the AttachedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetAttachedToOk() (*string, bool) {
	if o == nil || o.AttachedTo == nil {
		return nil, false
	}
	return o.AttachedTo, true
}

// HasAttachedTo returns a boolean if a field has been set.
func (o *Component) HasAttachedTo() bool {
	if o != nil && o.AttachedTo != nil {
		return true
	}

	return false
}

// SetAttachedTo gets a reference to the given string and assigns it to the AttachedTo field.
func (o *Component) SetAttachedTo(v string) {
	o.AttachedTo = &v
}

// GetType returns the Type field value
func (o *Component) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Component) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Component) SetType(v string) {
	o.Type = v
}

// GetPosition returns the Position field value
func (o *Component) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *Component) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *Component) SetPosition(v int32) {
	o.Position = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *Component) GetOrigin() Origin {
	if o == nil || o.Origin == nil {
		var ret Origin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetOriginOk() (*Origin, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *Component) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given Origin and assigns it to the Origin field.
func (o *Component) SetOrigin(v Origin) {
	o.Origin = &v
}

// GetFixedSpeed returns the FixedSpeed field value if set, zero value otherwise.
func (o *Component) GetFixedSpeed() float64 {
	if o == nil || o.FixedSpeed == nil {
		var ret float64
		return ret
	}
	return *o.FixedSpeed
}

// GetFixedSpeedOk returns a tuple with the FixedSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetFixedSpeedOk() (*float64, bool) {
	if o == nil || o.FixedSpeed == nil {
		return nil, false
	}
	return o.FixedSpeed, true
}

// HasFixedSpeed returns a boolean if a field has been set.
func (o *Component) HasFixedSpeed() bool {
	if o != nil && o.FixedSpeed != nil {
		return true
	}

	return false
}

// SetFixedSpeed gets a reference to the given float64 and assigns it to the FixedSpeed field.
func (o *Component) SetFixedSpeed(v float64) {
	o.FixedSpeed = &v
}

// GetPositionDescription returns the PositionDescription field value if set, zero value otherwise.
func (o *Component) GetPositionDescription() string {
	if o == nil || o.PositionDescription == nil {
		var ret string
		return ret
	}
	return *o.PositionDescription
}

// GetPositionDescriptionOk returns a tuple with the PositionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetPositionDescriptionOk() (*string, bool) {
	if o == nil || o.PositionDescription == nil {
		return nil, false
	}
	return o.PositionDescription, true
}

// HasPositionDescription returns a boolean if a field has been set.
func (o *Component) HasPositionDescription() bool {
	if o != nil && o.PositionDescription != nil {
		return true
	}

	return false
}

// SetPositionDescription gets a reference to the given string and assigns it to the PositionDescription field.
func (o *Component) SetPositionDescription(v string) {
	o.PositionDescription = &v
}

// GetDesignation returns the Designation field value if set, zero value otherwise.
func (o *Component) GetDesignation() string {
	if o == nil || o.Designation == nil {
		var ret string
		return ret
	}
	return *o.Designation
}

// GetDesignationOk returns a tuple with the Designation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetDesignationOk() (*string, bool) {
	if o == nil || o.Designation == nil {
		return nil, false
	}
	return o.Designation, true
}

// HasDesignation returns a boolean if a field has been set.
func (o *Component) HasDesignation() bool {
	if o != nil && o.Designation != nil {
		return true
	}

	return false
}

// SetDesignation gets a reference to the given string and assigns it to the Designation field.
func (o *Component) SetDesignation(v string) {
	o.Designation = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *Component) GetManufacturer() string {
	if o == nil || o.Manufacturer == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetManufacturerOk() (*string, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *Component) HasManufacturer() bool {
	if o != nil && o.Manufacturer != nil {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *Component) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *Component) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *Component) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *Component) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetRotatingRing returns the RotatingRing field value if set, zero value otherwise.
func (o *Component) GetRotatingRing() string {
	if o == nil || o.RotatingRing == nil {
		var ret string
		return ret
	}
	return *o.RotatingRing
}

// GetRotatingRingOk returns a tuple with the RotatingRing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetRotatingRingOk() (*string, bool) {
	if o == nil || o.RotatingRing == nil {
		return nil, false
	}
	return o.RotatingRing, true
}

// HasRotatingRing returns a boolean if a field has been set.
func (o *Component) HasRotatingRing() bool {
	if o != nil && o.RotatingRing != nil {
		return true
	}

	return false
}

// SetRotatingRing gets a reference to the given string and assigns it to the RotatingRing field.
func (o *Component) SetRotatingRing(v string) {
	o.RotatingRing = &v
}

// GetShaftSide returns the ShaftSide field value if set, zero value otherwise.
func (o *Component) GetShaftSide() string {
	if o == nil || o.ShaftSide == nil {
		var ret string
		return ret
	}
	return *o.ShaftSide
}

// GetShaftSideOk returns a tuple with the ShaftSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetShaftSideOk() (*string, bool) {
	if o == nil || o.ShaftSide == nil {
		return nil, false
	}
	return o.ShaftSide, true
}

// HasShaftSide returns a boolean if a field has been set.
func (o *Component) HasShaftSide() bool {
	if o != nil && o.ShaftSide != nil {
		return true
	}

	return false
}

// SetShaftSide gets a reference to the given string and assigns it to the ShaftSide field.
func (o *Component) SetShaftSide(v string) {
	o.ShaftSide = &v
}

func (o Component) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if o.AttachedTo != nil {
		toSerialize["attachedTo"] = o.AttachedTo
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["position"] = o.Position
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.FixedSpeed != nil {
		toSerialize["fixedSpeed"] = o.FixedSpeed
	}
	if o.PositionDescription != nil {
		toSerialize["positionDescription"] = o.PositionDescription
	}
	if o.Designation != nil {
		toSerialize["designation"] = o.Designation
	}
	if o.Manufacturer != nil {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if o.SerialNumber != nil {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if o.RotatingRing != nil {
		toSerialize["rotatingRing"] = o.RotatingRing
	}
	if o.ShaftSide != nil {
		toSerialize["shaftSide"] = o.ShaftSide
	}
	return json.Marshal(toSerialize)
}

type NullableComponent struct {
	value *Component
	isSet bool
}

func (v NullableComponent) Get() *Component {
	return v.value
}

func (v *NullableComponent) Set(val *Component) {
	v.value = val
	v.isSet = true
}

func (v NullableComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponent(val *Component) *NullableComponent {
	return &NullableComponent{value: val, isSet: true}
}

func (v NullableComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


