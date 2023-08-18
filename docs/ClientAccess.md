# ClientAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ClientId** | Pointer to **int64** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**ServiceAccessId** | Pointer to **int64** |  | [optional] 

## Methods

### NewClientAccess

`func NewClientAccess() *ClientAccess`

NewClientAccess instantiates a new ClientAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAccessWithDefaults

`func NewClientAccessWithDefaults() *ClientAccess`

NewClientAccessWithDefaults instantiates a new ClientAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientAccess) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientAccess) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientAccess) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClientAccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *ClientAccess) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientAccess) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientAccess) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientAccess) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetPrivileges

`func (o *ClientAccess) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *ClientAccess) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *ClientAccess) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *ClientAccess) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetServiceAccessId

`func (o *ClientAccess) GetServiceAccessId() int64`

GetServiceAccessId returns the ServiceAccessId field if non-nil, zero value otherwise.

### GetServiceAccessIdOk

`func (o *ClientAccess) GetServiceAccessIdOk() (*int64, bool)`

GetServiceAccessIdOk returns a tuple with the ServiceAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccessId

`func (o *ClientAccess) SetServiceAccessId(v int64)`

SetServiceAccessId sets ServiceAccessId field to given value.

### HasServiceAccessId

`func (o *ClientAccess) HasServiceAccessId() bool`

HasServiceAccessId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


