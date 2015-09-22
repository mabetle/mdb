package main

import (
	"github.com/mabetle/mdb/demo/ioc"
	"fmt"
	. "github.com/mabetle/mcore/mtest"
)

var(
	sql = ioc.GetSql()
)

func Demo(){
	sql.PrintDatabases()
	println(sql.IsHasDatabase("demo"))
	println(sql.IsHasDatabase("demo_no"))
	sql.PrintDBTables("demo")
	sql.PrintDBTablesDetail("demo")
	sql.PrintTableColumns("demo", "user_info")
	sql.PrintTable("demo_table")
}

func TableDemo(){
	AssertFalse(sql.IsExistTable("none"))
	AssertTrue(sql.IsExistTable("demo"))
	AssertFalse(sql.IsExistColumn("none","none"))
	AssertTrue(sql.IsExistColumn("demo","id"))
	//AssertFalse(sql.IsExistColumn("demo","id"))
}

func TestDB(){
	AssertTrue(sql.IsHasDatabase("demo"))
	AssertTrue(sql.IsDbExistTable("demo","demo"))
	AssertFalse(sql.IsDbExistTable("demo","none"))
	AssertTrue(sql.IsDbTableExistColumn("demo","demo", "id"))
	AssertFalse(sql.IsDbTableExistColumn("demo","demo", "none"))
}

func main() {
	fmt.Println()
	//Demo()
	//TableDemo()
	TestDB()
}


