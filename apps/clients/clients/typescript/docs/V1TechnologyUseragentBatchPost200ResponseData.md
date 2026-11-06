# V1TechnologyUseragentBatchPost200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**results** | **Array&lt;any&gt;** | List of parsed user agents results returned in the same order as the input user agents. | [optional] [default to undefined]
**results_user_agent** | **string** | The original user agent string from the input. | [optional] [default to undefined]
**results_data** | **object** | Parsed user agent details. | [optional] [default to undefined]
**results_data_browser** | **string** | Detected browser name (e.g. Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other). Empty if not detected. | [optional] [default to undefined]
**results_data_browser_version** | **string** | Detected browser version (major.minor). Empty if not detected. | [optional] [default to undefined]
**results_data_os** | **string** | Detected operating system (e.g. Windows, macOS, Linux, Android, iOS, ChromeOS, Other). Empty if not detected. | [optional] [default to undefined]
**results_data_os_version** | **string** | Detected OS version (format varies by platform). Empty if not detected. | [optional] [default to undefined]
**results_data_device** | **string** | Device type — one of desktop, mobile, tablet, bot, or unknown | [optional] [default to undefined]
**results_data_is_bot** | **boolean** | True when the user agent matches a known bot or crawler pattern | [optional] [default to undefined]
**total** | **number** | Total number of parsed user agent results returned. | [optional] [default to undefined]

## Example

```typescript
import { V1TechnologyUseragentBatchPost200ResponseData } from 'api-client';

const instance: V1TechnologyUseragentBatchPost200ResponseData = {
    results,
    results_user_agent,
    results_data,
    results_data_browser,
    results_data_browser_version,
    results_data_os,
    results_data_os_version,
    results_data_device,
    results_data_is_bot,
    total,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
