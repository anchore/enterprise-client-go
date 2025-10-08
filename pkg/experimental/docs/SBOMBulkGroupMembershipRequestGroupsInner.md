# SBOMBulkGroupMembershipRequestGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupUuid** | **string** |  | 
**RemoveAll** | Pointer to **bool** | Setting this to True will cause ALL SBOMs to be removed from the group. Cannot be used in combination with the add_sboms or remove_sboms properties. | [optional] 
**AddSboms** | Pointer to **[]string** | A list of SBOM UUIDs to be added to the Group | [optional] 
**RemoveSboms** | Pointer to **[]string** | A list of SBOM UUIDs to be removed from the Group | [optional] 

## Methods

### NewSBOMBulkGroupMembershipRequestGroupsInner

`func NewSBOMBulkGroupMembershipRequestGroupsInner(groupUuid string, ) *SBOMBulkGroupMembershipRequestGroupsInner`

NewSBOMBulkGroupMembershipRequestGroupsInner instantiates a new SBOMBulkGroupMembershipRequestGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSBOMBulkGroupMembershipRequestGroupsInnerWithDefaults

`func NewSBOMBulkGroupMembershipRequestGroupsInnerWithDefaults() *SBOMBulkGroupMembershipRequestGroupsInner`

NewSBOMBulkGroupMembershipRequestGroupsInnerWithDefaults instantiates a new SBOMBulkGroupMembershipRequestGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupUuid

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) GetGroupUuid() string`

GetGroupUuid returns the GroupUuid field if non-nil, zero value otherwise.

### GetGroupUuidOk

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) GetGroupUuidOk() (*string, bool)`

GetGroupUuidOk returns a tuple with the GroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUuid

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) SetGroupUuid(v string)`

SetGroupUuid sets GroupUuid field to given value.


### GetRemoveAll

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) GetRemoveAll() bool`

GetRemoveAll returns the RemoveAll field if non-nil, zero value otherwise.

### GetRemoveAllOk

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) GetRemoveAllOk() (*bool, bool)`

GetRemoveAllOk returns a tuple with the RemoveAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAll

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) SetRemoveAll(v bool)`

SetRemoveAll sets RemoveAll field to given value.

### HasRemoveAll

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) HasRemoveAll() bool`

HasRemoveAll returns a boolean if a field has been set.

### GetAddSboms

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) GetAddSboms() []string`

GetAddSboms returns the AddSboms field if non-nil, zero value otherwise.

### GetAddSbomsOk

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) GetAddSbomsOk() (*[]string, bool)`

GetAddSbomsOk returns a tuple with the AddSboms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSboms

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) SetAddSboms(v []string)`

SetAddSboms sets AddSboms field to given value.

### HasAddSboms

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) HasAddSboms() bool`

HasAddSboms returns a boolean if a field has been set.

### GetRemoveSboms

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) GetRemoveSboms() []string`

GetRemoveSboms returns the RemoveSboms field if non-nil, zero value otherwise.

### GetRemoveSbomsOk

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) GetRemoveSbomsOk() (*[]string, bool)`

GetRemoveSbomsOk returns a tuple with the RemoveSboms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveSboms

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) SetRemoveSboms(v []string)`

SetRemoveSboms sets RemoveSboms field to given value.

### HasRemoveSboms

`func (o *SBOMBulkGroupMembershipRequestGroupsInner) HasRemoveSboms() bool`

HasRemoveSboms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


