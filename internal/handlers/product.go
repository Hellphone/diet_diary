package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"diet_diary/internal/database"
	"diet_diary/internal/domain"
	"diet_diary/internal/services"
)

func getProducts(c *gin.Context) {
	product, err := services.GetProducts(nil)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": product})
}

func getProductById(c *gin.Context) {
	id, ok := getInt64Param(c, "id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param 'id'"})
		return
	}

	product, err := services.GetProductById(id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func productGrid(c *gin.Context) {
	var reqBody struct {
		Filter *database.Filter `json:"filter"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := services.GetProducts(reqBody.Filter)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": product})
}

func insertProduct(c *gin.Context) {
	var reqBody struct {
		Product *domain.Product `json:"product"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.InsertProduct(reqBody.Product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Product with id = %d inserted", id)})
}

func updateProduct(c *gin.Context) {
	var reqBody struct {
		Product *domain.Product `json:"product"`
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

	reqBody.Product.ID = id
	_, err := services.UpdateProduct(reqBody.Product)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Product with id = %d updated", id)})
}

func deleteProduct(c *gin.Context) {
	id, ok := getInt64Param(c, "id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param 'id'"})
		return
	}

	id, err := services.DeleteProduct(id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Product with id = %d deleted", id)})
}
