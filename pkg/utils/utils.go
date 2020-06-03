package utils

import (
	"database/sql"
)

// RowsToJSON transform result of sql query to []map[string]interface{} to allow JSON convertion
func RowsToJSON(rows *sql.Rows) ([]map[string]interface{}, error) {
	// an array of JSON objects
	// the map key is the field name
	var objects []map[string]interface{}

	for rows.Next() {
		// figure out what columns were returned
		// the column names will be the JSON object field keys
		columns, err := rows.ColumnTypes()
		if err != nil {
			return nil, err
		}

		// Scan needs an array of pointers to the values it is setting
		// This creates the object and sets the values correctly
		values := make([]interface{}, len(columns))
		object := map[string]interface{}{}
		for i, column := range columns {
			var v interface{}
			dataType := column.DatabaseTypeName()
			// Special types cases to handle
			switch dataType {
			case "text":
			case "BPCHAR":
				// Avoid VARCHAR types that give base64 encrypted JSON
				v = new(string)
			default:
				v = new(interface{})
			}

			object[column.Name()] = v
			values[i] = v
		}

		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		objects = append(objects, object)
	}
	rows.Close()
	return objects, rows.Err()
}
