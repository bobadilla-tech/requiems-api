# V1TechnologyPasswordBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **string** | List of password generation requests (1–50 items). | [default to undefined]
**items_length** | **number** | Password length (8–128). Defaults to 16 when omitted. | [optional] [default to undefined]
**items_uppercase** | **boolean** | Include uppercase letters (A–Z). Defaults to false. | [optional] [default to undefined]
**items_numbers** | **boolean** | Include numbers (0–9). Defaults to false. | [optional] [default to undefined]
**items_symbols** | **boolean** | Include special characters. Defaults to false. | [optional] [default to undefined]

## Example

```typescript
import { V1TechnologyPasswordBatchPostRequest } from 'api-client';

const instance: V1TechnologyPasswordBatchPostRequest = {
    items,
    items_length,
    items_uppercase,
    items_numbers,
    items_symbols,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
