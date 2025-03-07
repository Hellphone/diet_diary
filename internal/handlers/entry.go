package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"diet_diary/internal/database"
	"diet_diary/internal/domain"
	"diet_diary/internal/services"
)

func getEntries(c *gin.Context) {
	entry, err := services.GetEntries(nil)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"entrys": entry})
}

func getEntryById(c *gin.Context) {
	id, ok := getInt64Param(c, "id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param 'id'"})
		return
	}

	entry, err := services.GetEntryById(id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"entry": entry})
}

func getEntryTotalByDate(c *gin.Context) {
	var reqBody struct {
		Date *time.Time `json:"date"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entryTotal, err := services.GetEntryTotalByDate(reqBody.Date)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": entryTotal})
}

func entryGrid(c *gin.Context) {
	var reqBody struct {
		Filter *database.Filter `json:"filter"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entries, err := services.GetEntries(reqBody.Filter)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"entries": entries})
}

func insertEntry(c *gin.Context) {
	var reqBody struct {
		Entry *domain.Entry `json:"entry"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.InsertEntry(reqBody.Entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Entry with id = %d inserted", id)})
}

func updateEntry(c *gin.Context) {
	var reqBody struct {
		Entry *domain.Entry `json:"entry"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, ok := getInt64Param(c, "id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param 'id'"})
		return
	}

	reqBody.Entry.ID = id
	_, err := services.UpdateEntry(reqBody.Entry)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Entry with id = %d updated", id)})
}

func deleteEntry(c *gin.Context) {
	id, ok := getInt64Param(c, "id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param 'id'"})
		return
	}

	id, err := services.DeleteEntry(id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Entry with id = %d deleted", id)})
}
