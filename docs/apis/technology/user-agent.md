# User Agent API

## Status

✅ **Live** — Available at `GET /v1/technology/useragent`

## Overview

Parse and analyze user agent strings to extract browser name, version, operating
system, device type, and bot detection.

## Endpoint

### Parse User Agent

`GET /v1/technology/useragent?ua=<encoded-ua-string>`

### Query Parameters

| Parameter | Required | Description                   |
| --------- | -------- | ----------------------------- |
| `ua`      | Yes      | URL-encoded user agent string |

### Response

```json
{
  "data": {
    "browser": "Chrome",
    "browser_version": "120.0",
    "os": "Windows",
    "os_version": "10/11",
    "device": "desktop",
    "is_bot": false
  },
  "metadata": {
    "timestamp": "2026-01-01T00:00:00Z"
  }
}
```

### Device Values

| Value     | Description                                   |
| --------- | --------------------------------------------- |
| `desktop` | Desktop browser on Windows, macOS, or Linux   |
| `mobile`  | Mobile browser (iPhone, Android Mobile)       |
| `tablet`  | Tablet browser (iPad, Android without Mobile) |
| `bot`     | Known bot or crawler                          |
| `unknown` | Empty user agent string                       |

### Response Fields

| Field           | Type    | Description                                                                                  |
| --------------- | ------- | -------------------------------------------------------------------------------------------- |
| browser         | string  | Detected browser name (e.g., Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other) |
| browser_version | string  | Browser version in "major.minor" format (e.g., "120.0"), empty if not detected               |
| os              | string  | Detected operating system (e.g., Windows, macOS, Linux, Android, iOS, ChromeOS, Other)       |
| os_version      | string  | OS version (format varies by platform)                                                       |
| device          | string  | Device type: desktop, mobile, tablet, bot, or unknown                                        |
| is_bot          | boolean | True when the user agent matches a known bot or crawler pattern                              |

### Error Codes

| Code          | Status | When                            |
| ------------- | ------ | ------------------------------- |
| `bad_request` | 400    | `ua` query parameter is missing |

## Code Examples

### cURL

```bash
curl "https://requiems.xyz/v1/technology/useragent?ua=Mozilla%2F5.0+%28Windows+NT+10.0%29+Chrome%2F120.0.0.0" \
  -H "requiems-api-key: YOUR_API_KEY"
```

### Python

```python
import requests
from urllib.parse import quote

ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0"
url = f"https://requiems.xyz/v1/technology/useragent?ua={quote(ua)}"
headers = {"requiems-api-key": "YOUR_API_KEY"}

response = requests.get(url, headers=headers)
result = response.json()['data']
print(f"Browser: {result['browser']} {result['browser_version']}")
print(f"OS: {result['os']} {result['os_version']}")
print(f"Device: {result['device']}")
```

### JavaScript

```javascript
const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0";
const url = `https://requiems.xyz/v1/technology/useragent?ua=${
  encodeURIComponent(ua)
}`;

const response = await fetch(url, {
  headers: { "requiems-api-key": "YOUR_API_KEY" },
});

const { data } = await response.json();
console.log(`Browser: ${data.browser} ${data.browser_version}`);
console.log(`OS: ${data.os} ${data.os_version}`);
console.log(`Device: ${data.device}`);
```

### Ruby

```ruby
require 'net/http'
require 'json'

ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0"
uri = URI("https://requiems.xyz/v1/technology/useragent?ua=#{URI.encode_www_form_component(ua)}")
request = Net::HTTP::Get.new(uri)
request['requiems-api-key'] = 'YOUR_API_KEY'

response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
  http.request(request)
end

data = JSON.parse(response.body)['data']
puts "Browser: #{data['browser']} #{data['browser_version']}"
puts "OS: #{data['os']} #{data['os_version']}"
puts "Device: #{data['device']}"
```

## Use Cases

- **Analytics Dashboards** - Track browser and device usage patterns
- **Serving Device-Specific Content** - Deliver optimized experiences based on
  device type
- **Bot Filtering** - Identify and filter bot traffic from analytics
- **Browser Compatibility Reporting** - Understand which browsers your users are
  on

## FAQ

**Which browsers are detected?** Chrome, Firefox, Safari, Edge, Opera, and
Internet Explorer. Anything else returns "Other".

**How is bot detection handled?** The API checks for known bot/crawler keywords
in the user agent string, including Googlebot, Bingbot, curl, wget,
python-requests, and many others.

**What device types are returned?** desktop, mobile, tablet, bot, or unknown.
Tablets are identified by iPad or Android without 'Mobile' in the UA.
