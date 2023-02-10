package controllers

import (
	"go-rest-example/infra/database"
	"go-rest-example/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Check if the user exists in the database
	var dbUser models.User
	query := "SELECT username, password FROM users WHERE username = ? AND password = ?"
	err := database.DB.QueryRow(query, user.Username, user.Password).Scan(&dbUser.Username, &dbUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or password is incorrect"})
		return
	}

	// Create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token and get the final encoded token as a string
	tokenString, err := token.SignedString([]byte(viper.GetString("JWT_SACRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while signing the token"})
		return
	}

	// Return the token
	c.JSON(http.StatusOK, gin.H{"success": true, "token": tokenString})
}
