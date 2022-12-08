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

// checks if the ApiDocsRawReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDocsRawReport{}

// ApiDocsRawReport struct for ApiDocsRawReport
type ApiDocsRawReport struct {
	Payload *string `json:"payload,omitempty"`
}

// NewApiDocsRawReport instantiates a new ApiDocsRawReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDocsRawReport() *ApiDocsRawReport {
	this := ApiDocsRawReport{}
	return &this
}

// NewApiDocsRawReportWithDefaults instantiates a new ApiDocsRawReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDocsRawReportWithDefaults() *ApiDocsRawReport {
	this := ApiDocsRawReport{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ApiDocsRawReport) GetPayload() string {
	if o == nil || isNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDocsRawReport) GetPayloadOk() (*string, bool) {
	if o == nil || isNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ApiDocsRawReport) HasPayload() bool {
	if o != nil && !isNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *ApiDocsRawReport) SetPayload(v string) {
	o.Payload = &v
}

func (o ApiDocsRawReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDocsRawReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableApiDocsRawReport struct {
	value *ApiDocsRawReport
	isSet bool
}

func (v NullableApiDocsRawReport) Get() *ApiDocsRawReport {
	return v.value
}

func (v *NullableApiDocsRawReport) Set(val *ApiDocsRawReport) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDocsRawReport) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDocsRawReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDocsRawReport(val *ApiDocsRawReport) *NullableApiDocsRawReport {
	return &NullableApiDocsRawReport{value: val, isSet: true}
}

func (v NullableApiDocsRawReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDocsRawReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


