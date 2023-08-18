# AuthorizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** |  | [optional] 
**ResultMessage** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Service** | Pointer to [**Service**](Service.md) |  | [optional] 
**Client** | Pointer to [**Client**](Client.md) |  | [optional] 
**Display** | Pointer to **string** |  | [optional] 
**MaxAge** | Pointer to **int32** |  | [optional] 
**Scopes** | Pointer to [**[]Scope**](Scope.md) |  | [optional] 
**DynamicScopes** | Pointer to [**[]DynamicScope**](DynamicScope.md) |  | [optional] 
**UiLocales** | Pointer to **[]string** |  | [optional] 
**ClaimsLocales** | Pointer to **[]string** |  | [optional] 
**Claims** | Pointer to **[]string** |  | [optional] 
**ClaimsAtUserInfo** | Pointer to **[]string** |  | [optional] 
**AcrEssential** | Pointer to **bool** |  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** |  | [optional] 
**ClientEntityIdUsed** | Pointer to **bool** |  | [optional] 
**Acrs** | Pointer to **[]string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**LoginHint** | Pointer to **string** |  | [optional] 
**LowestPrompt** | Pointer to **string** |  | [optional] 
**Prompts** | Pointer to **[]string** |  | [optional] 
**RequestObjectPayload** | Pointer to **string** |  | [optional] 
**IdTokenClaims** | Pointer to **string** |  | [optional] 
**UserInfoClaims** | Pointer to **string** |  | [optional] 
**TransformedClaims** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to **[]string** |  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthzDetails**](AuthzDetails.md) |  | [optional] 
**Purpose** | Pointer to **string** |  | [optional] 
**GmAction** | Pointer to **string** |  | [optional] 
**GrantId** | Pointer to **string** |  | [optional] 
**GrantSubject** | Pointer to **string** |  | [optional] 
**Grant** | Pointer to [**Grant**](Grant.md) |  | [optional] 
**RequestedClaimsForTx** | Pointer to **[]string** |  | [optional] 
**RequestedVerifiedClaimsForTx** | Pointer to [**[]StringArray**](StringArray.md) |  | [optional] 
**ResponseContent** | Pointer to **string** |  | [optional] 
**Ticket** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthorizationResponse

`func NewAuthorizationResponse() *AuthorizationResponse`

NewAuthorizationResponse instantiates a new AuthorizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationResponseWithDefaults

`func NewAuthorizationResponseWithDefaults() *AuthorizationResponse`

NewAuthorizationResponseWithDefaults instantiates a new AuthorizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *AuthorizationResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *AuthorizationResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *AuthorizationResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *AuthorizationResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *AuthorizationResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *AuthorizationResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *AuthorizationResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *AuthorizationResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *AuthorizationResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuthorizationResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuthorizationResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuthorizationResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetService

`func (o *AuthorizationResponse) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AuthorizationResponse) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AuthorizationResponse) SetService(v Service)`

SetService sets Service field to given value.

### HasService

`func (o *AuthorizationResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetClient

`func (o *AuthorizationResponse) GetClient() Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *AuthorizationResponse) GetClientOk() (*Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *AuthorizationResponse) SetClient(v Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *AuthorizationResponse) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetDisplay

`func (o *AuthorizationResponse) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *AuthorizationResponse) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *AuthorizationResponse) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *AuthorizationResponse) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetMaxAge

`func (o *AuthorizationResponse) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *AuthorizationResponse) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *AuthorizationResponse) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *AuthorizationResponse) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetScopes

