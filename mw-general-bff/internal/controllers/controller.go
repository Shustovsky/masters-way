package controllers

import "mw-general-bff/internal/services"

type Controller struct {
	FileController *FileController
}

func NewController(services *services.Service) *Controller {
	return &Controller{
		FileController: NewFileController(services.GeneralService, services.FileService),
	}
}
