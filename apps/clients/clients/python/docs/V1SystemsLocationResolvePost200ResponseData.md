# V1SystemsLocationResolvePost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** | Normalized full address returned by the geocoder | [optional] 
**city** | **str** | City name | [optional] 
**country** | **str** | Full country name | [optional] 
**country_code** | **str** | ISO 3166-1 alpha-2 country code | [optional] 
**coordinates** | **object** | Resolved lat/lng coordinates | [optional] 
**timezone** | **str** | IANA timezone identifier (e.g. Europe/Berlin). Null when timezone resolution fails. | [optional] 
**utc_offset** | **str** | UTC offset string (e.g. +02:00). Null when timezone resolution fails. | [optional] 
**current_time** | **str** | Current local time in RFC3339 format. Null when timezone resolution fails. | [optional] 
**is_holiday_today** | **bool** | True when today is a public holiday in the resolved country | [optional] 
**working_days_this_month** | **int** | Number of working days (excluding weekends and public holidays) in the current calendar month | [optional] 
**next_holiday** | **str** | Next public holiday within 90 days. Shape: {date: YYYY-MM-DD, name: string, type: national} | [optional] 
**flags** | **List[object]** | Non-fatal warnings. Possible values: timezone_unavailable, calendar_unavailable | [optional] 

## Example

```python
from openapi_client.models.v1_systems_location_resolve_post200_response_data import V1SystemsLocationResolvePost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1SystemsLocationResolvePost200ResponseData from a JSON string
v1_systems_location_resolve_post200_response_data_instance = V1SystemsLocationResolvePost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1SystemsLocationResolvePost200ResponseData.to_json())

# convert the object into a dict
v1_systems_location_resolve_post200_response_data_dict = v1_systems_location_resolve_post200_response_data_instance.to_dict()
# create an instance of V1SystemsLocationResolvePost200ResponseData from a dict
v1_systems_location_resolve_post200_response_data_from_dict = V1SystemsLocationResolvePost200ResponseData.from_dict(v1_systems_location_resolve_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


