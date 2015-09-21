package sql_mysqld

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mabetle/mdb/dbconf"
)

// NewDBFromDBConf
func NewDBFromDBConf(c *dbconf.DBConf) (*sql.DB, error) {
	logger.Tracef("NewDBFromDBConf() Host:%s Database:%s User:%s ", c.Host, c.Database, c.User)
	db, errDB := sql.Open(c.Driver, c.GetMySqlConnURL())
	if logger.CheckError(errDB) {
		logger.Errorf("Really cannot go on, open sql error.Error:%v Host:%s Database:%s", errDB, c.Host, c.Database)
		//panic(errDB)
		return db, errDB
	}

	// infact if db not work, app still can go on.
	if err := db.Ping(); logger.CheckError(err) {
		logger.Errorf("Really cannot go on, Ping db error. Host:%s Database:%s", c.Host, c.Database)
		// app should know how to do with this err.
		//panic(err)
		return db, err
	}
	return db, nil
}
