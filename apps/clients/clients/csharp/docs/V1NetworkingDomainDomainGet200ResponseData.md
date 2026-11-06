# Org.OpenAPITools.Model.V1NetworkingDomainDomainGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain that was looked up | [optional] 
**Available** | **bool** | True when the domain appears to be unregistered (NS lookup returns NXDOMAIN). False when name servers are delegated. | [optional] 
**DnsA** | **string** | IPv4 addresses (A records) | [optional] 
**DnsAaaa** | **string** | IPv6 addresses (AAAA records) | [optional] 
**DnsMx** | **string** | Mail exchange records, each with host and priority fields | [optional] 
**DnsNs** | **string** | Authoritative name server hostnames | [optional] 
**DnsTxt** | **string** | TXT record values (SPF, DKIM, verification tokens, etc.) | [optional] 
**DnsCname** | **string** | CNAME alias target, if the domain is an alias. Empty string when no alias exists. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

