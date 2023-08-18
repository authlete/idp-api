# ServiceAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AllClientPrivileges** | Pointer to **[]string** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**ClientAccess** | Pointer to [**[]ClientAccess**](ClientAccess.md) |  | [optional] 
**ApiServerAccessId** | Pointer to **int64** |  | [optional] 
**ServiceId** | Pointer to [**ServiceInstance**](ServiceInstance.md) |  | [optional] 

## Methods

### NewServiceAccess

`func NewServiceAccess() *ServiceAccess`

NewServiceAccess instantiates a new ServiceAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccessWithDefaults

`func NewServiceAccessWithDefaults() *ServiceAccess`

NewServiceAccessWithDefaults instantiates a new ServiceAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceAccess) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccess) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccess) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceAccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllClientPrivileges

`func (o *ServiceAccess) GetAllClientPrivileges() []string`

GetAllClientPrivileges returns the AllClientPrivileges field if non-nil, zero value otherwise.

### GetAllClientPrivilegesOk

`func (o *ServiceAccess) GetAllClientPrivilegesOk() (*[]string, bool)`

GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClientPrivileges

`func (o *ServiceAccess) SetAllClientPrivileges(v []string)`

SetAllClientPrivileges sets AllClientPrivileges field to given value.

### HasAllClientPrivileges

`func (o *ServiceAccess) HasAllClientPrivileges() bool`

HasAllClientPrivileges returns a boolean if a field has been set.

### GetPrivileges

`func (o *ServiceAccess) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *ServiceAccess) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *ServiceAccess) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *ServiceAccess) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetClientAccess

`func (o *ServiceAccess) GetClientAccess() []ClientAccess`

GetClientAccess returns the ClientAccess field if non-nil, zero value otherwise.

### GetClientAccessOk

`func (o *ServiceAccess) GetClientAccessOk() (*[]ClientAccess, bool)`

GetClientAccessOk returns a tuple with the ClientAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAccess

`func (o *ServiceAccess) SetClientAccess(v []ClientAccess)`

SetClientAccess sets ClientAccess field to given value.

### HasClientAccess

`func (o *ServiceAccess) HasClientAccess() bool`

HasClientAccess returns a boolean if a field has been set.

### GetApiServerAccessId

`func (o *ServiceAccess) GetApiServerAccessId() int64`

GetApiServerAccessId returns the ApiServerAccessId field if non-nil, zero value otherwise.

### GetApiServerAccessIdOk

`func (o *ServiceAccess) GetApiServerAccessIdOk() (*int64, bool)`

GetApiServerAccessIdOk returns a tuple with the ApiServerAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerAccessId

`func (o *ServiceAccess) SetApiServerAccessId(v int64)`

SetApiServerAccessId sets ApiServerAccessId field to given value.

### HasApiServerAccessId

`func (o *ServiceAccess) HasApiServerAccessId() bool`

HasApiServerAccessId returns a boolean if a field has been set.

### GetServiceId

`func (o *ServiceAccess) GetServiceId() ServiceInstance`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceAccess) GetServiceIdOk() (*ServiceInstance, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceAccess) SetServiceId(v ServiceInstance)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceAccess) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


