# V1TechnologyUseragentGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**browser** | **str** | Detected browser name (e.g. Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other) | [optional] 
**browser_version** | **str** | Detected browser version (major.minor) | [optional] 
**os** | **str** | Detected operating system (e.g. Windows, macOS, Linux, Android, iOS, ChromeOS, Other) | [optional] 
**os_version** | **str** | Detected OS version (format varies by platform) | [optional] 
**device** | **str** | Device type — one of desktop, mobile, tablet, bot, or unknown | [optional] 
**is_bot** | **bool** | True when the user agent matches a known bot or crawler pattern | [optional] 

## Example

```python
from openapi_client.models.v1_technology_useragent_get200_response_data import V1TechnologyUseragentGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyUseragentGet200ResponseData from a JSON string
v1_technology_useragent_get200_response_data_instance = V1TechnologyUseragentGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyUseragentGet200ResponseData.to_json())

# convert the object into a dict
v1_technology_useragent_get200_response_data_dict = v1_technology_useragent_get200_response_data_instance.to_dict()
# create an instance of V1TechnologyUseragentGet200ResponseData from a dict
v1_technology_useragent_get200_response_data_from_dict = V1TechnologyUseragentGet200ResponseData.from_dict(v1_technology_useragent_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


