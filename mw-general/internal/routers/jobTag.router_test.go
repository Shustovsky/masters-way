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

func TestCreateJobTag(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"

	type RequestBody struct {
		Color       string
		Description string
		Name        string
		WayUuid     string
	}

	requestData := RequestBody{
		Color:       "red",
		Description: "A description",
		Name:        "Sample Name",
		WayUuid:     "9972552a-c0b3-41f3-b464-284d36a36964",
	}

	t.Run("should get all ways successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)

		requestCreateJobTag := openapiGeneral.SchemasCreateJobTagPayload{
			Color:       requestData.Color,
			Description: requestData.Description,
			Name:        requestData.Name,
			WayUuid:     requestData.WayUuid,
		}
		jobTag, response, err := generalApi.JobTagAPI.CreateJobTag(ctx).Request(requestCreateJobTag).Execute()
		if err != nil {
			t.Fatalf("Failed to create job tag: %v", err)
		}

		expectedData := &openapiGeneral.SchemasJobTagResponse{
			Color:       requestData.Color,
			Description: requestData.Description,
			Name:        requestData.Name,
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData.Color, jobTag.Color)
		assert.Equal(t, expectedData.Description, jobTag.Description)
		assert.Equal(t, expectedData.Name, jobTag.Name)
	})
}
func TestUpdateJobTag(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"

	type RequestBody struct {
		Color       string
		Description string
		Name        string
		JobTagID    string
	}

	requestData := RequestBody{
		Color:       "red",
		Description: "A description",
		Name:        "Sample Name",
		JobTagID:    "bf63a158-c3d9-40aa-bc0f-e6686e96de2c",
	}

	t.Run("should get all ways successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)

		request := openapiGeneral.SchemasUpdateJobTagPayload{
			Color:       &requestData.Color,
			Description: &requestData.Description,
			Name:        &requestData.Name,
		}
		jobTag, response, err := generalApi.JobTagAPI.UpdateJobTag(ctx, requestData.JobTagID).Request(request).Execute()
		if err != nil {
			t.Fatalf("Failed to create job tag: %v", err)
		}

		expectedData := &openapiGeneral.SchemasJobTagResponse{
			Color:       requestData.Color,
			Description: requestData.Description,
			Name:        requestData.Name,
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData.Color, jobTag.Color)
		assert.Equal(t, expectedData.Description, jobTag.Description)
		assert.Equal(t, expectedData.Name, jobTag.Name)
	})
}

func TestDeleteJobTagById(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"
	jobTagID := "d569aa06-452c-4602-a788-2ffca4c959a8"
	wayID := "32cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"

	t.Run("should get all ways successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		response, err := generalApi.JobTagAPI.DeleteJobTag(ctx, jobTagID).Execute()
		if err != nil {
			t.Fatalf("Failed to delete job tag: %v", err)
		}

		assert.Equal(t, http.StatusNoContent, response.StatusCode)

		report, response, err := generalApi.DayReportAPI.GetDayReports(ctx, wayID).Execute()
		if err != nil {
			t.Fatalf("Failed to get report: %v", err)
		}

		isJobTagExists := false
		for _, dayReport := range report.DayReports {
			for _, jobDone := range dayReport.JobsDone {
				for _, tag := range jobDone.Tags {
					if tag.Uuid == jobTagID {
						isJobTagExists = true
						break
					}
				}
			}
		}

		if isJobTagExists {
			t.Fatalf("Found job tag: %v", jobTagID)
		}
	})
}