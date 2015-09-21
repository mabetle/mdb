package dbconf

import (
	"github.com/mabetle/mcore/mconf"
	"github.com/mabetle/mcore/mconf/ini"
)

//DBConf
type DBConf struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
	Driver   string
	ConnVar  string
}

// init default values
var (
	host     string = "127.0.0.1"
	user     string = "demo"
	password string = "demo"
	database string = "demo"
	port     string = "3306"
	driver   string = "mysql"
	connVar  string = "charset=utf8&keepalive=1200"
)

const (
	KEY_HOST     = "db.host"
	KEY_USER     = "db.user"
	KEY_PASSWORD = "db.password"
	KEY_DATABASE = "db.database"
	KEY_PORT     = "db.port"
	KEY_DRIVER   = "db.driver"
	KEY_CONN_VAR = "db.connVar"
)

// NewDBConf
func NewDBConf() *DBConf {
	logger.Tracef("NewDBConf() Host:%s DataBase:%s", host, database)
	return &DBConf{
		Host:     host,
		User:     user,
		Password: password,
		Database: database,
		Port:     port,
		Driver:   driver,
		ConnVar:  connVar}
}

// Config is a interface.
func NewDBConfFromConfig(c mconf.Config) *DBConf {
	logger.Tracef("NewDBConfFromConfig()")
	// check args
	if nil == c {
		logger.Error("Build DBConf error, config is nil")
		return nil
	}

	//load values from location
	host = c.GetStringWithDefault(KEY_HOST, host)
	user = c.GetStringWithDefault(KEY_USER, user)
	password = c.GetStringWithDefault(KEY_PASSWORD, password)
	database = c.GetStringWithDefault(KEY_DATABASE, database)
	port = c.GetStringWithDefault(KEY_PORT, port)
	driver = c.GetStringWithDefault(KEY_DRIVER, driver)
	connVar = c.GetStringWithDefault(KEY_CONN_VAR, connVar)
	return NewDBConf()
}

// NewDBConfFromFile
// file in ini format.
func NewDBConfFromFile(location string) *DBConf {
	logger.Trace("NewDBConfFromFile() locaton: ", location)
	loader := ini.NewIniConfig(location)
	c := mconf.NewConfig(loader)
	return NewDBConfFromConfig(c)
}

//GetConnURL
func (c DBConf) GetConnURL() string {
	return "tcp:" + c.Host + ":" + c.Port + "*" + c.Database + "/" + c.User + "/" + c.Password
}

// GetMySqlConnURL
// connURL example: xxx:xxx@tcp(xxxx:3306)/xxx
func (c DBConf) GetMySqlConnURL() string {
	connURL := c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database + "?charset=utf8"
	//logger.Trace("DB Connect Info, Host: ", c.Host, ", Database: ", c.Database, ", User: ", c.User)
	return connURL
}
