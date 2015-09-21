package mdb

import (
	"encoding/json"
)

// QueryForJsonData
// returns Map key is rows id, value stores row marshal json data.
func (s Sql)QueryForJsonData(sql string, args ... interface{})(map[string]string){
	rows, _ := s.Query(sql, args...)
	defer rows.Close()

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values	 := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	//its a map
	// strore id and json values
	rowsData := map[string]string{}

	index:=0

	for rows.Next() {
		_ = rows.Scan(scanArgs...)

		record := make(map[string]interface{})

		for i, col := range values {
			if col != nil {
				//v:=fmt.Sprintf("%s", string(col.([]byte)))
				v:=GetString(col)
				record[columns[i]] = v
			}
		}

		s, _ := json.Marshal(record)

		//FIXME
		// index should put Row ID
		rowsData[string(index)] = string(s)

		index++
	}
	return rowsData
}

