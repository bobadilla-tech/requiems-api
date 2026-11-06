# Org.OpenAPITools.Model.V1SystemsUserVerifyPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verified** | **bool** | True when risk score is below 0.3, confidence is above 0.5, email is valid, domain has MX, and domain is registered | [optional] 
**Confidence** | **string** | Confidence in the verification result, based on how many signals resolved | [optional] 
**RiskScore** | **string** | Composite risk score from 0.0 to 1.0 | [optional] 
**Flags** | **List&lt;Object&gt;** | Triggered risk flags. Possible values: email_invalid, disposable_email, no_mx, domain_not_registered, young_domain, whois_unavailable, ip_risk | [optional] 
**SignalsEmail** | **Object** | Email validation result | [optional] 
**SignalsDomain** | **Object** | Domain intelligence including WHOIS age, MX and A record presence, and registration status | [optional] 
**SignalsIp** | **string** | IP risk signals. Null when ip_address was not provided | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

