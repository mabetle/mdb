package mdb

import (
	"encoding/json"
	"fmt"
)

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
