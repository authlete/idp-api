# WebAuthnChallengeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenge** | Pointer to [**Challenge**](Challenge.md) |  | [optional] 
**Authenticators** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWebAuthnChallengeResponse

`func NewWebAuthnChallengeResponse() *WebAuthnChallengeResponse`

NewWebAuthnChallengeResponse instantiates a new WebAuthnChallengeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAuthnChallengeResponseWithDefaults

`func NewWebAuthnChallengeResponseWithDefaults() *WebAuthnChallengeResponse`

NewWebAuthnChallengeResponseWithDefaults instantiates a new WebAuthnChallengeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenge

`func (o *WebAuthnChallengeResponse) GetChallenge() Challenge`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *WebAuthnChallengeResponse) GetChallengeOk() (*Challenge, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *WebAuthnChallengeResponse) SetChallenge(v Challenge)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *WebAuthnChallengeResponse) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetAuthenticators

`func (o *WebAuthnChallengeResponse) GetAuthenticators() []string`

GetAuthenticators returns the Authenticators field if non-nil, zero value otherwise.

### GetAuthenticatorsOk

`func (o *WebAuthnChallengeResponse) GetAuthenticatorsOk() (*[]string, bool)`

GetAuthenticatorsOk returns a tuple with the Authenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticators

`func (o *WebAuthnChallengeResponse) SetAuthenticators(v []string)`

SetAuthenticators sets Authenticators field to given value.

### HasAuthenticators

`func (o *WebAuthnChallengeResponse) HasAuthenticators() bool`

HasAuthenticators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


