package mdb

import (
	"encoding/json"
	"fmt"
)

// QueryForMaps
func (s Sql) QueryForMaps(
	sql string,
	args ...interface{},
) (jsonDatas []map[string]interface{}, columns []string, err error) {
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
		jsonDatas = append(jsonDatas, record)
	}
	return
}

// QueryForJsonData
// returns Map key is rows id, value stores row marshal json data.
func (s Sql) QueryForJsonData(sql string, args ...interface{}) map[string]string {
	jsonDatas, _, _ := s.QueryForMaps(sql, args...)
	rowsData := make(map[string]string)
	for _, v := range jsonDatas {
		s, _ := json.Marshal(v)
		id := fmt.Sprint(v["Id"])
		rowsData[id] = string(s)
	}
	return rowsData
}
