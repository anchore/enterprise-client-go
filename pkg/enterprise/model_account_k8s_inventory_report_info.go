/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AccountK8sInventoryReportInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountK8sInventoryReportInfo{}

// AccountK8sInventoryReportInfo Information about sent K8s inventory report for an account
type AccountK8sInventoryReportInfo struct {
	// Timestamp for the inventory report
	ReportTimestamp string `json:"report_timestamp"`
	// Account to which the inventory reports belong
	AccountName string `json:"account_name"`
	// User that the integration instance sent the inventory report as
	SentAsUser string `json:"sent_as_user"`
	// Number of batches that the inventory report was sent in
	BatchSize int32 `json:"batch_size"`
	// Index of last successfully sent batch, -1 if none were successful
	LastSuccessfulIndex int32 `json:"last_successful_index"`
	// true if one or more of the batches coult not be sent, false otherwise
	HasErrors bool `json:"has_errors"`
	// List of information about inventory report batches
	Batches []AccountK8sInventoryReportInfoBatchesInner `json:"batches"`
}

type _AccountK8sInventoryReportInfo AccountK8sInventoryReportInfo

// NewAccountK8sInventoryReportInfo instantiates a new AccountK8sInventoryReportInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountK8sInventoryReportInfo(reportTimestamp string, accountName string, sentAsUser string, batchSize int32, lastSuccessfulIndex int32, hasErrors bool, batches []AccountK8sInventoryReportInfoBatchesInner) *AccountK8sInventoryReportInfo {
	this := AccountK8sInventoryReportInfo{}
	this.ReportTimestamp = reportTimestamp
	this.AccountName = accountName
	this.SentAsUser = sentAsUser
	this.BatchSize = batchSize
	this.LastSuccessfulIndex = lastSuccessfulIndex
	this.HasErrors = hasErrors
	this.Batches = batches
	return &this
}

// NewAccountK8sInventoryReportInfoWithDefaults instantiates a new AccountK8sInventoryReportInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountK8sInventoryReportInfoWithDefaults() *AccountK8sInventoryReportInfo {
	this := AccountK8sInventoryReportInfo{}
	return &this
}

// GetReportTimestamp returns the ReportTimestamp field value
func (o *AccountK8sInventoryReportInfo) GetReportTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportTimestamp
}

// GetReportTimestampOk returns a tuple with the ReportTimestamp field value
// and a boolean to check if the value has been set.
func (o *AccountK8sInventoryReportInfo) GetReportTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportTimestamp, true
}

// SetReportTimestamp sets field value
func (o *AccountK8sInventoryReportInfo) SetReportTimestamp(v string) {
	o.ReportTimestamp = v
}

// GetAccountName returns the AccountName field value
func (o *AccountK8sInventoryReportInfo) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *AccountK8sInventoryReportInfo) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *AccountK8sInventoryReportInfo) SetAccountName(v string) {
	o.AccountName = v
}

// GetSentAsUser returns the SentAsUser field value
func (o *AccountK8sInventoryReportInfo) GetSentAsUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SentAsUser
}

// GetSentAsUserOk returns a tuple with the SentAsUser field value
// and a boolean to check if the value has been set.
func (o *AccountK8sInventoryReportInfo) GetSentAsUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SentAsUser, true
}

// SetSentAsUser sets field value
func (o *AccountK8sInventoryReportInfo) SetSentAsUser(v string) {
	o.SentAsUser = v
}

// GetBatchSize returns the BatchSize field value
func (o *AccountK8sInventoryReportInfo) GetBatchSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value
// and a boolean to check if the value has been set.
func (o *AccountK8sInventoryReportInfo) GetBatchSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchSize, true
}

// SetBatchSize sets field value
func (o *AccountK8sInventoryReportInfo) SetBatchSize(v int32) {
	o.BatchSize = v
}

// GetLastSuccessfulIndex returns the LastSuccessfulIndex field value
func (o *AccountK8sInventoryReportInfo) GetLastSuccessfulIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LastSuccessfulIndex
}

// GetLastSuccessfulIndexOk returns a tuple with the LastSuccessfulIndex field value
// and a boolean to check if the value has been set.
func (o *AccountK8sInventoryReportInfo) GetLastSuccessfulIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSuccessfulIndex, true
}

// SetLastSuccessfulIndex sets field value
func (o *AccountK8sInventoryReportInfo) SetLastSuccessfulIndex(v int32) {
	o.LastSuccessfulIndex = v
}

// GetHasErrors returns the HasErrors field value
func (o *AccountK8sInventoryReportInfo) GetHasErrors() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasErrors
}

// GetHasErrorsOk returns a tuple with the HasErrors field value
// and a boolean to check if the value has been set.
func (o *AccountK8sInventoryReportInfo) GetHasErrorsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasErrors, true
}

// SetHasErrors sets field value
func (o *AccountK8sInventoryReportInfo) SetHasErrors(v bool) {
	o.HasErrors = v
}

// GetBatches returns the Batches field value
func (o *AccountK8sInventoryReportInfo) GetBatches() []AccountK8sInventoryReportInfoBatchesInner {
	if o == nil {
		var ret []AccountK8sInventoryReportInfoBatchesInner
		return ret
	}

	return o.Batches
}

// GetBatchesOk returns a tuple with the Batches field value
// and a boolean to check if the value has been set.
func (o *AccountK8sInventoryReportInfo) GetBatchesOk() ([]AccountK8sInventoryReportInfoBatchesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Batches, true
}

// SetBatches sets field value
func (o *AccountK8sInventoryReportInfo) SetBatches(v []AccountK8sInventoryReportInfoBatchesInner) {
	o.Batches = v
}

func (o AccountK8sInventoryReportInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountK8sInventoryReportInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["report_timestamp"] = o.ReportTimestamp
	toSerialize["account_name"] = o.AccountName
	toSerialize["sent_as_user"] = o.SentAsUser
	toSerialize["batch_size"] = o.BatchSize
	toSerialize["last_successful_index"] = o.LastSuccessfulIndex
	toSerialize["has_errors"] = o.HasErrors
	toSerialize["batches"] = o.Batches
	return toSerialize, nil
}

func (o *AccountK8sInventoryReportInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"report_timestamp",
		"account_name",
		"sent_as_user",
		"batch_size",
		"last_successful_index",
		"has_errors",
		"batches",
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

	varAccountK8sInventoryReportInfo := _AccountK8sInventoryReportInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountK8sInventoryReportInfo)

	if err != nil {
		return err
	}

	*o = AccountK8sInventoryReportInfo(varAccountK8sInventoryReportInfo)

	return err
}

type NullableAccountK8sInventoryReportInfo struct {
	value *AccountK8sInventoryReportInfo
	isSet bool
}

func (v NullableAccountK8sInventoryReportInfo) Get() *AccountK8sInventoryReportInfo {
	return v.value
}

func (v *NullableAccountK8sInventoryReportInfo) Set(val *AccountK8sInventoryReportInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountK8sInventoryReportInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountK8sInventoryReportInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountK8sInventoryReportInfo(val *AccountK8sInventoryReportInfo) *NullableAccountK8sInventoryReportInfo {
	return &NullableAccountK8sInventoryReportInfo{value: val, isSet: true}
}

func (v NullableAccountK8sInventoryReportInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountK8sInventoryReportInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


