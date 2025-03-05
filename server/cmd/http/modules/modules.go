package modules

import (
	cfg "alter-io-go/config"
	ctrl "alter-io-go/controllers"
	"alter-io-go/helpers/http/middlewares"
	"alter-io-go/repositories/postgresql"
	"alter-io-go/routes"
	svc "alter-io-go/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// shouldSkipLogging determines if logging should be skipped for a given request.
func shouldSkipLogging(_, _ string) bool {
	return false
}

func RegisterModules(dbConn *pgxpool.Pool, router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "health"})
	})

	// Logging Middleware
	router.Use(middlewares.SlogLoggerWithSkipper(func(c *gin.Context) bool {
		return shouldSkipLogging(c.FullPath(), c.Request.Method)
	}))

	// CORS middleware configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins: cfg.GetConfig().App.CorsOrigin,
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	queries := postgresql.New(dbConn)
	service := svc.NewService(queries)
	handler := ctrl.NewController(service)

	routes.NewRegisterRoutes(router, handler)
}
