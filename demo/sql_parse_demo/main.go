package main

import (
	"fmt"
	"github.com/mabetle/mdb"
)

func demo(){
	q:="select * from demo"
	t:=mdb.ParseSqlTableName(q)
	fmt.Printf("table: %s\n", t)
}

func main() {
	demo()
}
