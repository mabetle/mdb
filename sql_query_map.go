package mdb

// QueryForMaps query for maps
// JSONDatas holds rows key and value
// columns holds include columns
// good ways to store query result.
// because map has no sequence concepts, so columns is needed when access
// columns.
func (s Sql) QueryForMaps(sql string, args ...interface{}) (JSONDatas []map[string]interface{}, columns []string, err error) {
	rows, errQ := s.Query(sql, args...)
	if errQ != nil {
		err = errQ
		return
	}
	defer rows.Close()
	columns, err = rows.Columns()
	if err != nil {
		return
	}
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		rows.Scan(scanArgs...)
		record := make(map[string]interface{})
		for i, col := range values {
			record[columns[i]] = col
		}
		JSONDatas = append(JSONDatas, record)
	}
	return
}
