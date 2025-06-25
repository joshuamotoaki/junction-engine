package application

import (
	"github.com/gin-gonic/gin"
	"github.com/tigerappsorg/junction-engine/auth"
	"github.com/tigerappsorg/junction-engine/config"
	"github.com/tigerappsorg/junction-engine/database"
	"github.com/tigerappsorg/junction-engine/handlers"
	"github.com/tigerappsorg/junction-engine/middleware"
)

type router struct {
	engine        *gin.Engine
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

func NewRouter(cfg *config.Config, db database.Neo4jDB, casService auth.CASService) Router {
	engine := gin.Default()

	return &router{
		engine:        engine,
		authHandler:   handlers.NewAuthHandler(casService, db, cfg),
		userHandler:   handlers.NewUserHandler(db),
		healthHandler: handlers.NewHealthHandler(db),
		casService:    casService,
	}
}

func (r *router) SetupRoutes() {
	r.setupPublicRoutes()
	r.setupAuthRoutes()
	r.setupProtectedRoutes()
}

func (r *router) setupPublicRoutes() {
	public := r.engine.Group("/")
	{
		public.GET("/health", r.healthHandler.Check)
		public.GET("/health/database", r.healthHandler.DatabaseStatus)
	}
}

func (r *router) setupAuthRoutes() {
	auth := r.engine.Group("/auth")
	{
		auth.GET("/login", r.authHandler.Login)
		auth.GET("/callback", r.authHandler.Callback)
		auth.GET("/logout", r.authHandler.Logout)
	}
}

func (r *router) setupProtectedRoutes() {
	protected := r.engine.Group("/api")
	protected.Use(middleware.AuthMiddleware(r.casService))
	{
		// Auth endpoints
		protected.GET("/profile", r.authHandler.Profile)
	}
}

func (r *router) GetEngine() *gin.Engine {
	return r.engine
}

func (r *router) Run() error {
	return r.engine.Run()
}
