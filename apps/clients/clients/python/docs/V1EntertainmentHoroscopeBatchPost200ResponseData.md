# V1EntertainmentHoroscopeBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **List[object]** | List of daily horoscopes returned in the same order as the input signs. | [optional] 
**results_sign** | **str** | Zodiac sign the horoscope corresponds to. | [optional] 
**results_date** | **str** | Date the horoscope is for, in YYYY-MM-DD format. | [optional] 
**results_horoscope** | **str** | Daily horoscope reading for the given sign. | [optional] 
**results_lucky_number** | **int** | Lucky number for the given sign on that date. | [optional] 
**results_mood** | **str** | Predicted mood for the given sign on that date. | [optional] 
**total** | **int** | Total number of horoscopes returned. | [optional] 

## Example

```python
from openapi_client.models.v1_entertainment_horoscope_batch_post200_response_data import V1EntertainmentHoroscopeBatchPost200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentHoroscopeBatchPost200ResponseData from a JSON string
v1_entertainment_horoscope_batch_post200_response_data_instance = V1EntertainmentHoroscopeBatchPost200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentHoroscopeBatchPost200ResponseData.to_json())

# convert the object into a dict
v1_entertainment_horoscope_batch_post200_response_data_dict = v1_entertainment_horoscope_batch_post200_response_data_instance.to_dict()
# create an instance of V1EntertainmentHoroscopeBatchPost200ResponseData from a dict
v1_entertainment_horoscope_batch_post200_response_data_from_dict = V1EntertainmentHoroscopeBatchPost200ResponseData.from_dict(v1_entertainment_horoscope_batch_post200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


