# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | The id of the application | [optional] 
**Name** | **string** | The name of the application | 
**Description** | **string** | The description of the application | [optional] 
**Version** | **string** | The version of the application | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | RFC 3339 formatted UTC timestamp when the application was created | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) | RFC 3339 formatted UTC timestamp when the application was last updated | [optional] 
**SourceSbomIds** | **[]string** | The list of repo artifact ids associated with this application | [optional] 
**ImageSbomIds** | **[]string** | The list of image artifact ids associated with this application | [optional] 
**ContainerSbomIds** | **[]string** | The list of container artifact ids associated with this application | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


