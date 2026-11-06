# V1TechnologyUseragentBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | List of parsed user agents results returned in the same order as the input user agents. | [optional] 
**results_user_agent** | **str** | The original user agent string from the input. | [optional] 
**results_data** | **object** | Parsed user agent details. | [optional] 
**results_data_browser** | **str** | Detected browser name (e.g. Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other). Empty if not detected. | [optional] 
**results_data_browser_version** | **str** | Detected browser version (major.minor). Empty if not detected. | [optional] 
**results_data_os** | **str** | Detected operating system (e.g. Windows, macOS, Linux, Android, iOS, ChromeOS, Other). Empty if not detected. | [optional] 
**results_data_os_version** | **str** | Detected OS version (format varies by platform). Empty if not detected. | [optional] 
**results_data_device** | **str** | Device type — one of desktop, mobile, tablet, bot, or unknown | [optional] 
**results_data_is_bot** | **bool** | True when the user agent matches a known bot or crawler pattern | [optional] 
**total** | **int** | Total number of parsed user agent results returned. | [optional] 

## Example

```python
from openapi_client.models.v1_technology_useragent_batch_post200_response_data import V1TechnologyUseragentBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyUseragentBatchPost200ResponseData from a JSON string
v1_technology_useragent_batch_post200_response_data_instance = V1TechnologyUseragentBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyUseragentBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_technology_useragent_batch_post200_response_data_dict = v1_technology_useragent_batch_post200_response_data_instance.to_dict()
# create an instance of V1TechnologyUseragentBatchPost200ResponseData from a dict
v1_technology_useragent_batch_post200_response_data_from_dict = V1TechnologyUseragentBatchPost200ResponseData.from_dict(v1_technology_useragent_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


