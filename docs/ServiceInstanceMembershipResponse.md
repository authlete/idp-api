# ServiceInstanceMembershipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **int64** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**Clients** | Pointer to [**[]ClientAccessResponse**](ClientAccessResponse.md) |  | [optional] 
**ClientPrivileges** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceInstanceMembershipResponse

`func NewServiceInstanceMembershipResponse() *ServiceInstanceMembershipResponse`

NewServiceInstanceMembershipResponse instantiates a new ServiceInstanceMembershipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstanceMembershipResponseWithDefaults

`func NewServiceInstanceMembershipResponseWithDefaults() *ServiceInstanceMembershipResponse`

NewServiceInstanceMembershipResponseWithDefaults instantiates a new ServiceInstanceMembershipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ServiceInstanceMembershipResponse) GetServiceId() int64`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceInstanceMembershipResponse) GetServiceIdOk() (*int64, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceInstanceMembershipResponse) SetServiceId(v int64)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceInstanceMembershipResponse) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPrivileges

`func (o *ServiceInstanceMembershipResponse) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *ServiceInstanceMembershipResponse) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *ServiceInstanceMembershipResponse) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *ServiceInstanceMembershipResponse) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetClients

`func (o *ServiceInstanceMembershipResponse) GetClients() []ClientAccessResponse`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ServiceInstanceMembershipResponse) GetClientsOk() (*[]ClientAccessResponse, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ServiceInstanceMembershipResponse) SetClients(v []ClientAccessResponse)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ServiceInstanceMembershipResponse) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetClientPrivileges

`func (o *ServiceInstanceMembershipResponse) GetClientPrivileges() []string`

GetClientPrivileges returns the ClientPrivileges field if non-nil, zero value otherwise.

### GetClientPrivilegesOk

`func (o *ServiceInstanceMembershipResponse) GetClientPrivilegesOk() (*[]string, bool)`

GetClientPrivilegesOk returns a tuple with the ClientPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivileges

`func (o *ServiceInstanceMembershipResponse) SetClientPrivileges(v []string)`

SetClientPrivileges sets ClientPrivileges field to given value.

### HasClientPrivileges

`func (o *ServiceInstanceMembershipResponse) HasClientPrivileges() bool`

HasClientPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


