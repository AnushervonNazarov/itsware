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

func CreateCabinet(ctx *gin.Context) {
	var cabinet models.Cabinet
	if err := ctx.BindJSON(&cabinet); err != nil {
		logger.Error.Printf("[controllers.CreateCabinet] error creating cabinet %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errors := utils.ValidateStruct(cabinet); errors != nil {
		logger.Error.Printf("[controllers.CreateCabinet] error validate cabinet %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	if err := services.CreateCabinet(ctx.Request.Context(), cabinet); err != nil {
		logger.Error.Printf("[controllers.CreateCabinet] error creating cabinet %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "cabinet created successfully")
}

func GetCabinet(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetCabinet] error getting cabinet %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	cabinet, err := services.GetCabinet(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.GetCabinet] error getting cabinet %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, cabinet)
}

func GetAllCabinets(ctx *gin.Context) {
	cabinets, err := services.GetAllCabinets(ctx.Request.Context())
	if err != nil {
		logger.Error.Printf("[controllers.GetAllCabinets] error getting all cabinets %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, cabinets)
}

func UpdateCabinet(ctx *gin.Context) {
	var input models.UpdateCabinet
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Error.Printf("[controllers.UpdateCabinet] error updating cabinet %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	if errors := utils.ValidateStruct(input); errors != nil {
		logger.Error.Printf("[controllers.UpdateCabinet] error validate cabinet %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	err := services.UpdateCabinet(ctx.Request.Context(), input)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateCabinet] error updating cabinet %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
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
		logger.Error.Printf("[controllers.DeleteCabinet] error deleting cabinet %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = services.DeleteCabinet(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteCabinet] error deleting cabinet %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "cabinet deleted successfully",
	})
}

func AddCabinetToTeam(ctx *gin.Context) {
	var teamCabinet models.TeamCabinet
	if err := ctx.BindJSON(&teamCabinet); err != nil {
		logger.Error.Printf("[controllers.AddCabinetToTeam] error adding cabinet to team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := services.AddCabinetToTeam(ctx.Request.Context(), teamCabinet); err != nil {
		logger.Error.Printf("[controllers.AddCabinetToTeam] error adding cabinet to team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "cabinet added to team successfully")
}

func RemoveCabinetFromTeam(ctx *gin.Context) {
	cabinet_id, err := strconv.Atoi(ctx.Param("cabinet_id"))
	if err != nil {
		logger.Error.Printf("[controllers.RemoveCabinetFromTeam] error removing cabinet from team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid cabinet_id",
		})
		return
	}

	team_id, err := strconv.Atoi(ctx.Param("team_id"))
	if err != nil {
		logger.Error.Printf("[controllers.RemoveCabinetFromTeam] error getting team id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid team_id",
		})
		return
	}

	err = services.RemoveCabinetFromTeam(ctx.Request.Context(), cabinet_id, team_id)
	if err != nil {
		logger.Error.Printf("[controllers.RemoveCabinetFromTeam] error removing cabinet from team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "cabinet removed from team successfully",
	})
}
