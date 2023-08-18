# DeleteServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServerId** | **int64** |  | 
**OrganizationId** | **int64** |  | 
**ServiceId** | **int64** |  | 

## Methods

### NewDeleteServiceRequest

`func NewDeleteServiceRequest(apiServerId int64, organizationId int64, serviceId int64, ) *DeleteServiceRequest`

NewDeleteServiceRequest instantiates a new DeleteServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteServiceRequestWithDefaults

`func NewDeleteServiceRequestWithDefaults() *DeleteServiceRequest`

NewDeleteServiceRequestWithDefaults instantiates a new DeleteServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServerId

`func (o *DeleteServiceRequest) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *DeleteServiceRequest) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *DeleteServiceRequest) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.


### GetOrganizationId

`func (o *DeleteServiceRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DeleteServiceRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DeleteServiceRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetServiceId

`func (o *DeleteServiceRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DeleteServiceRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DeleteServiceRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


