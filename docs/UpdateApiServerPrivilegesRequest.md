# UpdateApiServerPrivilegesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServerId** | **int64** |  | 
**OrganizationId** | **int64** |  | 
**UserId** | **int64** |  | 
**Privileges** | Pointer to **[]string** |  | [optional] 
**AllServicePrivileges** | Pointer to **[]string** |  | [optional] 
**AllClientPrivileges** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateApiServerPrivilegesRequest

`func NewUpdateApiServerPrivilegesRequest(apiServerId int64, organizationId int64, userId int64, ) *UpdateApiServerPrivilegesRequest`

NewUpdateApiServerPrivilegesRequest instantiates a new UpdateApiServerPrivilegesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApiServerPrivilegesRequestWithDefaults

`func NewUpdateApiServerPrivilegesRequestWithDefaults() *UpdateApiServerPrivilegesRequest`

NewUpdateApiServerPrivilegesRequestWithDefaults instantiates a new UpdateApiServerPrivilegesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServerId

`func (o *UpdateApiServerPrivilegesRequest) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *UpdateApiServerPrivilegesRequest) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *UpdateApiServerPrivilegesRequest) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.


### GetOrganizationId

`func (o *UpdateApiServerPrivilegesRequest) GetOrganizationId() int64`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UpdateApiServerPrivilegesRequest) GetOrganizationIdOk() (*int64, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UpdateApiServerPrivilegesRequest) SetOrganizationId(v int64)`

SetOrganizationId sets OrganizationId field to given value.


### GetUserId

`func (o *UpdateApiServerPrivilegesRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateApiServerPrivilegesRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateApiServerPrivilegesRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetPrivileges

`func (o *UpdateApiServerPrivilegesRequest) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *UpdateApiServerPrivilegesRequest) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *UpdateApiServerPrivilegesRequest) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *UpdateApiServerPrivilegesRequest) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetAllServicePrivileges

`func (o *UpdateApiServerPrivilegesRequest) GetAllServicePrivileges() []string`

GetAllServicePrivileges returns the AllServicePrivileges field if non-nil, zero value otherwise.

### GetAllServicePrivilegesOk

`func (o *UpdateApiServerPrivilegesRequest) GetAllServicePrivilegesOk() (*[]string, bool)`

GetAllServicePrivilegesOk returns a tuple with the AllServicePrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllServicePrivileges

`func (o *UpdateApiServerPrivilegesRequest) SetAllServicePrivileges(v []string)`

SetAllServicePrivileges sets AllServicePrivileges field to given value.

### HasAllServicePrivileges

`func (o *UpdateApiServerPrivilegesRequest) HasAllServicePrivileges() bool`

HasAllServicePrivileges returns a boolean if a field has been set.

### GetAllClientPrivileges

`func (o *UpdateApiServerPrivilegesRequest) GetAllClientPrivileges() []string`

GetAllClientPrivileges returns the AllClientPrivileges field if non-nil, zero value otherwise.

### GetAllClientPrivilegesOk

`func (o *UpdateApiServerPrivilegesRequest) GetAllClientPrivilegesOk() (*[]string, bool)`

GetAllClientPrivilegesOk returns a tuple with the AllClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClientPrivileges

`func (o *UpdateApiServerPrivilegesRequest) SetAllClientPrivileges(v []string)`

SetAllClientPrivileges sets AllClientPrivileges field to given value.

### HasAllClientPrivileges

`func (o *UpdateApiServerPrivilegesRequest) HasAllClientPrivileges() bool`

HasAllClientPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


