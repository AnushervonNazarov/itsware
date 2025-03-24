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

func CreateDevice(ctx *gin.Context) {
	var device models.Device
	if err := ctx.BindJSON(&device); err != nil {
		logger.Error.Printf("[controllers.CreateDevice] error creating device %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errors := utils.ValidateStruct(device); errors != nil {
		logger.Error.Printf("[controllers.CreateDevice] error validate device %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	if err := services.CreateDevice(ctx.Request.Context(), device); err != nil {
		logger.Error.Printf("[controllers.CreateDevice] error creating device %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "device created successfully")
}

func GetDevice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetDevice] error: invalid id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	device, err := services.GetDevice(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.GetDevice] error getting device %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, device)
}

func GetAllDevices(ctx *gin.Context) {
	devices, err := services.GetAllDevices(ctx.Request.Context())
	if err != nil {
		logger.Error.Printf("[controllers.GetAllDevices] error getting all devices %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, devices)
}

func UpdateDevice(ctx *gin.Context) {
	var input models.UpdateDevice
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Error.Printf("[controllers.UpdateDevice] error: invalid input data %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	if errors := utils.ValidateStruct(input); errors != nil {
		logger.Error.Printf("[controllers.UpdateDevice] error validate device %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	err := services.UpdateDevice(ctx.Request.Context(), input)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateDevice] error updating device %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "device updated successfully",
	})
}

func DeleteDevice(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteDevice] error; invalid id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = services.DeleteDevice(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteDevice] error deleting device %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "device deleted successfully",
	})
}
