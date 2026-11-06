# OpenapiClient::V1ValidationEmailBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | List of validation results for each email, preserving input order | [optional] |
| **results_email** | **String** | Original email input (null if invalid syntax) | [optional] |
| **results_valid** | **Boolean** | Overall validity (syntax + MX record) | [optional] |
| **results_syntax_valid** | **Boolean** | Whether the email is syntactically valid (RFC 5322) | [optional] |
| **results_mx_valid** | **Boolean** | Whether the domain has valid MX records | [optional] |
| **results_disposable** | **Boolean** | Whether the email comes from a disposable domain | [optional] |
| **results_normalized** | **String** | Canonical normalized email (lowercase, alias handling, etc.) | [optional] |
| **results_domain** | **String** | Extracted domain from email address | [optional] |
| **results_suggestion** | **String** | Suggested correction for common domain typos | [optional] |
| **total** | **Integer** | Number of emails processed in the batch | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1ValidationEmailBatchPost200ResponseData.new(
  results: null,
  results_email: null,
  results_valid: null,
  results_syntax_valid: null,
  results_mx_valid: null,
  results_disposable: null,
  results_normalized: null,
  results_domain: null,
  results_suggestion: null,
  total: null
)
```

