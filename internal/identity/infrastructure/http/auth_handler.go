package http

import (
	"net/http"

	"prueba-tecnica-back/internal/identity/application"
	"prueba-tecnica-back/internal/identity/infrastructure/http/dto"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	registerUseCase *application.RegisterUseCase
	loginUseCase    *application.LoginUseCase
}

func NewAuthHandler(
	registerUseCase *application.RegisterUseCase,
	loginUseCase *application.LoginUseCase,
) *AuthHandler {
	return &AuthHandler{
		registerUseCase: registerUseCase,
		loginUseCase:    loginUseCase,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.registerUseCase.Execute(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, token, err := h.loginUseCase.Execute(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	resp := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  resp,
		"token": token,
	})
}
