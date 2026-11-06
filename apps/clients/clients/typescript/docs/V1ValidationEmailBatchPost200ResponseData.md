# V1ValidationEmailBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | List of validation results for each email, preserving input order | [optional] [default to undefined]
**results_email** | **string** | Original email input (null if invalid syntax) | [optional] [default to undefined]
**results_valid** | **boolean** | Overall validity (syntax + MX record) | [optional] [default to undefined]
**results_syntax_valid** | **boolean** | Whether the email is syntactically valid (RFC 5322) | [optional] [default to undefined]
**results_mx_valid** | **boolean** | Whether the domain has valid MX records | [optional] [default to undefined]
**results_disposable** | **boolean** | Whether the email comes from a disposable domain | [optional] [default to undefined]
**results_normalized** | **string** | Canonical normalized email (lowercase, alias handling, etc.) | [optional] [default to undefined]
**results_domain** | **string** | Extracted domain from email address | [optional] [default to undefined]
**results_suggestion** | **string** | Suggested correction for common domain typos | [optional] [default to undefined]
**total** | **number** | Number of emails processed in the batch | [optional] [default to undefined]

## Example

```typescript
import { V1ValidationEmailBatchPost200ResponseData } from 'api-client';

const instance: V1ValidationEmailBatchPost200ResponseData = {
    results,
    results_email,
    results_valid,
    results_syntax_valid,
    results_mx_valid,
    results_disposable,
    results_normalized,
    results_domain,
    results_suggestion,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
