package mdb

import (
	"github.com/mabetle/mcore"
)

// inspect database.
// most of meta relation functions depends on ROOT user or has select rights
// for database INFORMATION_SCHEMA.
//
// Only for mysql.
type MetaDatabase struct {
	SchemaName string
}

type MetaTable struct {
	SchemaName string
	TableName  string

	// when create a table, can set comment in mysql.
	TableComment string //2048
	Columns      []MetaColumn
}

type MetaColumn struct {
	SchemaName string
	TableName  string
	ColumnName string

	ColumnDefault string // longtext in database
	Nullable      string // yes or no, varchar(3) in database
	DataType      string //	varchar long ...varchar(64) in database
	ColumnType    string // varchar(50), decimal(19, 2) ...
	ColumnKey     string // PRI or null
	ColumnComment string // 1024

	OrdinalPosition    int64 // col seq
	CharacterMaxLength int64 // varchar(500) is 500
	NumericPrecision   int64 // long(19) is 19
	NumericScale       int64 // number or null
	DatetimePrecision  int64 //
}

func (c MetaColumn) IsNullable() (r bool) {
	r = mcore.String(c.Nullable).IsEqualIgnoreCase("yes")
	return
}

func (c MetaColumn) IsPrimaryKey() (r bool) {
	r = mcore.String(c.ColumnKey).IsEqualIgnoreCase("pri")
	return
}

func (c MetaColumn) HasDefault() (r bool) {
	r = (c.ColumnDefault != "null")
	return
}
