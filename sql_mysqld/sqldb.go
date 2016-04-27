package sql_mysqld

import (
	"github.com/mabetle/mdb"
	"github.com/mabetle/mdb/dbconf"
)

// NewSql
func NewSql(conf *dbconf.DBConf) (*mdb.Sql, error) {
	logger.Infof("Create new mdb.Sql. Host:%s Schema:%s", conf.Host, conf.Database)
	db, err := NewDBFromDBConf(conf)
	if logger.CheckError(err) {
		return nil, err
	}
	sql := mdb.NewSql(db)
	return sql, nil
}
