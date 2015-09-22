package mdb

import (
//"fmt"
//"github.com/mabetle/mcore"
)

// use []][]string array to store sql query result.
// if withHeader is true, the result contain column names in array first element.
func (s Sql) QueryForArray(withHeader bool, sql string, args ...interface{}) (result [][]string) {
	rows, _ := s.Query(sql, args...)
	defer rows.Close()

	columns, _ := rows.Columns()

	if withHeader {
		result = append(result, columns)
	}

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		_ = rows.Scan(scanArgs...)
		rowData := []string{}
		for _, col := range values {
			if col != nil {
				//v:=fmt.Sprintf("%s", string(col.([]byte)))
				v := GetString(col)
				rowData = append(rowData, v)
			}
		}
		result = append(result, rowData)
	}
	return
}

// result [][]string include header
func (s Sql) QueryForArrayWithHeader(query string, args ...interface{}) [][]string {
	return s.QueryForArray(true, query, args...)
}

// result [][]string not include header
func (s Sql) QueryForArrayWithoutHeader(query string, args ...interface{}) [][]string {
	return s.QueryForArray(false, query, args...)
}
