package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// SQLInstance implement database interface
type SQLInstance struct {
	*sql.DB
}

// Connect open connection to postgresql instance
func Connect(host string, port int, user string, pwd string, dbName string) (*SQLInstance, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pwd, dbName)
	db, err := sql.Open("postgres", psqlInfo)

	// Establish connection with db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &SQLInstance{db}, err
}

// SQLQuery execute a query
func (i *SQLInstance) SQLQuery(query string) ([]string, error) {
	rows, err := i.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return nil, err
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, err
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		fmt.Printf("%#v\n", result)
	}

	// Catch errors happening during iterations
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CloseConnection close connection
func (i *SQLInstance) CloseConnection() {
	i.Close()
}
