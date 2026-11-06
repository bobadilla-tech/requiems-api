# V1NetworkingMxDomainGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domain** | **string** | The domain that was queried | [optional] [default to undefined]
**records** | **Array&lt;any&gt;** | List of MX records, sorted by priority ascending (lowest number &#x3D; highest priority) | [optional] [default to undefined]
**records_host** | **string** | Fully-qualified hostname of the mail server (typically ends with a trailing dot) | [optional] [default to undefined]
**records_priority** | **number** | MX priority value. Lower values have higher delivery priority per RFC 5321. | [optional] [default to undefined]

## Example

```typescript
import { V1NetworkingMxDomainGet200ResponseData } from 'api-client';

const instance: V1NetworkingMxDomainGet200ResponseData = {
    domain,
    records,
    records_host,
    records_priority,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
