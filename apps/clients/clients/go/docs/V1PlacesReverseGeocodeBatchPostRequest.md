# V1PlacesReverseGeocodeBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **string** | List of coordinate pairs to reverse-geocode (1–20 items). Max 20 due to Nominatim usage policy. | 
**ItemsLat** | **float32** | Latitude (-90 to 90). | 
**ItemsLon** | **float32** | Longitude (-180 to 180). | 

## Methods

### NewV1PlacesReverseGeocodeBatchPostRequest

`func NewV1PlacesReverseGeocodeBatchPostRequest(items string, itemsLat float32, itemsLon float32, ) *V1PlacesReverseGeocodeBatchPostRequest`

NewV1PlacesReverseGeocodeBatchPostRequest instantiates a new V1PlacesReverseGeocodeBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesReverseGeocodeBatchPostRequestWithDefaults

`func NewV1PlacesReverseGeocodeBatchPostRequestWithDefaults() *V1PlacesReverseGeocodeBatchPostRequest`

NewV1PlacesReverseGeocodeBatchPostRequestWithDefaults instantiates a new V1PlacesReverseGeocodeBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1PlacesReverseGeocodeBatchPostRequest) GetItems() string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1PlacesReverseGeocodeBatchPostRequest) GetItemsOk() (*string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1PlacesReverseGeocodeBatchPostRequest) SetItems(v string)`

SetItems sets Items field to given value.


### GetItemsLat

`func (o *V1PlacesReverseGeocodeBatchPostRequest) GetItemsLat() float32`

GetItemsLat returns the ItemsLat field if non-nil, zero value otherwise.

### GetItemsLatOk

`func (o *V1PlacesReverseGeocodeBatchPostRequest) GetItemsLatOk() (*float32, bool)`

GetItemsLatOk returns a tuple with the ItemsLat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsLat

`func (o *V1PlacesReverseGeocodeBatchPostRequest) SetItemsLat(v float32)`

SetItemsLat sets ItemsLat field to given value.


### GetItemsLon

`func (o *V1PlacesReverseGeocodeBatchPostRequest) GetItemsLon() float32`

GetItemsLon returns the ItemsLon field if non-nil, zero value otherwise.

### GetItemsLonOk

`func (o *V1PlacesReverseGeocodeBatchPostRequest) GetItemsLonOk() (*float32, bool)`

GetItemsLonOk returns a tuple with the ItemsLon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsLon

`func (o *V1PlacesReverseGeocodeBatchPostRequest) SetItemsLon(v float32)`

SetItemsLon sets ItemsLon field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


