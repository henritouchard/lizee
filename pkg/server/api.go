package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	productAvailabilityQuery = `
	SELECT COUNT(*) 
		FROM rental_order 
		WHERE product=$1 
			AND (start_date, end_date + INTERVAL '4 DAYS') 
				OVERLAPS ($2::date - INTERVAL '2 DAYS', $3::date) 
			IS FALSE;`
	categoryQuery = `
	SELECT p.id , p.name, p.stock - 
	(
		SELECT COUNT(*) 
			FROM rental_order as r 
			WHERE r.product = p.id
				AND (r.start_date, r.end_date + INTERVAL '5 DAYS') OVERLAPS ($1::date - INTERVAL '2 DAYS', $2::date) 
				IS TRUE
	) AS availability, p.picture
		FROM product as p WHERE cstr_category=$3
	`
)

type productType struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Availability int    `json:"availability"`
	Picture      string `json:"picture"`
}
type categoryType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type apiHandler interface {
	AddAPI(string, string, func(*gin.Context)) error
}

// --> This is what we call an interface, in go any structure that satisfy following functions
// can be provided to InitAPIStorage(). It allows to change db when needed without modifying this package
// and avoid to share postgresql package dependencies with this package.
type strorageInstance interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Prepare(string) (*sql.Stmt, error)
	QueryRow(string, ...interface{}) *sql.Row
}

var db strorageInstance

// InitAPIStorage will keep a pointer to db to interact with external packages
func (router *Server) InitAPIStorage(s strorageInstance) {
	db = s
}

func listCategories(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM category")
	if err != nil {
		c.JSON(http.StatusBadRequest, rows)
	}
	result, err := rowsToJSON(rows)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)
}

func checkProductAvailability(c *gin.Context) {
	productID := c.Query("productID")
	begin := c.Query("fromDate")
	end := c.Query("toDate")

	stmt, err := db.Prepare(productAvailabilityQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	start := time.Now()
	row := stmt.QueryRow(productID, begin, end)
	fmt.Printf("call took %v\n", time.Since(start))

	var r int
	err = row.Scan(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"availability": r,
	})
}

func checkCategoryAvailability(c *gin.Context) {
	categoryID := c.Query("categoryID")
	begin := c.Query("fromDate")
	end := c.Query("toDate")

	stmt, err := db.Prepare(categoryQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	rows, err := stmt.Query(begin, end, categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	result, err := rowsToJSON(rows)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)
}

func post(c *gin.Context) {
	// var postMessage postMessage
	// if err := c.ShouldBindJSON(&postMessage); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": "thank for your message",
		"error":   nil,
	})
}
