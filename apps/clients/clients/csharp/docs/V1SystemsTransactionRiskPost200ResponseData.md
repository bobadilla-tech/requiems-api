# Org.OpenAPITools.Model.V1SystemsTransactionRiskPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskScore** | **string** | Transaction risk score from 0.0 to 1.0 | [optional] 
**IsSafe** | **bool** | True when score is below 0.5 and no TOR or proxy detected | [optional] 
**Flags** | **List&lt;Object&gt;** | Triggered flags. Possible values: ip_country_mismatch, bin_country_mismatch, country_mismatch, vpn_detected, proxy_detected, tor_detected, high_value_vpn | [optional] 
**SignalsIpCountry** | **string** | Country code resolved from the IP address | [optional] 
**SignalsBillingCountry** | **string** | Billing country provided in the request, or null if omitted | [optional] 
**SignalsBinCountry** | **string** | Country of the card issuer resolved from the BIN | [optional] 
**SignalsCountryMismatch** | **bool** | True when any country comparison produced a mismatch | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

