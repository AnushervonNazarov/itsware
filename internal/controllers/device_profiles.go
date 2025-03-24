package controllers

import (
	"itsware/internal/models"
	"itsware/internal/services"
	"itsware/logger"
	"itsware/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateDeviceProfile(ctx *gin.Context) {
	var deviceProfile models.DeviceProfile
	if err := ctx.BindJSON(&deviceProfile); err != nil {
		logger.Error.Printf("[controllers.CreateDeviceProfile] error creating device profile %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errors := utils.ValidateStruct(deviceProfile); errors != nil {
		logger.Error.Printf("[controllers.CreateDeviceProfile] error validate device profile %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	if err := services.CreateDeviceProfile(ctx.Request.Context(), deviceProfile); err != nil {
		logger.Error.Printf("[controllers.CreateDeviceProfile] error creating device profile %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "device profile created successfully")
}

func GetDeviceProfile(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetDeviceProfile] error getting id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	deviceProfile, err := services.GetDeviceProfile(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.GetDeviceProfile] error getting device profile %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, deviceProfile)
}

func GetAllDeviceProfiles(ctx *gin.Context) {
	deviceProfiles, err := services.GetAllDeviceProfiles(ctx.Request.Context())
	if err != nil {
		logger.Error.Printf("[controllers.GetAllDeviceProfiles] error getting all device profiles %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, deviceProfiles)
}

func UpdateDeviceProfile(ctx *gin.Context) {
	var input models.UpdateDeviceProfile
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Error.Printf("[controllers.UpdateDeviceProfile] error: invalid input data %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	if errors := utils.ValidateStruct(input); errors != nil {
		logger.Error.Printf("[controllers.UpdateDeviceProfile] error validate device profile %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	err := services.UpdateDeviceProfile(ctx.Request.Context(), input)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateDeviceProfile] error updating device profile %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "device profile updated successfully",
	})
}

func DeleteDeviceProfile(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteDeviceProfile] error: invalid id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = services.DeleteDeviceProfile(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteDeviceProfile] error: deleting device profile %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "device profile deleted successfully",
	})
}
