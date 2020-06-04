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
		return
	}
	c.JSON(http.StatusOK, categories)
}

// checkProductAvailability will check if product is available
// by using it's ID
func checkProductAvailability(c *gin.Context) {
	p := products.ProductQuery{}
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result, err := products.CheckAvailabilityByProduct(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"availability": result,
	})
}

// checkProductsAvailability will return all available
// products id and quantity available
func checkProductsAvailability(c *gin.Context) {
	p := products.ProductQuery{}
	// Deserialize json order to array of productOrder structures
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := products.AllAvailable(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// checkCategoryAvailability will return all available products corresponding to
// the category asked
func checkCategoryAvailability(c *gin.Context) {
	categoryQ := products.CategoryQuery{}
	if err := c.Bind(&categoryQ); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := products.CheckAvailabilityByCategory(&categoryQ)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// postOrder is a handler used when user posts it's rental order
// it will control availabilty of products and reserve it if is available
func postOrder(c *gin.Context) {
	var p []products.ProductOrder
	// Deserialize json order to array of productOrder structures
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if product, err := products.ProcessRentalOrder(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "product": product})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

// modifyQuantity modify quantity of provided product
func modifyQuantity(c *gin.Context) {
	p := products.ProductQuantity{}
	// Deserialize json order to array of productOrder structures
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := products.ModifyProductQuantity(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
