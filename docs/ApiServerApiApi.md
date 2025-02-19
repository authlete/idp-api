# \ApiServerApiAPI

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiServer**](ApiServerApiAPI.md#CreateApiServer) | **Post** /api/apiserver | 
[**DeleteApiServer**](ApiServerApiAPI.md#DeleteApiServer) | **Delete** /api/apiserver/{id} | 
[**GetAll1**](ApiServerApiAPI.md#GetAll1) | **Get** /api/apiserver | 
[**GetApiServer**](ApiServerApiAPI.md#GetApiServer) | **Get** /api/apiserver/{id} | 
[**UpdateApiServer**](ApiServerApiAPI.md#UpdateApiServer) | **Post** /api/apiserver/{id} | 



## CreateApiServer

> AuthleteApiServerUpdateResponse CreateApiServer(ctx).CreateApiServerRequest(createApiServerRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	createApiServerRequest := *openapiclient.NewCreateApiServerRequest("ApiServerUrl_example", "Description_example") // CreateApiServerRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiServerApiAPI.CreateApiServer(context.Background()).CreateApiServerRequest(createApiServerRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiServerApiAPI.CreateApiServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiServer`: AuthleteApiServerUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiServerApiAPI.CreateApiServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiServerRequest** | [**CreateApiServerRequest**](CreateApiServerRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**AuthleteApiServerUpdateResponse**](AuthleteApiServerUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiServer

> DeleteApiServer(ctx, id).Authorization(authorization).DPoP(dPoP).Execute()



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
	r, err := apiClient.ApiServerApiAPI.DeleteApiServer(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiServerApiAPI.DeleteApiServer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteApiServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

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


## GetAll1

> []AuthleteApiServerUpdateResponse GetAll1(ctx).Authorization(authorization).DPoP(dPoP).Execute()



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
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiServerApiAPI.GetAll1(context.Background()).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiServerApiAPI.GetAll1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAll1`: []AuthleteApiServerUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiServerApiAPI.GetAll1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAll1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**[]AuthleteApiServerUpdateResponse**](AuthleteApiServerUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiServer

> AuthleteApiServerUpdateResponse GetApiServer(ctx, id).Authorization(authorization).DPoP(dPoP).Execute()



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
	resp, r, err := apiClient.ApiServerApiAPI.GetApiServer(context.Background(), id).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiServerApiAPI.GetApiServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiServer`: AuthleteApiServerUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiServerApiAPI.GetApiServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**AuthleteApiServerUpdateResponse**](AuthleteApiServerUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiServer

> AuthleteApiServerUpdateResponse UpdateApiServer(ctx, id).UpdateApiServerRequest(updateApiServerRequest).Authorization(authorization).DPoP(dPoP).Execute()



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
	updateApiServerRequest := *openapiclient.NewUpdateApiServerRequest("ApiServerUrl_example", "Description_example") // UpdateApiServerRequest | 
	authorization := "authorization_example" // string |  (optional)
	dPoP := "dPoP_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiServerApiAPI.UpdateApiServer(context.Background(), id).UpdateApiServerRequest(updateApiServerRequest).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiServerApiAPI.UpdateApiServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiServer`: AuthleteApiServerUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ApiServerApiAPI.UpdateApiServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateApiServerRequest** | [**UpdateApiServerRequest**](UpdateApiServerRequest.md) |  | 
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**AuthleteApiServerUpdateResponse**](AuthleteApiServerUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

