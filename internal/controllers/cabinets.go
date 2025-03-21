package controllers

import (
	"itsware/internal/models"
	"itsware/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCabinet(ctx *gin.Context) {
	var cabinet models.Cabinet
	if err := ctx.BindJSON(&cabinet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := services.CreateCabinet(ctx.Request.Context(), cabinet); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "cabinet created successfully")
}

func GetCabinet(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	cabinet, err := services.GetCabinet(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, cabinet)
}

func GetAllCabinets(ctx *gin.Context) {
	cabinets, err := services.GetAllCabinets(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, cabinets)
}

func UpdateCabinet(ctx *gin.Context) {
	var input models.Cabinet
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	err := services.UpdateCabinet(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "cabinet updated successfully",
	})
}

func DeleteCabinet(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = services.DeleteCabinet(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "cabinet deleted successfully",
	})
}
