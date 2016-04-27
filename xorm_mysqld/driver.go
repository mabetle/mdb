package xorm_mysqld

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/mabetle/mdb"
	"github.com/mabetle/mdb/dbconf"
)

// NewXorm
func NewXorm(conf *dbconf.DBConf) (*mdb.Xorm, error) {
	logger.Tracef("Create new mdb.Xorm, DB: %s ", conf.Database)

	// check args
	if conf == nil {
		return nil, fmt.Errorf("DBConf is nil")
	}

	engine, err := xorm.NewEngine("mysql", conf.GetMySqlConnURL())
	if logger.CheckError(err) {
		return nil, err
	}

	// set engine default config
	//engine.ShowSQL = false
	//engine.ShowDebug = false
	//engine.ShowErr = true
	//engine.ShowWarn = false

	engine.SetTableMapper(core.SnakeMapper{})

	//engine.SetColumnMapper(core.SnakeMapper{})
	engine.SetColumnMapper(core.SameMapper{})

	return mdb.NewXorm(engine), nil
}

// NewXorm
func NewDevXorm(conf *dbconf.DBConf) (*mdb.Xorm, error) {
	logger.Tracef("Create new mdb.Xorm, DB: %s ", conf.Database)

	// check args
	if conf == nil {
		return nil, fmt.Errorf("DBConf is nil")
	}

	engine, err := xorm.NewEngine("mysql", conf.GetMySqlConnURL())
	if logger.CheckError(err) {
		return nil, err
	}

	// set engine default config
	//engine.ShowSQL = true
	//engine.ShowDebug = true
	//engine.ShowErr = true
	//engine.ShowWarn = true

	engine.SetTableMapper(core.SnakeMapper{})
	engine.SetColumnMapper(core.SameMapper{})

	return mdb.NewXorm(engine), nil
}
