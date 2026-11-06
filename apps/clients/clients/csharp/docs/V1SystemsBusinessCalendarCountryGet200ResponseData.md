# Org.OpenAPITools.Model.V1SystemsBusinessCalendarCountryGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | ISO 3166-1 alpha-2 country code echoed from the path parameter | [optional] 
**Year** | **int** | Year the calendar covers | [optional] 
**Month** | **string** | Month scope. Present only when the month query parameter was provided. | [optional] 
**WorkingDays** | **int** | Number of working days (excluding weekends and public holidays) in the scope | [optional] 
**TotalDays** | **string** | Total calendar days in the month. Present only when month is provided. | [optional] 
**WeekendDays** | **string** | Number of weekend days in the month. Present only when month is provided. | [optional] 
**Holidays** | **List&lt;Object&gt;** | List of public holidays in the scope. Each item: {date: YYYY-MM-DD, name: string, type: national} | [optional] 
**HolidayCount** | **int** | Count of holidays in the holidays array | [optional] 
**NextHoliday** | **string** | Next public holiday within 90 days from today. Shape: {date, name, type} | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

