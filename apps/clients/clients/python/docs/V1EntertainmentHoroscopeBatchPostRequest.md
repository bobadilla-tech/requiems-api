# V1EntertainmentHoroscopeBatchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**signs** | **List[object]** | Array of signs to get (min: 1, max: 12). Allowed values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces. | 

## Example

```python
from openapi_client.models.v1_entertainment_horoscope_batch_post_request import V1EntertainmentHoroscopeBatchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1EntertainmentHoroscopeBatchPostRequest from a JSON string
v1_entertainment_horoscope_batch_post_request_instance = V1EntertainmentHoroscopeBatchPostRequest.from_json(json)
# print the JSON string representation of the object
print(V1EntertainmentHoroscopeBatchPostRequest.to_json())

# convert the object into a dict
v1_entertainment_horoscope_batch_post_request_dict = v1_entertainment_horoscope_batch_post_request_instance.to_dict()
# create an instance of V1EntertainmentHoroscopeBatchPostRequest from a dict
v1_entertainment_horoscope_batch_post_request_from_dict = V1EntertainmentHoroscopeBatchPostRequest.from_dict(v1_entertainment_horoscope_batch_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


