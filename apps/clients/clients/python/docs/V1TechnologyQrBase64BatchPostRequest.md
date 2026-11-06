# V1TechnologyQrBase64BatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **str** | List of QR generation requests (1–50 items). | 
**items_data** | **str** | The text or URL to encode. | 
**items_size** | **int** | Image size in pixels (50–1000). Defaults to 256. | [optional] 
**items_recovery** | **str** | Error-correction level: low, medium, high, highest. Defaults to medium. | [optional] 

## Example

```python
from openapi_client.models.v1_technology_qr_base64_batch_post_request import V1TechnologyQrBase64BatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyQrBase64BatchPostRequest from a JSON string
v1_technology_qr_base64_batch_post_request_instance = V1TechnologyQrBase64BatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyQrBase64BatchPostRequest.to_json())

# convert the object into a dict
v1_technology_qr_base64_batch_post_request_dict = v1_technology_qr_base64_batch_post_request_instance.to_dict()
# create an instance of V1TechnologyQrBase64BatchPostRequest from a dict
v1_technology_qr_base64_batch_post_request_from_dict = V1TechnologyQrBase64BatchPostRequest.from_dict(v1_technology_qr_base64_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


