package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/health", health)

	r.GET("/product", getProducts)
	r.GET("/product/:id", getProductById)
	r.POST("/product/grid", productGrid)
	r.POST("/product", insertProduct)
	r.PUT("/product/:id", updateProduct)
	r.DELETE("/product/:id", deleteProduct)

	r.GET("/entry", getEntries)
	r.GET("/entry/:id", getEntryById)
	r.POST("/entry/grid", entryGrid)
	r.POST("/entry", insertEntry)
	r.PUT("/entry/:id", updateEntry)
	r.DELETE("/entry/:id", deleteEntry)
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