`func (o *AuthorizationResponse) GetScopes() []Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthorizationResponse) GetScopesOk() (*[]Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthorizationResponse) SetScopes(v []Scope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthorizationResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetDynamicScopes

`func (o *AuthorizationResponse) GetDynamicScopes() []DynamicScope`

GetDynamicScopes returns the DynamicScopes field if non-nil, zero value otherwise.

### GetDynamicScopesOk

`func (o *AuthorizationResponse) GetDynamicScopesOk() (*[]DynamicScope, bool)`

GetDynamicScopesOk returns a tuple with the DynamicScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicScopes

`func (o *AuthorizationResponse) SetDynamicScopes(v []DynamicScope)`

SetDynamicScopes sets DynamicScopes field to given value.

### HasDynamicScopes

`func (o *AuthorizationResponse) HasDynamicScopes() bool`

HasDynamicScopes returns a boolean if a field has been set.

### GetUiLocales

`func (o *AuthorizationResponse) GetUiLocales() []string`

GetUiLocales returns the UiLocales field if non-nil, zero value otherwise.

### GetUiLocalesOk

`func (o *AuthorizationResponse) GetUiLocalesOk() (*[]string, bool)`

GetUiLocalesOk returns a tuple with the UiLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiLocales

`func (o *AuthorizationResponse) SetUiLocales(v []string)`

SetUiLocales sets UiLocales field to given value.

### HasUiLocales

`func (o *AuthorizationResponse) HasUiLocales() bool`

HasUiLocales returns a boolean if a field has been set.

### GetClaimsLocales

`func (o *AuthorizationResponse) GetClaimsLocales() []string`

GetClaimsLocales returns the ClaimsLocales field if non-nil, zero value otherwise.

### GetClaimsLocalesOk

`func (o *AuthorizationResponse) GetClaimsLocalesOk() (*[]string, bool)`

GetClaimsLocalesOk returns a tuple with the ClaimsLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsLocales

`func (o *AuthorizationResponse) SetClaimsLocales(v []string)`

SetClaimsLocales sets ClaimsLocales field to given value.

### HasClaimsLocales

`func (o *AuthorizationResponse) HasClaimsLocales() bool`

HasClaimsLocales returns a boolean if a field has been set.

### GetClaims

`func (o *AuthorizationResponse) GetClaims() []string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *AuthorizationResponse) GetClaimsOk() (*[]string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *AuthorizationResponse) SetClaims(v []string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *AuthorizationResponse) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetClaimsAtUserInfo

`func (o *AuthorizationResponse) GetClaimsAtUserInfo() []string`

GetClaimsAtUserInfo returns the ClaimsAtUserInfo field if non-nil, zero value otherwise.

### GetClaimsAtUserInfoOk

`func (o *AuthorizationResponse) GetClaimsAtUserInfoOk() (*[]string, bool)`

GetClaimsAtUserInfoOk returns a tuple with the ClaimsAtUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsAtUserInfo

`func (o *AuthorizationResponse) SetClaimsAtUserInfo(v []string)`

SetClaimsAtUserInfo sets ClaimsAtUserInfo field to given value.

### HasClaimsAtUserInfo

`func (o *AuthorizationResponse) HasClaimsAtUserInfo() bool`

HasClaimsAtUserInfo returns a boolean if a field has been set.

### GetAcrEssential

`func (o *AuthorizationResponse) GetAcrEssential() bool`

GetAcrEssential returns the AcrEssential field if non-nil, zero value otherwise.

### GetAcrEssentialOk

`func (o *AuthorizationResponse) GetAcrEssentialOk() (*bool, bool)`

GetAcrEssentialOk returns a tuple with the AcrEssential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrEssential

`func (o *AuthorizationResponse) SetAcrEssential(v bool)`

SetAcrEssential sets AcrEssential field to given value.

### HasAcrEssential

`func (o *AuthorizationResponse) HasAcrEssential() bool`

HasAcrEssential returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *AuthorizationResponse) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *AuthorizationResponse) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *AuthorizationResponse) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *AuthorizationResponse) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetClientEntityIdUsed

`func (o *AuthorizationResponse) GetClientEntityIdUsed() bool`

GetClientEntityIdUsed returns the ClientEntityIdUsed field if non-nil, zero value otherwise.

### GetClientEntityIdUsedOk

`func (o *AuthorizationResponse) GetClientEntityIdUsedOk() (*bool, bool)`

GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityIdUsed

`func (o *AuthorizationResponse) SetClientEntityIdUsed(v bool)`

SetClientEntityIdUsed sets ClientEntityIdUsed field to given value.

### HasClientEntityIdUsed

`func (o *AuthorizationResponse) HasClientEntityIdUsed() bool`

HasClientEntityIdUsed returns a boolean if a field has been set.

### GetAcrs

`func (o *AuthorizationResponse) GetAcrs() []string`

GetAcrs returns the Acrs field if non-nil, zero value otherwise.

### GetAcrsOk

`func (o *AuthorizationResponse) GetAcrsOk() (*[]string, bool)`

GetAcrsOk returns a tuple with the Acrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrs

`func (o *AuthorizationResponse) SetAcrs(v []string)`

SetAcrs sets Acrs field to given value.

### HasAcrs

`func (o *AuthorizationResponse) HasAcrs() bool`

HasAcrs returns a boolean if a field has been set.

### GetSubject

