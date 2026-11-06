# V1PlacesTimezoneBatchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **string** | List of timezone lookup requests (1–50 items). | 
**ItemsTimezone** | Pointer to **string** | IANA timezone name (e.g. &#39;America/New_York&#39;). Highest priority. | [optional] 
**ItemsCity** | Pointer to **string** | City name for city-based lookup. Used when timezone is not set. | [optional] 
**ItemsLat** | Pointer to **float32** | Latitude (-90 to 90). Used when neither timezone nor city is set. | [optional] 
**ItemsLon** | Pointer to **float32** | Longitude (-180 to 180). Used alongside lat for coordinate-based lookup. | [optional] 

## Methods

### NewV1PlacesTimezoneBatchPostRequest

`func NewV1PlacesTimezoneBatchPostRequest(items string, ) *V1PlacesTimezoneBatchPostRequest`

NewV1PlacesTimezoneBatchPostRequest instantiates a new V1PlacesTimezoneBatchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PlacesTimezoneBatchPostRequestWithDefaults

`func NewV1PlacesTimezoneBatchPostRequestWithDefaults() *V1PlacesTimezoneBatchPostRequest`

NewV1PlacesTimezoneBatchPostRequestWithDefaults instantiates a new V1PlacesTimezoneBatchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *V1PlacesTimezoneBatchPostRequest) GetItems() string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1PlacesTimezoneBatchPostRequest) GetItemsOk() (*string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1PlacesTimezoneBatchPostRequest) SetItems(v string)`

SetItems sets Items field to given value.


### GetItemsTimezone

`func (o *V1PlacesTimezoneBatchPostRequest) GetItemsTimezone() string`

GetItemsTimezone returns the ItemsTimezone field if non-nil, zero value otherwise.

### GetItemsTimezoneOk

`func (o *V1PlacesTimezoneBatchPostRequest) GetItemsTimezoneOk() (*string, bool)`

GetItemsTimezoneOk returns a tuple with the ItemsTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTimezone

`func (o *V1PlacesTimezoneBatchPostRequest) SetItemsTimezone(v string)`

SetItemsTimezone sets ItemsTimezone field to given value.

### HasItemsTimezone

`func (o *V1PlacesTimezoneBatchPostRequest) HasItemsTimezone() bool`

HasItemsTimezone returns a boolean if a field has been set.

### GetItemsCity

`func (o *V1PlacesTimezoneBatchPostRequest) GetItemsCity() string`

GetItemsCity returns the ItemsCity field if non-nil, zero value otherwise.

### GetItemsCityOk

`func (o *V1PlacesTimezoneBatchPostRequest) GetItemsCityOk() (*string, bool)`

GetItemsCityOk returns a tuple with the ItemsCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsCity

`func (o *V1PlacesTimezoneBatchPostRequest) SetItemsCity(v string)`

SetItemsCity sets ItemsCity field to given value.

### HasItemsCity

`func (o *V1PlacesTimezoneBatchPostRequest) HasItemsCity() bool`

HasItemsCity returns a boolean if a field has been set.

### GetItemsLat

`func (o *V1PlacesTimezoneBatchPostRequest) GetItemsLat() float32`

GetItemsLat returns the ItemsLat field if non-nil, zero value otherwise.

### GetItemsLatOk

`func (o *V1PlacesTimezoneBatchPostRequest) GetItemsLatOk() (*float32, bool)`

GetItemsLatOk returns a tuple with the ItemsLat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsLat

`func (o *V1PlacesTimezoneBatchPostRequest) SetItemsLat(v float32)`

SetItemsLat sets ItemsLat field to given value.

### HasItemsLat

`func (o *V1PlacesTimezoneBatchPostRequest) HasItemsLat() bool`

HasItemsLat returns a boolean if a field has been set.

### GetItemsLon

`func (o *V1PlacesTimezoneBatchPostRequest) GetItemsLon() float32`

GetItemsLon returns the ItemsLon field if non-nil, zero value otherwise.

### GetItemsLonOk

`func (o *V1PlacesTimezoneBatchPostRequest) GetItemsLonOk() (*float32, bool)`

GetItemsLonOk returns a tuple with the ItemsLon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsLon

`func (o *V1PlacesTimezoneBatchPostRequest) SetItemsLon(v float32)`

SetItemsLon sets ItemsLon field to given value.

### HasItemsLon

`func (o *V1PlacesTimezoneBatchPostRequest) HasItemsLon() bool`

HasItemsLon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


