# V1TechnologyQrBase64BatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **string** | List of QR generation requests (1–50 items). | [default to undefined]
**items_data** | **string** | The text or URL to encode. | [default to undefined]
**items_size** | **number** | Image size in pixels (50–1000). Defaults to 256. | [optional] [default to undefined]
**items_recovery** | **string** | Error-correction level: low, medium, high, highest. Defaults to medium. | [optional] [default to undefined]

## Example

```typescript
import { V1TechnologyQrBase64BatchPostRequest } from 'api-client';

const instance: V1TechnologyQrBase64BatchPostRequest = {
    items,
    items_data,
    items_size,
    items_recovery,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
