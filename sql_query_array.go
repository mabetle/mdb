package mdb

// QueryForArray use [][]string array to store sql query result.
// if withHeader is true, the result contain column names in array first element.
func (s Sql) QueryForArray(withHeader bool, sql string, args ...interface{}) (result [][]string) {
	data, header := s.QueryForDataArray(sql, args...)
	// with header
	if withHeader {
		// has header
		if len(header) > 0 {
			result = append(result, header)
		}
		// has datas
		if len(data) > 0 {
			result = append(result, data...)
		}
		return result
	}
	// no header
	return data
}

// QueryForArrayData
func (s Sql) QueryForDataArray(q string, args ...interface{}) (data [][]string, columns []string) {
	rows, err := s.Query(q, args...)
	// query err
	if err != nil {
		return
	}
	defer rows.Close()
	// hold columns
	columns, _ = rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		_ = rows.Scan(scanArgs...)
		rowData := make([]string, len(columns))
		// interface to string
		for index, col := range values {
			if col != nil {
				v := GetString(col)
				rowData[index] = v
			}
		}
		data = append(data, rowData)
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
