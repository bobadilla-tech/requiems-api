# Org.OpenAPITools.Model.V1SystemsTransactionRiskPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardBin** | **string** | Card BIN (first 6–8 digits) to look up card issuing country. | 
**IpAddress** | **string** | Customer&#39;s IPv4 or IPv6 address. Must be a valid IP. | 
**BillingCountry** | **string** | ISO 3166-1 alpha-2 billing country code. When provided, country mismatch checks are more precise. | [optional] 
**AmountUsd** | **decimal** | Transaction amount in USD. High-value transactions via VPN incur an additional risk score boost. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

