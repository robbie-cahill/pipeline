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

// PackageReference A record of a software item which is vulnerable or carries a fix for a vulnerability
type PackageReference struct {
	// Package name
	Name *string `json:"name,omitempty"`
	// A version for the package. If null, then references all versions
	Version NullableString `json:"version,omitempty"`
	// Package type (e.g. package, rpm, deb, apk, jar, npm, gem, ...)
	Type *string `json:"type,omitempty"`
	// Whether a vendor will or will not fix a vulnerabitlity
	WillNotFix *bool `json:"will_not_fix,omitempty"`
}

// NewPackageReference instantiates a new PackageReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageReference() *PackageReference {
	this := PackageReference{}
	return &this
}

// NewPackageReferenceWithDefaults instantiates a new PackageReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageReferenceWithDefaults() *PackageReference {
	this := PackageReference{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PackageReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageReference) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PackageReference) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PackageReference) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageReference) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageReference) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *PackageReference) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *PackageReference) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *PackageReference) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *PackageReference) UnsetVersion() {
	o.Version.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PackageReference) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageReference) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PackageReference) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PackageReference) SetType(v string) {
	o.Type = &v
}

// GetWillNotFix returns the WillNotFix field value if set, zero value otherwise.
func (o *PackageReference) GetWillNotFix() bool {
	if o == nil || o.WillNotFix == nil {
		var ret bool
		return ret
	}
	return *o.WillNotFix
}

// GetWillNotFixOk returns a tuple with the WillNotFix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageReference) GetWillNotFixOk() (*bool, bool) {
	if o == nil || o.WillNotFix == nil {
		return nil, false
	}
	return o.WillNotFix, true
}

// HasWillNotFix returns a boolean if a field has been set.
func (o *PackageReference) HasWillNotFix() bool {
	if o != nil && o.WillNotFix != nil {
		return true
	}

	return false
}

// SetWillNotFix gets a reference to the given bool and assigns it to the WillNotFix field.
func (o *PackageReference) SetWillNotFix(v bool) {
	o.WillNotFix = &v
}

func (o PackageReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.WillNotFix != nil {
		toSerialize["will_not_fix"] = o.WillNotFix
	}
	return json.Marshal(toSerialize)
}

type NullablePackageReference struct {
	value *PackageReference
	isSet bool
}

func (v NullablePackageReference) Get() *PackageReference {
	return v.value
}

func (v *NullablePackageReference) Set(val *PackageReference) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageReference) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageReference(val *PackageReference) *NullablePackageReference {
	return &NullablePackageReference{value: val, isSet: true}
}

func (v NullablePackageReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


