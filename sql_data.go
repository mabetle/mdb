package mdb

import (
	"fmt"
)

// QueryForJsonData
func (s Sql)QueryForData(sql string, args ... interface{})(rowsData [][]string){
	return s.QueryForArray(false, sql, args ... )
}

func (s Sql)PrintQueryData(sql string, args ... interface{}){
	data:=s.QueryForData(sql, args ... )
	fmt.Println("====Begin Print Query Data====")
	PrintData(data)
	fmt.Println("====Begin Print Query Data====")
}


