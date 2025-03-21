package controllers

import (
	"itsware/internal/models"
	"itsware/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTenant(ctx *gin.Context) {
	var tenant models.Tenant
	if err := ctx.BindJSON(&tenant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := services.CreateTenant(ctx.Request.Context(), tenant); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "tenant created successfully")
}

func GetTenant(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	tenant, err := services.GetTenant(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tenant)
}

func GetAllTenants(ctx *gin.Context) {
	tenants, err := services.GetAllTenants(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tenants)
}

func UpdateTenant(ctx *gin.Context) {
	var input models.Tenant
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	err := services.UpdateTenant(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = services.DeleteTenant(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "tenant deleted successfully",
	})
}
