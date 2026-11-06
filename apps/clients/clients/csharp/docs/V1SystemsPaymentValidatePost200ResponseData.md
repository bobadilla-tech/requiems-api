# Org.OpenAPITools.Model.V1SystemsPaymentValidatePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | **string** | BIN result. Null when bin was not provided. Contains valid, scheme, card_type, card_level, country_code, issuer, prepaid, luhn. | [optional] 
**Iban** | **string** | IBAN result. Null when iban was not provided. Contains valid, country_code, bank_code, account_number. | [optional] 
**Swift** | **string** | SWIFT result. Null when swift was not provided. Contains valid, institution, country, branch. | [optional] 
**ConsistencyOk** | **bool** | True when no cross-field mismatches were detected | [optional] 
**ConsistencyFlags** | **List&lt;Object&gt;** | Mismatch flags. Possible values: country_mismatch_bin_iban, country_mismatch_bin_swift, country_mismatch_iban_swift, bank_mismatch_iban_swift | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

