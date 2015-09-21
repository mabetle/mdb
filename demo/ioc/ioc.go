package ioc


import (
	"github.com/mabetle/mdb"
	"github.com/mabetle/mdb/sql_mysqld"
)
var sqlDB *mdb.Sql
func GetSql()*mdb.Sql{
	if sqlDB == nil {
		sqlDB = sql_mysqld.NewSqlFromSchema("demo")
	}
	return sqlDB
}

