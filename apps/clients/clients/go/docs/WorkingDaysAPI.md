# \WorkingDaysAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PlacesWorkingDaysGet**](WorkingDaysAPI.md#V1PlacesWorkingDaysGet) | **Get** /v1/places/working-days | Calculate Working Days



## V1PlacesWorkingDaysGet

> V1PlacesWorkingDaysGet200Response V1PlacesWorkingDaysGet(ctx).From(from).To(to).Country(country).Subdivision(subdivision).Execute()

Calculate Working Days



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
	from := "2024-02-23" // string | Start date in YYYY-MM-DD format (ISO 8601)
	to := "2024-02-28" // string | End date in YYYY-MM-DD format (ISO 8601). Must be >= from date.
	country := "US" // string | ISO 3166-1 alpha-2 country code (e.g., \"US\", \"GB\", \"FR\"). When provided, country-specific holidays are excluded from working days count. (optional)
	subdivision := "NY" // string | ISO 3166-2 subdivision code for state/region within the country (e.g., \"NY\" for New York, \"CA\" for California). Only used when country is provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkingDaysAPI.V1PlacesWorkingDaysGet(context.Background()).From(from).To(to).Country(country).Subdivision(subdivision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkingDaysAPI.V1PlacesWorkingDaysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PlacesWorkingDaysGet`: V1PlacesWorkingDaysGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkingDaysAPI.V1PlacesWorkingDaysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PlacesWorkingDaysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Start date in YYYY-MM-DD format (ISO 8601) | 
 **to** | **string** | End date in YYYY-MM-DD format (ISO 8601). Must be &gt;&#x3D; from date. | 
 **country** | **string** | ISO 3166-1 alpha-2 country code (e.g., \&quot;US\&quot;, \&quot;GB\&quot;, \&quot;FR\&quot;). When provided, country-specific holidays are excluded from working days count. | 
 **subdivision** | **string** | ISO 3166-2 subdivision code for state/region within the country (e.g., \&quot;NY\&quot; for New York, \&quot;CA\&quot; for California). Only used when country is provided. | 

### Return type

[**V1PlacesWorkingDaysGet200Response**](V1PlacesWorkingDaysGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

