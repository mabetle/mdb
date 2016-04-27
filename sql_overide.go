package mdb

import (
	"database/sql"
)

// Exec override sql.DB.Exec(), add some logs, print sql and args, when occur error print it.
func (s Sql) Exec(query string, v ...interface{}) (r sql.Result, err error) {
	s.Log(query, v...)
	r, err = s.DB.Exec(query, v...)
	if logger.CheckError(err) {
		return
	}
	affected, e := r.RowsAffected()
	if !logger.CheckError(e) {
		logger.Debug("Succeed:", affected, " rows affected.")
	}
	return
}

// Query  overide sql.DB Query, add logs
func (s Sql) Query(query string, args ...interface{}) (*sql.Rows, error) {
	s.Log(query, args...)
	r, err := s.DB.Query(query, args...)
	logger.CheckError(err)
	return r, err
}

// QueryRow overide sql.DB QueryRow, add logs
func (s Sql) QueryRow(query string, args ...interface{}) *sql.Row {
	s.Log(query, args...)
	r := s.DB.QueryRow(query, args...)
	//logger.Debug()
	return r
}
