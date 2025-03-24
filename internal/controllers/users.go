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

type User struct {
	Service *services.User
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		logger.Error.Printf("[controllers.CreateUser] error creating user %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errors := utils.ValidateStruct(user); errors != nil {
		logger.Error.Printf("[controllers.CreateUser] error validate user %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	if err := services.CreateUser(ctx.Request.Context(), user); err != nil {
		logger.Error.Printf("[controllers.CreateUser] error creating user %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "user created successfully")
}

func GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUser] error: invalid id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := services.GetUser(ctx.Request.Context(), id)
	logger.Error.Printf("[controllers.GetUser] error getting user %v\n", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func GetAllUsers(ctx *gin.Context) {
	users, err := services.GetAllUsers(ctx.Request.Context())
	if err != nil {
		logger.Error.Printf("[controllers.GetAllUsers] error getting all users %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func UpdateUser(ctx *gin.Context) {
	var userInput models.UpdateUser
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		logger.Error.Printf("[controllers.UpdateUser] error: invalid input data %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	if errors := utils.ValidateStruct(userInput); errors != nil {
		logger.Error.Printf("[controllers.UpdateUser] error validate user %v\n", errors)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	err := services.UpdateUser(ctx.Request.Context(), userInput)
	if err != nil {
		logger.Error.Printf("[controllers.UpdateUser] error updating user %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
	})
}

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.DeleteUser] error: invalid id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = services.DeleteUser(ctx.Request.Context(), id)
	if err != nil {
		logger.Error.Printf("[controllers.DeleteUser] error deleting user %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}

func AddUserToTeam(ctx *gin.Context) {
	var teamUser models.TeamUser
	if err := ctx.BindJSON(&teamUser); err != nil {
		logger.Error.Printf("[controllers.AddUserToTeam] error adding user to team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := services.AddUserToTeam(ctx.Request.Context(), teamUser); err != nil {
		logger.Error.Printf("[controllers.AddUserToTeam] error adding user to team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "user added to team successfully")
}

func RemoveUserFromTeam(ctx *gin.Context) {
	user_id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		logger.Error.Printf("[controllers.RemoveUserFromTeam] error: invalid user_id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user_id",
		})
		return
	}

	team_id, err := strconv.Atoi(ctx.Param("team_id"))
	if err != nil {
		logger.Error.Printf("[controllers.RemoveUserFromTeam] error: invalid team_id %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid team_id",
		})
		return
	}

	err = services.RemoveUserFromTeam(ctx.Request.Context(), user_id, team_id)
	if err != nil {
		logger.Error.Printf("[controllers.RemoveUserFromTeam] error removing user from team %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user removed from team successfully",
	})
}
