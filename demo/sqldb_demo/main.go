package main

import (
	"github.com/mabetle/mdb/demo"
	"github.com/mabetle/mdb/sql_mysqld"
)

var(
	sql=sql_mysqld.NewSqlFromSchema("demo")
)

// SqlDemo
func Demo(){
	demo.InitDB(sql)
}

func main() {
	Demo()
}


