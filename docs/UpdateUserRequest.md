# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**GivenName** | Pointer to **string** |  | [optional] 
**FamilyName** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest(email string, ) *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGivenName

`func (o *UpdateUserRequest) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *UpdateUserRequest) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *UpdateUserRequest) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *UpdateUserRequest) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetFamilyName

`func (o *UpdateUserRequest) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *UpdateUserRequest) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *UpdateUserRequest) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *UpdateUserRequest) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


