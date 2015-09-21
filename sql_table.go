package mdb

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// table relate funcs
// table name as arg of func

// Sql RemoveRow
// no row id?
// no table?
// db conn error?
func (s Sql) RemoveTableRow(table string, id interface{}) error {
	sql := GetRemoveRowSql(table)
	_, err := s.Exec(sql, id)
	CheckErrorWithSucceedMsg(err, "Remove row succeed. Table:%s ID:%v", table, id)
	return err
}

// if table has specific id
func (s Sql) IsTableHasID(table string, id interface{}) bool {
	sql := GetIsHasIDSql(table)
	return s.IsQueryHasRows(sql)
}

// IsTableQueryHasRows
func (s Sql) IsTableQueryHasRows(table, query string, args ...interface{}) (r bool) {
	sql := "select count(*) from " + table + " "
	sql = BuildWhereQuery(sql, query)
	return s.IsQueryHasRows(sql, args...)
}

// IsTableHasRowsByColumns(table, []string{"OrgName","Flag"}, "一局")
func (s Sql) IsTableHasRowsByColumns(table string, columns []string, values ...interface{}) (b bool) {
	if !IsValidTableName(table) {
		return
	}
	where := s.BuildColumnsWhere(columns)
	q := fmt.Sprintf("select count(*) from %s %s", table, where)
	rows, _ := s.QueryForInt(q, values...)
	return rows > 0
}

// Sql ClearTable
func (s Sql) ClearTable(table string) error {
	sql := GetClearTableSql(table)
	_, err := s.Exec(sql)
	return err
}

func (s Sql) ClearTables(tables ...string) error {
	var err error
	for _, v := range tables {
		e := s.ClearTable(v)
		if e != nil {
			err = e
		}
	}
	return err
}

// Sql DropTable
func (s Sql) DropTable(table string) error {
	sql := GetDropTableSql(table)
	_, err := s.Exec(sql)
	return err
}

// DropTables
func (s Sql) DropTables(tables ...string) error {
	var err error
	for _, v := range tables {
		e := s.DropTable(v)
		if e != nil {
			err = e
		}
	}
	return err
}

// Sql CountTableRows
func (s Sql) CountTableRows(table string) (int64, error) {
	sql := GetCountRowsSql(table)
	return s.QueryForInt(sql)
}

func (s Sql) IsExistTable(table string) (r bool) {
	q := "select count(*) from " + table
	_, err := s.Query(q)
	if err == nil {
		r = true
	}
	return
}

func (s Sql) IsExistColumn(table, column string) (r bool) {
	columns := s.GetTableColumns(table)
	return mcore.String(column).IsInArrayIgnoreCase(columns)
}

// CountTableColumns
func (s Sql) CountTableColumns(table string) int {
	return len(s.GetTableColumns(table))
}

// GetTableColumns
func (s Sql) GetTableColumns(table string) []string {
	sql := GetCountColumnsSql(table)
	r, err := s.GetQueryColumns(sql)
	if err != nil {
		return nil
	}
	return r
}

// GetTableRowsJsonData
func (s Sql) GetTableRowsJsonData(table string) map[string]string {
	sql := GetSelectAllSql(table)
	return s.QueryForJsonData(sql)
}
