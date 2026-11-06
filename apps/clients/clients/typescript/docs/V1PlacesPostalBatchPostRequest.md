# V1PlacesPostalBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **string** | List of postal code lookups (1–50 items). | [default to undefined]
**items_code** | **string** | The postal code to look up. | [default to undefined]
**items_country** | **string** | ISO 3166-1 alpha-2 country code (default: US). | [optional] [default to undefined]

## Example

```typescript
import { V1PlacesPostalBatchPostRequest } from 'api-client';

const instance: V1PlacesPostalBatchPostRequest = {
    items,
    items_code,
    items_country,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
