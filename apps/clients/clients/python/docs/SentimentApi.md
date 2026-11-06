# openapi_client.SentimentApi

All URIs are relative to *https://api.requiems.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_text_sentiment_post**](SentimentApi.md#v1_text_sentiment_post) | **POST** /v1/text/sentiment | Analyze Sentiment


# **v1_text_sentiment_post**
> V1TextSentimentPost200Response v1_text_sentiment_post(v1_text_sentiment_post_request)

Analyze Sentiment

Analyzes the sentiment of the provided text and returns a classification, confidence score, and a full breakdown across all three sentiment classes.

### Example

* Api Key Authentication (requiems-api-key):

```python
import openapi_client
from openapi_client.models.v1_text_sentiment_post200_response import V1TextSentimentPost200Response
from openapi_client.models.v1_text_sentiment_post_request import V1TextSentimentPostRequest
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
    api_instance = openapi_client.SentimentApi(api_client)
    v1_text_sentiment_post_request = openapi_client.V1TextSentimentPostRequest() # V1TextSentimentPostRequest | 

    try:
        # Analyze Sentiment
        api_response = api_instance.v1_text_sentiment_post(v1_text_sentiment_post_request)
        print("The response of SentimentApi->v1_text_sentiment_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SentimentApi->v1_text_sentiment_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_text_sentiment_post_request** | [**V1TextSentimentPostRequest**](V1TextSentimentPostRequest.md)|  | 

### Return type

[**V1TextSentimentPost200Response**](V1TextSentimentPost200Response.md)

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful response |  -  |
**422** | unprocessable_entity |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

