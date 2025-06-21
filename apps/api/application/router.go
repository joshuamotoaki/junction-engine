package application

import (
	"github.com/gin-gonic/gin"
	"github.com/tigerappsorg/junction-engine/auth"
	"github.com/tigerappsorg/junction-engine/config"
	"github.com/tigerappsorg/junction-engine/database"
	"github.com/tigerappsorg/junction-engine/handlers"
	"github.com/tigerappsorg/junction-engine/middleware"
)

type Router struct {
	engine        *gin.Engine
	authHandler   *handlers.AuthHandler
	userHandler   *handlers.UserHandler
	healthHandler *handlers.HealthHandler
	casService    *auth.CASService
}

func NewRouter(cfg *config.Config, db *database.Neo4jDB, casService *auth.CASService) *Router {
	engine := gin.Default()

	return &Router{
		engine:        engine,
		authHandler:   handlers.NewAuthHandler(casService, db, cfg),
		userHandler:   handlers.NewUserHandler(db),
		healthHandler: handlers.NewHealthHandler(db),
		casService:    casService,
	}
}

func (r *Router) SetupRoutes() {
	r.setupPublicRoutes()
	r.setupAuthRoutes()
	r.setupProtectedRoutes()
}

func (r *Router) setupPublicRoutes() {
	public := r.engine.Group("/")
	{
		public.GET("/health", r.healthHandler.Check)
		public.GET("/health/database", r.healthHandler.DatabaseStatus)
	}
}

func (r *Router) setupAuthRoutes() {
	auth := r.engine.Group("/auth")
	{
		auth.GET("/login", r.authHandler.Login)
		auth.GET("/callback", r.authHandler.Callback)
		auth.GET("/logout", r.authHandler.Logout)
	}
}

func (r *Router) setupProtectedRoutes() {
	protected := r.engine.Group("/api")
	protected.Use(middleware.AuthMiddleware(r.casService))
	{
		// Auth endpoints
		protected.GET("/profile", r.authHandler.Profile)
	}
}

func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

func (r *Router) Run() error {
	return r.engine.Run()
}
