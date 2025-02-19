# \UserApiAPI

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindTotpCredentials**](UserApiAPI.md#BindTotpCredentials) | **Post** /api/user/{id}/totp | 
[**BindWebAuthnChallenge**](UserApiAPI.md#BindWebAuthnChallenge) | **Post** /api/user/{id}/webauthn | 
[**CreateTotpCredentials**](UserApiAPI.md#CreateTotpCredentials) | **Get** /api/user/{id}/totp | 
[**CreateUser**](UserApiAPI.md#CreateUser) | **Post** /api/user | 
[**CreateWebAuthnChallenge**](UserApiAPI.md#CreateWebAuthnChallenge) | **Get** /api/user/{id}/webauthn | 
[**DeleteUser**](UserApiAPI.md#DeleteUser) | **Delete** /api/user/{id} | 
[**SetAdmin**](UserApiAPI.md#SetAdmin) | **Post** /api/user/{id}/admin | 
[**UnsetAdmin**](UserApiAPI.md#UnsetAdmin) | **Delete** /api/user/{id}/admin | 
[**UpdateUser**](UserApiAPI.md#UpdateUser) | **Post** /api/user/{id} | 



## BindTotpCredentials

> BindTotpCredentials(ctx, id).BindTotpCredentialsRequest(bindTotpCredentialsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 
	bindTotpCredentialsRequest := *openapiclient.NewBindTotpCredentialsRequest(int32(123)) // BindTotpCredentialsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserApiAPI.BindTotpCredentials(context.Background(), id).BindTotpCredentialsRequest(bindTotpCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApiAPI.BindTotpCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindTotpCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindTotpCredentialsRequest** | [**BindTotpCredentialsRequest**](BindTotpCredentialsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BindWebAuthnChallenge

> string BindWebAuthnChallenge(ctx, id).WebAuthnChallengeRequest(webAuthnChallengeRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 
	webAuthnChallengeRequest := *openapiclient.NewWebAuthnChallengeRequest() // WebAuthnChallengeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApiAPI.BindWebAuthnChallenge(context.Background(), id).WebAuthnChallengeRequest(webAuthnChallengeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApiAPI.BindWebAuthnChallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindWebAuthnChallenge`: string
	fmt.Fprintf(os.Stdout, "Response from `UserApiAPI.BindWebAuthnChallenge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindWebAuthnChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webAuthnChallengeRequest** | [**WebAuthnChallengeRequest**](WebAuthnChallengeRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTotpCredentials

> GoogleAuthenticatorKey CreateTotpCredentials(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApiAPI.CreateTotpCredentials(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApiAPI.CreateTotpCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTotpCredentials`: GoogleAuthenticatorKey
	fmt.Fprintf(os.Stdout, "Response from `UserApiAPI.CreateTotpCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTotpCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GoogleAuthenticatorKey**](GoogleAuthenticatorKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> CreateUserResponse CreateUser(ctx).CreateUserRequest(createUserRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	createUserRequest := *openapiclient.NewCreateUserRequest("Email_example", "Password_example") // CreateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApiAPI.CreateUser(context.Background()).CreateUserRequest(createUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApiAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: CreateUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserApiAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) |  | 

### Return type

[**CreateUserResponse**](CreateUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebAuthnChallenge

> Challenge CreateWebAuthnChallenge(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApiAPI.CreateWebAuthnChallenge(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApiAPI.CreateWebAuthnChallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebAuthnChallenge`: Challenge
	fmt.Fprintf(os.Stdout, "Response from `UserApiAPI.CreateWebAuthnChallenge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebAuthnChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Challenge**](Challenge.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserApiAPI.DeleteUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApiAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAdmin

> AuthleteUser SetAdmin(ctx, id).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApiAPI.SetAdmin(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApiAPI.SetAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAdmin`: AuthleteUser
	fmt.Fprintf(os.Stdout, "Response from `UserApiAPI.SetAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**AuthleteUser**](AuthleteUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetAdmin

> AuthleteUser UnsetAdmin(ctx, id).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApiAPI.UnsetAdmin(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApiAPI.UnsetAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsetAdmin`: AuthleteUser
	fmt.Fprintf(os.Stdout, "Response from `UserApiAPI.UnsetAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**AuthleteUser**](AuthleteUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> AuthleteUser UpdateUser(ctx, id).UpdateUserRequest(updateUserRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/idp-api"
)

func main() {
	id := int64(789) // int64 | 
	updateUserRequest := *openapiclient.NewUpdateUserRequest("Email_example") // UpdateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApiAPI.UpdateUser(context.Background(), id).UpdateUserRequest(updateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApiAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: AuthleteUser
	fmt.Fprintf(os.Stdout, "Response from `UserApiAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 

### Return type

[**AuthleteUser**](AuthleteUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

