package server

import (
	"net/http"

	"lizee/pkg/products"

	"github.com/gin-gonic/gin"
)

func listCategories(c *gin.Context) {
	categories, err := products.ListCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, categories)
}

// Correspond to the demand of exercise but to correct
func checkProductAvailability(c *gin.Context) {
	p := products.ProductQuery{}
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	r, err := products.CheckAvailabilityByProduct(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"availability": r,
	})
}

func checkCategoryAvailability(c *gin.Context) {
	categoryQ := products.CategoryQuery{}
	if err := c.Bind(&categoryQ); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	result, err := products.CheckAvailabilityByCategory(&categoryQ)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)
}

// postOrder is a handler used when user posts it's rental order
// it will control availabilty of products and reserve it if is available
func postOrder(c *gin.Context) {
	var n []products.ProductOrder
	// Deserialize json order to array of productOrder structures
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := products.ProcessRentalOrder(n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}
