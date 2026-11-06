# V1TechnologyBaseBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **string** | List of conversion requests (1–50 items). | [default to undefined]
**items_from** | **number** | Source base (2, 8, 10, or 16). | [default to undefined]
**items_to** | **number** | Target base (2, 8, 10, or 16). | [default to undefined]
**items_value** | **string** | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal). | [default to undefined]

## Example

```typescript
import { V1TechnologyBaseBatchPostRequest } from 'api-client';

const instance: V1TechnologyBaseBatchPostRequest = {
    items,
    items_from,
    items_to,
    items_value,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
