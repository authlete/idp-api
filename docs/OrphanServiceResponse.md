# OrphanServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnApiButNotOnIdp** | Pointer to [**[]OrphanService**](OrphanService.md) |  | [optional] 
**OnIdPButNotOnApi** | Pointer to [**[]OrphanService**](OrphanService.md) |  | [optional] 

## Methods

### NewOrphanServiceResponse

`func NewOrphanServiceResponse() *OrphanServiceResponse`

NewOrphanServiceResponse instantiates a new OrphanServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrphanServiceResponseWithDefaults

`func NewOrphanServiceResponseWithDefaults() *OrphanServiceResponse`

NewOrphanServiceResponseWithDefaults instantiates a new OrphanServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnApiButNotOnIdp

`func (o *OrphanServiceResponse) GetOnApiButNotOnIdp() []OrphanService`

GetOnApiButNotOnIdp returns the OnApiButNotOnIdp field if non-nil, zero value otherwise.

### GetOnApiButNotOnIdpOk

`func (o *OrphanServiceResponse) GetOnApiButNotOnIdpOk() (*[]OrphanService, bool)`

GetOnApiButNotOnIdpOk returns a tuple with the OnApiButNotOnIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnApiButNotOnIdp

`func (o *OrphanServiceResponse) SetOnApiButNotOnIdp(v []OrphanService)`

SetOnApiButNotOnIdp sets OnApiButNotOnIdp field to given value.

### HasOnApiButNotOnIdp

`func (o *OrphanServiceResponse) HasOnApiButNotOnIdp() bool`

HasOnApiButNotOnIdp returns a boolean if a field has been set.

### GetOnIdPButNotOnApi

`func (o *OrphanServiceResponse) GetOnIdPButNotOnApi() []OrphanService`

GetOnIdPButNotOnApi returns the OnIdPButNotOnApi field if non-nil, zero value otherwise.

### GetOnIdPButNotOnApiOk

`func (o *OrphanServiceResponse) GetOnIdPButNotOnApiOk() (*[]OrphanService, bool)`

GetOnIdPButNotOnApiOk returns a tuple with the OnIdPButNotOnApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnIdPButNotOnApi

`func (o *OrphanServiceResponse) SetOnIdPButNotOnApi(v []OrphanService)`

SetOnIdPButNotOnApi sets OnIdPButNotOnApi field to given value.

### HasOnIdPButNotOnApi

`func (o *OrphanServiceResponse) HasOnIdPButNotOnApi() bool`

HasOnIdPButNotOnApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


