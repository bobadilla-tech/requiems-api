# V1NetworkingWhoisDomainGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**domain** | **string** | The domain name that was looked up | [optional] [default to undefined]
**registrar** | **string** | The name of the registrar holding the domain registration | [optional] [default to undefined]
**name_servers** | **string** | List of authoritative name servers for the domain | [optional] [default to undefined]
**status** | **string** | EPP status codes for the domain (e.g. clientTransferProhibited) | [optional] [default to undefined]
**created_date** | **string** | Date the domain was first registered (ISO 8601) | [optional] [default to undefined]
**updated_date** | **string** | Date the domain record was last updated (ISO 8601) | [optional] [default to undefined]
**expiry_date** | **string** | Date the domain registration expires (ISO 8601) | [optional] [default to undefined]
**dnssec** | **boolean** | True when DNSSEC is enabled for the domain | [optional] [default to undefined]

## Example

```typescript
import { V1NetworkingWhoisDomainGet200ResponseData } from 'api-client';

const instance: V1NetworkingWhoisDomainGet200ResponseData = {
    domain,
    registrar,
    name_servers,
    status,
    created_date,
    updated_date,
    expiry_date,
    dnssec,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
