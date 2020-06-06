// Package errortypes allow to customise error messages for more readability
package errortypes

import (
	"fmt"
)

const (
	// UnavailableProduct is sent when a product is asked and not available
	UnavailableProduct = "This product is not available anymore"
	// DbConnection is sent when server can't connect to database
	DbConnection = "database connection error: "
	badDate      = "Bad date format, expected: 2020-04-06, got: %s."
)

// New creates custom error
func New(err string) error {
	return fmt.Errorf(err)
}

// BadDate indicates wrong date format
func BadDate(d string) error {
	return fmt.Errorf(badDate, d)
}
