package mdb

// Table
type Table struct {
	*Sql
	Name string
}

// NewTable
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

// Table RemoveRow
func (t Table) RemoveRow(id interface{}) error {
	return t.RemoveTableRow(t.Name, id)
}

// IsHasRow
func (t Table) IsHasRow(id interface{}) bool {
	return t.IsTableHasID(t.Name, id)
}

// DropTable
func (t Table) Drop() {
	t.DropTable(t.Name)
}

// ClearTable
func (t Table) Clear() {
	t.ClearTable(t.Name)
}

// CountRows
func (t Table) CountRows() (int64, error) {
	return t.CountTableRows(t.Name)
}

// CountColumns
func (t Table) CountColumns() int {
	return t.CountTableColumns(t.Name)
}

// GetColumns
func (t Table) GetColumns() []string {
	return t.GetTableColumns(t.Name)
}

// GetRowsJsonData
func (t Table) GetRowsJsonData() map[string]string {
	return t.GetTableRowsJsonData(t.Name)
}

// Print
func (t Table) Print() {
	q := "select * from " + t.Name
	t.PrintQuery(q)
}
