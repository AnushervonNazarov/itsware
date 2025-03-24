package router

import (
	"itsware/configs"
	"itsware/db"
	"itsware/internal/controllers"
	"itsware/internal/repositories"
	"itsware/internal/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiG := r.Group("/api")
	apiG.Use(controllers.AuthMiddleware())
	apiG.Use(controllers.SetDBSessionVariables())

	database := db.Pool
	userRepo := &repositories.User{DB: database}
	userService := &services.User{Repo: userRepo}
	userController := &controllers.User{Service: userService}

	auth := r.Group("/auth")
	{
		auth.POST("/sign-in", userController.SignIn)
	}

	userG := apiG.Group("/users")
	{
		userG.POST("", controllers.CreateUser)
		userG.GET("/:id", controllers.GetUser)
		userG.GET("", controllers.GetAllUsers)
		userG.PUT("", controllers.UpdateUser)
		userG.DELETE("/:id", controllers.DeleteUser)
		userG.POST("/", controllers.AddUserToTeam)
		userG.DELETE("/user/:user_id/team/:team_id", controllers.RemoveUserFromTeam)
	}

	tenantG := apiG.Group("/tenants")
	{
		tenantG.POST("", controllers.CreateTenant)
		tenantG.GET("/:id", controllers.GetTenant)
		tenantG.GET("", controllers.GetAllTenants)
		tenantG.PUT("", controllers.UpdateTenant)
		tenantG.DELETE("/:id", controllers.DeleteTenant)
	}

	teamG := apiG.Group("/teams")
	{
		teamG.POST("", controllers.CreateTeam)
		teamG.GET("/:id", controllers.GetTeam)
		teamG.GET("", controllers.GetAllTeams)
		teamG.PUT("", controllers.UpdateTeam)
		teamG.DELETE("/:id", controllers.DeleteTeam)
	}

	CabinetG := apiG.Group("/cabinets")
	{
		CabinetG.POST("", controllers.CreateCabinet)
		CabinetG.GET("/:id", controllers.GetCabinet)
		CabinetG.GET("", controllers.GetAllCabinets)
		CabinetG.PUT("", controllers.UpdateCabinet)
		CabinetG.DELETE("/:id", controllers.DeleteCabinet)
		CabinetG.POST("/", controllers.AddCabinetToTeam)
		CabinetG.DELETE("/cabinet/:cabinet_id/team/:team_id", controllers.RemoveCabinetFromTeam)
	}

	DeviceProfileG := apiG.Group("/deviceProfile")
	{
		DeviceProfileG.POST("", controllers.CreateDeviceProfile)
		DeviceProfileG.GET("/:id", controllers.GetDeviceProfile)
		DeviceProfileG.GET("", controllers.GetAllDeviceProfiles)
		DeviceProfileG.PUT("", controllers.UpdateDeviceProfile)
		DeviceProfileG.DELETE("/:id", controllers.DeleteDeviceProfile)
	}

	DeviceG := apiG.Group("devices")
	{
		DeviceG.POST("", controllers.CreateDevice)
		DeviceG.GET("/:id", controllers.GetDevice)
		DeviceG.GET("", controllers.GetAllDevices)
		DeviceG.PUT("", controllers.UpdateDevice)
		DeviceG.DELETE("/:id", controllers.DeleteDevice)
	}

	r.Run(":8080")

	return r
}
