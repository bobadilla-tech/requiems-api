# V1NetworkingIpGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip** | **string** | The IP address that was looked up (the requesting client\&#39;s IP) | [optional] [default to undefined]
**country** | **string** | Country name where the IP is located | [optional] [default to undefined]
**country_code** | **string** | Two-letter ISO country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;DE\&quot;) | [optional] [default to undefined]
**city** | **string** | City name where the IP is located | [optional] [default to undefined]
**isp** | **string** | Internet Service Provider providing the IP | [optional] [default to undefined]
**is_vpn** | **boolean** | True when the IP belongs to a known VPN | [optional] [default to undefined]

## Example

```typescript
import { V1NetworkingIpGet200ResponseData } from 'api-client';

const instance: V1NetworkingIpGet200ResponseData = {
    ip,
    country,
    country_code,
    city,
    isp,
    is_vpn,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
