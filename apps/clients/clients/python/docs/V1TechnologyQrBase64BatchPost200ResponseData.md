# V1TechnologyQrBase64BatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | Per-item results in input order. | [optional] 
**results_data** | **str** | The input text or URL that was encoded. | [optional] 
**results_image** | **str** | Base64-encoded PNG image data (omitted on error). | [optional] 
**results_width** | **int** | Width of the generated image in pixels (omitted on error). | [optional] 
**results_height** | **int** | Height of the generated image in pixels (omitted on error). | [optional] 
**results_error** | **str** | Error message if generation failed (omitted on success). | [optional] 
**total** | **int** | Total number of items processed. | [optional] 

## Example

```python
from openapi_client.models.v1_technology_qr_base64_batch_post200_response_data import V1TechnologyQrBase64BatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyQrBase64BatchPost200ResponseData from a JSON string
v1_technology_qr_base64_batch_post200_response_data_instance = V1TechnologyQrBase64BatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyQrBase64BatchPost200ResponseData.to_json())

# convert the object into a dict
v1_technology_qr_base64_batch_post200_response_data_dict = v1_technology_qr_base64_batch_post200_response_data_instance.to_dict()
# create an instance of V1TechnologyQrBase64BatchPost200ResponseData from a dict
v1_technology_qr_base64_batch_post200_response_data_from_dict = V1TechnologyQrBase64BatchPost200ResponseData.from_dict(v1_technology_qr_base64_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