`func (o *AuthorizationResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AuthorizationResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AuthorizationResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AuthorizationResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetLoginHint

`func (o *AuthorizationResponse) GetLoginHint() string`

GetLoginHint returns the LoginHint field if non-nil, zero value otherwise.

### GetLoginHintOk

`func (o *AuthorizationResponse) GetLoginHintOk() (*string, bool)`

GetLoginHintOk returns a tuple with the LoginHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginHint

`func (o *AuthorizationResponse) SetLoginHint(v string)`

SetLoginHint sets LoginHint field to given value.

### HasLoginHint

`func (o *AuthorizationResponse) HasLoginHint() bool`

HasLoginHint returns a boolean if a field has been set.

### GetLowestPrompt

`func (o *AuthorizationResponse) GetLowestPrompt() string`

GetLowestPrompt returns the LowestPrompt field if non-nil, zero value otherwise.

### GetLowestPromptOk

`func (o *AuthorizationResponse) GetLowestPromptOk() (*string, bool)`

GetLowestPromptOk returns a tuple with the LowestPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestPrompt

`func (o *AuthorizationResponse) SetLowestPrompt(v string)`

SetLowestPrompt sets LowestPrompt field to given value.

### HasLowestPrompt

`func (o *AuthorizationResponse) HasLowestPrompt() bool`

HasLowestPrompt returns a boolean if a field has been set.

### GetPrompts

`func (o *AuthorizationResponse) GetPrompts() []string`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *AuthorizationResponse) GetPromptsOk() (*[]string, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *AuthorizationResponse) SetPrompts(v []string)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *AuthorizationResponse) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### GetRequestObjectPayload

`func (o *AuthorizationResponse) GetRequestObjectPayload() string`

GetRequestObjectPayload returns the RequestObjectPayload field if non-nil, zero value otherwise.

### GetRequestObjectPayloadOk

`func (o *AuthorizationResponse) GetRequestObjectPayloadOk() (*string, bool)`

GetRequestObjectPayloadOk returns a tuple with the RequestObjectPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectPayload

`func (o *AuthorizationResponse) SetRequestObjectPayload(v string)`

SetRequestObjectPayload sets RequestObjectPayload field to given value.

### HasRequestObjectPayload

`func (o *AuthorizationResponse) HasRequestObjectPayload() bool`

HasRequestObjectPayload returns a boolean if a field has been set.

### GetIdTokenClaims

`func (o *AuthorizationResponse) GetIdTokenClaims() string`

GetIdTokenClaims returns the IdTokenClaims field if non-nil, zero value otherwise.

### GetIdTokenClaimsOk

`func (o *AuthorizationResponse) GetIdTokenClaimsOk() (*string, bool)`

GetIdTokenClaimsOk returns a tuple with the IdTokenClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenClaims

`func (o *AuthorizationResponse) SetIdTokenClaims(v string)`

SetIdTokenClaims sets IdTokenClaims field to given value.

### HasIdTokenClaims

`func (o *AuthorizationResponse) HasIdTokenClaims() bool`

HasIdTokenClaims returns a boolean if a field has been set.

### GetUserInfoClaims

`func (o *AuthorizationResponse) GetUserInfoClaims() string`

GetUserInfoClaims returns the UserInfoClaims field if non-nil, zero value otherwise.

### GetUserInfoClaimsOk

`func (o *AuthorizationResponse) GetUserInfoClaimsOk() (*string, bool)`

GetUserInfoClaimsOk returns a tuple with the UserInfoClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoClaims

`func (o *AuthorizationResponse) SetUserInfoClaims(v string)`

SetUserInfoClaims sets UserInfoClaims field to given value.

### HasUserInfoClaims

`func (o *AuthorizationResponse) HasUserInfoClaims() bool`

HasUserInfoClaims returns a boolean if a field has been set.

### GetTransformedClaims

`func (o *AuthorizationResponse) GetTransformedClaims() string`

GetTransformedClaims returns the TransformedClaims field if non-nil, zero value otherwise.

### GetTransformedClaimsOk

`func (o *AuthorizationResponse) GetTransformedClaimsOk() (*string, bool)`

GetTransformedClaimsOk returns a tuple with the TransformedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformedClaims

`func (o *AuthorizationResponse) SetTransformedClaims(v string)`

SetTransformedClaims sets TransformedClaims field to given value.

### HasTransformedClaims

`func (o *AuthorizationResponse) HasTransformedClaims() bool`

HasTransformedClaims returns a boolean if a field has been set.

### GetResources

`func (o *AuthorizationResponse) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AuthorizationResponse) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AuthorizationResponse) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AuthorizationResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *AuthorizationResponse) GetAuthorizationDetails() AuthzDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *AuthorizationResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *AuthorizationResponse) SetAuthorizationDetails(v AuthzDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *AuthorizationResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetPurpose

`func (o *AuthorizationResponse) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *AuthorizationResponse) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *AuthorizationResponse) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *AuthorizationResponse) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetGmAction

`func (o *AuthorizationResponse) GetGmAction() string`

GetGmAction returns the GmAction field if non-nil, zero value otherwise.

### GetGmActionOk

`func (o *AuthorizationResponse) GetGmActionOk() (*string, bool)`

GetGmActionOk returns a tuple with the GmAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmAction

`func (o *AuthorizationResponse) SetGmAction(v string)`

SetGmAction sets GmAction field to given value.

### HasGmAction

`func (o *AuthorizationResponse) HasGmAction() bool`

HasGmAction returns a boolean if a field has been set.

### GetGrantId

`func (o *AuthorizationResponse) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *AuthorizationResponse) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *AuthorizationResponse) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.

