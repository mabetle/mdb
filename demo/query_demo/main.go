package main

import (
	"github.com/mabetle/mdb/sql_mysqld"
)


var(
	sql = sql_mysqld.NewSqlFromSchema("demo")
)

func main() {
	sql.NewQuery().Table("demo_table").Print()

	sql.NewQuery().Table("demo_table").Columns("DemoName").Where("DemoAge > 0 ").Order("DemoName").Print()

}

