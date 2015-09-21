package main

import (
	"github.com/mabetle/mdb/demo/ioc"
	"github.com/mabetle/mdb/demo"
)

var (
	sql = ioc.GetSql()
)

func InitDB(){
	demo.InitDB(sql)
}

func main() {
	InitDB()
}

