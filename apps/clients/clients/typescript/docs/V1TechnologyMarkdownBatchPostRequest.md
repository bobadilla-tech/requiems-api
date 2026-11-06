# V1TechnologyMarkdownBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**markdowns** | **Array&lt;any&gt;** | Array of Markdown strings. Min: 1, Max: 50. | [default to undefined]
**sanitize** | **boolean** | When true, sanitizes HTML output to remove unsafe tags and attributes. | [optional] [default to undefined]

## Example

```typescript
import { V1TechnologyMarkdownBatchPostRequest } from 'api-client';

const instance: V1TechnologyMarkdownBatchPostRequest = {
    markdowns,
    sanitize,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
