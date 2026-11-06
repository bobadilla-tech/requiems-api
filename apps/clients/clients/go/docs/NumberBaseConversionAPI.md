# \NumberBaseConversionAPI

All URIs are relative to *https://requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyBaseBatchPost**](NumberBaseConversionAPI.md#V1TechnologyBaseBatchPost) | **Post** /v1/technology/base/batch | Batch Convert Base
[**V1TechnologyBaseGet**](NumberBaseConversionAPI.md#V1TechnologyBaseGet) | **Get** /v1/technology/base | Convert Base



## V1TechnologyBaseBatchPost

> V1TechnologyBaseBatchPost200Response V1TechnologyBaseBatchPost(ctx).V1TechnologyBaseBatchPostRequest(v1TechnologyBaseBatchPostRequest).Execute()

Batch Convert Base



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
	v1TechnologyBaseBatchPostRequest := *openapiclient.NewV1TechnologyBaseBatchPostRequest("Items_example", int32(123), int32(123), "ItemsValue_example") // V1TechnologyBaseBatchPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberBaseConversionAPI.V1TechnologyBaseBatchPost(context.Background()).V1TechnologyBaseBatchPostRequest(v1TechnologyBaseBatchPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberBaseConversionAPI.V1TechnologyBaseBatchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyBaseBatchPost`: V1TechnologyBaseBatchPost200Response
	fmt.Fprintf(os.Stdout, "Response from `NumberBaseConversionAPI.V1TechnologyBaseBatchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyBaseBatchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TechnologyBaseBatchPostRequest** | [**V1TechnologyBaseBatchPostRequest**](V1TechnologyBaseBatchPostRequest.md) |  | 

### Return type

[**V1TechnologyBaseBatchPost200Response**](V1TechnologyBaseBatchPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TechnologyBaseGet

> V1TechnologyBaseGet200Response V1TechnologyBaseGet(ctx).From(from).To(to).Value(value).Execute()

Convert Base



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
	from := int32(10) // int32 | Source base (2, 8, 10, or 16)
	to := int32(16) // int32 | Target base (2, 8, 10, or 16)
	value := "255" // string | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NumberBaseConversionAPI.V1TechnologyBaseGet(context.Background()).From(from).To(to).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NumberBaseConversionAPI.V1TechnologyBaseGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyBaseGet`: V1TechnologyBaseGet200Response
	fmt.Fprintf(os.Stdout, "Response from `NumberBaseConversionAPI.V1TechnologyBaseGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyBaseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int32** | Source base (2, 8, 10, or 16) | 
 **to** | **int32** | Target base (2, 8, 10, or 16) | 
 **value** | **string** | The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal). | 

### Return type

[**V1TechnologyBaseGet200Response**](V1TechnologyBaseGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

