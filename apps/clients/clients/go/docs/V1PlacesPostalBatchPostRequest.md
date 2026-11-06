# V1PlacesPostalBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **string** | List of postal code lookups (1–50 items). | 
**ItemsCode** | **string** | The postal code to look up. | 
**ItemsCountry** | Pointer to **string** | ISO 3166-1 alpha-2 country code (default: US). | [optional] 

## Methods

### NewV1PlacesPostalBatchPostRequest

`func NewV1PlacesPostalBatchPostRequest(items string, itemsCode string, ) *V1PlacesPostalBatchPostRequest`

NewV1PlacesPostalBatchPostRequest instantiates a new V1PlacesPostalBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesPostalBatchPostRequestWithDefaults

`func NewV1PlacesPostalBatchPostRequestWithDefaults() *V1PlacesPostalBatchPostRequest`

NewV1PlacesPostalBatchPostRequestWithDefaults instantiates a new V1PlacesPostalBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1PlacesPostalBatchPostRequest) GetItems() string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1PlacesPostalBatchPostRequest) GetItemsOk() (*string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1PlacesPostalBatchPostRequest) SetItems(v string)`

SetItems sets Items field to given value.


### GetItemsCode

`func (o *V1PlacesPostalBatchPostRequest) GetItemsCode() string`

GetItemsCode returns the ItemsCode field if non-nil, zero value otherwise.

### GetItemsCodeOk

`func (o *V1PlacesPostalBatchPostRequest) GetItemsCodeOk() (*string, bool)`

GetItemsCodeOk returns a tuple with the ItemsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsCode

`func (o *V1PlacesPostalBatchPostRequest) SetItemsCode(v string)`

SetItemsCode sets ItemsCode field to given value.


### GetItemsCountry

`func (o *V1PlacesPostalBatchPostRequest) GetItemsCountry() string`

GetItemsCountry returns the ItemsCountry field if non-nil, zero value otherwise.

### GetItemsCountryOk

`func (o *V1PlacesPostalBatchPostRequest) GetItemsCountryOk() (*string, bool)`

GetItemsCountryOk returns a tuple with the ItemsCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsCountry

`func (o *V1PlacesPostalBatchPostRequest) SetItemsCountry(v string)`

SetItemsCountry sets ItemsCountry field to given value.

### HasItemsCountry

`func (o *V1PlacesPostalBatchPostRequest) HasItemsCountry() bool`

HasItemsCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


