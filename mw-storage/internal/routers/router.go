package routers

import (
	"fmt"
	"mwstorage/internal/config"
	"mwstorage/internal/controllers"
	"net/http"

	_ "mwstorage/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Gin           *gin.Engine
	config        *config.Config
	messageRouter *messageRouter
	devRouter     *devRouter
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
		Gin:           ginRouter,
		config:        config,
		messageRouter: newMessageController(controller.MessagesController),
		devRouter:     newDevRouter(controller.DevController),
	}
}

func (r *Router) SetRoutes() {
	storage := r.Gin.Group("/storage")

	r.messageRouter.setMessageRoutes(storage)

	if r.config.EnvType != "prod" {
		r.devRouter.setDevRoutes(storage)
		r.Gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
