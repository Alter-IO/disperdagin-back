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

	// Routes only accessible by admin
	SARoute.PUT("/v1/users/reset-password/:username", handler.ResetPassword)

	PublicRoute := router.Group("")
	PublicRoute.POST("/v1/users", handler.CreateUser)
	PublicRoute.Use(middlewares.Guard())

	// User Route
	PublicRoute.GET("/v1/users", handler.GetAccounts)
	PublicRoute.GET("/v1/users/:id", handler.GetUserByID)
	PublicRoute.GET("/v1/users/username/:username", handler.GetUserByUsername)
	PublicRoute.PUT("/v1/users/password/:username", handler.UpdatePassword)
	PublicRoute.DELETE("/v1/users/:username", handler.DeleteUser)

	// Auth Route
	router.POST("/v1/auth/login", handler.LoginUser)

	// Commodity Type Routes
	PublicRoute.GET("/v1/commodity-types", handler.GetAllCommodityTypes)
	PublicRoute.GET("/v1/commodity-types/:id", handler.GetCommodityTypeByID)
	PublicRoute.POST("/v1/commodity-types", handler.CreateCommodityType)
	PublicRoute.PUT("/v1/commodity-types/:id", handler.UpdateCommodityType)
	PublicRoute.DELETE("/v1/commodity-types/:id", handler.DeleteCommodityType)

	// Commodity Routes
	PublicRoute.GET("/v1/commodities", handler.GetAllCommodities)
	PublicRoute.GET("/v1/commodities/:id", handler.GetCommodityByID)
	PublicRoute.POST("/v1/commodities", handler.CreateCommodity)
	PublicRoute.PUT("/v1/commodities/:id", handler.UpdateCommodity)
	PublicRoute.DELETE("/v1/commodities/:id", handler.DeleteCommodity)

	// Employee Routes
	PublicRoute.GET("/v1/employees", handler.GetAllEmployees)
	PublicRoute.GET("/v1/employees/active", handler.GetActiveEmployees)
	PublicRoute.GET("/v1/employees/position/:position", handler.GetEmployeesByPosition)
	PublicRoute.GET("/v1/employees/:id", handler.GetEmployeeByID)
	PublicRoute.POST("/v1/employees", handler.CreateEmployee)
	PublicRoute.PUT("/v1/employees/:id", handler.UpdateEmployee)
	PublicRoute.DELETE("/v1/employees/:id", handler.DeleteEmployee)

	// Upload Routes
	SARoute.POST("/v1/uploads", handler.UploadFile)
}
