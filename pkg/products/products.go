package products

import (
	"database/sql"
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

var db strorageInstance

// InitAPIStorage will keep a pointer to db to interact with external packages
func InitAPIStorage(s strorageInstance) {
	db = s
}

// Product defines basic informations of on product
type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Availability int    `json:"availability"`
	Picture      string `json:"picture"`
}

// ProductQuery defines query we get for products availability corresponding to category
type ProductQuery struct {
	ProductID int    `form:"productID"`
	FromDate  string `form:"fromDate"`
	ToDate    string `form:"toDate"`
}

// ProductOrder defines basic informations of product rental order
type ProductOrder struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
	FromDate string  `json:"fromDate"`
	ToDate   string  `json:"toDate"`
}

// CheckAvailabilityByProduct query all products available
// at given date corresponding to this category.
func CheckAvailabilityByProduct(c *ProductQuery) (int, error) {
	stmt, err := db.Prepare(queryProductAvailability)
	if err != nil {
		return 0, err
	}

	row := stmt.QueryRow(c.ProductID, c.FromDate, c.ToDate)

	var available int
	err = row.Scan(&available)
	return available, err
}

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

// ProcessRentalOrder Check if a product is available
// in base and execute rental if possible.
func ProcessRentalOrder(orders []ProductOrder) (interface{}, error) {
	var err error
	for _, order := range orders {
		order.FromDate, order.ToDate, err = parseDates(order.FromDate, order.ToDate)
		if err != nil {
			return nil, err
		}
		query, err := buildRentalQuery(order)
		if err != nil {
			return nil, err
		}

		r := db.QueryRow(query)
		var t interface{}
		r.Scan(&t)
		if t != nil {
			fmt.Printf("%v", t)
		}
	}

	return nil, nil
}
