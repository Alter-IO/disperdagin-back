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

	// Greeting Routes
	PublicRoute.GET("/v1/greetings", handler.GetAllGreetings)
	PublicRoute.GET("/v1/greetings/latest", handler.GetLatestGreeting)
	PublicRoute.GET("/v1/greetings/:id", handler.GetGreetingByID)
	AdminRoute.POST("/v1/greetings", handler.CreateGreeting)
	AdminRoute.PUT("/v1/greetings/:id", handler.UpdateGreeting)
	AdminRoute.DELETE("/v1/greetings/:id", handler.DeleteGreeting)

	// IKM Type Routes
	PublicRoute.GET("/v1/ikm-types", handler.GetAllIKMTypes)
	PublicRoute.GET("/v1/ikm-types/:id", handler.GetIKMTypeByID)
	PublicRoute.GET("/v1/ikm-types/info-type/:infoType", handler.GetIKMTypesByInfoType)
	AdminRoute.POST("/v1/ikm-types", handler.CreateIKMType)
	AdminRoute.PUT("/v1/ikm-types/:id", handler.UpdateIKMType)
	AdminRoute.DELETE("/v1/ikm-types/:id", handler.DeleteIKMType)

	// IKM Routes
	PublicRoute.GET("/v1/ikms", handler.GetAllIKMs)
	PublicRoute.GET("/v1/ikms/:id", handler.GetIKMByID)
	PublicRoute.GET("/v1/ikms/village/:villageId", handler.GetIKMsByVillage)
	PublicRoute.GET("/v1/ikms/business-type/:businessType", handler.GetIKMsByBusinessType)
	AdminRoute.POST("/v1/ikms", handler.CreateIKM)
	AdminRoute.PUT("/v1/ikms/:id", handler.UpdateIKM)
	AdminRoute.DELETE("/v1/ikms/:id", handler.DeleteIKM)

	// Legal Document Routes
	PublicRoute.GET("/v1/legal-documents", handler.GetAllLegalDocuments)
	PublicRoute.GET("/v1/legal-documents/:id", handler.GetLegalDocumentByID)
	PublicRoute.GET("/v1/legal-documents/type/:docType", handler.GetLegalDocumentsByType)
	AdminRoute.POST("/v1/legal-documents", handler.CreateLegalDocument)
	AdminRoute.PUT("/v1/legal-documents/:id", handler.UpdateLegalDocument)
	AdminRoute.DELETE("/v1/legal-documents/:id", handler.DeleteLegalDocument)

	// Legal Document Type Routes
	PublicRoute.GET("/v1/legal-doc-types", handler.GetAllLegalDocTypes)
	PublicRoute.GET("/v1/legal-doc-types/:id", handler.GetLegalDocTypeByID)
	AdminRoute.POST("/v1/legal-doc-types", handler.CreateLegalDocType)
	AdminRoute.PUT("/v1/legal-doc-types/:id", handler.UpdateLegalDocType)
	AdminRoute.DELETE("/v1/legal-doc-types/:id", handler.DeleteLegalDocType)

	// Market Routes
	PublicRoute.GET("/v1/markets", handler.GetAllMarkets)
	PublicRoute.GET("/v1/markets/:id", handler.GetMarketByID)
	AdminRoute.POST("/v1/markets", handler.CreateMarket)
	AdminRoute.PUT("/v1/markets/:id", handler.UpdateMarket)
	AdminRoute.DELETE("/v1/markets/:id", handler.DeleteMarket)

	// Market Fee Routes
	PublicRoute.GET("/v1/market-fees", handler.GetAllMarketFees)
	PublicRoute.GET("/v1/market-fees/:id", handler.GetMarketFeeByID)
	PublicRoute.GET("/v1/market-fees/market/:marketId", handler.GetMarketFeesByMarket)
	PublicRoute.GET("/v1/market-fees/year/:year", handler.GetMarketFeesByYear)
	PublicRoute.GET("/v1/market-fees/semester/:semester/year/:year", handler.GetMarketFeesBySemesterAndYear)
	AdminRoute.POST("/v1/market-fees", handler.CreateMarketFee)
	AdminRoute.PUT("/v1/market-fees/:id", handler.UpdateMarketFee)
	AdminRoute.DELETE("/v1/market-fees/:id", handler.DeleteMarketFee)

	// Photo Category Routes
	PublicRoute.GET("/v1/photo-categories", handler.GetAllPhotoCategories)
	PublicRoute.GET("/v1/photo-categories/:id", handler.GetPhotoCategoryByID)
	AdminRoute.POST("/v1/photo-categories", handler.CreatePhotoCategory)
	AdminRoute.PUT("/v1/photo-categories/:id", handler.UpdatePhotoCategory)
	AdminRoute.DELETE("/v1/photo-categories/:id", handler.DeletePhotoCategory)

	// Photo Routes
	PublicRoute.GET("/v1/photos", handler.GetAllPhotos)
	PublicRoute.GET("/v1/photos/:id", handler.GetPhotoByID)
	PublicRoute.GET("/v1/photos/category/:categoryId", handler.GetPhotosByCategory)
	AdminRoute.POST("/v1/photos", handler.CreatePhoto)
	AdminRoute.PUT("/v1/photos/:id", handler.UpdatePhoto)
	AdminRoute.DELETE("/v1/photos/:id", handler.DeletePhoto)

	// Public Information Type Routes
	PublicRoute.GET("/v1/public-info-types", handler.GetAllPublicInfoTypes)
	PublicRoute.GET("/v1/public-info-types/:id", handler.GetPublicInfoTypeByID)
	AdminRoute.POST("/v1/public-info-types", handler.CreatePublicInfoType)
	AdminRoute.PUT("/v1/public-info-types/:id", handler.UpdatePublicInfoType)
	AdminRoute.DELETE("/v1/public-info-types/:id", handler.DeletePublicInfoType)

	// Public Information Routes
	PublicRoute.GET("/v1/public-info", handler.GetAllPublicInfo)
	PublicRoute.GET("/v1/public-info/:id", handler.GetPublicInfoByID)
	PublicRoute.GET("/v1/public-info/type/:infoType", handler.GetPublicInfoByType)
	AdminRoute.POST("/v1/public-info", handler.CreatePublicInfo)
	AdminRoute.PUT("/v1/public-info/:id", handler.UpdatePublicInfo)
	AdminRoute.DELETE("/v1/public-info/:id", handler.DeletePublicInfo)

	// Subdistrict Routes
	PublicRoute.GET("/v1/subdistricts", handler.GetAllSubdistricts)
	PublicRoute.GET("/v1/subdistricts/:id", handler.GetSubdistrictByID)
	AdminRoute.POST("/v1/subdistricts", handler.CreateSubdistrict)
	AdminRoute.PUT("/v1/subdistricts/:id", handler.UpdateSubdistrict)
	AdminRoute.DELETE("/v1/subdistricts/:id", handler.DeleteSubdistrict)

	// Village Routes
	PublicRoute.GET("/v1/villages", handler.GetAllVillages)
	PublicRoute.GET("/v1/villages/:id", handler.GetVillageByID)
	PublicRoute.GET("/v1/villages/subdistrict/:subdistrictId", handler.GetVillagesBySubdistrict)
	AdminRoute.POST("/v1/villages", handler.CreateVillage)
	AdminRoute.PUT("/v1/villages/:id", handler.UpdateVillage)
	AdminRoute.DELETE("/v1/villages/:id", handler.DeleteVillage)

	// Vision Mission Routes
	PublicRoute.GET("/v1/vision-missions", handler.GetAllVisionMissions)
	PublicRoute.GET("/v1/vision-missions/latest", handler.GetLatestVisionMission)
	PublicRoute.GET("/v1/vision-missions/:id", handler.GetVisionMissionByID)
	AdminRoute.POST("/v1/vision-missions", handler.CreateVisionMission)
	AdminRoute.PUT("/v1/vision-missions/:id", handler.UpdateVisionMission)
	AdminRoute.DELETE("/v1/vision-missions/:id", handler.DeleteVisionMission)

	// Upload Routes
	AdminRoute.POST("/v1/uploads", handler.UploadFile)
	PublicRoute.Static("/v1/uploads", "./uploads")
}
