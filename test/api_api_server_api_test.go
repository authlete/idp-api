/*
OpenAPI definition

Testing ApiServerApiApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_openapi_ApiServerApiApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ApiServerApiApiService CreateApiServer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ApiServerApiApi.CreateApiServer(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ApiServerApiApiService DeleteApiServer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id int64

        resp, httpRes, err := apiClient.ApiServerApiApi.DeleteApiServer(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ApiServerApiApiService GetAll1", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ApiServerApiApi.GetAll1(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ApiServerApiApiService GetApiServer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id int64

        resp, httpRes, err := apiClient.ApiServerApiApi.GetApiServer(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ApiServerApiApiService UpdateApiServer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id int64

        resp, httpRes, err := apiClient.ApiServerApiApi.UpdateApiServer(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}