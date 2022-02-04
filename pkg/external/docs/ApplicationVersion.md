# ApplicationVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationVersionId** | Pointer to **string** | The id of the application version | [optional] 
**ApplicationId** | Pointer to **string** | The id of the application | [optional] 
**VersionName** | **string** | The name of the application | 
**CreatedAt** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the application was created | [optional] 
**LastUpdated** | Pointer to **time.Time** | RFC 3339 formatted UTC timestamp when the application was last updated | [optional] 

## Methods

### NewApplicationVersion

`func NewApplicationVersion(versionName string, ) *ApplicationVersion`

NewApplicationVersion instantiates a new ApplicationVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationVersionWithDefaults

`func NewApplicationVersionWithDefaults() *ApplicationVersion`

NewApplicationVersionWithDefaults instantiates a new ApplicationVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationVersionId

`func (o *ApplicationVersion) GetApplicationVersionId() string`

GetApplicationVersionId returns the ApplicationVersionId field if non-nil, zero value otherwise.

### GetApplicationVersionIdOk

`func (o *ApplicationVersion) GetApplicationVersionIdOk() (*string, bool)`

GetApplicationVersionIdOk returns a tuple with the ApplicationVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersionId

`func (o *ApplicationVersion) SetApplicationVersionId(v string)`

SetApplicationVersionId sets ApplicationVersionId field to given value.

### HasApplicationVersionId

`func (o *ApplicationVersion) HasApplicationVersionId() bool`

HasApplicationVersionId returns a boolean if a field has been set.

### GetApplicationId

`func (o *ApplicationVersion) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationVersion) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationVersion) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApplicationVersion) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetVersionName

`func (o *ApplicationVersion) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *ApplicationVersion) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *ApplicationVersion) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetCreatedAt

`func (o *ApplicationVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ApplicationVersion) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ApplicationVersion) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ApplicationVersion) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ApplicationVersion) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


