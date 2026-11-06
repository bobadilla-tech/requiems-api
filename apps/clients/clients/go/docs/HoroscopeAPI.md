# \HoroscopeAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EntertainmentHoroscopeSignGet**](HoroscopeAPI.md#V1EntertainmentHoroscopeSignGet) | **Get** /v1/entertainment/horoscope/{sign} | Get Daily Horoscope



## V1EntertainmentHoroscopeSignGet

> V1EntertainmentHoroscopeSignGet200Response V1EntertainmentHoroscopeSignGet(ctx, sign).Execute()

Get Daily Horoscope



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
	sign := "sign_example" // string | Zodiac sign (case-insensitive). Supported values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HoroscopeAPI.V1EntertainmentHoroscopeSignGet(context.Background(), sign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoroscopeAPI.V1EntertainmentHoroscopeSignGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EntertainmentHoroscopeSignGet`: V1EntertainmentHoroscopeSignGet200Response
	fmt.Fprintf(os.Stdout, "Response from `HoroscopeAPI.V1EntertainmentHoroscopeSignGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sign** | **string** | Zodiac sign (case-insensitive). Supported values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EntertainmentHoroscopeSignGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1EntertainmentHoroscopeSignGet200Response**](V1EntertainmentHoroscopeSignGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

