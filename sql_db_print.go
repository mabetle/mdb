package mdb

// db related Print

// PrintDatabases TODO only work for mysql
func (s Sql)PrintDatabases(){
	sql:="SELECT SCHEMA_NAME,DEFAULT_CHARACTER_SET_NAME,DEFAULT_COLLATION_NAME  from `INFORMATION_SCHEMA`.`SCHEMATA`"
	s.PrintQuery(sql)
}

// PrintDBTables TODO only work for mysql
func (s Sql)PrintDBTables(db string){
	sql:="SELECT TABLE_NAME,ENGINE,TABLE_ROWS,TABLE_COLLATION  from `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA`=? "
	s.PrintQuery(sql, db)
}

// PrintDBTablesDetail TODO only work for mysql
func (s Sql)PrintDBTablesDetail(db string){
	sql:="SELECT *  from `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA`=? "
	s.PrintQueryVertical(sql, db)
}


// PrintTableColumns TODO only work for mysql
func (s Sql)PrintTableColumns(db string, table string){
	sql:="SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`=? and TABLE_NAME= ?"
	s.PrintQuery(sql, db, table)
}

// PrintTableColumnsDetail  TODO only work for mysql
func (s Sql)PrintTableColumnsDetail(db string, table string){
	sql:="SELECT * from `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA`=? and TABLE_NAME= ?"
	s.PrintQueryVertical(sql, db, table)
}


