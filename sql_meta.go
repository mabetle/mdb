package mdb

import (
	"fmt"
	"github.com/mabetle/mcore"
	"strings"
)

// get all database metas for loop.
// not used often, because user should know which database to process.
func (s Sql) GetMetaDatabases() (r []MetaDatabase) {
	schemas := s.GetSchemas()
	for _, schema := range schemas {
		r = append(r, MetaDatabase{SchemaName: schema})
	}
	return
}

// get all table meta from db. schema equal database.
// used for loop all tables.
func (s Sql) GetMetaTables(schema string) (r []MetaTable) {
	q := "select TABLE_NAME,TABLE_COMMENT from `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA`=?"
	rows, err := s.Query(q, schema)
	if err != nil {
		logger.Error(err)
		return
	}
	for rows.Next() {
		var table string
		var comment string
		err = rows.Scan(&table, &comment)
		if err != nil {
			logger.Error(err)
			continue
		}
		item := MetaTable{
			TableName:    table,
			TableComment: comment,
			//Columns:GetMetaColumns(db, table),
		}
		r = append(r, item)
	}
	return
}

// get one Table meta.
// used for process one table.
func (s Sql) GetMetaTable(db, table string) (m MetaTable) {
	q := "select TABLE_COMMENT from `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA`=? and TABLE_NAME = ? "
	m.TableComment = s.QueryForStringNoError(q, db, table)
	m.Columns = s.GetMetaColumns(db, table)
	return
}

// get table one column meta.
// FIXME too much query, should run more fast.
func (s Sql) GetMetaColumn(db, table, col string) MetaColumn {
	metaCols := []string{
		"COLUMN_NAME",    //1
		"COLUMN_DEFAULT", //2
		"IS_NULLABLE",    //3
		"DATA_TYPE",      //4
		"COLUMN_TYPE",    //5

		"COLUMN_KEY",               //6
		"COLUMN_COMMENT",           //7
		"ORDINAL_POSITION",         //8
		"CHARACTER_MAXIMUM_LENGTH", //9
		"NUMERIC_PRECISION",        //10

		"NUMERIC_SCALE",      //11
		"DATETIME_PRECISION", //12
	}
	var columnDefault string // longtext in database
	var nullable string      // yes or no, varchar(3) in database
	var dataType string      //varchar long ...varchar(64) in database
	var columnType string    // varchar(50), decimal(19, 2) ...

	var columnKey string         //PRI or null
	var columnComment string     // 1024
	var ordinalPosition int64    // col seq
	var characterMaxLength int64 // varchar(500) is 500
	var numericPrecision int64   // long(19) is 19

	var numericScale int64      // number or null
	var datetimePrecision int64 //

	q := "select %s from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`=? and TABLE_NAME= ? and COLUMN_NAME = ? "
	// if error set null
	var err error
	columnDefault, err = s.QueryForString(fmt.Sprintf(q, metaCols[1]), db, table, col)
	if err != nil {
		columnDefault = "null"
	}

	nullable = s.QueryForStringNoError(fmt.Sprintf(q, metaCols[2]), db, table, col)
	dataType = s.QueryForStringNoError(fmt.Sprintf(q, metaCols[3]), db, table, col)
	columnType = s.QueryForStringNoError(fmt.Sprintf(q, metaCols[4]), db, table, col)

	columnKey = s.QueryForStringNoError(fmt.Sprintf(q, metaCols[5]), db, table, col)
	columnComment = s.QueryForStringNoError(fmt.Sprintf(q, metaCols[6]), db, table, col)
	ordinalPosition = s.QueryForIntNoError(fmt.Sprintf(q, metaCols[7]), db, table, col)
	characterMaxLength = s.QueryForIntNoError(fmt.Sprintf(q, metaCols[8]), db, table, col)
	numericPrecision = s.QueryForIntNoError(fmt.Sprintf(q, metaCols[9]), db, table, col)

	numericScale = s.QueryForIntNoError(fmt.Sprintf(q, metaCols[10]), db, table, col)
	datetimePrecision = s.QueryForIntNoError(fmt.Sprintf(q, metaCols[11]), db, table, col)

	item := MetaColumn{
		SchemaName: db,
		TableName:  table,

		ColumnName:    col,
		ColumnDefault: columnDefault,
		Nullable:      nullable,
		DataType:      dataType,
		ColumnType:    columnType,

		ColumnKey:          columnKey,
		ColumnComment:      columnComment,
		OrdinalPosition:    ordinalPosition,
		CharacterMaxLength: characterMaxLength,
		NumericPrecision:   numericPrecision,

		NumericScale:      numericScale,
		DatetimePrecision: datetimePrecision,
	}
	return item
}

// get table all columns meta.
// loop all columns.
func (s Sql) GetMetaColumns(db, table string) (r []MetaColumn) {
	cols := s.GetDbTableColumns(db, table)
	for _, col := range cols {
		item := s.GetMetaColumn(db, table, col)
		r = append(r, item)
	}
	return
}

// not specific table, return all columns match conditions
func (s Sql) QueryMetaColumns(db, where string, args ...interface{}) (r []MetaColumn) {
	q := "select TABLE_NAME, COLUMN_NAME from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`='%s' "
	q = fmt.Sprintf(q, db)
	where = strings.TrimSpace(where)
	where = strings.TrimPrefix(where, "where")
	if where != "" {
		q = q + " and " + where
	}

	rows, err := s.Query(q, args...)
	if err != nil {
		logger.Warn(err)
		return
	}
	for rows.Next() {
		var tableName string
		var colName string
		rows.Scan(&tableName, &colName)
		item := s.GetMetaColumn(db, tableName, colName)
		r = append(r, item)
	}
	return
}

// Generate create table sql form database table.
// used for export exist tables struct.
func (s Sql) GenTableCreateSql(db, table string) string {
	sb := mcore.NewStringBuffer()
	sb.Appendf("-- generate create sql for table : %s\n", table)
	sb.Appendf("create table %s (\n", table)
	cols := s.GetMetaColumns(db, table)
	n := len(cols)
	for k, v := range cols {
		var nullDesc string
		if v.IsNullable() {
			nullDesc = "null"
		} else {
			nullDesc = "not null"
		}
		var priDesc string
		if v.IsPrimaryKey() {
			priDesc = "primary key"
		}

		var end string = ","
		if k == n-1 {
			end = ""
		}

		var defaultDesc string
		if (!v.IsNullable()) && v.HasDefault() {
			switch mcore.String(v.DataType).ToLower() {
			case "char", "varchar", "text", "tinytext", "mediumtext", "longtext":
				defaultDesc = fmt.Sprintf("default '%s'", v.ColumnDefault)
			case "int", "integer", "tinyint", "bigint", "smallint", "mediumint", "float", "double", "real", "decimal", "numeric":
				defaultDesc = fmt.Sprintf("default %s", v.ColumnDefault)
			default:
			}
		}

		// name varchar(50) null default '',
		//
		sb.Appendf("\t%s %s %s %s %s %s\n",
			v.ColumnName,
			v.ColumnType,
			nullDesc,
			defaultDesc,
			priDesc,
			end)
	}
	sb.Append(")\n")
	//print it before return
	mcore.Println(sb)
	return sb.String()
}
