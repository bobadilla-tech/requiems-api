# Org.OpenAPITools.Model.V1SystemsLocationResolvePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Normalized full address returned by the geocoder | [optional] 
**City** | **string** | City name | [optional] 
**Country** | **string** | Full country name | [optional] 
**CountryCode** | **string** | ISO 3166-1 alpha-2 country code | [optional] 
**Coordinates** | **Object** | Resolved lat/lng coordinates | [optional] 
**Timezone** | **string** | IANA timezone identifier (e.g. Europe/Berlin). Null when timezone resolution fails. | [optional] 
**UtcOffset** | **string** | UTC offset string (e.g. +02:00). Null when timezone resolution fails. | [optional] 
**CurrentTime** | **string** | Current local time in RFC3339 format. Null when timezone resolution fails. | [optional] 
**IsHolidayToday** | **bool** | True when today is a public holiday in the resolved country | [optional] 
**WorkingDaysThisMonth** | **int** | Number of working days (excluding weekends and public holidays) in the current calendar month | [optional] 
**NextHoliday** | **string** | Next public holiday within 90 days. Shape: {date: YYYY-MM-DD, name: string, type: national} | [optional] 
**Flags** | **List&lt;Object&gt;** | Non-fatal warnings. Possible values: timezone_unavailable, calendar_unavailable | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

