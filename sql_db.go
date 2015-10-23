package mdb

import (
	"fmt"

	"github.com/mabetle/mcore"
)

// IsHasDatabase check db exists
// TODO only work for mysql
func (s Sql) IsHasDatabase(db string) bool {
	q := "select count(*) from `INFORMATION_SCHEMA`.`SCHEMATA` where SCHEMA_NAME = ? "
	return s.IsQueryHasRows(q, db)
}

// CreateDatabase create db
// TODO only work for mysql
func (s Sql) CreateDatabase(db string) (err error) {
	if s.IsHasDatabase(db) {
		logger.Warn("Database exists: ", db)
		return fmt.Errorf("database %s exsits", db)
	}
	sql := "create database " + db + " default character set utf8 default collate utf8_general_ci"
	_, err = s.Exec(sql)
	return
}

// DropDatabase drop db
func (s Sql) DropDatabase(db string) error {
	sql := "drop database " + db
	_, err := s.Exec(sql)
	return err
}

// IsDbExistTable check db tabke exists.
func (s Sql) IsDbExistTable(db, table string) bool {
	q := "SELECT count(*) from `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA`=? and TABLE_NAME = ? "
	return s.IsQueryHasRows(q, db, table)
}

// IsDbTableExistColumn check db table column exists.
func (s Sql) IsDbTableExistColumn(db, table, column string) bool {
	q := "SELECT count(*) from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`=? and TABLE_NAME= ? and COLUMN_NAME = ?"
	return s.IsQueryHasRows(q, db, table, column)
}

// GetColumnDefault returns table column define
func (s Sql) GetColumnDefault(db, table, column string) (r string, err error) {
	q := "SELECT COLUMN_DEFAULT from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`=? and TABLE_NAME= ? and COLUMN_NAME = ?"
	r, err = s.QueryForString(q, db, table, column)
	return
}

// GetColumnDataType return column datatype, varchar,	int etc.
func (s Sql) GetColumnDataType(db, table, column string) (r string, err error) {
	q := "SELECT DATA_TYPE from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`=? and TABLE_NAME= ? and COLUMN_NAME = ?"
	r, err = s.QueryForString(q, db, table, column)
	return
}

// GetColumnType return column type. varchar(60), decimal(18,2), etc.
func (s Sql) GetColumnType(db, table, column string) (r string, err error) {
	q := "SELECT COLUMN_TYPE from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`=? and TABLE_NAME= ? and COLUMN_NAME = ?"
	r, err = s.QueryForString(q, db, table, column)
	return
}

// IsColumnNullable returns ture or false
func (s Sql) IsColumnNullable(db, table, column string) (r bool, err error) {
	q := "SELECT IS_NULLABLE from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`=? and TABLE_NAME= ? and COLUMN_NAME = ?"
	var t string
	t, err = s.QueryForString(q, db, table, column)
	r = mcore.NewString(t).ToBool()
	return r, err
}

// IsColumnPrimary returns true or false
func (s Sql) IsColumnPrimary(db, table, column string) (r bool, err error) {
	q := "SELECT COLUMN_KEY from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`=? and TABLE_NAME= ? and COLUMN_NAME = ?"
	var t string
	t, err = s.QueryForString(q, db, table, column)
	return t == "PRI", err
}

// GetSchemas return all dbs
func (s Sql) GetSchemas() []string {
	q := "select SCHEMA_NAME from `INFORMATION_SCHEMA`.`SCHEMATA`"
	return s.QueryColumnForArray(q)
}

// GetTables if db not exists, return blank array
func (s Sql) GetTables(db string) []string {
	q := "SELECT TABLE_NAME from `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA`=?"
	return s.QueryColumnForArray(q, db)
}

// GetDbTableColumns return table columns
func (s Sql) GetDbTableColumns(dbName, table string) []string {
	q := "SELECT COLUMN_NAME from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`=? and TABLE_NAME= ?"
	return s.QueryColumnForArray(q, dbName, table)
}

// TableExec format include 1 string place holder
func (s Sql) TableExec(table string, format string) error {
	q := fmt.Sprintf(format, table)
	_, err := s.Exec(q)
	return err
}

// ColumnExec format include 2 string place holder
func (s Sql) ColumnExec(table, column, format string) error {
	q := fmt.Sprintf(format, table, column)
	_, err := s.Exec(q)
	return err
}

// DbTablesExec loop all tables in db, format include 1 string place holder.
func (s Sql) DbTablesExec(db string, format string) error {
	errs := mcore.NewErrors()
	ts := s.GetTables(db)
	for _, t := range ts {
		e := s.TableExec(t, format)
		errs.Record(e)
	}
	return errs.Error()
}

// IsHasTable check table exists.
func (s Sql) IsHasTable(table string) bool {
	q := "select * from %s where 1 != 1"
	q = fmt.Sprintf(q, table)
	_, err := s.Exec(q)
	if err != nil {
		return false
	}
	return true
}
