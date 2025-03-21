package controllers

import (
	"itsware/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (s *User) SignIn(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, err := s.Service.SignIn(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, accessTokenResponse{accessToken})
}
