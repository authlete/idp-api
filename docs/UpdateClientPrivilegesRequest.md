# UpdateClientPrivilegesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServerId** | **int64** |  | 
**OrganizationId** | **int64** |  | 
**UserId** | **int64** |  | 
**ServiceId** | **int64** |  | 
**ClientId** | **int64** |  | 
**Privileges** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateClientPrivilegesRequest

`func NewUpdateClientPrivilegesRequest(apiServerId int64, organizationId int64, userId int64, serviceId int64, clientId int64, ) *UpdateClientPrivilegesRequest`

NewUpdateClientPrivilegesRequest instantiates a new UpdateClientPrivilegesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClientPrivilegesRequestWithDefaults

`func NewUpdateClientPrivilegesRequestWithDefaults() *UpdateClientPrivilegesRequest`

NewUpdateClientPrivilegesRequestWithDefaults instantiates a new UpdateClientPrivilegesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServerId

`func (o *UpdateClientPrivilegesRequest) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *UpdateClientPrivilegesRequest) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *UpdateClientPrivilegesRequest) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.


### GetOrganizationId

`func (o *UpdateClientPrivilegesRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UpdateClientPrivilegesRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UpdateClientPrivilegesRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetUserId

`func (o *UpdateClientPrivilegesRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateClientPrivilegesRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateClientPrivilegesRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetServiceId

`func (o *UpdateClientPrivilegesRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateClientPrivilegesRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateClientPrivilegesRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetClientId

`func (o *UpdateClientPrivilegesRequest) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UpdateClientPrivilegesRequest) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UpdateClientPrivilegesRequest) SetClientId(v int64)`

SetClientId sets ClientId field to given value.


### GetPrivileges

`func (o *UpdateClientPrivilegesRequest) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *UpdateClientPrivilegesRequest) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *UpdateClientPrivilegesRequest) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *UpdateClientPrivilegesRequest) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


