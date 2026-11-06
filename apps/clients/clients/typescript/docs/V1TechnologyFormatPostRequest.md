# V1TechnologyFormatPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**from** | **string** | Source format. One of: json, yaml, csv, xml, toml | [default to undefined]
**to** | **string** | Target format. One of: json, yaml, csv, xml, toml | [default to undefined]
**content** | **string** | The content to convert, serialized as a string in the source format. | [default to undefined]

## Example

```typescript
import { V1TechnologyFormatPostRequest } from 'api-client';

const instance: V1TechnologyFormatPostRequest = {
    from,
    to,
    content,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
