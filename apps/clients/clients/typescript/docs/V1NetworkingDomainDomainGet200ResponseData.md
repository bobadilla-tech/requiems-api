# V1NetworkingDomainDomainGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domain** | **string** | The domain that was looked up | [optional] [default to undefined]
**available** | **boolean** | True when the domain appears to be unregistered (NS lookup returns NXDOMAIN). False when name servers are delegated. | [optional] [default to undefined]
**dns_a** | **string** | IPv4 addresses (A records) | [optional] [default to undefined]
**dns_aaaa** | **string** | IPv6 addresses (AAAA records) | [optional] [default to undefined]
**dns_mx** | **string** | Mail exchange records, each with host and priority fields | [optional] [default to undefined]
**dns_ns** | **string** | Authoritative name server hostnames | [optional] [default to undefined]
**dns_txt** | **string** | TXT record values (SPF, DKIM, verification tokens, etc.) | [optional] [default to undefined]
**dns_cname** | **string** | CNAME alias target, if the domain is an alias. Empty string when no alias exists. | [optional] [default to undefined]

## Example

```typescript
import { V1NetworkingDomainDomainGet200ResponseData } from 'api-client';

const instance: V1NetworkingDomainDomainGet200ResponseData = {
    domain,
    available,
    dns_a,
    dns_aaaa,
    dns_mx,
    dns_ns,
    dns_txt,
    dns_cname,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
