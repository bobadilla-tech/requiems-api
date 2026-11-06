# V1NetworkingMxDomainGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | The domain that was queried | [optional] 
**Records** | Pointer to **[]interface{}** | List of MX records, sorted by priority ascending (lowest number &#x3D; highest priority) | [optional] 
**RecordsHost** | Pointer to **string** | Fully-qualified hostname of the mail server (typically ends with a trailing dot) | [optional] 
**RecordsPriority** | Pointer to **int32** | MX priority value. Lower values have higher delivery priority per RFC 5321. | [optional] 

## Methods

### NewV1NetworkingMxDomainGet200ResponseData

`func NewV1NetworkingMxDomainGet200ResponseData() *V1NetworkingMxDomainGet200ResponseData`

NewV1NetworkingMxDomainGet200ResponseData instantiates a new V1NetworkingMxDomainGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NetworkingMxDomainGet200ResponseDataWithDefaults

`func NewV1NetworkingMxDomainGet200ResponseDataWithDefaults() *V1NetworkingMxDomainGet200ResponseData`

NewV1NetworkingMxDomainGet200ResponseDataWithDefaults instantiates a new V1NetworkingMxDomainGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *V1NetworkingMxDomainGet200ResponseData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *V1NetworkingMxDomainGet200ResponseData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *V1NetworkingMxDomainGet200ResponseData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *V1NetworkingMxDomainGet200ResponseData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRecords

`func (o *V1NetworkingMxDomainGet200ResponseData) GetRecords() []interface{}`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *V1NetworkingMxDomainGet200ResponseData) GetRecordsOk() (*[]interface{}, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *V1NetworkingMxDomainGet200ResponseData) SetRecords(v []interface{})`

SetRecords sets Records field to given value.

### HasRecords

`func (o *V1NetworkingMxDomainGet200ResponseData) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetRecordsHost

`func (o *V1NetworkingMxDomainGet200ResponseData) GetRecordsHost() string`

GetRecordsHost returns the RecordsHost field if non-nil, zero value otherwise.

### GetRecordsHostOk

`func (o *V1NetworkingMxDomainGet200ResponseData) GetRecordsHostOk() (*string, bool)`

GetRecordsHostOk returns a tuple with the RecordsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsHost

`func (o *V1NetworkingMxDomainGet200ResponseData) SetRecordsHost(v string)`

SetRecordsHost sets RecordsHost field to given value.

### HasRecordsHost

`func (o *V1NetworkingMxDomainGet200ResponseData) HasRecordsHost() bool`

HasRecordsHost returns a boolean if a field has been set.

### GetRecordsPriority

`func (o *V1NetworkingMxDomainGet200ResponseData) GetRecordsPriority() int32`

GetRecordsPriority returns the RecordsPriority field if non-nil, zero value otherwise.

### GetRecordsPriorityOk

`func (o *V1NetworkingMxDomainGet200ResponseData) GetRecordsPriorityOk() (*int32, bool)`

GetRecordsPriorityOk returns a tuple with the RecordsPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsPriority

`func (o *V1NetworkingMxDomainGet200ResponseData) SetRecordsPriority(v int32)`

SetRecordsPriority sets RecordsPriority field to given value.

### HasRecordsPriority

`func (o *V1NetworkingMxDomainGet200ResponseData) HasRecordsPriority() bool`

HasRecordsPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


