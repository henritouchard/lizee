// Package products contains all products concerns, from special types to allow
// JSON serialize/deserialize and execute queries to database
package products

import (
	"lizee/pkg/errortypes"
	"lizee/pkg/utils"

	"database/sql"
	"errors"
	"fmt"
	"time"
)

// --> This is what we call an interface, in go any structure that satisfy following functions
// can be provided to InitAPIStorage(). It allows to change db when needed without modifying this package
// and avoid to share postgresql package dependencies with this package.
type strorageInstance interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Prepare(string) (*sql.Stmt, error)
	QueryRow(string, ...interface{}) *sql.Row
}

// Global variable scoped to the package to make database queries easier
var db strorageInstance

// InitAPIStorage will keep a pointer to db to interact with external packages
func InitAPIStorage(s strorageInstance) {
	db = s
}

// Product defines basic informations concerning product
type Product struct {
	ID           int    `json:"id"` // This allows automatic json deserialization
	Name         string `json:"name"`
	Availability int    `json:"availability"`
	Picture      string `json:"picture"`
}

// ProductQuery defines data we get from API for products availability
type ProductQuery struct {
	ProductID int    `form:"product_id" json:"product_id"`
	FromDate  string `form:"from" json:"from"`
	ToDate    string `form:"to" json:"to"`
}

// ProductOrder defines basic informations concerning product rental order
type ProductOrder struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
	FromDate string  `json:"from"`
	ToDate   string  `json:"to"`
}

// UnAvailable is returned when you asked unavailable products quantity
type UnAvailable struct {
	Product   Product `json:"product"`
	Asked     int     `json:"quantity"`
	Available int     `json:"availability"`
}

// Available is returned when you query alla available products
type Available struct {
	ProductID string `json:"product_id"`
	Available int    `json:"available"`
}

// CheckAvailabilityByProduct query all products available
// at given date corresponding to this category.
func CheckAvailabilityByProduct(c *ProductQuery) (int, error) {
	// Prepare query statement
	stmt, err := db.Prepare(queryProductAvailability)
	if err != nil {
		return 0, err
	}
	// Execute query
	row := stmt.QueryRow(c.ProductID, c.FromDate, c.ToDate)
	// Extract result
	var available int
	err = row.Scan(&available)
	return available, err
}

// parseDates ensure that structure contains dates by checking it's validity (and avoid sql injections)
func parseDates(fromDate, toDate string) (string, string, error) {
	convert := func(date string) (string, error) {
		jsTime, err := time.Parse("2006-01-02", date)
		if err != nil {
			fmt.Printf("Error %v\n", err)
			return "", err
		}
		return jsTime.Format("2006-01-02"), nil
	}
	fd, err := convert(fromDate)
	td, err := convert(toDate)
	return fd, td, err
}

// ProcessRentalOrder Check if a product is available in base and
// execute rental if possible. in case it's not, no insert are executed,
// and a list of unavailable products is returned
func ProcessRentalOrder(orders []ProductOrder) ([]UnAvailable, error) {
	unavailables := []UnAvailable{}
	// check if products are still available
	for _, order := range orders {
		o := ProductQuery{order.Product.ID, order.FromDate, order.ToDate}
		available, err := CheckAvailabilityByProduct(&o)
		if err != nil {
			return nil, err
		}
		if available < order.Quantity {
			unavailables = append(unavailables, UnAvailable{order.Product, order.Quantity, available})
		}
	}

	if len(unavailables) > 0 {
		return unavailables, errors.New(errortypes.UnavailableProduct)
	}

	// No error found, prepare db statement
	insertStmt, err := db.Prepare(queryInsert)
	if err != nil {
		return nil, err
	}
	// Execute queries.
	// TODO: find a way to aggregate queries and execute everything in once
	for _, order := range orders {
		order.FromDate, order.ToDate, err = parseDates(order.FromDate, order.ToDate)
		_, err := insertStmt.Exec(order.Product.ID, order.Quantity, order.FromDate, order.ToDate)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// AllAvailable query to db all products availables from/to dates and
// returns an array of { “product_id”: “1234”, “available”: 6 }
func AllAvailable(c *ProductQuery) ([]map[string]interface{}, error) {
	// Prepare statement
	stmt, err := db.Prepare(queryCheckAllAvailable)
	if err != nil {
		return nil, err
	}

	// Check if dates are corrects
	c.FromDate, c.ToDate, err = parseDates(c.FromDate, c.ToDate)
	if err != nil {
		return nil, err
	}

	// Execute query
	rows, err := stmt.Query(c.FromDate, c.ToDate)
	if err != nil {
		return nil, err
	}

	return utils.RowsToJSON(rows)
}
