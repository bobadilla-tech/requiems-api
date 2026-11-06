# V1TechnologyMarkdownPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**markdown** | **str** | The Markdown text to convert. | 
**sanitize** | **bool** | When true, sanitizes the HTML output to remove unsafe tags and attributes. | [optional] 

## Example

```python
from openapi_client.models.v1_technology_markdown_post_request import V1TechnologyMarkdownPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyMarkdownPostRequest from a JSON string
v1_technology_markdown_post_request_instance = V1TechnologyMarkdownPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyMarkdownPostRequest.to_json())

# convert the object into a dict
v1_technology_markdown_post_request_dict = v1_technology_markdown_post_request_instance.to_dict()
# create an instance of V1TechnologyMarkdownPostRequest from a dict
v1_technology_markdown_post_request_from_dict = V1TechnologyMarkdownPostRequest.from_dict(v1_technology_markdown_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


