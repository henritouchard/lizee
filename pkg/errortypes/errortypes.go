// Package errortypes allow to customise error messages for more readability
package errortypes

import "errors"

const (
	// UnavailableProduct is sent when a product is asked and not available
	UnavailableProduct = "This product is not available anymore"
	// DbConnection is sent when server can't connect to database
	DbConnection = "database connection error: "
)

// New creates custom error
func New(err string) error {
	return errors.New(err)
}
