# JiraEndpointConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The instance identifier for the configuration | [optional] 
**Description** | **string** | User friendly name or description for the configuration | [optional] 
**VerifyTls** | **bool** | Verify the cert if using tls for connecting externally. Defaults to true if not specified | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp for last modification to the record | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) | Timestamp for last modification to the record | [optional] 
**Url** | **string** | Jira endpoint URL including host and port, should begin with &#39;http://&#39; or &#39;https://&#39; | 
**Username** | **string** | Jira username for creating issues | 
**Password** | **string** | Jira password for creating issues | 
**ProjectKey** | **string** | Key of the Jira project for creating issues | 
**IssueType** | **string** | Type associated with the issue | 
**Priority** | **string** | Priority assigned to the issue | [optional] 
**Assignee** | **string** | Jira user to associate with the issue | [optional] 
**Labels** | **[]string** | List of labels to associate with the issue | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


