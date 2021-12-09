# SourceRepositoryMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A system-assigned identifier unique for each source analysis | [optional] 
**Url** | **string** | The url hosting the source tree (e.g. https://github.com/anchore/syft.git) | [optional] 
**Protocol** | **string** | The type of source control system used (e.g. git, subversion, mercurial) | [optional] 
**Revision** | **string** | The revision of the source tree analyzed (e.g. commit hash in git) | [optional] 
**Metadata** | [**SourceRepositoryMetadataMetadata**](SourceRepositoryMetadata_metadata.md) |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


