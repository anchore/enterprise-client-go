# RbacManagerApiErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**ErrorType** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **interface{}** | Details structure for additional information about the error if available. Content and structure will be error specific. | [optional] 

## Methods

### NewRbacManagerApiErrorResponse

`func NewRbacManagerApiErrorResponse() *RbacManagerApiErrorResponse`

NewRbacManagerApiErrorResponse instantiates a new RbacManagerApiErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacManagerApiErrorResponseWithDefaults

`func NewRbacManagerApiErrorResponseWithDefaults() *RbacManagerApiErrorResponse`

NewRbacManagerApiErrorResponseWithDefaults instantiates a new RbacManagerApiErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RbacManagerApiErrorResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RbacManagerApiErrorResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RbacManagerApiErrorResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *RbacManagerApiErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetErrorType

`func (o *RbacManagerApiErrorResponse) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *RbacManagerApiErrorResponse) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *RbacManagerApiErrorResponse) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *RbacManagerApiErrorResponse) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetMessage

`func (o *RbacManagerApiErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RbacManagerApiErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RbacManagerApiErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RbacManagerApiErrorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *RbacManagerApiErrorResponse) GetDetail() interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RbacManagerApiErrorResponse) GetDetailOk() (*interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RbacManagerApiErrorResponse) SetDetail(v interface{})`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *RbacManagerApiErrorResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


