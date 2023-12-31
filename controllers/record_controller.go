package controllers

import (
	"fake-discogs-api/database"
	"fake-discogs-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var allowedFields = []string{"genre", "release_year", "status"}
var allowedStatuses = []string{"available", "reserved", "sold"}
var allowedGenres = []string{"rock", "pop", "jazz", "hip-hop", "electronic", "classical", "metal", "country", "folk", "blues", "reggae", "latin", "punk", "indie", "r&b", "soul", "funk", "dance", "world", "experimental", "new age", "spoken", "children's", "comedy", "other"}

func isFieldAllowed(field string) bool {
	for _, allowedField := range allowedFields {
		if field == allowedField {
			return true
		}
	}
	return false
}

func CreateRecord(c *gin.Context) {
	var record models.Record
	if err := c.BindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	SellerID := uint(c.GetInt("user_id"))
	newRecord, err := models.CreateRecord(database.GetDB(), record.Title, record.Artist, record.Genre, record.ReleaseYear, record.Description, record.Price, record.Status, SellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newRecord)
}

func UpdateRecord(c *gin.Context) {
	recordID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	var record models.Record
	if err := c.BindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRecord, err := models.UpdateRecord(database.GetDB(), uint(recordID), record.Title, record.Artist, record.Genre, record.ReleaseYear, record.Description, record.Price, record.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedRecord)
}

func DeleteRecord(c *gin.Context) {
	recordID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}
	userID := c.GetInt("user_id")
	username, err := models.GetUsernameByID(database.GetDB(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	record, err := models.GetRecordByID(database.GetDB(), uint(recordID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	if record.SellerName != username {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	err = models.DeleteRecord(database.GetDB(), uint(recordID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func DeleteAllRecordsBySellerName(c *gin.Context) {
	userID := c.GetInt("user_id")
	username, err := models.GetUsernameByID(database.GetDB(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteAllRecordsBySellerName(database.GetDB(), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func GetRecordByID(c *gin.Context) {
	recordID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	record, err := models.GetRecordByID(database.GetDB(), uint(recordID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, record)
}

func GetAllRecords(c *gin.Context) {
	var filters []models.Filter

	if err := c.BindJSON(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, filter := range filters {
		if !isFieldAllowed(filter.Field) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid field"})
			return
		}
	}
	records, err := models.GetAllRecords(database.GetDB(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

func GetRecordsBySellerName(c *gin.Context) {
	sellerName := c.Param("name")
	var filters []models.Filter

	if err := c.BindJSON(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, filter := range filters {
		if !isFieldAllowed(filter.Field) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid field"})
			return
		}
	}
	records, err := models.GetRecordsBySellerName(database.GetDB(), filters, sellerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

func SearchRecordsWithPagination(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	searchTerm := c.Query("search_term")
	var filters []models.Filter
	if err := c.BindJSON(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, filter := range filters {
		if !isFieldAllowed(filter.Field) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid field"})
			return
		}
	}
	records, err := models.SearchRecordsWithPagination(database.GetDB(), filters, searchTerm, page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}
