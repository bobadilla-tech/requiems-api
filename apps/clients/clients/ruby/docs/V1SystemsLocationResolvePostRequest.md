# OpenapiClient::V1SystemsLocationResolvePostRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **address** | **String** | Human-readable address string to geocode. Either address or coordinates is required. | [optional] |
| **coordinates** | **Object** | Coordinates object for reverse geocoding. Either address or coordinates is required. Shape: {lat: float, lng: float} | [optional] |
| **country_code** | **String** | ISO 3166-1 alpha-2 country code hint. Overrides the geocoder-inferred country code when provided. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1SystemsLocationResolvePostRequest.new(
  address: Alexanderplatz 1, Berlin,
  coordinates: {&quot;lat&quot;: 52.52, &quot;lng&quot;: 13.405},
  country_code: DE
)
```

