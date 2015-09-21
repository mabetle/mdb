package main

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"github.com/mabetle/mdb/sql_mysqld"
)

func checkError(err error){
	if nil == err {
		return
	}
	fmt.Println(err)
}

func Demo(db *sql.DB, table string){
	q := "select * from " + table
	rows, err := db.Query(q)
	checkError(err)
	defer rows.Close()

	columns, err := rows.Columns()
	checkError(err)

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	//define a map store json values
	datas:=map[int]string{}

	i:=0
	for rows.Next() {
		i++
		err = rows.Scan(scanArgs...)
		checkError(err)

		fmt.Printf("%v\n",scanArgs)

		record := make(map[string]interface{})

		for i, col := range values {
			if col != nil {
				v:=fmt.Sprintf("%s", string(col.([]byte)))
				record[columns[i]] = v
			}
		}
		s, _ := json.Marshal(record)

		datas[i]=string(s)
		//fmt.Printf("%s\n", s)
	}

	//print data
	for _,d:=range datas{
		fmt.Printf("JSON:%v\n", d)
	}
}


func Demo1(){
	sql:=sql_mysqld.NewDBFromSchema("demo")
	Demo(sql, "demo_table")
}

func Demo2(){
	sql:=sql_mysqld.NewDBFromSchema("web_common")
	Demo(sql, "user_info")
}

func main() {
	Demo2()
}


