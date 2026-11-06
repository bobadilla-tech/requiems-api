# DadJokesApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1EntertainmentJokesDadGet**](#v1entertainmentjokesdadget) | **GET** /v1/entertainment/jokes/dad | Get Random Dad Joke|

# **v1EntertainmentJokesDadGet**
> V1EntertainmentJokesDadGet200Response v1EntertainmentJokesDadGet()

Returns a randomly selected dad joke from the collection.

### Example

```typescript
import {
    DadJokesApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new DadJokesApi(configuration);

const { status, data } = await apiInstance.v1EntertainmentJokesDadGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1EntertainmentJokesDadGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**401** | Missing API key |  -  |
|**403** | Invalid or revoked API key |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

