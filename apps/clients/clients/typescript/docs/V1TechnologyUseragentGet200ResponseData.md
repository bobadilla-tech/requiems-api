# V1TechnologyUseragentGet200ResponseData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**browser** | **string** | Detected browser name (e.g. Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other) | [optional] [default to undefined]
**browser_version** | **string** | Detected browser version (major.minor) | [optional] [default to undefined]
**os** | **string** | Detected operating system (e.g. Windows, macOS, Linux, Android, iOS, ChromeOS, Other) | [optional] [default to undefined]
**os_version** | **string** | Detected OS version (format varies by platform) | [optional] [default to undefined]
**device** | **string** | Device type — one of desktop, mobile, tablet, bot, or unknown | [optional] [default to undefined]
**is_bot** | **boolean** | True when the user agent matches a known bot or crawler pattern | [optional] [default to undefined]

## Example

```typescript
import { V1TechnologyUseragentGet200ResponseData } from 'api-client';

const instance: V1TechnologyUseragentGet200ResponseData = {
    browser,
    browser_version,
    os,
    os_version,
    device,
    is_bot,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
