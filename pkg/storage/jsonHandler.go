package storage

import "log"

// QueryToJSON returns  direct formatted json from db query
func (db *SQLInstance) QueryToJSON(query string, args ...interface{}) ([]map[string]interface{}, error) {
	// an array of JSON objects
	// the map key is the field name
	var objects []map[string]interface{}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
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
				// Avoid VARCHAR types to give encrypted JSON
				v = new(string)
			default:
				v = new(interface{}) // destination must be a pointer
			}

			object[column.Name()] = v
			values[i] = v
		}

		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		// use to see what is produced
		log.Printf("%#v", object)

		objects = append(objects, object)
	}

	return objects, nil
}
