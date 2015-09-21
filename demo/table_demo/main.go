package main

import (
	//"github.com/mabetle/mdb"
	"github.com/mabetle/mdb/sql_mysqld"
	"fmt"
)


var(
	db		=sql_mysqld.NewSqlFromSchema("demo")
	table	=db.NewTable("demo_table")
)

func Demo(){
	fmt.Println(table.CountRows())
	fmt.Println(table.CountColumns())
	fmt.Println(table.GetColumns())
	table.Print()
}

func main() {
	Demo()
}


