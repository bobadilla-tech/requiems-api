# V1TechnologyUseragentBatchPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]interface{}** | List of parsed user agents results returned in the same order as the input user agents. | [optional] 
**ResultsUserAgent** | Pointer to **string** | The original user agent string from the input. | [optional] 
**ResultsData** | Pointer to **map[string]interface{}** | Parsed user agent details. | [optional] 
**ResultsDataBrowser** | Pointer to **string** | Detected browser name (e.g. Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other). Empty if not detected. | [optional] 
**ResultsDataBrowserVersion** | Pointer to **string** | Detected browser version (major.minor). Empty if not detected. | [optional] 
**ResultsDataOs** | Pointer to **string** | Detected operating system (e.g. Windows, macOS, Linux, Android, iOS, ChromeOS, Other). Empty if not detected. | [optional] 
**ResultsDataOsVersion** | Pointer to **string** | Detected OS version (format varies by platform). Empty if not detected. | [optional] 
**ResultsDataDevice** | Pointer to **string** | Device type — one of desktop, mobile, tablet, bot, or unknown | [optional] 
**ResultsDataIsBot** | Pointer to **bool** | True when the user agent matches a known bot or crawler pattern | [optional] 
**Total** | Pointer to **int32** | Total number of parsed user agent results returned. | [optional] 

## Methods

### NewV1TechnologyUseragentBatchPost200ResponseData

`func NewV1TechnologyUseragentBatchPost200ResponseData() *V1TechnologyUseragentBatchPost200ResponseData`

NewV1TechnologyUseragentBatchPost200ResponseData instantiates a new V1TechnologyUseragentBatchPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TechnologyUseragentBatchPost200ResponseDataWithDefaults

`func NewV1TechnologyUseragentBatchPost200ResponseDataWithDefaults() *V1TechnologyUseragentBatchPost200ResponseData`

NewV1TechnologyUseragentBatchPost200ResponseDataWithDefaults instantiates a new V1TechnologyUseragentBatchPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *V1TechnologyUseragentBatchPost200ResponseData) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *V1TechnologyUseragentBatchPost200ResponseData) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsUserAgent

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsUserAgent() string`

GetResultsUserAgent returns the ResultsUserAgent field if non-nil, zero value otherwise.

### GetResultsUserAgentOk

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsUserAgentOk() (*string, bool)`

GetResultsUserAgentOk returns a tuple with the ResultsUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsUserAgent

`func (o *V1TechnologyUseragentBatchPost200ResponseData) SetResultsUserAgent(v string)`

SetResultsUserAgent sets ResultsUserAgent field to given value.

### HasResultsUserAgent

`func (o *V1TechnologyUseragentBatchPost200ResponseData) HasResultsUserAgent() bool`

HasResultsUserAgent returns a boolean if a field has been set.

### GetResultsData

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsData() map[string]interface{}`

GetResultsData returns the ResultsData field if non-nil, zero value otherwise.

### GetResultsDataOk

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataOk() (*map[string]interface{}, bool)`

GetResultsDataOk returns a tuple with the ResultsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsData

`func (o *V1TechnologyUseragentBatchPost200ResponseData) SetResultsData(v map[string]interface{})`

SetResultsData sets ResultsData field to given value.

### HasResultsData

`func (o *V1TechnologyUseragentBatchPost200ResponseData) HasResultsData() bool`

HasResultsData returns a boolean if a field has been set.

### GetResultsDataBrowser

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataBrowser() string`

GetResultsDataBrowser returns the ResultsDataBrowser field if non-nil, zero value otherwise.

### GetResultsDataBrowserOk

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataBrowserOk() (*string, bool)`

GetResultsDataBrowserOk returns a tuple with the ResultsDataBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataBrowser

`func (o *V1TechnologyUseragentBatchPost200ResponseData) SetResultsDataBrowser(v string)`

SetResultsDataBrowser sets ResultsDataBrowser field to given value.

### HasResultsDataBrowser

