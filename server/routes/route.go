package routes

import (
	"alter-io-go/controllers"
	"alter-io-go/helpers/http/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRegisterRoutes(router *gin.Engine, handler *controllers.Controller) {

	// SARoute = SuperAdmin Route
	SARoute := router.Group("")
	SARoute.Use(middlewares.Guard(), middlewares.CheckUserRoles([]string{"superadmin"}))

	PublicRoute := router.Group("")
	PublicRoute.POST("/v1/users", handler.CreateUser)
	PublicRoute.Use(middlewares.Guard())

	// User Route
	PublicRoute.GET("/v1/users", handler.GetAccounts)

	// Auth Route
	router.POST("/v1/auth/login", handler.LoginUser)
}
