package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/usecase"
	"net/http"
)

type AuthController struct {
	AuthUsecase usecase.AuthUsecase
}

func NewAuthController(authUsecase usecase.AuthUsecase) AuthController {
	return AuthController{
		AuthUsecase: authUsecase,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := ac.AuthUsecase.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
