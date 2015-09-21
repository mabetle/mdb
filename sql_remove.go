package mdb

import (
	"fmt"
)

// Equal to RemoveTableRow()
func (s Sql)RemoveRow(table string, id interface{}){
	s.RemoveTableRow(table, id)
}

// Remove rows by query
func (s Sql)RemoveWhereRows(table string, where string, whereArgs ... interface{}){
	if !IsValidTableName(table){
		return
	}
	where = QulifyWhere(where)
	q:=fmt.Sprintf("delete from %s %s", table, where)
	_, err:=s.Exec(q, whereArgs ... )
	CheckErrorWithSucceedMsg(err,"Remove rows succeed. Table:%s, Where:%s WhereArgs:%v", table, where, whereArgs)
}

// Remove rows by conditions
func (s Sql)RemoveRows(table string, columns []string, args... interface{}){
	if !IsValidTableName(table){
		return
	}
	where :=BuildColumnsWhere(columns)
	q:=fmt.Sprintf("delete from %s %s ", table, where)
	_, err:=s.Exec(q, args ...)
	CheckErrorWithSucceedMsg(err,"Remove rows succeed. Table:%s, Where:%s WhereArgs:%v", table, where, args)
}

