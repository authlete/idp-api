# UpdateServicePrivilegesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServerId** | **int64** |  | 
**OrganizationId** | **int64** |  | 
**UserId** | **int64** |  | 
**ServiceId** | **int64** |  | 
**Privileges** | Pointer to **[]string** |  | [optional] 
**AllClientPrivileges** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateServicePrivilegesRequest

`func NewUpdateServicePrivilegesRequest(apiServerId int64, organizationId int64, userId int64, serviceId int64, ) *UpdateServicePrivilegesRequest`

NewUpdateServicePrivilegesRequest instantiates a new UpdateServicePrivilegesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServicePrivilegesRequestWithDefaults

`func NewUpdateServicePrivilegesRequestWithDefaults() *UpdateServicePrivilegesRequest`

NewUpdateServicePrivilegesRequestWithDefaults instantiates a new UpdateServicePrivilegesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServerId

`func (o *UpdateServicePrivilegesRequest) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *UpdateServicePrivilegesRequest) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *UpdateServicePrivilegesRequest) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.


### GetOrganizationId

`func (o *UpdateServicePrivilegesRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UpdateServicePrivilegesRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UpdateServicePrivilegesRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetUserId

`func (o *UpdateServicePrivilegesRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateServicePrivilegesRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateServicePrivilegesRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetServiceId

`func (o *UpdateServicePrivilegesRequest) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateServicePrivilegesRequest) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateServicePrivilegesRequest) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.


### GetPrivileges

`func (o *UpdateServicePrivilegesRequest) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *UpdateServicePrivilegesRequest) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *UpdateServicePrivilegesRequest) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *UpdateServicePrivilegesRequest) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetAllClientPrivileges

`func (o *UpdateServicePrivilegesRequest) GetAllClientPrivileges() []string`

GetAllClientPrivileges returns the AllClientPrivileges field if non-nil, zero value otherwise.

### GetAllClientPrivilegesOk

`func (o *UpdateServicePrivilegesRequest) GetAllClientPrivilegesOk() (*[]string, bool)`

GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClientPrivileges

`func (o *UpdateServicePrivilegesRequest) SetAllClientPrivileges(v []string)`

SetAllClientPrivileges sets AllClientPrivileges field to given value.

### HasAllClientPrivileges

`func (o *UpdateServicePrivilegesRequest) HasAllClientPrivileges() bool`

HasAllClientPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


