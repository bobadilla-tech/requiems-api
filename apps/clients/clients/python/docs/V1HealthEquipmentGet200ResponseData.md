# V1HealthEquipmentGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | **List[object]** | Sorted list of all distinct equipment names | [optional] 
**total** | **int** | Total number of distinct equipment types | [optional] 

## Example

```python
from openapi_client.models.v1_health_equipment_get200_response_data import V1HealthEquipmentGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1HealthEquipmentGet200ResponseData from a JSON string
v1_health_equipment_get200_response_data_instance = V1HealthEquipmentGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1HealthEquipmentGet200ResponseData.to_json())

# convert the object into a dict
v1_health_equipment_get200_response_data_dict = v1_health_equipment_get200_response_data_instance.to_dict()
# create an instance of V1HealthEquipmentGet200ResponseData from a dict
v1_health_equipment_get200_response_data_from_dict = V1HealthEquipmentGet200ResponseData.from_dict(v1_health_equipment_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


