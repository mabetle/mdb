package main


import (
	"github.com/mabetle/mdb"
	"github.com/mabetle/mdb/demo/ioc"
)

var(
	db  =  ioc.GetSql()
)

func main() {
	mdb.NewLinkSql(db).SetTable("demo_table").Print()
}

