package mdb

import (
//"strings"
)

// return where ColumnA = ? and CloumnB = ?
func (s Sql) BuildColumnsWhere(columns []string) string {
	return BuildColumnsWhere(columns)
}

// return ?,?,?
func (s Sql) BuildValueHoder(args ...interface{}) string {
	return BuildValueHoder(args...)
}

// Join columns , return "ColumnA, ColumnB"
func (s Sql) BuildSelectColumns(columns []string) string {
	return BuildSelectColumns(columns)
}

func (s Sql) IsValidTableName(table string) bool {
	return IsValidTableName(table)
}

func (s Sql) CheckErrorWithSucceedMsg(err error, format string, args ...interface{}) {
	CheckErrorWithSucceedMsg(err, format, args...)
}

func (s Sql) IsBlank(v interface{}, name string) bool {
	return IsBlank(v, name)
}

func (s Sql) IsZero(v interface{}, name string) bool {
	return IsZero(v, name)
}
