# Org.OpenAPITools.Model.V1SystemsSignupProtectPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskScore** | **string** | Composite risk score from 0.0 (clean) to 1.0 (high risk) | [optional] 
**IsSafe** | **bool** | True when risk_score is below the safety threshold with no critical flags | [optional] 
**Confidence** | **string** | Confidence in the decision from 0.0 to 1.0, based on which signals resolved successfully | [optional] 
**Flags** | **List&lt;Object&gt;** | Array of triggered risk flags. Possible values: disposable_email, vpn_detected, proxy_detected, tor_detected, is_hosting | [optional] 
**SignalsEmail** | **Object** | Email validation result including disposable detection and MX check | [optional] 
**SignalsPhone** | **Object** | Phone number validation result with country code and VoIP/virtual detection | [optional] 
**SignalsIp** | **Object** | IP intelligence including VPN/proxy/TOR status and fraud score | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

