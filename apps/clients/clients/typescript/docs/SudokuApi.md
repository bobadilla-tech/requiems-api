# SudokuApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1EntertainmentSudokuBatchPost**](#v1entertainmentsudokubatchpost) | **POST** /v1/entertainment/sudoku/batch | Batch Generate Sudoku Puzzles|
|[**v1EntertainmentSudokuGet**](#v1entertainmentsudokuget) | **GET** /v1/entertainment/sudoku | Get Sudoku Puzzle|

# **v1EntertainmentSudokuBatchPost**
> V1EntertainmentSudokuBatchPost200Response v1EntertainmentSudokuBatchPost(v1EntertainmentSudokuBatchPostRequest)

Generate up to 20 Sudoku puzzles in a single request. Results are returned in the same order as the input array. Each puzzle in the batch counts as one unit of API usage.

### Example

```typescript
import {
    SudokuApi,
    Configuration,
    V1EntertainmentSudokuBatchPostRequest
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new SudokuApi(configuration);

let v1EntertainmentSudokuBatchPostRequest: V1EntertainmentSudokuBatchPostRequest; //

const { status, data } = await apiInstance.v1EntertainmentSudokuBatchPost(
    v1EntertainmentSudokuBatchPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **v1EntertainmentSudokuBatchPostRequest** | **V1EntertainmentSudokuBatchPostRequest**|  | |


### Return type

**V1EntertainmentSudokuBatchPost200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The request body is missing or contains malformed JSON. |  -  |
|**401** | Missing API key |  -  |
|**403** | Invalid or revoked API key |  -  |
|**422** | The puzzles array is missing, empty, exceeds 20 items, or contains a value other than easy, medium, or hard. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1EntertainmentSudokuGet**
> V1EntertainmentSudokuGet200Response v1EntertainmentSudokuGet()

Returns a randomly generated Sudoku puzzle and its solution. Difficulty defaults to medium when not specified.

### Example

```typescript
import {
    SudokuApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new SudokuApi(configuration);

let difficulty: string; //Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium. (optional) (default to undefined)

const { status, data } = await apiInstance.v1EntertainmentSudokuGet(
    difficulty
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **difficulty** | [**string**] | Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium. | (optional) defaults to undefined|


### Return type

**V1EntertainmentSudokuGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The difficulty parameter is not one of easy, medium, or hard |  -  |
|**401** | Missing API key |  -  |
|**403** | Invalid or revoked API key |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

