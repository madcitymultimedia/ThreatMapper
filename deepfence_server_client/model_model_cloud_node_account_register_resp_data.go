/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deepfence_server_client

import (
	"encoding/json"
)

// checks if the ModelCloudNodeAccountRegisterRespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeAccountRegisterRespData{}

// ModelCloudNodeAccountRegisterRespData struct for ModelCloudNodeAccountRegisterRespData
type ModelCloudNodeAccountRegisterRespData struct {
	CloudtrailTrails []ModelCloudNodeCloudtrailTrail `json:"cloudtrail_trails,omitempty"`
	Refresh *string `json:"refresh,omitempty"`
	Scans map[string]ModelCloudComplianceScanDetails `json:"scans,omitempty"`
}

// NewModelCloudNodeAccountRegisterRespData instantiates a new ModelCloudNodeAccountRegisterRespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeAccountRegisterRespData() *ModelCloudNodeAccountRegisterRespData {
	this := ModelCloudNodeAccountRegisterRespData{}
	return &this
}

// NewModelCloudNodeAccountRegisterRespDataWithDefaults instantiates a new ModelCloudNodeAccountRegisterRespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeAccountRegisterRespDataWithDefaults() *ModelCloudNodeAccountRegisterRespData {
	this := ModelCloudNodeAccountRegisterRespData{}
	return &this
}

// GetCloudtrailTrails returns the CloudtrailTrails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudNodeAccountRegisterRespData) GetCloudtrailTrails() []ModelCloudNodeCloudtrailTrail {
	if o == nil {
		var ret []ModelCloudNodeCloudtrailTrail
		return ret
	}
	return o.CloudtrailTrails
}

// GetCloudtrailTrailsOk returns a tuple with the CloudtrailTrails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudNodeAccountRegisterRespData) GetCloudtrailTrailsOk() ([]ModelCloudNodeCloudtrailTrail, bool) {
	if o == nil || isNil(o.CloudtrailTrails) {
		return nil, false
	}
	return o.CloudtrailTrails, true
}

// HasCloudtrailTrails returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountRegisterRespData) HasCloudtrailTrails() bool {
	if o != nil && isNil(o.CloudtrailTrails) {
		return true
	}

	return false
}

// SetCloudtrailTrails gets a reference to the given []ModelCloudNodeCloudtrailTrail and assigns it to the CloudtrailTrails field.
func (o *ModelCloudNodeAccountRegisterRespData) SetCloudtrailTrails(v []ModelCloudNodeCloudtrailTrail) {
	o.CloudtrailTrails = v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountRegisterRespData) GetRefresh() string {
	if o == nil || isNil(o.Refresh) {
		var ret string
		return ret
	}
	return *o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountRegisterRespData) GetRefreshOk() (*string, bool) {
	if o == nil || isNil(o.Refresh) {
		return nil, false
	}
	return o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountRegisterRespData) HasRefresh() bool {
	if o != nil && !isNil(o.Refresh) {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given string and assigns it to the Refresh field.
func (o *ModelCloudNodeAccountRegisterRespData) SetRefresh(v string) {
	o.Refresh = &v
}

// GetScans returns the Scans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudNodeAccountRegisterRespData) GetScans() map[string]ModelCloudComplianceScanDetails {
	if o == nil {
		var ret map[string]ModelCloudComplianceScanDetails
		return ret
	}
	return o.Scans
}

// GetScansOk returns a tuple with the Scans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudNodeAccountRegisterRespData) GetScansOk() (*map[string]ModelCloudComplianceScanDetails, bool) {
	if o == nil || isNil(o.Scans) {
		return nil, false
	}
	return &o.Scans, true
}

// HasScans returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountRegisterRespData) HasScans() bool {
	if o != nil && isNil(o.Scans) {
		return true
	}

	return false
}

// SetScans gets a reference to the given map[string]ModelCloudComplianceScanDetails and assigns it to the Scans field.
func (o *ModelCloudNodeAccountRegisterRespData) SetScans(v map[string]ModelCloudComplianceScanDetails) {
	o.Scans = v
}

func (o ModelCloudNodeAccountRegisterRespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeAccountRegisterRespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudtrailTrails != nil {
		toSerialize["cloudtrail_trails"] = o.CloudtrailTrails
	}
	if !isNil(o.Refresh) {
		toSerialize["refresh"] = o.Refresh
	}
	if o.Scans != nil {
		toSerialize["scans"] = o.Scans
	}
	return toSerialize, nil
}

type NullableModelCloudNodeAccountRegisterRespData struct {
	value *ModelCloudNodeAccountRegisterRespData
	isSet bool
}

func (v NullableModelCloudNodeAccountRegisterRespData) Get() *ModelCloudNodeAccountRegisterRespData {
	return v.value
}

func (v *NullableModelCloudNodeAccountRegisterRespData) Set(val *ModelCloudNodeAccountRegisterRespData) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeAccountRegisterRespData) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeAccountRegisterRespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeAccountRegisterRespData(val *ModelCloudNodeAccountRegisterRespData) *NullableModelCloudNodeAccountRegisterRespData {
	return &NullableModelCloudNodeAccountRegisterRespData{value: val, isSet: true}
}

func (v NullableModelCloudNodeAccountRegisterRespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeAccountRegisterRespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


