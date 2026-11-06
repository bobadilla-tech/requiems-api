# EmojiApi

All URIs are relative to *https://api.requiems.xyz*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**v1EntertainmentEmojiNameGet**](#v1entertainmentemojinameget) | **GET** /v1/entertainment/emoji/{name} | Get Emoji by Name|
|[**v1EntertainmentEmojiRandomGet**](#v1entertainmentemojirandomget) | **GET** /v1/entertainment/emoji/random | Get Random Emoji|
|[**v1EntertainmentEmojiSearchGet**](#v1entertainmentemojisearchget) | **GET** /v1/entertainment/emoji/search | Search Emoji|

# **v1EntertainmentEmojiNameGet**
> V1EntertainmentEmojiNameGet200Response v1EntertainmentEmojiNameGet()

Returns a specific emoji by its CLDR snake_case name. The name is case-insensitive.

### Example

```typescript
import {
    EmojiApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new EmojiApi(configuration);

let name: string; //CLDR snake_case emoji name (e.g. grinning_face, thumbs_up) (default to undefined)

const { status, data } = await apiInstance.v1EntertainmentEmojiNameGet(
    name
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **name** | [**string**] | CLDR snake_case emoji name (e.g. grinning_face, thumbs_up) | defaults to undefined|


### Return type

**V1EntertainmentEmojiNameGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**404** | No emoji found with the given name. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1EntertainmentEmojiRandomGet**
> V1EntertainmentEmojiRandomGet200Response v1EntertainmentEmojiRandomGet()

Returns a randomly selected emoji with its full metadata.

### Example

```typescript
import {
    EmojiApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new EmojiApi(configuration);

const { status, data } = await apiInstance.v1EntertainmentEmojiRandomGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**V1EntertainmentEmojiRandomGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1EntertainmentEmojiSearchGet**
> V1EntertainmentEmojiSearchGet200Response v1EntertainmentEmojiSearchGet()

Search for emojis whose name or category contains the given query string (case-insensitive). Returns a list of all matches.

### Example

```typescript
import {
    EmojiApi,
    Configuration
} from 'api-client';

const configuration = new Configuration();
const apiInstance = new EmojiApi(configuration);

let q: string; //Search term to match against emoji names and categories (e.g. happy, heart, food) (default to undefined)

const { status, data } = await apiInstance.v1EntertainmentEmojiSearchGet(
    q
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **q** | [**string**] | Search term to match against emoji names and categories (e.g. happy, heart, food) | defaults to undefined|


### Return type

**V1EntertainmentEmojiSearchGet200Response**

### Authorization

[requiems-api-key](../README.md#requiems-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |
|**400** | The q query parameter is missing or empty. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

