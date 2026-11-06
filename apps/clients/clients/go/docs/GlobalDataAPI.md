# \GlobalDataAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SystemsBusinessCalendarCountryGet**](GlobalDataAPI.md#V1SystemsBusinessCalendarCountryGet) | **Get** /v1/systems/business-calendar/{country} | Business Calendar
[**V1SystemsLocationResolvePost**](GlobalDataAPI.md#V1SystemsLocationResolvePost) | **Post** /v1/systems/location/resolve | Resolve Location
[**V1SystemsTimezoneFromIpIpGet**](GlobalDataAPI.md#V1SystemsTimezoneFromIpIpGet) | **Get** /v1/systems/timezone/from-ip/{ip} | Timezone from IP



## V1SystemsBusinessCalendarCountryGet

> V1SystemsBusinessCalendarCountryGet200Response V1SystemsBusinessCalendarCountryGet(ctx, country).Year(year).Month(month).Execute()

Business Calendar



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	country := "DE" // string | ISO 3166-1 alpha-2 country code (e.g. DE, US, FR)
	year := int32(2026) // int32 | Year to retrieve the calendar for. Defaults to the current year. (optional)
	month := int32(6) // int32 | Month (1–12). When provided, returns month-scoped working day count, total days, weekend days, and filtered holidays. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalDataAPI.V1SystemsBusinessCalendarCountryGet(context.Background(), country).Year(year).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalDataAPI.V1SystemsBusinessCalendarCountryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SystemsBusinessCalendarCountryGet`: V1SystemsBusinessCalendarCountryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalDataAPI.V1SystemsBusinessCalendarCountryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**country** | **string** | ISO 3166-1 alpha-2 country code (e.g. DE, US, FR) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemsBusinessCalendarCountryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **year** | **int32** | Year to retrieve the calendar for. Defaults to the current year. | 
 **month** | **int32** | Month (1–12). When provided, returns month-scoped working day count, total days, weekend days, and filtered holidays. | 

### Return type

[**V1SystemsBusinessCalendarCountryGet200Response**](V1SystemsBusinessCalendarCountryGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SystemsLocationResolvePost

> V1SystemsLocationResolvePost200Response V1SystemsLocationResolvePost(ctx).V1SystemsLocationResolvePostRequest(v1SystemsLocationResolvePostRequest).Execute()

Resolve Location



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v1SystemsLocationResolvePostRequest := *openapiclient.NewV1SystemsLocationResolvePostRequest() // V1SystemsLocationResolvePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalDataAPI.V1SystemsLocationResolvePost(context.Background()).V1SystemsLocationResolvePostRequest(v1SystemsLocationResolvePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalDataAPI.V1SystemsLocationResolvePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SystemsLocationResolvePost`: V1SystemsLocationResolvePost200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalDataAPI.V1SystemsLocationResolvePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemsLocationResolvePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1SystemsLocationResolvePostRequest** | [**V1SystemsLocationResolvePostRequest**](V1SystemsLocationResolvePostRequest.md) |  | 

### Return type

[**V1SystemsLocationResolvePost200Response**](V1SystemsLocationResolvePost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SystemsTimezoneFromIpIpGet

> V1SystemsTimezoneFromIpIpGet200Response V1SystemsTimezoneFromIpIpGet(ctx, ip).Execute()

Timezone from IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ip := "92.168.1.1" // string | IPv4 or IPv6 address to resolve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalDataAPI.V1SystemsTimezoneFromIpIpGet(context.Background(), ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalDataAPI.V1SystemsTimezoneFromIpIpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SystemsTimezoneFromIpIpGet`: V1SystemsTimezoneFromIpIpGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GlobalDataAPI.V1SystemsTimezoneFromIpIpGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | IPv4 or IPv6 address to resolve | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SystemsTimezoneFromIpIpGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1SystemsTimezoneFromIpIpGet200Response**](V1SystemsTimezoneFromIpIpGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

