package handlers

import (
	"net/http"

	"github.com/ctrlaltpat/skate-events/internal/models"
	"github.com/ctrlaltpat/skate-events/internal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Service services.UserService
}

type registerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required,min=2"`
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var register registerRequest
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	user := models.User{
		Email:    register.Email,
		Password: string(hashedPassword),
		Name:     register.Name,
	}

	createdUser, err := h.Service.Register(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdUser.Password = ""
	c.JSON(http.StatusCreated, createdUser)
}
