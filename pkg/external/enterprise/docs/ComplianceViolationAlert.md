# ComplianceViolationAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Identifier for the alert | [optional] 
**Type** | **string** | Type of alert generated | [optional] 
**State** | **string** | Current state of the alert | [optional] 
**Resource** | [**ComplianceResource**](ComplianceResource.md) |  | [optional] 
**ClosedBy** | **string** | Account that closed the alert | [optional] 
**ClosedReason** | **string** | Reason for closing the alert | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | RFC 3339 formatted UTC timestamp when the alert was generated | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) | RFC 3339 formatted UTC timestamp when the alert was last modified | [optional] 
**ComplianceStatusReason** | **string** | Reason for compliance check status. Compliance check could fail due to policy evaluation or blacklisting or errors evaluating compliance | [optional] 
**ViolationsCount** | **int32** | Number of STOP action results in the compliance check report | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


