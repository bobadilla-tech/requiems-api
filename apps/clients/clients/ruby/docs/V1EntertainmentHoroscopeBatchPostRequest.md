# OpenapiClient::V1EntertainmentHoroscopeBatchPostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **signs** | **Array&lt;Object&gt;** | Array of signs to get (min: 1, max: 12). Allowed values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1EntertainmentHoroscopeBatchPostRequest.new(
  signs: [&quot;taurus&quot;,&quot;cancer&quot;,&quot;libra&quot;]

)
```

