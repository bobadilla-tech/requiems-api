# \CommoditiesAPI

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1FinanceCommoditiesCommodityGet**](CommoditiesAPI.md#V1FinanceCommoditiesCommodityGet) | **Get** /v1/finance/commodities/{commodity} | Get Commodity Price



## V1FinanceCommoditiesCommodityGet

> V1FinanceCommoditiesCommodityGet200Response V1FinanceCommoditiesCommodityGet(ctx, commodity).Execute()

Get Commodity Price



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
	commodity := "gold" // string | Commodity slug (e.g. gold, silver, oil). See supported slugs below.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommoditiesAPI.V1FinanceCommoditiesCommodityGet(context.Background(), commodity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommoditiesAPI.V1FinanceCommoditiesCommodityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1FinanceCommoditiesCommodityGet`: V1FinanceCommoditiesCommodityGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CommoditiesAPI.V1FinanceCommoditiesCommodityGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commodity** | **string** | Commodity slug (e.g. gold, silver, oil). See supported slugs below. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1FinanceCommoditiesCommodityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1FinanceCommoditiesCommodityGet200Response**](V1FinanceCommoditiesCommodityGet200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

