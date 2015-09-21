package main

import (
	"github.com/mabetle/mdb/dbconf"
)

func DemoConf(conf *dbconf.DBConf) {
	println("==========Begin show db conf ==============")
	println(conf.GetConnURL())
	println(conf.GetMySqlConnURL())
	println("==========End.. show db conf ==============")
}

func DemoFile() {
	c := dbconf.NewDBConfFromFile("/conf/common/db_demo.conf")
	DemoConf(c)
}

func main() {
	DemoFile()
}
