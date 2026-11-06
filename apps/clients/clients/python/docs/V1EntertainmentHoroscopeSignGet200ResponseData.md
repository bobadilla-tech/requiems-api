# V1EntertainmentHoroscopeSignGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sign** | **str** | Normalized zodiac sign (lowercase) | [optional] 
**var_date** | **str** | Today&#39;s date in YYYY-MM-DD format (UTC) | [optional] 
**horoscope** | **str** | Daily horoscope reading | [optional] 
**lucky_number** | **int** | Lucky number for the day (1-99) | [optional] 
**mood** | **str** | Suggested mood for the day | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_horoscope_sign_get200_response_data import V1EntertainmentHoroscopeSignGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentHoroscopeSignGet200ResponseData from a JSON string
v1_entertainment_horoscope_sign_get200_response_data_instance = V1EntertainmentHoroscopeSignGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentHoroscopeSignGet200ResponseData.to_json())

# convert the object into a dict
v1_entertainment_horoscope_sign_get200_response_data_dict = v1_entertainment_horoscope_sign_get200_response_data_instance.to_dict()
# create an instance of V1EntertainmentHoroscopeSignGet200ResponseData from a dict
v1_entertainment_horoscope_sign_get200_response_data_from_dict = V1EntertainmentHoroscopeSignGet200ResponseData.from_dict(v1_entertainment_horoscope_sign_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


