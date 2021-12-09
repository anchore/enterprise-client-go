# SamlConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name to use for referencing this IDP configuration. This will configured as part of the url string the Idp must have the client POST the saml assertion to. | 
**Enabled** | **bool** | If this IDP configuration should be enabled for user logins | 
**SpEntityId** | **string** | The entity ID for this SP. Can be the same for all IDP configurations in this installation or unique to each. This is typically a URL, but you can use any value as long as you also configure the IDP to expect this value. | 
**AcsUrl** | **string** | The URL the IDP can use to access the Assertion Consumer Service to provide the token for sso. This is the way to reach the rbac manager services /v1/saml/sso/{IDP_name} route externally | 
**AcsHttpsPort** | **int32** | The port number to use for https if not 443. If omitted or -1, 443 is assumed and used as a default | [optional] 
**IdpMetadataUrl** | **string** | The url where the SP (anchore) can retrieve the metadata about the Identity Provider. Only one of this or metadata_xml should be set. This is typically provided by the IDP. | [optional] 
**IdpMetadataXml** | **string** | The direct metadata xml payload, if a url is not available. Only one of this or metadata_url should be set. | [optional] 
**IdpUsernameAttribute** | **string** | The SAML attribute to use from the response assertions to determine the anchore username. If unset, the subject is used. | [optional] 
**IdpAccountAttribute** | **string** | The SAML attribute to use from the response assertions to determine the anchore account to use. If unset, the default is used. | [optional] 
**IdpRoleAttribute** | **string** | The SAML attribute to use from the response assertions to determine the anchore role(s) to assign a new user in the specified account. If unset, the default is used. | [optional] 
**DefaultAccount** | **string** | The existing anchore account to assign all users to from this IDP if no account attribute is mapped or present. | [optional] 
**DefaultRole** | **string** | The default role to apply to new users from this IDP if no attribute is mapped or found in the SAML assertions. | [optional] 
**RequireSignedAssertions** | **bool** | Require assertions in to be signed from the IDP | [optional] [default to true]
**RequireSignedResponse** | **bool** | Require the authn response to be signed by the IDP | [optional] [default to true]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**LastUpdated** | [**time.Time**](time.Time.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


