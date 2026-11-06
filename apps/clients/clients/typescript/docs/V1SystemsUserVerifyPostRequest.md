# V1SystemsUserVerifyPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **string** | Email address to verify. Domain-level WHOIS, MX, and DNS checks are run automatically. | [default to undefined]
**ip_address** | **string** | Optional IPv4 or IPv6 address. When provided, VPN and proxy signals are added to the result. | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsUserVerifyPostRequest } from 'api-client';

const instance: V1SystemsUserVerifyPostRequest = {
    email,
    ip_address,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
