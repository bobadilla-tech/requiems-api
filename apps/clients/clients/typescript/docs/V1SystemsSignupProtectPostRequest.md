# V1SystemsSignupProtectPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **string** | Email address to validate and score. At least one of email, phone, or ip_address is required. | [optional] [default to undefined]
**phone** | **string** | Phone number in E.164 format to validate and check for VoIP or virtual numbers. | [optional] [default to undefined]
**ip_address** | **string** | IPv4 or IPv6 address to check for VPN, proxy, TOR, and hosting status. | [optional] [default to undefined]

## Example

```typescript
import { V1SystemsSignupProtectPostRequest } from 'api-client';

const instance: V1SystemsSignupProtectPostRequest = {
    email,
    phone,
    ip_address,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
