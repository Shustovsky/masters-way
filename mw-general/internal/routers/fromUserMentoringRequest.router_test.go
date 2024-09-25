package routers

import (
	"context"
	"mwgeneral/internal/auth"
	"mwgeneral/internal/config"
	"mwgeneral/internal/openapi"
	"net/http"
	"testing"

	openapiGeneral "mwgeneral/apiAutogenerated/general"

	"github.com/stretchr/testify/assert"
)

func TestCreateFromUserMentoringRequest(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	currentUserID := "1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1"
	wayID := "dce03ca6-f626-4c33-a44b-5a1b4ff62aa7"

	t.Run("should create FromUserMentoringRequest and return it successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(currentUserID, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		request := openapiGeneral.SchemasCreateFromUserMentoringRequestPayload{
			UserUuid: currentUserID,
			WayUuid:  wayID,
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		fromUserMentoringRequest, response, err := generalApi.FromUserMentoringRequestAPI.CreateFromUserMentoringRequest(ctx).Request(request).Execute()
		if err != nil {
			t.Fatalf("Failed to create FromUserMentoringRequest: %v", err)
		}

		expectedData := &openapiGeneral.SchemasFromUserMentoringRequestResponse{
			UserId: currentUserID,
			WayId:  wayID,
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData.WayId, fromUserMentoringRequest.WayId)
		assert.Equal(t, expectedData.UserId, fromUserMentoringRequest.UserId)
	})
}

func TestDeleteFromUserMentoringRequestById(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	userID := "3d922e8a-5d58-4b82-9a3d-83e2e73b3f91"
	wayID := "32cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"

	t.Run("should delete FromUserMentoringRequest and return Way without deleted FromUserMentoringRequest", func(t *testing.T) {
		token, err := auth.GenerateJWT(userID, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		response, err := generalApi.FromUserMentoringRequestAPI.DeleteFromUserMentoringRequest(ctx, userID, wayID).Execute()
		if err != nil {
			t.Fatalf("Failed to delete FromUserMentoringRequest: %v", err)
		}

		assert.Equal(t, http.StatusNoContent, response.StatusCode)

		way, wayResponse, err := generalApi.WayAPI.GetWayByUuid(ctx, wayID).Execute()
		if err != nil {
			t.Fatalf("Failed to get Way: %v", err)
		}

		assert.Equal(t, wayResponse.StatusCode, wayResponse.StatusCode)

		for _, mentorRequest := range way.MentorRequests {
			if mentorRequest.Uuid == userID {
				t.Fatalf("FromUserMentoringRequest from User %s for Way %s wasn't removed", userID, wayID)
			}
		}
	})
}