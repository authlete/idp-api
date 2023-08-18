# \AccessMapApiApi

All URIs are relative to *https://devidp.authlete.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessMap**](AccessMapApiApi.md#GetAccessMap) | **Get** /api/accessmap | 



## GetAccessMap

> UserServiceMembershipResponse GetAccessMap(ctx).Authorization(authorization).DPoP(dPoP).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string |  (optional)
    dPoP := "dPoP_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessMapApiApi.GetAccessMap(context.Background()).Authorization(authorization).DPoP(dPoP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessMapApiApi.GetAccessMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessMap`: UserServiceMembershipResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessMapApiApi.GetAccessMap`: %v\n", resp)
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

