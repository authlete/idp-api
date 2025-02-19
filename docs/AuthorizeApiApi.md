# \AuthorizeApiAPI

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentAuthorizationResponse**](AuthorizeApiAPI.md#GetCurrentAuthorizationResponse) | **Get** /api/authorize | 
[**SubmitUserDecision**](AuthorizeApiAPI.md#SubmitUserDecision) | **Post** /api/authorize | 



## GetCurrentAuthorizationResponse

> PendingAuthorization GetCurrentAuthorizationResponse(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizeApiAPI.GetCurrentAuthorizationResponse(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeApiAPI.GetCurrentAuthorizationResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentAuthorizationResponse`: PendingAuthorization
	fmt.Fprintf(os.Stdout, "Response from `AuthorizeApiAPI.GetCurrentAuthorizationResponse`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentAuthorizationResponseRequest struct via the builder pattern


### Return type

[**PendingAuthorization**](PendingAuthorization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitUserDecision

> map[string]interface{} SubmitUserDecision(ctx).Approval(approval).Execute()



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
	approval := *openapiclient.NewApproval(false, "Ticket_example") // Approval | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizeApiAPI.SubmitUserDecision(context.Background()).Approval(approval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeApiAPI.SubmitUserDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitUserDecision`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthorizeApiAPI.SubmitUserDecision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUserDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **approval** | [**Approval**](Approval.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

