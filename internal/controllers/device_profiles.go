package controllers

import (
	"itsware/internal/models"
	"itsware/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateDeviceProfile(ctx *gin.Context) {
	var deviceProfile models.DeviceProfile
	if err := ctx.BindJSON(&deviceProfile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := services.CreateDeviceProfile(ctx.Request.Context(), deviceProfile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "device profile created successfully")
}

func GetDeviceProfile(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	deviceProfile, err := services.GetDeviceProfile(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, deviceProfile)
}

func GetAllDeviceProfiles(ctx *gin.Context) {
	deviceProfiles, err := services.GetAllDeviceProfiles(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, deviceProfiles)
}

func UpdateDeviceProfile(ctx *gin.Context) {
	var input models.DeviceProfile
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	err := services.UpdateDeviceProfile(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = services.DeleteDeviceProfile(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "device profile deleted successfully",
	})
}
