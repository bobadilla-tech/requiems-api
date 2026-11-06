# OpenapiClient::V1NetworkingIpAsnBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | Per-item results in input order. | [optional] |
| **results_ip** | **String** | The IP address that was looked up. | [optional] |
| **results_result** | **Object** | ASN data (omitted on error). Same fields as the single-item endpoint. | [optional] |
| **results_error** | **String** | Error message if the lookup failed (omitted on success). | [optional] |
| **total** | **Integer** | Total number of items processed. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1NetworkingIpAsnBatchPost200ResponseData.new(
  results: null,
  results_ip: null,
  results_result: null,
  results_error: null,
  total: null
)
```

