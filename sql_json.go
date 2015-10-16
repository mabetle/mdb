package mdb

import (
	"encoding/json"
	"fmt"
)

// QueryForMaps
func (s Sql) QueryForMaps(
	sql string,
	include string,
	exclude string,
	args ...interface{},
) ([]map[string]interface{}, error) {
	jsonDatas := make([]map[string]interface{}, 0)
	rows, err := s.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, errC := rows.Columns()
	if errC != nil {
		return nil, errC
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
	return jsonDatas, nil
}

// QueryForJsonData
// returns Map key is rows id, value stores row marshal json data.
func (s Sql) QueryForJsonData(sql string, args ...interface{}) map[string]string {
	jsonDatas, _ := s.QueryForMaps(sql, "", "", args...)
	rowsData := make(map[string]string)
	for _, v := range jsonDatas {
		s, _ := json.Marshal(v)
		id := fmt.Sprint(v["Id"])
		rowsData[id] = string(s)
	}
	return rowsData
}
