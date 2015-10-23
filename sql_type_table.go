package mdb

// Table defines Table
type Table struct {
	*Sql
	Name string
}

// NewTable creates Table
func NewTable(sql *Sql, table string) *Table {
	return &Table{
		Sql:  sql,
		Name: table,
	}
}

// NewTable create a new Table instance. Table extends from Sql
func (s *Sql) NewTable(table string) *Table {
	return NewTable(s, table)
}

// RemoveRow table id row
func (t Table) RemoveRow(id interface{}) error {
	return t.RemoveTableRow(t.Name, id)
}

// IsHasRow checks table id row exists.
func (t Table) IsHasRow(id interface{}) bool {
	return t.IsTableHasID(t.Name, id)
}

// Drop drops table
func (t Table) Drop() {
	t.DropTable(t.Name)
}

// Clear clears table rows.
func (t Table) Clear() {
	t.ClearTable(t.Name)
}

// CountRows returns table rows number.
func (t Table) CountRows() (int64, error) {
	return t.CountTableRows(t.Name)
}

// CountColumns return table column numbers.
func (t Table) CountColumns() int {
	return t.CountTableColumns(t.Name)
}

// GetColumns returns table column names.
func (t Table) GetColumns() ([]string, error) {
	return t.GetTableColumns(t.Name)
}

// GetRowsJSONData returns table rows JSON data.
func (t Table) GetRowsJSONData() map[string]string {
	return t.GetTableRowsJSONData(t.Name)
}

// Print prints table datas.
func (t Table) Print() {
	q := "select * from " + t.Name
	t.PrintQuery(q)
}
