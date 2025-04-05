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

	AdminRoute := router.Group("")
	AdminRoute.Use(middlewares.Guard(), middlewares.CheckUserRoles([]string{"admin"}))

	// User Route
	SARoute.GET("/v1/users", handler.GetAccounts)
	SARoute.GET("/v1/users/:id", handler.GetUserByID)
	SARoute.POST("/v1/users", handler.CreateUser)
	SARoute.GET("/v1/users/username/:username", handler.GetUserByUsername)
	SARoute.PUT("/v1/users/password/:username", handler.UpdatePassword)
	SARoute.DELETE("/v1/users/:username", handler.DeleteUser)

	// Auth Route
	PublicRoute.POST("/v1/auth/login", handler.LoginUser)

	// Commodity Type Routes
	PublicRoute.GET("/v1/commodity-types", handler.GetAllCommodityTypes)
	PublicRoute.GET("/v1/commodity-types/:id", handler.GetCommodityTypeByID)
	AdminRoute.POST("/v1/commodity-types", handler.CreateCommodityType)
	AdminRoute.PUT("/v1/commodity-types/:id", handler.UpdateCommodityType)
	AdminRoute.DELETE("/v1/commodity-types/:id", handler.DeleteCommodityType)

	// Commodity Routes
	PublicRoute.GET("/v1/commodities", handler.GetAllCommodities)
	PublicRoute.GET("/v1/commodities/:id", handler.GetCommodityByID)
	AdminRoute.POST("/v1/commodities", handler.CreateCommodity)
	AdminRoute.PUT("/v1/commodities/:id", handler.UpdateCommodity)
	AdminRoute.DELETE("/v1/commodities/:id", handler.DeleteCommodity)

	// Employee Routes
	PublicRoute.GET("/v1/employees", handler.GetAllEmployees)
	PublicRoute.GET("/v1/employees/active", handler.GetActiveEmployees)
	PublicRoute.GET("/v1/employees/position/:position", handler.GetEmployeesByPosition)
	PublicRoute.GET("/v1/employees/:id", handler.GetEmployeeByID)
	AdminRoute.POST("/v1/employees", handler.CreateEmployee)
	AdminRoute.PUT("/v1/employees/:id", handler.UpdateEmployee)
	AdminRoute.DELETE("/v1/employees/:id", handler.DeleteEmployee)

	// News Routes
	PublicRoute.GET("/v1/news", handler.GetAllNews)
	PublicRoute.GET("/v1/news/:id", handler.GetNewsById)
	AdminRoute.POST("/v1/news", handler.CreateNews)
	AdminRoute.PUT("/v1/news/:id", handler.UpdateNews)
	AdminRoute.DELETE("/v1/news/:id", handler.DeleteNews)

	// Sector Routes
	PublicRoute.GET("/v1/sectors", handler.GetSectors)
	PublicRoute.GET("/v1/sectors/:id", handler.GetSector)
	AdminRoute.POST("/v1/sectors", handler.CreateSector)
	AdminRoute.PUT("/v1/sectors/:id", handler.UpdateSector)
	AdminRoute.DELETE("/v1/sectors/:id", handler.DeleteSector)

	// Upload Routes
	AdminRoute.POST("/v1/uploads", handler.UploadFile)
	PublicRoute.Static("/v1/uploads", "./uploads")
}
