package mdb

import (
	"github.com/mabetle/mcell/wxlsx"
	"io"
)

// QueryToExcel
// NoLocal support
func (s Sql)QueryToExcel(out io.Writer,
	include string,
	exclude string,
	q string, args ... interface{}) error{
	
	table:=""
	locale:=""
	enableLocale:=false
	return s.QueryToExcelWithLocale(out,table,include,exclude,locale,enableLocale,q,args...)
}


// QueryToExcelWithLocale
// locale
// enableLocaleHeader
func (s Sql)QueryToExcelWithLocale(out io.Writer, 
	table string,
	include string,
	exclude string,
	locale string,
	enableLocale bool,
	q string, args ... interface{}) error{
	
	rows,err:=s.Query(q, args ...)
	
	if err!=nil{
		return err
	}

	// parse table name from sql
	if table == "" {
		table = ParseSqlTableName(q)
	}

	if table == ""{
		table = "common"
	}

	file, errFile := wxlsx.SqlRowsToExcelWithLocale("", table, rows, include, exclude,locale,enableLocale)
	
	if errFile !=nil{
		return errFile
	}
	
	return file.Write(out)
}
