# UserServiceMembershipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizations** | Pointer to [**[]OrganizationMembershipResponse**](OrganizationMembershipResponse.md) |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserServiceMembershipResponse

`func NewUserServiceMembershipResponse() *UserServiceMembershipResponse`

NewUserServiceMembershipResponse instantiates a new UserServiceMembershipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserServiceMembershipResponseWithDefaults

`func NewUserServiceMembershipResponseWithDefaults() *UserServiceMembershipResponse`

NewUserServiceMembershipResponseWithDefaults instantiates a new UserServiceMembershipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizations

`func (o *UserServiceMembershipResponse) GetOrganizations() []OrganizationMembershipResponse`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *UserServiceMembershipResponse) GetOrganizationsOk() (*[]OrganizationMembershipResponse, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *UserServiceMembershipResponse) SetOrganizations(v []OrganizationMembershipResponse)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *UserServiceMembershipResponse) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetUserId

`func (o *UserServiceMembershipResponse) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserServiceMembershipResponse) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserServiceMembershipResponse) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserServiceMembershipResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAdmin

`func (o *UserServiceMembershipResponse) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *UserServiceMembershipResponse) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *UserServiceMembershipResponse) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *UserServiceMembershipResponse) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


