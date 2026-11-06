# V1TechnologyConvertUnitsGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**length** | **List[object]** | Available length units: millimeter (mm), centimeter (cm), meter (m), kilometer (km), inch (in), foot (ft), yard (yd), mile (miles), nautical mile (nmi) | [optional] 
**weight** | **List[object]** | Available weight units: milligram (mg), gram (g), kilogram (kg), metric ton (t), ounce (oz), pound (lb), stone (stone) | [optional] 
**volume** | **List[object]** | Available volume units: milliliter (ml), liter (l), teaspoon (tsp), tablespoon (tbsp), fluid ounce (fl_oz), cup (cup), pint (pt), quart (qt), gallon (gal) | [optional] 
**temperature** | **List[object]** | Available temperature units: celsius (c), fahrenheit (f), kelvin (k) | [optional] 
**area** | **List[object]** | Available area units: square millimeter (mm2), square centimeter (cm2), square meter (m2), square kilometer (km2), square inch (in2), square foot (ft2), square yard (yd2), acre (acre), hectare (ha) | [optional] 
**speed** | **List[object]** | Available speed units: meters per second (m_s), kilometers per hour (km_h), miles per hour (mph), knots (knots) | [optional] 

## Example

```python
from openapi_client.models.v1_technology_convert_units_get200_response_data import V1TechnologyConvertUnitsGet200ResponseData

# TODO update the JSON string below
json = "{}"
# create an instance of V1TechnologyConvertUnitsGet200ResponseData from a JSON string
v1_technology_convert_units_get200_response_data_instance = V1TechnologyConvertUnitsGet200ResponseData.from_json(json)
# print the JSON string representation of the object
print(V1TechnologyConvertUnitsGet200ResponseData.to_json())

# convert the object into a dict
v1_technology_convert_units_get200_response_data_dict = v1_technology_convert_units_get200_response_data_instance.to_dict()
# create an instance of V1TechnologyConvertUnitsGet200ResponseData from a dict
v1_technology_convert_units_get200_response_data_from_dict = V1TechnologyConvertUnitsGet200ResponseData.from_dict(v1_technology_convert_units_get200_response_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


