package mdb

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// some method about update.

// update a row column value.
func (s Sql) UpdateTableRowColumn(table string, id interface{}, column string, newValue interface{}) {
	if !IsValidTableName(table) {
		return
	}
	if IsBlank(id, "ID") || IsBlank(column, "Column") {
		return
	}
	q := fmt.Sprintf("update %s set %s = ? where %s <> ? and id = ?", table, column, newValue, newValue)
	_, err := s.Exec(q, id)
	CheckErrorWithSucceedMsg(err, "Update row column value succeed: Table:%s ID:%v Column:%s NewValue:%v", table, id, column, newValue)
}

// Column value should be 0 or 1
func (s Sql) EnableTableRowColumn(table string, id interface{}, column string) {
	s.UpdateTableRowColumn(table, id, column, 1)
}

func (s Sql) DisableTableRowColumn(table string, id interface{}, column string) {
	s.UpdateTableRowColumn(table, id, column, 0)
}

// update column value
func (s Sql) UpdateTableColumn(table string, column string, newValue interface{}, where string, whereArgs ...interface{}) {
	if (!IsValidTableName(table)) || (!IsValidColumnName(column)) {
		return
	}
	where = QulifyWhere(where)
	q := fmt.Sprintf("update %s set %s = ? %s", table, column, where)
	args := mcore.AppendBefore(whereArgs, newValue)
	_, err := s.Exec(q, args...)
	CheckErrorWithSucceedMsg(err, "Update column succeed.Table:%s Column:%s NewValue:%v Where:%s WhereArgs:%v", table, column, newValue, where, whereArgs)
}
