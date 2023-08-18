# CreateServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServerId** | **int64** |  | 
**OrganizationId** | **int64** |  | 
**Service** | Pointer to [**Service**](Service.md) |  | [optional] 

## Methods

### NewCreateServiceRequest

`func NewCreateServiceRequest(apiServerId int64, organizationId int64, ) *CreateServiceRequest`

NewCreateServiceRequest instantiates a new CreateServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceRequestWithDefaults

`func NewCreateServiceRequestWithDefaults() *CreateServiceRequest`

NewCreateServiceRequestWithDefaults instantiates a new CreateServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServerId

`func (o *CreateServiceRequest) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *CreateServiceRequest) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *CreateServiceRequest) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.


### GetOrganizationId

`func (o *CreateServiceRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateServiceRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateServiceRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetService

`func (o *CreateServiceRequest) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateServiceRequest) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateServiceRequest) SetService(v Service)`

SetService sets Service field to given value.

### HasService

`func (o *CreateServiceRequest) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


