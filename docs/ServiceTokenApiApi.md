# \ServiceTokenApiAPI

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTokenForService**](ServiceTokenApiAPI.md#CreateTokenForService) | **Post** /api/servicetoken/create | 
[**GetTokensForService**](ServiceTokenApiAPI.md#GetTokensForService) | **Post** /api/servicetoken/all | 
[**RenameServiceToken**](ServiceTokenApiAPI.md#RenameServiceToken) | **Post** /api/servicetoken/update | 
[**RevokeServiceToken**](ServiceTokenApiAPI.md#RevokeServiceToken) | **Post** /api/servicetoken/revoke | 
[**RotateServiceToken**](ServiceTokenApiAPI.md#RotateServiceToken) | **Post** /api/servicetoken/rotate | 



## CreateTokenForService

> ServiceTokenResponse CreateTokenForService(ctx).CreateServiceTokenRequest(createServiceTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	createServiceTokenRequest := *openapiclient.NewCreateServiceTokenRequest(int64(123), int64(123), int64(123), "Description_example") // CreateServiceTokenRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokenApiAPI.CreateTokenForService(context.Background()).CreateServiceTokenRequest(createServiceTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokenApiAPI.CreateTokenForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTokenForService`: ServiceTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokenApiAPI.CreateTokenForService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServiceTokenRequest** | [**CreateServiceTokenRequest**](CreateServiceTokenRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**ServiceTokenResponse**](ServiceTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokensForService

> []ServiceTokenResponse GetTokensForService(ctx).GetServiceTokensRequest(getServiceTokensRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	getServiceTokensRequest := *openapiclient.NewGetServiceTokensRequest(int64(123), int64(123), int64(123)) // GetServiceTokensRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokenApiAPI.GetTokensForService(context.Background()).GetServiceTokensRequest(getServiceTokensRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokenApiAPI.GetTokensForService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokensForService`: []ServiceTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokenApiAPI.GetTokensForService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensForServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getServiceTokensRequest** | [**GetServiceTokensRequest**](GetServiceTokensRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**[]ServiceTokenResponse**](ServiceTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameServiceToken

> ServiceTokenResponse RenameServiceToken(ctx).RenameServiceTokenRequest(renameServiceTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	renameServiceTokenRequest := *openapiclient.NewRenameServiceTokenRequest(int64(123), int64(123), int64(123), "TokenId_example", "Description_example") // RenameServiceTokenRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokenApiAPI.RenameServiceToken(context.Background()).RenameServiceTokenRequest(renameServiceTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokenApiAPI.RenameServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameServiceToken`: ServiceTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokenApiAPI.RenameServiceToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameServiceTokenRequest** | [**RenameServiceTokenRequest**](RenameServiceTokenRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**ServiceTokenResponse**](ServiceTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeServiceToken

> RevokeServiceToken(ctx).DeleteServiceTokenRequest(deleteServiceTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	deleteServiceTokenRequest := *openapiclient.NewDeleteServiceTokenRequest() // DeleteServiceTokenRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceTokenApiAPI.RevokeServiceToken(context.Background()).DeleteServiceTokenRequest(deleteServiceTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokenApiAPI.RevokeServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteServiceTokenRequest** | [**DeleteServiceTokenRequest**](DeleteServiceTokenRequest.md) |  | 
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


## RotateServiceToken

> ServiceTokenResponse RotateServiceToken(ctx).RotateServiceTokenRequest(rotateServiceTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	rotateServiceTokenRequest := *openapiclient.NewRotateServiceTokenRequest(int64(123), int64(123), int64(123), "TokenId_example") // RotateServiceTokenRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceTokenApiAPI.RotateServiceToken(context.Background()).RotateServiceTokenRequest(rotateServiceTokenRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceTokenApiAPI.RotateServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateServiceToken`: ServiceTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceTokenApiAPI.RotateServiceToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRotateServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rotateServiceTokenRequest** | [**RotateServiceTokenRequest**](RotateServiceTokenRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**ServiceTokenResponse**](ServiceTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

