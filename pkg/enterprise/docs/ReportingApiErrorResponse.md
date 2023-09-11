# ReportingApiErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **interface{}** | Details structure for additional information about the error if available. Content and structure will be error specific. | [optional] 

## Methods

### NewReportingApiErrorResponse

`func NewReportingApiErrorResponse() *ReportingApiErrorResponse`

NewReportingApiErrorResponse instantiates a new ReportingApiErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingApiErrorResponseWithDefaults

`func NewReportingApiErrorResponseWithDefaults() *ReportingApiErrorResponse`

NewReportingApiErrorResponseWithDefaults instantiates a new ReportingApiErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ReportingApiErrorResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReportingApiErrorResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReportingApiErrorResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReportingApiErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetErrorType

`func (o *ReportingApiErrorResponse) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *ReportingApiErrorResponse) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *ReportingApiErrorResponse) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *ReportingApiErrorResponse) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetMessage

`func (o *ReportingApiErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReportingApiErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReportingApiErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ReportingApiErrorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *ReportingApiErrorResponse) GetDetail() interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ReportingApiErrorResponse) GetDetailOk() (*interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ReportingApiErrorResponse) SetDetail(v interface{})`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ReportingApiErrorResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


