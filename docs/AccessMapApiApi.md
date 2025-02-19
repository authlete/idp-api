# \AccessMapApiAPI

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessMap**](AccessMapApiAPI.md#GetAccessMap) | **Get** /api/accessmap | 



## GetAccessMap

> UserServiceMembershipResponse GetAccessMap(ctx).Authorization(authorization).DPoP(dPoP).Execute()



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
	resp, r, err := apiClient.AccessMapApiAPI.GetAccessMap(context.Background()).Authorization(authorization).DPoP(dPoP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessMapApiAPI.GetAccessMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessMap`: UserServiceMembershipResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessMapApiAPI.GetAccessMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **dPoP** | **string** |  | 

### Return type

[**UserServiceMembershipResponse**](UserServiceMembershipResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

