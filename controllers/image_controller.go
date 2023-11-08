package controllers

import (
	"fake-discogs-api/database"
	"fake-discogs-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	imageUrl := generateUUID()
	imagePath := "images/" + imageUrl + ".jpg"
	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	recordIDstr := c.Param("id")
	recordID, err := strconv.Atoi(recordIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newImage, err := models.CreateImage(database.GetDB(), uint(recordID), imageUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newImage)
}

func generateUUID() string {
	return uuid.New().String()
}
