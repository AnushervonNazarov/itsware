package controllers

import (
	"itsware/configs"
	"itsware/db"
	"itsware/internal/repositories"
	"itsware/internal/services"

	"github.com/gin-gonic/gin"
)

func RunRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	apiG := r.Group("/api")
	apiG.Use(AuthMiddleware())
	apiG.Use(SetDBSessionVariables())

	database := db.Pool
	userRepo := &repositories.User{DB: database}
	userService := &services.User{Repo: userRepo}
	userController := &User{Service: userService}

	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", userController.SignIn)
	}

	userG := apiG.Group("/users")
	{
		userG.POST("", CreateUser)
		userG.GET("/:id", GetUser)
		userG.GET("", GetAllUsers)
		userG.PUT("", UpdateUser)
		userG.DELETE("/:id", DeleteUser)
	}

	tenantG := apiG.Group("/tenants")
	{
		tenantG.POST("", CreateTenant)
		tenantG.GET("/:id", GetTenant)
		tenantG.GET("", GetAllTenants)
		tenantG.PUT("", UpdateTenant)
		tenantG.DELETE("/:id", DeleteTenant)
	}

	teamG := apiG.Group("/teams")
	{
		teamG.POST("", CreateTeam)
		teamG.GET("/:id", GetTeam)
		teamG.GET("", GetAllTeams)
		teamG.PUT("", UpdateTeam)
		teamG.DELETE("/:id", DeleteTeam)
	}

	CabinetG := apiG.Group("/cabinets")
	{
		CabinetG.POST("", CreateCabinet)
		CabinetG.GET("/:id", GetCabinet)
		CabinetG.GET("", GetAllCabinets)
		CabinetG.PUT("", UpdateCabinet)
		CabinetG.DELETE("/:id", DeleteCabinet)
	}

	DeviceProfileG := apiG.Group("/deviceProfile")
	{
		DeviceProfileG.POST("", CreateDeviceProfile)
		DeviceProfileG.GET("/:id", GetDeviceProfile)
		DeviceProfileG.GET("", GetAllDeviceProfiles)
		DeviceProfileG.PUT("", UpdateDeviceProfile)
		DeviceProfileG.DELETE("/:id", DeleteDeviceProfile)
	}

	DeviceG := apiG.Group("devices")
	{
		DeviceG.POST("", CreateDevice)
		DeviceG.GET("/:id", GetDevice)
		DeviceG.GET("", GetAllDevices)
		DeviceG.PUT("", UpdateDevice)
		DeviceG.DELETE("/:id", DeleteDevice)
	}

	r.Run(":8080")

	return r
}
