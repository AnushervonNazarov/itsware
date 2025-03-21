package controllers

import (
	"itsware/internal/models"
	"itsware/internal/services"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := services.CreateUser(ctx.Request.Context(), user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, "user created successfully")
}

func GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := services.GetUser(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func GetAllUsers(ctx *gin.Context) {
	users, err := services.GetAllUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func UpdateUser(ctx *gin.Context) {
	var userInput models.User
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input data",
		})
		return
	}

	err := services.UpdateUser(ctx.Request.Context(), userInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = services.DeleteUser(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}
