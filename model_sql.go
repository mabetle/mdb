package mdb

import (
	"database/sql"
	"reflect"
)

// CountModelRows
func (s Sql) CountModelRows(model interface{}) (int64, error) {
	return s.CountTableRows(TableName(model))
}

// ClearModel
func (s Sql) ClearModel(model interface{}) error {
	return s.ClearTable(TableName(model))
}

// ClearModels
func (s Sql) ClearModels(models ...interface{}) error {
	var err error
	for _, model := range models {
		e := s.ClearModel(model)
		if logger.CheckError(e) {
			err = e
		}
	}
	return err
}

// DropModelTable
func (s Sql) DropModelTable(model interface{}) error {
	return s.DropTable(TableName(model))
}

// DropModelsTable
func (s Sql) DropModelsTable(models ...interface{}) error {
	var err error
	for _, model := range models {
		e := s.DropModelTable(model)
		if logger.CheckError(e) {
			err = e
		}
	}
	return err
}

// IsModelQueryHasRows
func (s Sql) IsModelQueryHasRows(model interface{}, query string, args ...interface{}) bool {
	table := TableName(model)
	return s.IsTableQueryHasRows(table, query, args...)
}

// IsModelHasRows
func (s Sql) IsModelHasRows(model interface{}) (r bool, err error) {
	var n int64
	n, err = s.CountModelRows(model)
	if n > 0 {
		r = true
	}
	return
}

// ModelQuery
func (s Sql) ModelQuery(modelType interface{}, qs string, args ...interface{}) interface{} {
	q := "select " + GetModelFields(modelType) + " from " + TableName(modelType)

	q = BuildWhereQuery(q, qs)

	rows, err := s.Query(q, args...)
	checkError(err)
	defer rows.Close()

	columns, err := rows.Columns()
	checkError(err)

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	var list []interface{}

	t, err := ToType(modelType)
	checkError(err)

	for rows.Next() {
		rows.Scan(scanArgs...)
		v := reflect.New(t)

		dest := make([]interface{}, len(columns))
		for x := range columns {
			f := v.Elem()
			var dummy sql.RawBytes
			dest[x] = &dummy
			target := f.Addr().Interface()
			dest[x] = target
		}

		list = append(list, dest)
	}
	return list
}
