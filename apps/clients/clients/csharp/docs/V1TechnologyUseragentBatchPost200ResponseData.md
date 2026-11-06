# Org.OpenAPITools.Model.V1TechnologyUseragentBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **List&lt;Object&gt;** | List of parsed user agents results returned in the same order as the input user agents. | [optional] 
**ResultsUserAgent** | **string** | The original user agent string from the input. | [optional] 
**ResultsData** | **Object** | Parsed user agent details. | [optional] 
**ResultsDataBrowser** | **string** | Detected browser name (e.g. Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other). Empty if not detected. | [optional] 
**ResultsDataBrowserVersion** | **string** | Detected browser version (major.minor). Empty if not detected. | [optional] 
**ResultsDataOs** | **string** | Detected operating system (e.g. Windows, macOS, Linux, Android, iOS, ChromeOS, Other). Empty if not detected. | [optional] 
**ResultsDataOsVersion** | **string** | Detected OS version (format varies by platform). Empty if not detected. | [optional] 
**ResultsDataDevice** | **string** | Device type — one of desktop, mobile, tablet, bot, or unknown | [optional] 
**ResultsDataIsBot** | **bool** | True when the user agent matches a known bot or crawler pattern | [optional] 
**Total** | **int** | Total number of parsed user agent results returned. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

