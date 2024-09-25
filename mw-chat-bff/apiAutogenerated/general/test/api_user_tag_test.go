/*
Masters way general API

Testing UserTagAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_UserTagAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserTagAPIService CreateUserTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserTagAPI.CreateUserTag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserTagAPIService DeleteUserTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userTagId string
		var userId string

		httpRes, err := apiClient.UserTagAPI.DeleteUserTag(context.Background(), userTagId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
