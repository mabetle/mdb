package mdb

import (
	"fmt"
	"github.com/mabetle/mcore"
)

type ConnErr struct {
	mcore.RuntimeException
}

func NewConnErr(host, schema string) ConnErr {
	e := mcore.NewRuntimeExceptionf("Can't connect to database. Host:%s Schema:%s", host, schema)
	return ConnErr{RuntimeException: e}
}

// not a fatal error.
type SqlRunError struct {
	mcore.RuntimeException
	sql string
}

func NewSqlRunError(sql string, err error) SqlRunError {
	e := mcore.NewRuntimeException(err)
	return SqlRunError{RuntimeException: e, sql: sql}
}

// overide default Error.
func (e SqlRunError) Error() string {
	return fmt.Sprintf("SqlRunError: %s \nSql:%s", e.ErrorString(), e.sql)
}
