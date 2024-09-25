package routers

import (
	"context"
	openapiGeneral "mwgeneral/apiAutogenerated/general"
	"mwgeneral/internal/auth"
	"mwgeneral/internal/config"
	"mwgeneral/internal/openapi"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUserTagByName(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	tagName := "Test user tag"
	userID := "7cdb041b-4574-4f7b-a500-c53e74c72e94"

	t.Run("should successfully create a new user tag and return the correct userTag name", func(t *testing.T) {
		token, err := auth.GenerateJWT(userID, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		request := openapiGeneral.SchemasCreateUserTagPayload{
			Name:      tagName,
			OwnerUuid: userID,
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		allUsersResponse, response, err := generalApi.UserTagAPI.CreateUserTag(ctx).Request(request).Execute()
		if err != nil {
			t.Fatalf("Failed to create UserTag: %v", err)
		}

		expectedData := &openapiGeneral.SchemasUserTagResponse{
			Name: tagName,
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData.Name, allUsersResponse.Name)
	})
}

func TestDeleteUserTagByFromUserByTag(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	userID := "1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1"
	userTagID := "8749d799-0a89-4ffd-b1bd-02ada9234e5a"

	t.Run("should delete a user tag and ensure it no longer exists in the user's tags", func(t *testing.T) {
		token, err := auth.GenerateJWT(userID, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		response, err := generalApi.UserTagAPI.DeleteUserTag(ctx, userTagID, userID).Execute()
		if err != nil {
			t.Fatalf("Failed to delete MentorUserWay: %v", err)
		}

		assert.Equal(t, http.StatusNoContent, response.StatusCode)

		user, userResponse, err := generalApi.UserAPI.GetUserByUuid(ctx, userID).Execute()
		if err != nil {
			t.Fatalf("Failed to get user: %v", err)
		}

		assert.Equal(t, http.StatusOK, userResponse.StatusCode)

		tagFound := false
		for _, tag := range user.Tags {
			if tag.Uuid == userTagID {
				tagFound = true
				break
			}
		}

		if tagFound {
			t.Fatalf("UserTag was not deleted successfully; it still exists in the user's tags")
		}
	})
}