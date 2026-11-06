# V1SystemsTimezoneFromIpIpGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip** | **str** | The IP address as provided | [optional] 
**city** | **str** | City resolved from the IP | [optional] 
**country_code** | **str** | ISO 3166-1 alpha-2 country code | [optional] 
**timezone** | **str** | IANA timezone identifier. Null when timezone resolution fails. | [optional] 
**utc_offset** | **str** | UTC offset string. Null when timezone resolution fails. | [optional] 
**dst_active** | **str** | True when daylight saving time is currently active. Null when timezone resolution fails. | [optional] 
**current_time** | **str** | Current local time in RFC3339 format | [optional] 

## Example

```python
from openapi_client.models.v1_systems_timezone_from_ip_ip_get200_response_data import V1SystemsTimezoneFromIpIpGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsTimezoneFromIpIpGet200ResponseData from a JSON string
v1_systems_timezone_from_ip_ip_get200_response_data_instance = V1SystemsTimezoneFromIpIpGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1SystemsTimezoneFromIpIpGet200ResponseData.to_json())

# convert the object into a dict
v1_systems_timezone_from_ip_ip_get200_response_data_dict = v1_systems_timezone_from_ip_ip_get200_response_data_instance.to_dict()
# create an instance of V1SystemsTimezoneFromIpIpGet200ResponseData from a dict
v1_systems_timezone_from_ip_ip_get200_response_data_from_dict = V1SystemsTimezoneFromIpIpGet200ResponseData.from_dict(v1_systems_timezone_from_ip_ip_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


