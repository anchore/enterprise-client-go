# AlertSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Identifier for the alert | [optional] 
**Type** | **string** | Type of the alert | [optional] 
**State** | **string** | Current state of the alert | [optional] 
**ResourceLabels** | [**[]ResourceLabel**](ResourceLabel.md) |  | [optional] 
**ClosedBy** | **string** | Account that closed the alert | [optional] 
**ClosedReason** | **string** | Reason for closing the alert | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | RFC 3339 formatted UTC timestamp when the alert was generated | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) | RFC 3339 formatted UTC timestamp when the alert was last modified | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


