# OpenapiClient::V1TechnologyConvertUnitsGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **length** | **Array&lt;Object&gt;** | Available length units: millimeter (mm), centimeter (cm), meter (m), kilometer (km), inch (in), foot (ft), yard (yd), mile (miles), nautical mile (nmi) | [optional] |
| **weight** | **Array&lt;Object&gt;** | Available weight units: milligram (mg), gram (g), kilogram (kg), metric ton (t), ounce (oz), pound (lb), stone (stone) | [optional] |
| **volume** | **Array&lt;Object&gt;** | Available volume units: milliliter (ml), liter (l), teaspoon (tsp), tablespoon (tbsp), fluid ounce (fl_oz), cup (cup), pint (pt), quart (qt), gallon (gal) | [optional] |
| **temperature** | **Array&lt;Object&gt;** | Available temperature units: celsius (c), fahrenheit (f), kelvin (k) | [optional] |
| **area** | **Array&lt;Object&gt;** | Available area units: square millimeter (mm2), square centimeter (cm2), square meter (m2), square kilometer (km2), square inch (in2), square foot (ft2), square yard (yd2), acre (acre), hectare (ha) | [optional] |
| **speed** | **Array&lt;Object&gt;** | Available speed units: meters per second (m_s), kilometers per hour (km_h), miles per hour (mph), knots (knots) | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyConvertUnitsGet200ResponseData.new(
  length: null,
  weight: null,
  volume: null,
  temperature: null,
  area: null,
  speed: null
)
```

