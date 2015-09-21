package mdb

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
	logger.Trace("Query:", query)
	if len(args) > 0 {
		logger.Trace("Args:")
		for i, arg := range args {
			logger.Tracef("Arg %d : %v", i, arg)
		}
	}
}
