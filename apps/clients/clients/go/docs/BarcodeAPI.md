# \BarcodeAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TechnologyBarcodeBase64Get**](BarcodeAPI.md#V1TechnologyBarcodeBase64Get) | **Get** /v1/technology/barcode/base64 | Generate Barcode (Base64 JSON)
[**V1TechnologyBarcodeGet**](BarcodeAPI.md#V1TechnologyBarcodeGet) | **Get** /v1/technology/barcode | Generate Barcode (PNG)



## V1TechnologyBarcodeBase64Get

> V1TechnologyBarcodeBase64Get200Response V1TechnologyBarcodeBase64Get(ctx).Data(data).Type_(type_).Execute()

Generate Barcode (Base64 JSON)



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
	data := "123456789" // string | The text or numeric string to encode in the barcode
	type_ := "code128" // string | Barcode format: code128, code93, code39, ean8, ean13

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BarcodeAPI.V1TechnologyBarcodeBase64Get(context.Background()).Data(data).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BarcodeAPI.V1TechnologyBarcodeBase64Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1TechnologyBarcodeBase64Get`: V1TechnologyBarcodeBase64Get200Response
	fmt.Fprintf(os.Stdout, "Response from `BarcodeAPI.V1TechnologyBarcodeBase64Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyBarcodeBase64GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | **string** | The text or numeric string to encode in the barcode | 
 **type_** | **string** | Barcode format: code128, code93, code39, ean8, ean13 | 

### Return type

[**V1TechnologyBarcodeBase64Get200Response**](V1TechnologyBarcodeBase64Get200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TechnologyBarcodeGet

> V1TechnologyBarcodeGet(ctx).Data(data).Type_(type_).Execute()

Generate Barcode (PNG)



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
	data := "123456789" // string | The text or numeric string to encode in the barcode
	type_ := "code128" // string | Barcode format: code128, code93, code39, ean8, ean13

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BarcodeAPI.V1TechnologyBarcodeGet(context.Background()).Data(data).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BarcodeAPI.V1TechnologyBarcodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TechnologyBarcodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | **string** | The text or numeric string to encode in the barcode | 
 **type_** | **string** | Barcode format: code128, code93, code39, ean8, ean13 | 

### Return type

 (empty response body)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

