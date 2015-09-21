package mdb

import (
	"fmt"
	"github.com/mabetle/mcore"
	"github.com/mabetle/mcore/mmodel"
	"reflect"
	"strings"
)

// GetString
func GetString(value interface{}) string {
	if v, ok := value.(string); ok {
		return v
	}

	if v, ok := value.([]byte); ok {
		return string(v)
	}

	return fmt.Sprintf("%v", value)
}

// GetSqlBeginFrom
func GetSqlBeginFrom(sql string) string {
	fromIndex := strings.Index(sql, "from")
	//TODO not include "from"
	s := mcore.SubRight(sql, fromIndex+4)
	return s
}

// TableName
func TableName(model interface{}) string {
	return mmodel.GetModelTableName(model)
}

// GetTableName equals to TableName
func GetTableName(model interface{}) string {
	return TableName(model)
}

// ToType
func ToType(i interface{}) (reflect.Type, error) {
	t := reflect.TypeOf(i)
	// If a Pointer to a type, follow
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Cannot SELECT into this type: %v", reflect.TypeOf(i))
	}
	return t, nil
}

// GetModelFields
func GetModelFields(model interface{}) string {
	m := GetModelFieldsMap(model)
	fields := ""
	index := 0
	for field := range m {
		if 0 == index {
			fields = field
		} else {
			fields = fields + "," + field
		}
		index++
	}
	return fields
}

// GetModelFieldsMap
func GetModelFieldsMap(model interface{}) map[string]string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fs := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		p := t.Field(i)
		if !p.Anonymous {
			field := p.Name
			fs[field] = field
		}
	}
	return fs
}

//checkError
func checkError(err error) {
	if nil == err {
		return
	}
	fmt.Println(err)
}

// PrintModels
func PrintModels(models ...interface{}) {
	for _, model := range models {
		PrintModel(model)
	}
}

//PrintModel
func PrintModel(model interface{}) {
	//FIXME should inpect model and print.
	fmt.Printf("%v\n", model)
}

//
func BuildWhereQuery(sql, where string) string {
	if where != "" {
		where = strings.TrimSpace(where)
		if !strings.HasPrefix(where, "where") {
			where = " where " + where
		}
		sql = sql + where
	}
	return sql
}

func PrintData(data [][]string) {
	for _, row := range data {
		PrintRowData(row)
		fmt.Println()
	}
}

func PrintRowData(row []string) {
	ncols := len(row) - 1
	for icol, col := range row {
		if icol == ncols {
			fmt.Printf("%v", col)
		} else {
			fmt.Printf("%v,", col)
		}
	}
}

// Used for check provide table name if blank.
func IsValidTableName(table string) bool {
	return !IsBlank(table, "Table name")
}

func IsValidColumnName(column string) bool {
	return !IsBlank(column, "Column name")
}

func IsBlank(value interface{}, name string) bool {
	v := strings.TrimSpace(fmt.Sprintf("%v", value))
	if v == "" {
		logger.Errorf("%s shouldnt be blank.", name)
		return true
	}
	return false
}

func IsZero(value interface{}, name string) bool {
	v := strings.TrimSpace(fmt.Sprintf("%v", value))
	if v == "0" || v == "0.00" || v == "0.0" || v == "0.000" {
		logger.Errorf("%s shouldnt be zero.", name)
		return true
	}
	return false
}

func QulifyWhere(where string) string {
	where = strings.TrimSpace(where)
	if !strings.HasPrefix(where, "where") {
		where = " where " + where
	}
	return where
}

// return where Column = ? and cloumnB = ?
func BuildColumnsWhere(columns []string) string {
	r := " where "
	for i, v := range columns {
		if i == 0 {
			r = r + v + " = ? "
			continue
		}
		r = r + " and " + v + " = ? "
	}
	return r
}

// Return ?,?,?
func BuildValueHoder(args ...interface{}) string {
	r := ""
	for k, _ := range args {
		if k == len(args) {
			r = r + "?"
			continue
		}
		r = r + "?,"
	}
	return r
}

// Join columns, return "ColumnA, ColumnB"
func BuildSelectColumns(columns []string) string {
	r := ""
	for k, v := range columns {
		if k == len(columns) {
			r = r + v
			continue
		}
		r = r + v + ","
	}
	return r
}

// print the succeed message.
func CheckErrorWithSucceedMsg(err error, format string, args ...interface{}) {
	if err == nil {
		logger.Infof(format, args...)
	}
}

// ParseSqLTableName gets table name from sql.
// from table
func ParseSqlTableName(sql string) string {
	sqlA := strings.Split(sql, " ")
	// no from key words
	if !mcore.NewString("from").IsInArrayIgnoreCase(sqlA) {
		return ""
	}
	table := ""
	start := false
	for _, v := range sqlA {
		if strings.ToLower(v) == "from" {
			start = true
			continue
		}
		if start && v != "" {
			table = v
		}
	}
	return table
}
