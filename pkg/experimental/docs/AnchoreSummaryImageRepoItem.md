# AnchoreSummaryImageRepoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **string** |  | [optional] 
**AnalysisStatus** | Pointer to **string** |  | [optional] 
**ImageCount** | Pointer to **int32** |  | [optional] 
**LatestTagDetectedAt** | Pointer to **time.Time** |  | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**InventoryType** | Pointer to **string** |  | [optional] 
**LastSeen** | Pointer to **time.Time** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**Alert** | Pointer to **bool** |  | [optional] 
**AlertId** | Pointer to **string** |  | [optional] 
**WatchTag** | Pointer to **bool** |  | [optional] 
**WatchTagId** | Pointer to **string** |  | [optional] 

## Methods

### NewAnchoreSummaryImageRepoItem

`func NewAnchoreSummaryImageRepoItem() *AnchoreSummaryImageRepoItem`

NewAnchoreSummaryImageRepoItem instantiates a new AnchoreSummaryImageRepoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnchoreSummaryImageRepoItemWithDefaults

`func NewAnchoreSummaryImageRepoItemWithDefaults() *AnchoreSummaryImageRepoItem`

NewAnchoreSummaryImageRepoItemWithDefaults instantiates a new AnchoreSummaryImageRepoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *AnchoreSummaryImageRepoItem) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AnchoreSummaryImageRepoItem) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AnchoreSummaryImageRepoItem) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AnchoreSummaryImageRepoItem) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetAnalysisStatus

`func (o *AnchoreSummaryImageRepoItem) GetAnalysisStatus() string`

GetAnalysisStatus returns the AnalysisStatus field if non-nil, zero value otherwise.

### GetAnalysisStatusOk

`func (o *AnchoreSummaryImageRepoItem) GetAnalysisStatusOk() (*string, bool)`

GetAnalysisStatusOk returns a tuple with the AnalysisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisStatus

`func (o *AnchoreSummaryImageRepoItem) SetAnalysisStatus(v string)`

SetAnalysisStatus sets AnalysisStatus field to given value.

### HasAnalysisStatus

`func (o *AnchoreSummaryImageRepoItem) HasAnalysisStatus() bool`

HasAnalysisStatus returns a boolean if a field has been set.

### GetImageCount

`func (o *AnchoreSummaryImageRepoItem) GetImageCount() int32`

GetImageCount returns the ImageCount field if non-nil, zero value otherwise.

### GetImageCountOk

`func (o *AnchoreSummaryImageRepoItem) GetImageCountOk() (*int32, bool)`

GetImageCountOk returns a tuple with the ImageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCount

`func (o *AnchoreSummaryImageRepoItem) SetImageCount(v int32)`

SetImageCount sets ImageCount field to given value.

### HasImageCount

`func (o *AnchoreSummaryImageRepoItem) HasImageCount() bool`

HasImageCount returns a boolean if a field has been set.

### GetLatestTagDetectedAt

`func (o *AnchoreSummaryImageRepoItem) GetLatestTagDetectedAt() time.Time`

GetLatestTagDetectedAt returns the LatestTagDetectedAt field if non-nil, zero value otherwise.

### GetLatestTagDetectedAtOk

`func (o *AnchoreSummaryImageRepoItem) GetLatestTagDetectedAtOk() (*time.Time, bool)`

GetLatestTagDetectedAtOk returns a tuple with the LatestTagDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTagDetectedAt

`func (o *AnchoreSummaryImageRepoItem) SetLatestTagDetectedAt(v time.Time)`

SetLatestTagDetectedAt sets LatestTagDetectedAt field to given value.

### HasLatestTagDetectedAt

`func (o *AnchoreSummaryImageRepoItem) HasLatestTagDetectedAt() bool`

HasLatestTagDetectedAt returns a boolean if a field has been set.

### GetImageDigest

`func (o *AnchoreSummaryImageRepoItem) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *AnchoreSummaryImageRepoItem) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *AnchoreSummaryImageRepoItem) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *AnchoreSummaryImageRepoItem) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetInventoryType

`func (o *AnchoreSummaryImageRepoItem) GetInventoryType() string`

GetInventoryType returns the InventoryType field if non-nil, zero value otherwise.

### GetInventoryTypeOk

`func (o *AnchoreSummaryImageRepoItem) GetInventoryTypeOk() (*string, bool)`

GetInventoryTypeOk returns a tuple with the InventoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryType

`func (o *AnchoreSummaryImageRepoItem) SetInventoryType(v string)`

SetInventoryType sets InventoryType field to given value.

### HasInventoryType

`func (o *AnchoreSummaryImageRepoItem) HasInventoryType() bool`

HasInventoryType returns a boolean if a field has been set.

### GetLastSeen

`func (o *AnchoreSummaryImageRepoItem) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *AnchoreSummaryImageRepoItem) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *AnchoreSummaryImageRepoItem) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *AnchoreSummaryImageRepoItem) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetContext

`func (o *AnchoreSummaryImageRepoItem) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AnchoreSummaryImageRepoItem) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AnchoreSummaryImageRepoItem) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AnchoreSummaryImageRepoItem) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAlert

`func (o *AnchoreSummaryImageRepoItem) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *AnchoreSummaryImageRepoItem) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *AnchoreSummaryImageRepoItem) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *AnchoreSummaryImageRepoItem) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetAlertId

`func (o *AnchoreSummaryImageRepoItem) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *AnchoreSummaryImageRepoItem) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *AnchoreSummaryImageRepoItem) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.

### HasAlertId

`func (o *AnchoreSummaryImageRepoItem) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### GetWatchTag

`func (o *AnchoreSummaryImageRepoItem) GetWatchTag() bool`

GetWatchTag returns the WatchTag field if non-nil, zero value otherwise.

### GetWatchTagOk

`func (o *AnchoreSummaryImageRepoItem) GetWatchTagOk() (*bool, bool)`

GetWatchTagOk returns a tuple with the WatchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchTag

`func (o *AnchoreSummaryImageRepoItem) SetWatchTag(v bool)`

SetWatchTag sets WatchTag field to given value.

### HasWatchTag

`func (o *AnchoreSummaryImageRepoItem) HasWatchTag() bool`

HasWatchTag returns a boolean if a field has been set.

### GetWatchTagId

`func (o *AnchoreSummaryImageRepoItem) GetWatchTagId() string`

GetWatchTagId returns the WatchTagId field if non-nil, zero value otherwise.

### GetWatchTagIdOk

`func (o *AnchoreSummaryImageRepoItem) GetWatchTagIdOk() (*string, bool)`

GetWatchTagIdOk returns a tuple with the WatchTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchTagId

`func (o *AnchoreSummaryImageRepoItem) SetWatchTagId(v string)`

SetWatchTagId sets WatchTagId field to given value.

### HasWatchTagId

`func (o *AnchoreSummaryImageRepoItem) HasWatchTagId() bool`

HasWatchTagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


