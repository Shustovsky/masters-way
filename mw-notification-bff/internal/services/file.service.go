package services

import (
	openapiNotification "mw-notification-bff/apiAutogenerated/mw-notification"
	"mw-notification-bff/internal/schemas"
	"net/http"
)

type FileService struct {
	fileAPI *openapiNotification.APIClient
}

func NewFileService(fileAPI *openapiNotification.APIClient) *FileService {
	return &FileService{fileAPI}
}

func (cs *FileService) UploadFile(request *http.Request, googleToken string) (*schemas.UploadFileResponse, error) {
	// r := openapiNotification.ApiGetNotificationListRequest{
	// 	ApiService: cs.fileAPI.NotificationAPI,
	// }

	// filePreviewRaw, _, err := cs.fileAPI.NotificationAPI.UploadFileStreamExecute(r, request, googleToken)
	// if err != nil {
	// 	return nil, err
	// }

	return &schemas.UploadFileResponse{
		// ID:         filePreviewRaw.Id,
		// OwnerID:    filePreviewRaw.OwnerId,
		// Name:       filePreviewRaw.Name,
		// PreviewURL: filePreviewRaw.PreviewUrl,
		// SrcURL:     filePreviewRaw.SrcUrl,
	}, nil
}
