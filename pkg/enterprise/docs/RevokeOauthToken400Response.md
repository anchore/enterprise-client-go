# RevokeOauthToken400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | ASCII error code from RFC6749 | [optional] 

## Methods

### NewRevokeOauthToken400Response

`func NewRevokeOauthToken400Response() *RevokeOauthToken400Response`

NewRevokeOauthToken400Response instantiates a new RevokeOauthToken400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeOauthToken400ResponseWithDefaults

`func NewRevokeOauthToken400ResponseWithDefaults() *RevokeOauthToken400Response`

NewRevokeOauthToken400ResponseWithDefaults instantiates a new RevokeOauthToken400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RevokeOauthToken400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RevokeOauthToken400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RevokeOauthToken400Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RevokeOauthToken400Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