### HasGrantId

`func (o *AuthorizationResponse) HasGrantId() bool`

HasGrantId returns a boolean if a field has been set.

### GetGrantSubject

`func (o *AuthorizationResponse) GetGrantSubject() string`

GetGrantSubject returns the GrantSubject field if non-nil, zero value otherwise.

### GetGrantSubjectOk

`func (o *AuthorizationResponse) GetGrantSubjectOk() (*string, bool)`

GetGrantSubjectOk returns a tuple with the GrantSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantSubject

`func (o *AuthorizationResponse) SetGrantSubject(v string)`

SetGrantSubject sets GrantSubject field to given value.

### HasGrantSubject

`func (o *AuthorizationResponse) HasGrantSubject() bool`

HasGrantSubject returns a boolean if a field has been set.

### GetGrant

`func (o *AuthorizationResponse) GetGrant() Grant`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *AuthorizationResponse) GetGrantOk() (*Grant, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *AuthorizationResponse) SetGrant(v Grant)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *AuthorizationResponse) HasGrant() bool`

HasGrant returns a boolean if a field has been set.

### GetRequestedClaimsForTx

`func (o *AuthorizationResponse) GetRequestedClaimsForTx() []string`

GetRequestedClaimsForTx returns the RequestedClaimsForTx field if non-nil, zero value otherwise.

### GetRequestedClaimsForTxOk

`func (o *AuthorizationResponse) GetRequestedClaimsForTxOk() (*[]string, bool)`

GetRequestedClaimsForTxOk returns a tuple with the RequestedClaimsForTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedClaimsForTx

`func (o *AuthorizationResponse) SetRequestedClaimsForTx(v []string)`

SetRequestedClaimsForTx sets RequestedClaimsForTx field to given value.

### HasRequestedClaimsForTx

`func (o *AuthorizationResponse) HasRequestedClaimsForTx() bool`

HasRequestedClaimsForTx returns a boolean if a field has been set.

### GetRequestedVerifiedClaimsForTx

`func (o *AuthorizationResponse) GetRequestedVerifiedClaimsForTx() []StringArray`

GetRequestedVerifiedClaimsForTx returns the RequestedVerifiedClaimsForTx field if non-nil, zero value otherwise.

### GetRequestedVerifiedClaimsForTxOk

`func (o *AuthorizationResponse) GetRequestedVerifiedClaimsForTxOk() (*[]StringArray, bool)`

GetRequestedVerifiedClaimsForTxOk returns a tuple with the RequestedVerifiedClaimsForTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedVerifiedClaimsForTx

`func (o *AuthorizationResponse) SetRequestedVerifiedClaimsForTx(v []StringArray)`

SetRequestedVerifiedClaimsForTx sets RequestedVerifiedClaimsForTx field to given value.

### HasRequestedVerifiedClaimsForTx

`func (o *AuthorizationResponse) HasRequestedVerifiedClaimsForTx() bool`

HasRequestedVerifiedClaimsForTx returns a boolean if a field has been set.

### GetResponseContent

`func (o *AuthorizationResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *AuthorizationResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *AuthorizationResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *AuthorizationResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetTicket

`func (o *AuthorizationResponse) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *AuthorizationResponse) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *AuthorizationResponse) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *AuthorizationResponse) HasTicket() bool`

HasTicket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


