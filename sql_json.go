package mdb

import (
	"encoding/json"
	"fmt"
)

// QueryForMaps query for maps
// JSONDatas holds rows key and value
// columns holds include columns
// good ways to store query result.
// because map has no sequence concepts, so columns is needed when access
// columns.
func (s Sql) QueryForMaps(
	sql string,
	args ...interface{},
) (JSONDatas []map[string]interface{}, columns []string, err error) {
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

// QueryForJSONData returns Map key is rows id, value stores row marshal JSON data.
// marshal row to string.
// not a good way.
func (s Sql) QueryForJSONData(sql string, args ...interface{}) map[string]string {
	JSONDatas, _, _ := s.QueryForMaps(sql, args...)
	rowsData := make(map[string]string)
	for _, v := range JSONDatas {
		s, _ := json.Marshal(v)
		id := fmt.Sprint(v["Id"])
		rowsData[id] = string(s)
	}
	return rowsData
}
