# V1SystemsBusinessCalendarCountryGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**country_code** | **str** | ISO 3166-1 alpha-2 country code echoed from the path parameter | [optional] 
**year** | **int** | Year the calendar covers | [optional] 
**month** | **str** | Month scope. Present only when the month query parameter was provided. | [optional] 
**working_days** | **int** | Number of working days (excluding weekends and public holidays) in the scope | [optional] 
**total_days** | **str** | Total calendar days in the month. Present only when month is provided. | [optional] 
**weekend_days** | **str** | Number of weekend days in the month. Present only when month is provided. | [optional] 
**holidays** | **List[object]** | List of public holidays in the scope. Each item: {date: YYYY-MM-DD, name: string, type: national} | [optional] 
**holiday_count** | **int** | Count of holidays in the holidays array | [optional] 
**next_holiday** | **str** | Next public holiday within 90 days from today. Shape: {date, name, type} | [optional] 

## Example

```python
from openapi_client.models.v1_systems_business_calendar_country_get200_response_data import V1SystemsBusinessCalendarCountryGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsBusinessCalendarCountryGet200ResponseData from a JSON string
v1_systems_business_calendar_country_get200_response_data_instance = V1SystemsBusinessCalendarCountryGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1SystemsBusinessCalendarCountryGet200ResponseData.to_json())

# convert the object into a dict
v1_systems_business_calendar_country_get200_response_data_dict = v1_systems_business_calendar_country_get200_response_data_instance.to_dict()
# create an instance of V1SystemsBusinessCalendarCountryGet200ResponseData from a dict
v1_systems_business_calendar_country_get200_response_data_from_dict = V1SystemsBusinessCalendarCountryGet200ResponseData.from_dict(v1_systems_business_calendar_country_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


