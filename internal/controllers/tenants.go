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

func CreateTenant(ctx *gin.Context) {
	var tenant models.Tenant
	if err := ctx.BindJSON(&tenant); err != nil {
		logger.Error.Printf("[controllers.CreateTenant] error creating tenant %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errors := utils.ValidateStruct(tenant); errors != nil {
		logger.Error.Printf("[controllers.CreateTenant] error validate tenant %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	if err := services.CreateTenant(ctx.Request.Context(), tenant); err != nil {
		logger.Error.Printf("[controllers.CreateTenant] error creating tenant %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "tenant created successfully")
}

func GetTenant(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetTenant] error: invalid id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	tenant, err := services.GetTenant(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.GetTenant] error getting tenant %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tenant)
}

func GetAllTenants(ctx *gin.Context) {
	tenants, err := services.GetAllTenants(ctx.Request.Context())
	if err != nil {
		logger.Error.Printf("[controllers.GetAllTenants] error getting all tenants %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tenants)
}

func UpdateTenant(ctx *gin.Context) {
	var input models.UpdateTenant
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Error.Printf("[controllers.UpdateTenant] error updating tenant %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	if errors := utils.ValidateStruct(input); errors != nil {
		logger.Error.Printf("[controllers.UpdateTenant] error validate tenant %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	err := services.UpdateTenant(ctx.Request.Context(), input)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateTenant] error updating tenant %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "tenant updated successfully",
	})
}

func DeleteTenant(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteTenant] error: invalid id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = services.DeleteTenant(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteTenant] error deleting tenant %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "tenant deleted successfully",
	})
}
