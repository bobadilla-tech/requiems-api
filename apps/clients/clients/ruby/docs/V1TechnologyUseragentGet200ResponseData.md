# OpenapiClient::V1TechnologyUseragentGet200ResponseData

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **browser** | **String** | Detected browser name (e.g. Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other) | [optional] |
| **browser_version** | **String** | Detected browser version (major.minor) | [optional] |
| **os** | **String** | Detected operating system (e.g. Windows, macOS, Linux, Android, iOS, ChromeOS, Other) | [optional] |
| **os_version** | **String** | Detected OS version (format varies by platform) | [optional] |
| **device** | **String** | Device type — one of desktop, mobile, tablet, bot, or unknown | [optional] |
| **is_bot** | **Boolean** | True when the user agent matches a known bot or crawler pattern | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::V1TechnologyUseragentGet200ResponseData.new(
  browser: null,
  browser_version: null,
  os: null,
  os_version: null,
  device: null,
  is_bot: null
)
```

