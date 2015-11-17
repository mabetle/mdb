package mdb

import (
	"fmt"

	"github.com/mabetle/mcore"
)

// RemoveTableRow removes table row with id.
func (s Sql) RemoveTableRow(table string, id interface{}) error {
	sql := GetRemoveRowSql(table)
	_, err := s.Exec(sql, id)
	CheckErrorWithSucceedMsg(err, "Remove row succeed. Table:%s ID:%v", table, id)
	return err
}

// IsTableHasID checks if table has specific id
func (s Sql) IsTableHasID(table string, id interface{}) bool {
	sql := GetIsHasIDSql(table)
	return s.IsQueryHasRows(sql)
}

// IsTableQueryHasRows checks table query has rows
func (s Sql) IsTableQueryHasRows(table, query string, args ...interface{}) (r bool) {
	sql := "select count(*) from " + table + " "
	sql = BuildWhereQuery(sql, query)
	return s.IsQueryHasRows(sql, args...)
}

// IsTableHasRowsByColumns checks table has
// example IsTableHasRowsByColumns(table, []string{"OrgName","Flag"}, "一局")
// TODO check
func (s Sql) IsTableHasRowsByColumns(table string, columns []string, values ...interface{}) (b bool) {
	if !IsValidTableName(table) {
		return
	}
	where := s.BuildColumnsWhere(columns)
	q := fmt.Sprintf("select count(*) from %s %s", table, where)
	rows, _ := s.QueryForInt(q, values...)
	return rows > 0
}

// ClearTable clear table rows
func (s Sql) ClearTable(table string) error {
	sql := GetClearTableSql(table)
	_, err := s.Exec(sql)
	return err
}

// ClearTables clears tables rows one by one
func (s Sql) ClearTables(tables ...string) error {
	errs := mcore.NewErrors()
	for _, v := range tables {
		e := s.ClearTable(v)
		errs.Record(e)
	}
	return errs.Error()
}

// DropTable drops table
func (s Sql) DropTable(table string) error {
	sql := GetDropTableSql(table)
	_, err := s.Exec(sql)
	return err
}

// DropTables drops tables
func (s Sql) DropTables(tables ...string) error {
	errs := mcore.NewErrors()
	for _, v := range tables {
		e := s.DropTable(v)
		errs.Record(e)
	}
	return errs.Error()
}

// CountTableRows returns table rows num.
func (s Sql) CountTableRows(table string) (int64, error) {
	sql := GetCountRowsSql(table)
	return s.QueryForInt(sql)
}

// IsExistTable checks table exists.
func (s Sql) IsExistTable(table string) bool {
	q := "select count(*) from " + table
	_, err := s.Query(q)
	return logger.CheckError(err)
}

// IsExistColumn checks table column exists.
func (s Sql) IsExistColumn(table, column string) (r bool) {
	columns, err := s.GetTableColumns(table)
	if err != nil {
		return false
	}
	return mcore.String(column).IsInArrayIgnoreCase(columns)
}

// CountTableColumns returns table rows number.
func (s Sql) CountTableColumns(table string) int {
	cols, err := s.GetTableColumns(table)
	if err != nil {
		return -1
	}
	return len(cols)
}

// GetTableColumns returns table column names
func (s Sql) GetTableColumns(table string) ([]string, error) {
	sql := GetCountColumnsSql(table)
	return s.GetQueryColumns(sql)
}

// IsTableHasColumn equals to IsExistColumn
func (s Sql) IsTableHasColumn(table string, column string) bool {
	return s.IsExistColumn(table, column)
}

// GetTableRowsJSONData gets table rows JSON data
func (s Sql) GetTableRowsJSONData(table string) map[string]string {
	sql := GetSelectAllSql(table)
	return s.QueryForJSONData(sql)
}

// GetTableRowsMap
func (s Sql) GetTableRowsMap(table string) ([]map[string]interface{}, []string, error) {
	q := GetSelectAllSql(table)
	return s.QueryForMaps(q)
}
