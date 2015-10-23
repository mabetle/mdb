package mdb

import (
	"fmt"
	"github.com/mabetle/mcore"
)

var (
	hasDefaultTextColumn = mcore.StringArray([]string{"varchar", "char", "text", "tinytext", "mediumtext", "longtext"})
	hasDefaultNumColumn  = mcore.StringArray([]string{"int", "integer", "bigint", "smallint", "mediumint", "double", "float", "real", "numeric"})
	hasDefaultColumn     = hasDefaultNumColumn.AppendStringArray(hasDefaultTextColumn)
)

// alter table add primary key to specific column
func (s Sql) AlterAddPrimaryKey(table string, colName string) error {
	sql := "ALTER TABLE `" + table + "` ADD PRIMARY KEY (`" + colName + "`)"
	_, err := s.Exec(sql)
	return err
}

// alter table add primay key to ID column
func (s Sql) AlterAddPrimaryKeyToIDColumn(table string) error {
	return s.AlterAddPrimaryKey(table, "ID")
}

// alter table, rename
func (s Sql) AlterRenameTable(table, newName string) error {
	q := fmt.Sprintf("rename %s table to %s", table, newName)
	_, err := s.Exec(q)
	return err
}

// alter table, drop column
func (s Sql) AlterDropColumn(table, column string) error {
	q := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", table, column)
	_, err := s.Exec(q)
	return err
}

// alter table, redifine column data type
func (s Sql) AlterColumnDataType(table, column, dataType string) error {
	q := fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s %s", table, column, dataType)
	_, err := s.Exec(q)
	return err
}

// alter table, add column
func (s Sql) AlterAddColumn(table, column, dataType string) error {
	q := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", table, column, dataType)
	_, err := s.Exec(q)
	return err
}

// alter table, drop primary key
func (s Sql) AlterDropPrimaryKey(table string) error {
	q := fmt.Sprintf("ALTER TABLE %s DROP PRIMARY KEY", table)
	_, err := s.Exec(q)
	return err
}

// alter table, rename column name
func (s Sql) AlterRenameColumn(table, column, newName string) error {
	var err error
	//TODO rename column name in mysql is really confused me.
	return err
}

// alter table column to decimal
func (s Sql) AlterColumnToDecimal(db, table, column, oldType string) error {
	dt, err := s.GetColumnDataType(db, table, column)
	if logger.CheckError(err) {
		return err
	}

	// data type is equal, dont need to udpate.
	if !mcore.String(dt).IsEqualIgnoreCase(oldType) {
		return nil
	}

	q := fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s decimal(19,2) NOT NULL DEFAULT 0.00", table, column)
	//dataType := "decimal(19,2) default 0"
	//s.AlterColumnDataType(table, column, dataType)
	_, err2 := s.Exec(q)
	return err2
}

// alter table all double column to decimal
func (s Sql) AlterTableToDecimal(db, table, oldType string) error {
	var err error
	cols := s.GetDbTableColumns(db, table)
	for _, col := range cols {
		e := s.AlterColumnToDecimal(db, table, col, oldType)
		if logger.CheckError(e) {
			err = e
		}
	}
	return err
}

// alter db all double column to decimal
func (s Sql) AlterDbToDecimal(db string, oldType string) {
	ts := s.GetTables(db)
	for _, t := range ts {
		s.AlterTableToDecimal(db, t, oldType)
	}
}

func (s Sql) AlterDbCharDefault(db string) {
	ts := s.GetTables(db)
	for _, t := range ts {
		s.AlterTableCharDefault(db, t)
	}
}

func (s Sql) AlterTableCharDefault(db, table string) {
	cols := s.GetDbTableColumns(db, table)
	for _, col := range cols {
		s.AlterColumnCharDefault(db, table, col)
	}
}

// AlterColumnCharDefault
func (s Sql) AlterColumnCharDefault(db, table, column string) {
	sc := mcore.String(column)
	c := s.GetMetaColumn(db, table, column)
	if sc.IsEqualIgnoreCase("id") {
		//skip id column
		return
	}
	if c.HasDefault() {
		// skip for has default column
		return
	}
	sct := mcore.String(c.DataType).ToLower()
	if !(sct.IsIn("char", "varchar", "text", "tinytext", "mediumtext", "longtext")) {
		//skip not char or varchar
		return
	}
	q := fmt.Sprintf("update %s	set %s = '' where %s is null", table, column, column)
	s.Exec(q)

	q = fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s %s NOT NULL DEFAULT ''", table, column, c.ColumnType)
	s.Exec(q)
}

