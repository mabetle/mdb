package mdb

import (
	"database/sql"
	"fmt"
	"github.com/mabetle/mcore"
)

// Sql wrap sqlDB and provide many functions.
type Sql struct {
	DB         *sql.DB
	SchemaName string
	Dialect    string // defautl to mysql.
	ShowSql    bool
}

// New creates sql
func New() *Sql {
	logger.Trace("Create new mdb.Sql ")
	return new(Sql)
}

// NewSql create a new Sql instance
// TODO add other dialect support
func NewSql(db *sql.DB) *Sql {
	logger.Trace("Create new mdb.Sql ")
	s := new(Sql)
	s.DB = db
	s.Dialect = "mysql"
	return s
}

// Use which schema/database.
func (s *Sql) Use(db *sql.DB, schemaName string) *Sql {
	s.DB = db
	//s.SchemaName = schemaName
	return s
}

// SetDialect set sql dialect.
func (s *Sql) SetDialect(dialect string) *Sql {
	s.Dialect = dialect
	return s
}

// GetSchemaName return schema name store in Sql.
func (s Sql) GetSchemaName() string {
	return s.SchemaName
}

// QueryForInt return query int result.
// Must only one column in sql.
// eg. 1.select count(*) from table
// eg. 2.select Age from table where Id = ?
// if more than one column, return err
// if not a int column return err
func (s Sql) QueryForInt(sql string, args ...interface{}) (r int64, err error) {
	row := s.QueryRow(sql, args...)
	err = row.Scan(&r)
	if logger.CheckError(err) {
		logger.Warn("Scan int error:", err)
	}
	return
}

// QueryForIntNoError returns int result.
func (s Sql) QueryForIntNoError(sql string, args ...interface{}) int64 {
	r, _ := s.QueryForInt(sql, args...)
	return r
}

// QueryForString returns string value of a column
// show only one column in sql, eg: selelct Name from table wehre id = ?
// if more than one clolumn or not found any rows returns error
// not restrict in string column, support all data types
func (s Sql) QueryForString(sql string, args ...interface{}) (r string, err error) {
	var rowValue interface{}
	row := s.QueryRow(sql, args...)
	//defer row.Close()
	err = row.Scan(&rowValue)
	if logger.CheckError(err) {
		logger.Warn("Scan string error:", err)
		return
	}
	r = fmt.Sprintf("%v", rowValue)
	logger.Debugf("Result: %s", r)
	return
}

// QueryForBool returns bool result.
func (s Sql) QueryForBool(sql string, args ...interface{}) bool {
	v := s.QueryForStringNoError(sql, args...)
	return mcore.NewString(v).ToBool()
}

// QueryForStringNoError returns string result.
func (s Sql) QueryForStringNoError(sql string, args ...interface{}) string {
	r, _ := s.QueryForString(sql, args...)
	return r
}

// QueryColumnForArray only one column in sql, mulity rows wrap to array
func (s Sql) QueryColumnForArray(sql string, args ...interface{}) (r []string) {
	rows, err := s.Query(sql, args...)
	// if error return init []string.
	if logger.CheckError(err) {
		return
	}
	// walk the rows.
	for rows.Next() {
		var t string
		err = rows.Scan(&t)
		if logger.CheckError(err) {
			continue
		}
		r = append(r, t)
	}
	return
}

// GetQueryColumns returns columns name in string array.
func (s Sql) GetQueryColumns(sql string, args ...interface{}) ([]string, error) {
	rows, err := s.Query(sql, args...)
	defer rows.Close()
	if logger.CheckError(err) {
		return nil, err
	}
	return rows.Columns()
}

// GetQueryRows returns query rows count
func (s Sql) GetQueryRows(sql string, args ...interface{}) (int64, error) {
	q := "select count(*) from " + GetSqlBeginFrom(sql)
	return s.QueryForInt(q, args...)
}

// IsQueryHasRows returns query rows exists.
// if error, return false
func (s Sql) IsQueryHasRows(sql string, args ...interface{}) bool {
	n, _ := s.GetQueryRows(sql, args...)
	return n > 0
}

// GetTableName returns tablename
func (s Sql) GetTableName(model interface{}) string {
	return TableName(model)
}
