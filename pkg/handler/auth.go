package handler

import (
	"mrs_project/pkg/models"
	"mrs_project/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"id": id,
	})
}

type loginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var loginCred loginUser

	if err := c.BindJSON(&loginCred); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, role, err := h.services.CheckUser(loginCred.Email, loginCred.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := utils.GenerateToken(userID, role)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
