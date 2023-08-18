# ClientExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestableScopesEnabled** | Pointer to **bool** |  | [optional] 
**RequestableScopes** | Pointer to **[]string** |  | [optional] 
**AccessTokenDuration** | Pointer to **int64** |  | [optional] 
**RefreshTokenDuration** | Pointer to **int64** |  | [optional] 
**TokenExchangePermitted** | Pointer to **bool** |  | [optional] 

## Methods

### NewClientExtension

`func NewClientExtension() *ClientExtension`

NewClientExtension instantiates a new ClientExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientExtensionWithDefaults

`func NewClientExtensionWithDefaults() *ClientExtension`

NewClientExtensionWithDefaults instantiates a new ClientExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestableScopesEnabled

`func (o *ClientExtension) GetRequestableScopesEnabled() bool`

GetRequestableScopesEnabled returns the RequestableScopesEnabled field if non-nil, zero value otherwise.

### GetRequestableScopesEnabledOk

`func (o *ClientExtension) GetRequestableScopesEnabledOk() (*bool, bool)`

GetRequestableScopesEnabledOk returns a tuple with the RequestableScopesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableScopesEnabled

`func (o *ClientExtension) SetRequestableScopesEnabled(v bool)`

SetRequestableScopesEnabled sets RequestableScopesEnabled field to given value.

### HasRequestableScopesEnabled

`func (o *ClientExtension) HasRequestableScopesEnabled() bool`

HasRequestableScopesEnabled returns a boolean if a field has been set.

### GetRequestableScopes

`func (o *ClientExtension) GetRequestableScopes() []string`

GetRequestableScopes returns the RequestableScopes field if non-nil, zero value otherwise.

### GetRequestableScopesOk

`func (o *ClientExtension) GetRequestableScopesOk() (*[]string, bool)`

GetRequestableScopesOk returns a tuple with the RequestableScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableScopes

`func (o *ClientExtension) SetRequestableScopes(v []string)`

SetRequestableScopes sets RequestableScopes field to given value.

### HasRequestableScopes

`func (o *ClientExtension) HasRequestableScopes() bool`

HasRequestableScopes returns a boolean if a field has been set.

### GetAccessTokenDuration

`func (o *ClientExtension) GetAccessTokenDuration() int64`

GetAccessTokenDuration returns the AccessTokenDuration field if non-nil, zero value otherwise.

### GetAccessTokenDurationOk

`func (o *ClientExtension) GetAccessTokenDurationOk() (*int64, bool)`

GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDuration

`func (o *ClientExtension) SetAccessTokenDuration(v int64)`

SetAccessTokenDuration sets AccessTokenDuration field to given value.

### HasAccessTokenDuration

`func (o *ClientExtension) HasAccessTokenDuration() bool`

HasAccessTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *ClientExtension) GetRefreshTokenDuration() int64`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *ClientExtension) GetRefreshTokenDurationOk() (*int64, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *ClientExtension) SetRefreshTokenDuration(v int64)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *ClientExtension) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetTokenExchangePermitted

`func (o *ClientExtension) GetTokenExchangePermitted() bool`

GetTokenExchangePermitted returns the TokenExchangePermitted field if non-nil, zero value otherwise.

### GetTokenExchangePermittedOk

`func (o *ClientExtension) GetTokenExchangePermittedOk() (*bool, bool)`

GetTokenExchangePermittedOk returns a tuple with the TokenExchangePermitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangePermitted

`func (o *ClientExtension) SetTokenExchangePermitted(v bool)`

SetTokenExchangePermitted sets TokenExchangePermitted field to given value.

### HasTokenExchangePermitted

`func (o *ClientExtension) HasTokenExchangePermitted() bool`

HasTokenExchangePermitted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


