# V1NetworkingIpAsnGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip** | **string** | The IP address that was looked up (the requesting client\&#39;s IP) | [optional] [default to undefined]
**asn** | **string** | Autonomous System Number in format \&quot;ASxxxx\&quot; (e.g., \&quot;AS15169\&quot;) | [optional] [default to undefined]
**org** | **string** | Organization name owning the IP address range | [optional] [default to undefined]
**isp** | **string** | Internet Service Provider providing the IP | [optional] [default to undefined]
**domain** | **string** | Domain name associated with the IP or IP range | [optional] [default to undefined]
**route** | **string** | CIDR notation of the network route (e.g., \&quot;8.8.8.0/24\&quot;) | [optional] [default to undefined]
**type** | **string** | Type of network (e.g., \&quot;hosting\&quot;, \&quot;isp\&quot;, \&quot;business\&quot;, \&quot;cdn\&quot;) | [optional] [default to undefined]

## Example

```typescript
import { V1NetworkingIpAsnGet200ResponseData } from 'api-client';

const instance: V1NetworkingIpAsnGet200ResponseData = {
    ip,
    asn,
    org,
    isp,
    domain,
    route,
    type,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
