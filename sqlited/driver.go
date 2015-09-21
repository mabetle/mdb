package sqlited

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mabetle/mdb"
)

// NewSqlDB
func NewSqlDB(location string) (*sql.DB, error) {
	logger.Trace("NewSqlDB() location:", location)
	return sql.Open("sqlite3", location)
}

// NewSql
func NewSql(location string) (*mdb.Sql, error) {
	logger.Trace("NewSql() location:", location)
	db, err := NewSqlDB(location)
	if err != nil {
		return nil, err
	}
	return mdb.NewSql(db), nil
}
