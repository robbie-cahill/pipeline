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

// GenericNotificationPayload Parent class for Notification Payloads
type GenericNotificationPayload struct {
	UserId *string `json:"userId,omitempty"`
	SubscriptionKey *string `json:"subscription_key,omitempty"`
	SubscriptionType *string `json:"subscription_type,omitempty"`
	NotificationId *string `json:"notificationId,omitempty"`
}

// NewGenericNotificationPayload instantiates a new GenericNotificationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericNotificationPayload() *GenericNotificationPayload {
	this := GenericNotificationPayload{}
	return &this
}

// NewGenericNotificationPayloadWithDefaults instantiates a new GenericNotificationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericNotificationPayloadWithDefaults() *GenericNotificationPayload {
	this := GenericNotificationPayload{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GenericNotificationPayload) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNotificationPayload) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GenericNotificationPayload) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *GenericNotificationPayload) SetUserId(v string) {
	o.UserId = &v
}

// GetSubscriptionKey returns the SubscriptionKey field value if set, zero value otherwise.
func (o *GenericNotificationPayload) GetSubscriptionKey() string {
	if o == nil || o.SubscriptionKey == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionKey
}

// GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNotificationPayload) GetSubscriptionKeyOk() (*string, bool) {
	if o == nil || o.SubscriptionKey == nil {
		return nil, false
	}
	return o.SubscriptionKey, true
}

// HasSubscriptionKey returns a boolean if a field has been set.
func (o *GenericNotificationPayload) HasSubscriptionKey() bool {
	if o != nil && o.SubscriptionKey != nil {
		return true
	}

	return false
}

// SetSubscriptionKey gets a reference to the given string and assigns it to the SubscriptionKey field.
func (o *GenericNotificationPayload) SetSubscriptionKey(v string) {
	o.SubscriptionKey = &v
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise.
func (o *GenericNotificationPayload) GetSubscriptionType() string {
	if o == nil || o.SubscriptionType == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionType
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNotificationPayload) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil || o.SubscriptionType == nil {
		return nil, false
	}
	return o.SubscriptionType, true
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *GenericNotificationPayload) HasSubscriptionType() bool {
	if o != nil && o.SubscriptionType != nil {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given string and assigns it to the SubscriptionType field.
func (o *GenericNotificationPayload) SetSubscriptionType(v string) {
	o.SubscriptionType = &v
}

// GetNotificationId returns the NotificationId field value if set, zero value otherwise.
func (o *GenericNotificationPayload) GetNotificationId() string {
	if o == nil || o.NotificationId == nil {
		var ret string
		return ret
	}
	return *o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericNotificationPayload) GetNotificationIdOk() (*string, bool) {
	if o == nil || o.NotificationId == nil {
		return nil, false
	}
	return o.NotificationId, true
}

// HasNotificationId returns a boolean if a field has been set.
func (o *GenericNotificationPayload) HasNotificationId() bool {
	if o != nil && o.NotificationId != nil {
		return true
	}

	return false
}

// SetNotificationId gets a reference to the given string and assigns it to the NotificationId field.
func (o *GenericNotificationPayload) SetNotificationId(v string) {
	o.NotificationId = &v
}

func (o GenericNotificationPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.SubscriptionKey != nil {
		toSerialize["subscription_key"] = o.SubscriptionKey
	}
	if o.SubscriptionType != nil {
		toSerialize["subscription_type"] = o.SubscriptionType
	}
	if o.NotificationId != nil {
		toSerialize["notificationId"] = o.NotificationId
	}
	return json.Marshal(toSerialize)
}

type NullableGenericNotificationPayload struct {
	value *GenericNotificationPayload
	isSet bool
}

func (v NullableGenericNotificationPayload) Get() *GenericNotificationPayload {
	return v.value
}

func (v *NullableGenericNotificationPayload) Set(val *GenericNotificationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericNotificationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericNotificationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericNotificationPayload(val *GenericNotificationPayload) *NullableGenericNotificationPayload {
	return &NullableGenericNotificationPayload{value: val, isSet: true}
}

func (v NullableGenericNotificationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericNotificationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


