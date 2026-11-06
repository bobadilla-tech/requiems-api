# Org.OpenAPITools.Model.V1NetworkingWhoisDomainGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain name that was looked up | [optional] 
**Registrar** | **string** | The name of the registrar holding the domain registration | [optional] 
**NameServers** | **string** | List of authoritative name servers for the domain | [optional] 
**Status** | **string** | EPP status codes for the domain (e.g. clientTransferProhibited) | [optional] 
**CreatedDate** | **string** | Date the domain was first registered (ISO 8601) | [optional] 
**UpdatedDate** | **string** | Date the domain record was last updated (ISO 8601) | [optional] 
**ExpiryDate** | **string** | Date the domain registration expires (ISO 8601) | [optional] 
**Dnssec** | **bool** | True when DNSSEC is enabled for the domain | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

