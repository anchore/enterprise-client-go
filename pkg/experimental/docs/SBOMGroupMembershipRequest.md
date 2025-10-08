# SBOMGroupMembershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoveAll** | Pointer to **bool** | Setting this to True will cause ALL SBOMs to be removed from the group. Cannot be used in combination with the add_sboms or remove_sboms properties. | [optional] 
**AddSboms** | Pointer to **[]string** | A list of SBOM UUIDs to be added to the Group | [optional] 
**RemoveSboms** | Pointer to **[]string** | A list of SBOM UUIDs to be removed from the Group | [optional] 

## Methods

### NewSBOMGroupMembershipRequest

`func NewSBOMGroupMembershipRequest() *SBOMGroupMembershipRequest`

NewSBOMGroupMembershipRequest instantiates a new SBOMGroupMembershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSBOMGroupMembershipRequestWithDefaults

`func NewSBOMGroupMembershipRequestWithDefaults() *SBOMGroupMembershipRequest`

NewSBOMGroupMembershipRequestWithDefaults instantiates a new SBOMGroupMembershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoveAll

`func (o *SBOMGroupMembershipRequest) GetRemoveAll() bool`

GetRemoveAll returns the RemoveAll field if non-nil, zero value otherwise.

### GetRemoveAllOk

`func (o *SBOMGroupMembershipRequest) GetRemoveAllOk() (*bool, bool)`

GetRemoveAllOk returns a tuple with the RemoveAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAll

`func (o *SBOMGroupMembershipRequest) SetRemoveAll(v bool)`

SetRemoveAll sets RemoveAll field to given value.

### HasRemoveAll

`func (o *SBOMGroupMembershipRequest) HasRemoveAll() bool`

HasRemoveAll returns a boolean if a field has been set.

### GetAddSboms

`func (o *SBOMGroupMembershipRequest) GetAddSboms() []string`

GetAddSboms returns the AddSboms field if non-nil, zero value otherwise.

### GetAddSbomsOk

`func (o *SBOMGroupMembershipRequest) GetAddSbomsOk() (*[]string, bool)`

GetAddSbomsOk returns a tuple with the AddSboms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSboms

`func (o *SBOMGroupMembershipRequest) SetAddSboms(v []string)`

SetAddSboms sets AddSboms field to given value.

### HasAddSboms

`func (o *SBOMGroupMembershipRequest) HasAddSboms() bool`

HasAddSboms returns a boolean if a field has been set.

### GetRemoveSboms

`func (o *SBOMGroupMembershipRequest) GetRemoveSboms() []string`

GetRemoveSboms returns the RemoveSboms field if non-nil, zero value otherwise.

### GetRemoveSbomsOk

`func (o *SBOMGroupMembershipRequest) GetRemoveSbomsOk() (*[]string, bool)`

GetRemoveSbomsOk returns a tuple with the RemoveSboms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveSboms

`func (o *SBOMGroupMembershipRequest) SetRemoveSboms(v []string)`

SetRemoveSboms sets RemoveSboms field to given value.

### HasRemoveSboms

`func (o *SBOMGroupMembershipRequest) HasRemoveSboms() bool`

HasRemoveSboms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


