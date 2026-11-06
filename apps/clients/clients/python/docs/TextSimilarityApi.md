# openapi_client.TextSimilarityApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_text_similarity_post**](TextSimilarityApi.md#v1_text_similarity_post) | **POST** /v1/text/similarity | Compare Text Similarity


# **v1_text_similarity_post**
> V1TextSimilarityPost200Response v1_text_similarity_post(v1_text_similarity_post_request)

Compare Text Similarity

Compares two texts and returns a cosine similarity score.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_text_similarity_post200_response import V1TextSimilarityPost200Response
from openapi_client.models.v1_text_similarity_post_request import V1TextSimilarityPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.requiems.xyz
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.requiems.xyz"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: requiems-api-key
configuration.api_key['requiems-api-key'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['requiems-api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.TextSimilarityApi(api_client)
    v1_text_similarity_post_request = openapi_client.V1TextSimilarityPostRequest() # V1TextSimilarityPostRequest | 

    try:
        # Compare Text Similarity
        api_response = api_instance.v1_text_similarity_post(v1_text_similarity_post_request)
        print("The response of TextSimilarityApi->v1_text_similarity_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TextSimilarityApi->v1_text_similarity_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_text_similarity_post_request** | [**V1TextSimilarityPostRequest**](V1TextSimilarityPostRequest.md)|  | 

### Return type

[**V1TextSimilarityPost200Response**](V1TextSimilarityPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**400** | The request body is missing or malformed. |  -  |
**422** | One or both text fields are missing or empty. |  -  |
**500** | Unexpected server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

