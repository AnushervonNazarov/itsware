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

func CreateTeam(ctx *gin.Context) {
	var team models.Team
	if err := ctx.BindJSON(&team); err != nil {
		logger.Error.Printf("[controllers.CreateTeam] error creating team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errors := utils.ValidateStruct(team); errors != nil {
		logger.Error.Printf("[controllers.CreateTeam] error validate team %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	if err := services.CreateTeam(ctx.Request.Context(), team); err != nil {
		logger.Error.Printf("[controllers.CreateTeam] error creating team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "team created successfully")
}

func GetTeam(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetTeam] error: invalid id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	team, err := services.GetTeam(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.GetTeam] error getting team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, team)
}

func GetAllTeams(ctx *gin.Context) {
	teams, err := services.GetAllTeams(ctx.Request.Context())
	if err != nil {
		logger.Error.Printf("[controllers.GetAllTeams] error getting all teams %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, teams)
}

func UpdateTeam(ctx *gin.Context) {
	var input models.UpdateTeam
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.Error.Printf("[controllers.UpdateTeam] error: invalid input data %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	if errors := utils.ValidateStruct(input); errors != nil {
		logger.Error.Printf("[controllers.UpdateTeam] error validate team %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	err := services.UpdateTeam(ctx.Request.Context(), input)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateTeam] error updating team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "team updated successfully",
	})
}

func DeleteTeam(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteTeam] error: invalid id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = services.DeleteTeam(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteTeam] error deleting team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "team deleted successfully",
	})
}
