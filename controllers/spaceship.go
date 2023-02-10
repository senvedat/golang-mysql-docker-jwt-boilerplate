package controllers

import (
	"go-rest-example/infra/database"
	"go-rest-example/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSpaceships(c *gin.Context) {
	var spaceships []models.Spaceship
	result, err := database.DB.Query("SELECT id, name, class, crew, image, value, status FROM spaceships")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for result.Next() {
		var spaceship models.Spaceship
		err := result.Scan(&spaceship.ID, &spaceship.Name, &spaceship.Class, &spaceship.Crew, &spaceship.Image, &spaceship.Value, &spaceship.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		spaceships = append(spaceships, spaceship)
	}
	c.JSON(http.StatusOK, gin.H{"data": spaceships})
}

func GetSpaceship(c *gin.Context) {
	id := c.Param("id")
	var spaceships []models.Spaceship
	result, err := database.DB.Query("SELECT * FROM spaceships WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for result.Next() {
		var spaceship models.Spaceship
		err := result.Scan(&spaceship.ID, &spaceship.Name, &spaceship.Class, &spaceship.Crew, &spaceship.Image, &spaceship.Value, &spaceship.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		spaceships = append(spaceships, spaceship)
	}
	c.JSON(http.StatusOK, gin.H{"data": spaceships})
}

func CreateSpaceship(c *gin.Context) {
	var spaceship models.Spaceship
	if err := c.ShouldBindJSON(&spaceship); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := "INSERT INTO spaceships (name, class, crew, image, value, status) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := database.DB.Exec(query, spaceship.Name, spaceship.Class, spaceship.Crew, spaceship.Image, spaceship.Value, spaceship.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "id": lastInsertId})
}

func UpdateSpaceship(c *gin.Context) {
	id := c.Param("id")
	var spaceship models.Spaceship
	if err := c.ShouldBindJSON(&spaceship); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := "UPDATE spaceships SET name=?, class=?, crew=?, image=?, value=?, status=? WHERE id = ?"
	result, err := database.DB.Exec(query, spaceship.Name, spaceship.Class, spaceship.Crew, spaceship.Image, spaceship.Value, spaceship.Status, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "id": lastInsertId})
}

func DeleteSpaceship(c *gin.Context) {
	id := c.Param("id")
	result, err := database.DB.Exec("DELETE FROM spaceships WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No spaceship found with the given id"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
