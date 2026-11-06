# \ColorConversionAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyColorGet**](ColorConversionAPI.md#V1TechnologyColorGet) | **Get** /v1/technology/color | Convert Color



## V1TechnologyColorGet

> V1TechnologyColorGet200Response V1TechnologyColorGet(ctx).From(from).To(to).Value(value).Execute()

Convert Color



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
	from := "hex" // string | Source color format: hex, rgb, hsl, or cmyk
	to := "hsl" // string | Target color format: hex, rgb, hsl, or cmyk
	value := "#ff5733" // string | Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%))

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ColorConversionAPI.V1TechnologyColorGet(context.Background()).From(from).To(to).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ColorConversionAPI.V1TechnologyColorGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyColorGet`: V1TechnologyColorGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ColorConversionAPI.V1TechnologyColorGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyColorGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Source color format: hex, rgb, hsl, or cmyk | 
 **to** | **string** | Target color format: hex, rgb, hsl, or cmyk | 
 **value** | **string** | Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%)) | 

### Return type

[**V1TechnologyColorGet200Response**](V1TechnologyColorGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

