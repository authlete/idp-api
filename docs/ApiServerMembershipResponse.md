# ApiServerMembershipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServerId** | Pointer to **int64** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Privileges** | Pointer to **[]string** |  | [optional] 
**Services** | Pointer to [**[]ServiceInstanceMembershipResponse**](ServiceInstanceMembershipResponse.md) |  | [optional] 

## Methods

### NewApiServerMembershipResponse

`func NewApiServerMembershipResponse() *ApiServerMembershipResponse`

NewApiServerMembershipResponse instantiates a new ApiServerMembershipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServerMembershipResponseWithDefaults

`func NewApiServerMembershipResponseWithDefaults() *ApiServerMembershipResponse`

NewApiServerMembershipResponseWithDefaults instantiates a new ApiServerMembershipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServerId

`func (o *ApiServerMembershipResponse) GetApiServerId() int64`

GetApiServerId returns the ApiServerId field if non-nil, zero value otherwise.

### GetApiServerIdOk

`func (o *ApiServerMembershipResponse) GetApiServerIdOk() (*int64, bool)`

GetApiServerIdOk returns a tuple with the ApiServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerId

`func (o *ApiServerMembershipResponse) SetApiServerId(v int64)`

SetApiServerId sets ApiServerId field to given value.

### HasApiServerId

`func (o *ApiServerMembershipResponse) HasApiServerId() bool`

HasApiServerId returns a boolean if a field has been set.

### GetUrl

`func (o *ApiServerMembershipResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiServerMembershipResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiServerMembershipResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApiServerMembershipResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDescription

`func (o *ApiServerMembershipResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiServerMembershipResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiServerMembershipResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiServerMembershipResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrivileges

`func (o *ApiServerMembershipResponse) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *ApiServerMembershipResponse) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *ApiServerMembershipResponse) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *ApiServerMembershipResponse) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetServices

`func (o *ApiServerMembershipResponse) GetServices() []ServiceInstanceMembershipResponse`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ApiServerMembershipResponse) GetServicesOk() (*[]ServiceInstanceMembershipResponse, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ApiServerMembershipResponse) SetServices(v []ServiceInstanceMembershipResponse)`

SetServices sets Services field to given value.

### HasServices

`func (o *ApiServerMembershipResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


