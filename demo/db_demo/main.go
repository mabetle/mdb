package main

import (
	"fmt"
	"mabetle/libs/hubs"
	"github.com/mabetle/mdb"
	"github.com/mabetle/mdb/demo"
)

var (
	sql   = hub.NewDemoSql()
	table = "demo_table"
)

func DemoCountTableRows() {
	v, _ := sql.CountTableRows(table)
	fmt.Println(v)
}

func DemoCountTableColumns() {
	v := sql.CountTableColumns(table)
	fmt.Println(v)
}

func DemoGetTableColumns() {
	v := sql.GetTableColumns(table)
	fmt.Println(v)
}

func DemoPrintTableData() {
	//util.GetTableRows(sql,table)
	sql.PrintTable(table)
}

func QueryArrayAndPrint() {
	q := "select * from demo_table"
	data := sql.QueryForArray(true, q)
	mdb.PrintRowsArray(data)
}

func InitDB() {
	demo.InitDB(sql)
}

func main() {
	InitDB()
	//DemoCountTableRows()
	//DemoGetTableColumns()
	//DemoCountTableColumns()
	//DemoPrintTableData()
	//util.PrintQuery(sql,"select UserName, Password, RealName from user_info where UserName = ?", "zsc")
	//sql:= "select UserName, Password, RealName from user_info where UserName = ?"
	q := "select * from demo_table"
	sql.PrintQuery(q)
}
