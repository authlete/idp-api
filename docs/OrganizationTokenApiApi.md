# \OrganizationTokenApiAPI

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTokenForOrganization**](OrganizationTokenApiAPI.md#CreateTokenForOrganization) | **Post** /api/organizationtoken/create | 
[**DeleteOrganizationToken**](OrganizationTokenApiAPI.md#DeleteOrganizationToken) | **Post** /api/organizationtoken/revoke | 
[**GetTokensForService1**](OrganizationTokenApiAPI.md#GetTokensForService1) | **Post** /api/organizationtoken/all | 
[**RenameTokenForService**](OrganizationTokenApiAPI.md#RenameTokenForService) | **Post** /api/organizationtoken/update | 
[**RotateTokenForOrganization**](OrganizationTokenApiAPI.md#RotateTokenForOrganization) | **Post** /api/organizationtoken/rotate | 



## CreateTokenForOrganization

> OrganizationTokenResponse CreateTokenForOrganization(ctx).CreateOrganizationTokenRequest(createOrganizationTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	createOrganizationTokenRequest := *openapiclient.NewCreateOrganizationTokenRequest(int64(123), "Description_example") // CreateOrganizationTokenRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTokenApiAPI.CreateTokenForOrganization(context.Background()).CreateOrganizationTokenRequest(createOrganizationTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTokenApiAPI.CreateTokenForOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTokenForOrganization`: OrganizationTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTokenApiAPI.CreateTokenForOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenForOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationTokenRequest** | [**CreateOrganizationTokenRequest**](CreateOrganizationTokenRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**OrganizationTokenResponse**](OrganizationTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationToken

> DeleteOrganizationToken(ctx).DeleteOrganizationTokenRequest(deleteOrganizationTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	deleteOrganizationTokenRequest := *openapiclient.NewDeleteOrganizationTokenRequest() // DeleteOrganizationTokenRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationTokenApiAPI.DeleteOrganizationToken(context.Background()).DeleteOrganizationTokenRequest(deleteOrganizationTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTokenApiAPI.DeleteOrganizationToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteOrganizationTokenRequest** | [**DeleteOrganizationTokenRequest**](DeleteOrganizationTokenRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

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


## GetTokensForService1

> []OrganizationTokenResponse GetTokensForService1(ctx).GetOrganizationTokensRequest(getOrganizationTokensRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	getOrganizationTokensRequest := *openapiclient.NewGetOrganizationTokensRequest(int64(123)) // GetOrganizationTokensRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTokenApiAPI.GetTokensForService1(context.Background()).GetOrganizationTokensRequest(getOrganizationTokensRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTokenApiAPI.GetTokensForService1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokensForService1`: []OrganizationTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTokenApiAPI.GetTokensForService1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensForService1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getOrganizationTokensRequest** | [**GetOrganizationTokensRequest**](GetOrganizationTokensRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**[]OrganizationTokenResponse**](OrganizationTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameTokenForService

> OrganizationTokenResponse RenameTokenForService(ctx).RenameOrganizationTokenRequest(renameOrganizationTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	renameOrganizationTokenRequest := *openapiclient.NewRenameOrganizationTokenRequest(int64(123), "TokenId_example", "Description_example") // RenameOrganizationTokenRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTokenApiAPI.RenameTokenForService(context.Background()).RenameOrganizationTokenRequest(renameOrganizationTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTokenApiAPI.RenameTokenForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameTokenForService`: OrganizationTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTokenApiAPI.RenameTokenForService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameTokenForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameOrganizationTokenRequest** | [**RenameOrganizationTokenRequest**](RenameOrganizationTokenRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**OrganizationTokenResponse**](OrganizationTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateTokenForOrganization

> OrganizationTokenResponse RotateTokenForOrganization(ctx).RotateOrganizationTokenRequest(rotateOrganizationTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	rotateOrganizationTokenRequest := *openapiclient.NewRotateOrganizationTokenRequest(int64(123), "TokenId_example") // RotateOrganizationTokenRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationTokenApiAPI.RotateTokenForOrganization(context.Background()).RotateOrganizationTokenRequest(rotateOrganizationTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationTokenApiAPI.RotateTokenForOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateTokenForOrganization`: OrganizationTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OrganizationTokenApiAPI.RotateTokenForOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRotateTokenForOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rotateOrganizationTokenRequest** | [**RotateOrganizationTokenRequest**](RotateOrganizationTokenRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**OrganizationTokenResponse**](OrganizationTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

