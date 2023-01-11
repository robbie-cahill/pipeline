/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.20
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package anchore

import (
	"encoding/json"
)

// VulnUpdateNotificationAllOf The Notification Object definition for Tag Update Notifications
type VulnUpdateNotificationAllOf struct {
	Data *VulnUpdateNotificationData `json:"data,omitempty"`
}

// NewVulnUpdateNotificationAllOf instantiates a new VulnUpdateNotificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnUpdateNotificationAllOf() *VulnUpdateNotificationAllOf {
	this := VulnUpdateNotificationAllOf{}
	return &this
}

// NewVulnUpdateNotificationAllOfWithDefaults instantiates a new VulnUpdateNotificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnUpdateNotificationAllOfWithDefaults() *VulnUpdateNotificationAllOf {
	this := VulnUpdateNotificationAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VulnUpdateNotificationAllOf) GetData() VulnUpdateNotificationData {
	if o == nil || o.Data == nil {
		var ret VulnUpdateNotificationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnUpdateNotificationAllOf) GetDataOk() (*VulnUpdateNotificationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VulnUpdateNotificationAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given VulnUpdateNotificationData and assigns it to the Data field.
func (o *VulnUpdateNotificationAllOf) SetData(v VulnUpdateNotificationData) {
	o.Data = &v
}

func (o VulnUpdateNotificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableVulnUpdateNotificationAllOf struct {
	value *VulnUpdateNotificationAllOf
	isSet bool
}

func (v NullableVulnUpdateNotificationAllOf) Get() *VulnUpdateNotificationAllOf {
	return v.value
}

func (v *NullableVulnUpdateNotificationAllOf) Set(val *VulnUpdateNotificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnUpdateNotificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnUpdateNotificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnUpdateNotificationAllOf(val *VulnUpdateNotificationAllOf) *NullableVulnUpdateNotificationAllOf {
	return &NullableVulnUpdateNotificationAllOf{value: val, isSet: true}
}

func (v NullableVulnUpdateNotificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnUpdateNotificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


