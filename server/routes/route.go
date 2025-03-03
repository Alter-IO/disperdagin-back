package routes

import (
	"alter-io-go/controllers"
	"alter-io-go/helpers/http/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRegisterRoutes(router *gin.Engine, handler *controllers.Controller) {
	// SARoute = SuperAdmin Route
	SARoute := router.Group("/v1")
	SARoute.Use(middlewares.Guard(), middlewares.CheckUserRoles([]string{"superadmin"}))

	PublicRoute := router.Group("/v1")
	PublicRoute.Use(middlewares.Guard())

	// User Route
	SARoute.POST("/users", handler.CreateUser)
	PublicRoute.GET("/users", handler.GetAccounts)
}
