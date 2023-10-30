package controllers

import (
	"NewApp/db"
	"NewApp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var allowedFields = []string{"title", "artist", "genre", "release_year", "description", "price", "status", "seller_id"}
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

	newRecord, err := models.CreateRecord(db.GetDB(), record.Title, record.Artist, record.Genre, record.ReleaseYear, record.Description, record.Price, record.Status, record.SellerID)
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

	updatedRecord, err := models.UpdateRecord(db.GetDB(), uint(recordID), record.Title, record.Artist, record.Genre, record.ReleaseYear, record.Description, record.Price, record.Status)
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

	err = models.DeleteRecord(db.GetDB(), uint(recordID))
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

	record, err := models.GetRecordByID(db.GetDB(), uint(recordID))
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
	records, err := models.GetAllRecords(db.GetDB(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

func GetRecordsBySellerID(c *gin.Context) {
	sellerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}
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
	records, err := models.GetRecordsBySellerID(db.GetDB(), filters, uint(sellerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}
