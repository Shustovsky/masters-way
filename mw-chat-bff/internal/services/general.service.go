package services

import (
	"fmt"
	openapiGeneral "mw-chat-bff/apiAutogenerated/general"
	utils "mw-chat-bff/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type GeneralService struct {
	generalAPI *openapiGeneral.APIClient
}

func NewGeneralService(generalAPI *openapiGeneral.APIClient) *GeneralService {
	return &GeneralService{generalAPI}
}

func (gs *GeneralService) GetPopulatedUsers(ctx *gin.Context, userIDs []string) (map[string]PopulatedUser, error) {
	chatUsersData, response, err := gs.generalAPI.UserAPI.GetUsersByIds(ctx).Request(userIDs).Execute()
	if err != nil {
		message, extractErr := utils.ExtractErrorMessageFromResponse(response)
		if extractErr != nil {
			return nil, fmt.Errorf("failed to extract error message: %w", extractErr)
		}
		return nil, fmt.Errorf(message)
	}

	userMap := lo.SliceToMap(chatUsersData, func(userData openapiGeneral.SchemasGetUsersByIDsResponse) (string, PopulatedUser) {
		return userData.UserId, PopulatedUser{
			UserID:   userData.UserId,
			Name:     userData.Name,
			ImageURL: userData.ImageUrl,
		}
	})

	return userMap, nil
}

func (gs *GeneralService) GetGoogleAccessTokenByID(ctx *gin.Context, userID string) (string, error) {
	googleToken, response, err := gs.generalAPI.AuthAPI.GetGoogleToken(ctx).Execute()
	if err != nil {
		message, extractErr := utils.ExtractErrorMessageFromResponse(response)
		if extractErr != nil {
			return "", fmt.Errorf("failed to extract error message: %w", extractErr)
		}
		return "", fmt.Errorf(message)
	}

	return googleToken.AccessToken, nil
}