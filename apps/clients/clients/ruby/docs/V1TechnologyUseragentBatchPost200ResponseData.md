# OpenapiClient::V1TechnologyUseragentBatchPost200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **results** | **Array&lt;Object&gt;** | List of parsed user agents results returned in the same order as the input user agents. | [optional] |
| **results_user_agent** | **String** | The original user agent string from the input. | [optional] |
| **results_data** | **Object** | Parsed user agent details. | [optional] |
| **results_data_browser** | **String** | Detected browser name (e.g. Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other). Empty if not detected. | [optional] |
| **results_data_browser_version** | **String** | Detected browser version (major.minor). Empty if not detected. | [optional] |
| **results_data_os** | **String** | Detected operating system (e.g. Windows, macOS, Linux, Android, iOS, ChromeOS, Other). Empty if not detected. | [optional] |
| **results_data_os_version** | **String** | Detected OS version (format varies by platform). Empty if not detected. | [optional] |
| **results_data_device** | **String** | Device type — one of desktop, mobile, tablet, bot, or unknown | [optional] |
| **results_data_is_bot** | **Boolean** | True when the user agent matches a known bot or crawler pattern | [optional] |
| **total** | **Integer** | Total number of parsed user agent results returned. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyUseragentBatchPost200ResponseData.new(
  results: null,
  results_user_agent: null,
  results_data: null,
  results_data_browser: null,
  results_data_browser_version: null,
  results_data_os: null,
  results_data_os_version: null,
  results_data_device: null,
  results_data_is_bot: null,
  total: null
)
```