func (s Sql) AlterColumnNumDefault(db, table, column string) {
	c := s.GetMetaColumn(db, table, column)
	s.AlterMetaColumnDefault(c)
}
func (s Sql) AlterDbNumDefault(db string) {
	ts := s.GetTables(db)
	for _, t := range ts {
		s.AlterTableNumDefault(db, t)
	}
}

func (s Sql) AlterTableNumDefault(db, table string) {
	cols := s.GetDbTableColumns(db, table)
	for _, col := range cols {
		s.AlterColumnNumDefault(db, table, col)
	}
}

func (s Sql) AlterMetaColumnDefault(col MetaColumn) {
	column := col.ColumnName
	table := col.TableName
	sc := mcore.String(column)
	if sc.IsEqualIgnoreCase("id") {
		//skip id column
		return
	}

	if !col.IsNullable() {
		//skip for not nullable
		logger.Debug(fmt.Sprintf("is not nulable, skip. column:%s.%s.%s, type: %s", col.SchemaName, col.TableName, col.ColumnName, col.DataType))
	}

	sct := mcore.String(col.DataType).ToLower()

	if sct.IsInArray(hasDefaultNumColumn) {
		//process num column
		q := fmt.Sprintf("update %s	set %s = %s where %s is null", table, column, col.ColumnDefault, column)
		s.Exec(q)

		q = fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s %s NOT NULL DEFAULT %s", table, column, col.ColumnType, col.ColumnDefault)
		s.Exec(q)
	} else if sct.IsInArray(hasDefaultTextColumn) {
		//process for char
		q := fmt.Sprintf("update %s	set %s = '%s' where %s is null", table, column, col.ColumnDefault, column)
		s.Exec(q)

		q = fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s %s NOT NULL DEFAULT '%s'", table, column, col.ColumnType, col.ColumnDefault)
		s.Exec(q)
	} else {
		//skip process
		logger.Debug(fmt.Sprintf("skip column:%s.%s.%s, type: %s", col.SchemaName, col.TableName, col.ColumnName, col.DataType))
	}
}

// set varchar char to ''
// set int bigint to 0
func (s Sql) AlterDbColumnDefault(db string) {
	cols := s.QueryMetaColumns(db, "DATA_TYPE in ("+hasDefaultColumn.RoundJoin("'", ",")+") and IS_NULLABLE = 'YES'")
	s.AlterMetaColumnsDefault(cols)
}

func (s Sql) AlterMetaColumnsDefault(cols []MetaColumn) {
	for _, col := range cols {
		s.AlterMetaColumnDefault(col)
	}
}

func (s Sql) AlterTableColumnDefault(db, table string) {
	cols := s.QueryMetaColumns(db, "DATA_TYPE in ("+hasDefaultColumn.RoundJoin("'", ",")+") and IS_NULLABLE = 'YES' and TABLE_NAME = ? ", table)
	s.AlterMetaColumnsDefault(cols)
}

func (s Sql) AlterColumnDefault(db, table, column string) {
	cols := s.QueryMetaColumns(db, "DATA_TYPE in ("+hasDefaultColumn.RoundJoin("'", ",")+") and IS_NULLABLE = 'YES' and TABLE_NAME = ? and COLUMN_NAME = ? ", table, column)
	s.AlterMetaColumnsDefault(cols)
}

// add unique constraint to column.
func (s Sql) AlterTableAddUniqueCloumn(table, column string) {
	q := fmt.Sprintf("ALTER TABLE %s ADD UNIQUE (%s)", table, column)
	s.Exec(q)
}

//add unique constraint to columns.
func (s Sql) AlterTableAddUniqueWithName(table, name string, columns ...string) {
	q := fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT %s UNIQUE (%s)",
		table,
		name,
		mcore.NewStringArray(columns...).Join(","))
	s.Exec(q)
}

// drop unique constraint need provide the constraint name.
func (s Sql) AlterTableDropUnique(table, constraintName string) {
	q := ""
	switch s.Dialect {
	case "mssql", "oracle":
		q = fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", table, constraintName)
	case "mysql":
		q = fmt.Sprintf("ALTER TABLE %s drop INDEX %s", table, constraintName)
	default:
	}
	if q == "" {
		return
	}
	s.Exec(q)
}
