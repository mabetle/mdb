package mdb

//SqlString

// GetDropTableSql
func GetDropTableSql(table string) string {
	return "drop table " + table
}

// GetClearTableSql
func GetClearTableSql(table string) string {
	return "delete from " + table
}

// GetRemoveRowSql
func GetRemoveRowSql(table string) string {
	return "delete from " + table + " where id= ? "
}

// GetCountRowsSql
func GetCountRowsSql(table string) string {
	return "select count(*) from " + table
}

// GetIsHasIDSql
func GetIsHasIDSql(table string) string {
	return "select count(*) from " + table + " where id = ?"
}

//GetCountColumnsSql
func GetCountColumnsSql(table string) string {
	return "select * from " + table + " where 1 != 1"
}

//GetSelectAllSql
func GetSelectAllSql(table string) string {
	return "select * from " + table
}
