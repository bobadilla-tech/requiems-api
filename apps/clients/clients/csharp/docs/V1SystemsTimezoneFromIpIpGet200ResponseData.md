# Org.OpenAPITools.Model.V1SystemsTimezoneFromIpIpGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | The IP address as provided | [optional] 
**City** | **string** | City resolved from the IP | [optional] 
**CountryCode** | **string** | ISO 3166-1 alpha-2 country code | [optional] 
**Timezone** | **string** | IANA timezone identifier. Null when timezone resolution fails. | [optional] 
**UtcOffset** | **string** | UTC offset string. Null when timezone resolution fails. | [optional] 
**DstActive** | **string** | True when daylight saving time is currently active. Null when timezone resolution fails. | [optional] 
**CurrentTime** | **string** | Current local time in RFC3339 format | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

