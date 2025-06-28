package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tigerappsorg/junction-engine/docs" // Import generated docs for Swagger
	"github.com/tigerappsorg/junction-engine/internal/api/handlers"
	"github.com/tigerappsorg/junction-engine/internal/api/middleware"
	"github.com/tigerappsorg/junction-engine/internal/database/neo4j"
	"github.com/tigerappsorg/junction-engine/internal/shared/auth"
	"github.com/tigerappsorg/junction-engine/internal/shared/config"
)

type router struct {
	engine        *gin.Engine
	port          string
	authHandler   handlers.AuthHandler
	userHandler   handlers.UserHandler
	healthHandler handlers.HealthHandler
	casService    auth.CASService
}

type Router interface {
	SetupRoutes()
	GetEngine() *gin.Engine
	Run() error
}

func NewRouter(cfg *config.Config, db neo4j.Neo4jDB, casService auth.CASService) Router {
	if cfg.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	engine.SetTrustedProxies(nil)

	return &router{
		engine:        engine,
		port:          cfg.Port,
		authHandler:   handlers.NewAuthHandler(casService, db, cfg),
		userHandler:   handlers.NewUserHandler(db),
		healthHandler: handlers.NewHealthHandler(db),
		casService:    casService,
	}
}

func (r *router) SetupRoutes() {
	api := r.engine.Group("/api/v1")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.GET("/health", r.healthHandler.Check)
		api.GET("/health/database", r.healthHandler.DatabaseStatus)

		auth := api.Group("/auth")
		{
			auth.GET("/login", r.authHandler.Login)
			auth.GET("/callback", r.authHandler.Callback)
			auth.GET("/logout", r.authHandler.Logout)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(r.casService))
		{
			protected.GET("/profile", r.authHandler.Profile)
		}
	}
}

func (r *router) GetEngine() *gin.Engine {
	return r.engine
}

func (r *router) Run() error {
	return r.engine.Run(":" + r.port)
}
