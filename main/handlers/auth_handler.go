package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	user, err := authenticateUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := generateToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: token})
}

func GetDashboard(c *gin.Context) {
	username, _ := c.Get(authContextUsernameKey)
	role, _ := c.Get(authContextRoleKey)

	response := gin.H{
		"message":  "Welcome, " + username.(string),
		"username": username,
	}

	if role != nil {
		response["role"] = role
	}

	c.JSON(http.StatusOK, response)
}
