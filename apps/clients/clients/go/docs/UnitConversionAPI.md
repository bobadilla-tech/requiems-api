# \UnitConversionAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyConvertGet**](UnitConversionAPI.md#V1TechnologyConvertGet) | **Get** /v1/technology/convert | Convert Units
[**V1TechnologyConvertUnitsGet**](UnitConversionAPI.md#V1TechnologyConvertUnitsGet) | **Get** /v1/technology/convert/units | List Available Units



## V1TechnologyConvertGet

> V1TechnologyConvertGet200Response V1TechnologyConvertGet(ctx).From(from).To(to).Value(value).Execute()

Convert Units



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
	from := "miles" // string | Source unit key (e.g. miles, kg, c)
	to := "km" // string | Target unit key (e.g. km, lb, f)
	value := float32(10) // float32 | Numeric value to convert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnitConversionAPI.V1TechnologyConvertGet(context.Background()).From(from).To(to).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitConversionAPI.V1TechnologyConvertGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyConvertGet`: V1TechnologyConvertGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UnitConversionAPI.V1TechnologyConvertGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyConvertGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Source unit key (e.g. miles, kg, c) | 
 **to** | **string** | Target unit key (e.g. km, lb, f) | 
 **value** | **float32** | Numeric value to convert | 

### Return type

[**V1TechnologyConvertGet200Response**](V1TechnologyConvertGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TechnologyConvertUnitsGet

> V1TechnologyConvertUnitsGet200Response V1TechnologyConvertUnitsGet(ctx).Execute()

List Available Units



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnitConversionAPI.V1TechnologyConvertUnitsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitConversionAPI.V1TechnologyConvertUnitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyConvertUnitsGet`: V1TechnologyConvertUnitsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UnitConversionAPI.V1TechnologyConvertUnitsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyConvertUnitsGetRequest struct via the builder pattern


### Return type

[**V1TechnologyConvertUnitsGet200Response**](V1TechnologyConvertUnitsGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

