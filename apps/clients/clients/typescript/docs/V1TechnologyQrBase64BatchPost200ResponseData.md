# V1TechnologyQrBase64BatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | Per-item results in input order. | [optional] [default to undefined]
**results_data** | **string** | The input text or URL that was encoded. | [optional] [default to undefined]
**results_image** | **string** | Base64-encoded PNG image data (omitted on error). | [optional] [default to undefined]
**results_width** | **number** | Width of the generated image in pixels (omitted on error). | [optional] [default to undefined]
**results_height** | **number** | Height of the generated image in pixels (omitted on error). | [optional] [default to undefined]
**results_error** | **string** | Error message if generation failed (omitted on success). | [optional] [default to undefined]
**total** | **number** | Total number of items processed. | [optional] [default to undefined]

## Example

```typescript
import { V1TechnologyQrBase64BatchPost200ResponseData } from 'api-client';

const instance: V1TechnologyQrBase64BatchPost200ResponseData = {
    results,
    results_data,
    results_image,
    results_width,
    results_height,
    results_error,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
