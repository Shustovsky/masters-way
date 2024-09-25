package routers

import (
	"fmt"
	"mw-notification-bff/internal/auth"
	"mw-notification-bff/internal/config"
	"mw-notification-bff/internal/controllers"
	"net/http"

	_ "mw-notification-bff/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Gin        *gin.Engine
	config     *config.Config
	fileRouter *fileRouter
}

func NewRouter(config *config.Config, controller *controllers.Controller) *Router {
	ginRouter := gin.Default()

	// Apply CORS middleware with custom options
	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.WebappBaseURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	ginRouter.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	return &Router{
		Gin:        ginRouter,
		config:     config,
		fileRouter: newFileRouter(controller.FileController),
	}
}

func (r *Router) SetRoutes() {
	chat := r.Gin.Group("/general", auth.HandleHeaders())

	r.fileRouter.setFileRoutes(chat)

	if r.config.EnvType != "prod" {
		r.Gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
