# SmtpEndpointConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The instance identifier for the configuration | [optional] 
**Description** | **string** | User friendly name or description for the configuration | [optional] 
**VerifyTls** | **bool** | Verify the cert if using tls for connecting externally. Defaults to true if not specified | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp for last modification to the record | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) | Timestamp for last modification to the record | [optional] 
**Host** | **string** |  | 
**Port** | **int32** |  | 
**Username** | **string** |  | [optional] 
**Password** | **string** |  | [optional] 
**UseTls** | **bool** | Encrypt the SMTP connection with TLS. Defaults to true | [optional] 
**From** | **string** | The from address to use for emails send by this configuration | 
**To** | **string** | The address to which the emails are sent | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


