# \HolidaysAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PlacesHolidaysBatchPost**](HolidaysAPI.md#V1PlacesHolidaysBatchPost) | **Post** /v1/places/holidays/batch | Batch Get Holidays
[**V1PlacesHolidaysGet**](HolidaysAPI.md#V1PlacesHolidaysGet) | **Get** /v1/places/holidays | Get Holidays



## V1PlacesHolidaysBatchPost

> V1PlacesHolidaysBatchPost200Response V1PlacesHolidaysBatchPost(ctx).V1PlacesHolidaysBatchPostRequest(v1PlacesHolidaysBatchPostRequest).Execute()

Batch Get Holidays



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
	v1PlacesHolidaysBatchPostRequest := *openapiclient.NewV1PlacesHolidaysBatchPostRequest([]interface{}{nil}) // V1PlacesHolidaysBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.V1PlacesHolidaysBatchPost(context.Background()).V1PlacesHolidaysBatchPostRequest(v1PlacesHolidaysBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.V1PlacesHolidaysBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PlacesHolidaysBatchPost`: V1PlacesHolidaysBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.V1PlacesHolidaysBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PlacesHolidaysBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1PlacesHolidaysBatchPostRequest** | [**V1PlacesHolidaysBatchPostRequest**](V1PlacesHolidaysBatchPostRequest.md) |  | 

### Return type

[**V1PlacesHolidaysBatchPost200Response**](V1PlacesHolidaysBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PlacesHolidaysGet

> V1PlacesHolidaysGet200Response V1PlacesHolidaysGet(ctx).Country(country).Year(year).Execute()

Get Holidays



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
	country := "US" // string | ISO 3166-1 alpha-2 country code (e.g., \"US\", \"GB\", \"DE\")
	year := int32(2025) // int32 | Year for which to retrieve holidays (e.g., 2025)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.V1PlacesHolidaysGet(context.Background()).Country(country).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.V1PlacesHolidaysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PlacesHolidaysGet`: V1PlacesHolidaysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.V1PlacesHolidaysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PlacesHolidaysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;DE\&quot;) | 
 **year** | **int32** | Year for which to retrieve holidays (e.g., 2025) | 

### Return type

[**V1PlacesHolidaysGet200Response**](V1PlacesHolidaysGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