`func (o *V1TechnologyUseragentBatchPost200ResponseData) HasResultsDataBrowser() bool`

HasResultsDataBrowser returns a boolean if a field has been set.

### GetResultsDataBrowserVersion

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataBrowserVersion() string`

GetResultsDataBrowserVersion returns the ResultsDataBrowserVersion field if non-nil, zero value otherwise.

### GetResultsDataBrowserVersionOk

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataBrowserVersionOk() (*string, bool)`

GetResultsDataBrowserVersionOk returns a tuple with the ResultsDataBrowserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataBrowserVersion

`func (o *V1TechnologyUseragentBatchPost200ResponseData) SetResultsDataBrowserVersion(v string)`

SetResultsDataBrowserVersion sets ResultsDataBrowserVersion field to given value.

### HasResultsDataBrowserVersion

`func (o *V1TechnologyUseragentBatchPost200ResponseData) HasResultsDataBrowserVersion() bool`

HasResultsDataBrowserVersion returns a boolean if a field has been set.

### GetResultsDataOs

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataOs() string`

GetResultsDataOs returns the ResultsDataOs field if non-nil, zero value otherwise.

### GetResultsDataOsOk

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataOsOk() (*string, bool)`

GetResultsDataOsOk returns a tuple with the ResultsDataOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataOs

`func (o *V1TechnologyUseragentBatchPost200ResponseData) SetResultsDataOs(v string)`

SetResultsDataOs sets ResultsDataOs field to given value.

### HasResultsDataOs

`func (o *V1TechnologyUseragentBatchPost200ResponseData) HasResultsDataOs() bool`

HasResultsDataOs returns a boolean if a field has been set.

### GetResultsDataOsVersion

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataOsVersion() string`

GetResultsDataOsVersion returns the ResultsDataOsVersion field if non-nil, zero value otherwise.

### GetResultsDataOsVersionOk

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataOsVersionOk() (*string, bool)`

GetResultsDataOsVersionOk returns a tuple with the ResultsDataOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataOsVersion

`func (o *V1TechnologyUseragentBatchPost200ResponseData) SetResultsDataOsVersion(v string)`

SetResultsDataOsVersion sets ResultsDataOsVersion field to given value.

### HasResultsDataOsVersion

`func (o *V1TechnologyUseragentBatchPost200ResponseData) HasResultsDataOsVersion() bool`

HasResultsDataOsVersion returns a boolean if a field has been set.

### GetResultsDataDevice

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataDevice() string`

GetResultsDataDevice returns the ResultsDataDevice field if non-nil, zero value otherwise.

### GetResultsDataDeviceOk

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataDeviceOk() (*string, bool)`

GetResultsDataDeviceOk returns a tuple with the ResultsDataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataDevice

`func (o *V1TechnologyUseragentBatchPost200ResponseData) SetResultsDataDevice(v string)`

SetResultsDataDevice sets ResultsDataDevice field to given value.

### HasResultsDataDevice

`func (o *V1TechnologyUseragentBatchPost200ResponseData) HasResultsDataDevice() bool`

HasResultsDataDevice returns a boolean if a field has been set.

### GetResultsDataIsBot

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataIsBot() bool`

GetResultsDataIsBot returns the ResultsDataIsBot field if non-nil, zero value otherwise.

### GetResultsDataIsBotOk

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetResultsDataIsBotOk() (*bool, bool)`

GetResultsDataIsBotOk returns a tuple with the ResultsDataIsBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDataIsBot

`func (o *V1TechnologyUseragentBatchPost200ResponseData) SetResultsDataIsBot(v bool)`

SetResultsDataIsBot sets ResultsDataIsBot field to given value.

### HasResultsDataIsBot

`func (o *V1TechnologyUseragentBatchPost200ResponseData) HasResultsDataIsBot() bool`

HasResultsDataIsBot returns a boolean if a field has been set.

### GetTotal

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1TechnologyUseragentBatchPost200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1TechnologyUseragentBatchPost200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1TechnologyUseragentBatchPost200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


