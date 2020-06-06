// Package products contains all products concerns, from special types to allow
// JSON serialize/deserialize and execute queries to database
package products

import (
	"lizee/pkg/errortypes"

	"database/sql"
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
	ID           int    `json:"id"` // Tags allow automatic json deserialization
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

// ProductQuantity defines data we get from API for modifying products quantity
type ProductQuantity struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
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
	Product   int `json:"product_id"`
	Asked     int `json:"quantity"`
	Available int `json:"availability"`
}

// Available is returned when you query alla available products
type Available struct {
	ProductID int `json:"product_id"`
	Available int `json:"available"`
}

// parseDates ensure that structure contains dates by checking it's validity (and avoid sql injections)
func parseDates(fromDate, toDate string) (string, string, error) {
	convert := func(date string) (string, error) {
		jsTime, err := time.Parse("2006-01-02", date)
		if err != nil {
			errortypes.BadDate(date)
			return "", err
		}
		return jsTime.Format("2006-01-02"), nil
	}

	fd, err := convert(fromDate)
	if err != nil {
		return fd, "", err
	}
	td, err := convert(toDate)
	if err != nil {
		return "", td, err
	}
	return fd, td, nil
}

// CheckAvailabilityByProduct query all products available
// at given date corresponding to this category.
func CheckAvailabilityByProduct(c *ProductQuery) (int, error) {
	// Prepare query statement
	stmt, err := db.Prepare(queryProductAvailability)
	if err != nil {
		return 0, err
	}

	c.FromDate, c.ToDate, err = parseDates(c.FromDate, c.ToDate)
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

// AllAvailable query to db all products availables from/to dates and
// returns an array of { “product_id”: “1234”, “available”: 6 }
func AllAvailable(c *ProductQuery) ([]Available, error) {
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

	availables := make([]Available, 0, 100)
	for rows.Next() {
		av := Available{}
		err := rows.Scan(&av.ProductID, &av.Available)
		if err != nil {
			return nil, err
		}
		availables = append(availables, av)
	}

	return availables, nil
}

// ModifyProductQuantity update product quantity
// it should not be used without caution if you don't
// want to get negatives quantity products on the iterface
func ModifyProductQuantity(p *ProductQuantity) error {
	// No error found, prepare db statement
	insertStmt, err := db.Prepare(queryModifyQuantity)
	if err != nil {
		return err
	}
	// Execute queries.
	// TODO: find a way to aggregate queries and execute everything in once

	_, err = insertStmt.Exec(p.ProductID, p.Quantity)

	return err
}
