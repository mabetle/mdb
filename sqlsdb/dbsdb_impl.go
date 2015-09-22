package sqlsdb

import (
	"fmt"
	"github.com/mabetle/mcore/msdb"
	"github.com/mabetle/mdb"
)

//implement for msdb.SimpleTable

type SqlTable struct {
	*msdb.BaseTable
	*mdb.Sql
	data [][]string
}

// NewSqlTableByQuery
func NewSqlTableByQuery(db *mdb.Sql, sql string, args ...interface{}) *SqlTable {
	logger.Trace("NewSqlTableByQuery()")
	table := new(SqlTable)
	bt := new(msdb.BaseTable)
	cu := new(msdb.Cusor)

	table.Sql = db

	rows, _ := table.GetQueryRows(sql, args...)
	cu.MaxIndex = int(rows)
	colsNames, err := table.GetQueryColumns(sql, args...)
	if err != nil {
		fmt.Println(err)
	}

	bt.Header = colsNames

	table.data = table.QueryForData(sql, args...)

	table.BaseTable = bt
	table.BaseTable.Cusor = cu

	table.StringGetter = table
	return table
}

// NewSqlTable
func NewSqlTable(db *mdb.Sql, table string) *SqlTable {
	sql := "select * from " + table
	return NewSqlTableByQuery(db, sql)
}

// overide BaseTable
func (t *SqlTable) GetString(col int) (value string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	row := t.GetRows() - 1
	value = t.data[row][col]
	return
}

// Random Access
func (t *SqlTable) GetRowColString(row, col int) (value string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	value = t.data[row][col]
	return
}
