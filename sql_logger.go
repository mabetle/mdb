package mdb

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// SqlLogger
type SqlLogger struct {
}

var (
	sqlLogger *SqlLogger
)

// GetSqlLogger
func GetSqlLogger() *SqlLogger {
	if nil == sqlLogger {
		sqlLogger = &SqlLogger{}
	}
	return sqlLogger
}

// SqlLogger Log
func (s SqlLogger) Log(query string, args ...interface{}) {
	msg := fmt.Sprintf("\nQuery:%s", query)
	if len(args) > 0 {
		argStr := mcore.Join(",", args...)
		msg = fmt.Sprintf("%s\n Args:%s", msg, argStr)
	}
	logger.Debug(msg)
}
