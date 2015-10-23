package mdb

import (
	"fmt"
)

// PrintTable prints table
func (s *Sql) PrintTable(table string) {
	s.PrintTableFriendly(table)
}

// PrintTableVertical prints table vertical
func (s *Sql) PrintTableVertical(table string) {
	s.PrintQueryVertical("select * from " + table)
}

// PrintTableInJSONFormat prints table in JSON format.
func (s *Sql) PrintTableInJSONFormat(table string) {
	fmt.Println("===Begin Print Table: ", table, "====")
	datas := s.GetTableRowsJSONData(table)
	PrintMap(datas)
	fmt.Println("===End.. Print Table: ", table, "====")
}

//PrintTableFriendly print table friendly.
func (s *Sql) PrintTableFriendly(table string) {
	s.PrintQueryFriendly("select * from " + table)
}

//PrintQuery print query.
func (s *Sql) PrintQuery(sql string, args ...interface{}) {
	s.PrintQueryFriendly(sql, args...)
}

// PrintTableQuery print table query.
func (s *Sql) PrintTableQuery(table string, ql string, args ...interface{}) {
	sql := "select * from " + table + " "
	sql = BuildWhereQuery(sql, ql)
	s.PrintQuery(sql, args...)
}

// PrintTableQueryVertical print table query vertical.
func (s *Sql) PrintTableQueryVertical(table string, ql string, args ...interface{}) {
	sql := "select * from " + table + " "
	sql = BuildWhereQuery(sql, ql)
	s.PrintQueryVertical(sql, args...)
}

// PrintQueryInJSONFormat prints query in JSON format.
func (s *Sql) PrintQueryInJSONFormat(sql string, args ...interface{}) {
	fmt.Println("============ Begin Print Query =================")
	fmt.Printf("SQL : %v \n", sql)
	fmt.Printf("Args: %v \n", args)

	datas := s.QueryForJSONData(sql, args...)

	PrintMap(datas)
	fmt.Println("============ End  Print Query =================")
}

// PrintQueryFriendly prints query friendly.
func (s *Sql) PrintQueryFriendly(sql string, args ...interface{}) {
	fmt.Println("============ Begin Print Query =================")
	rows, err := s.Query(sql, args...)
	defer rows.Close()

	// check error for wrong query sql.
	if nil != err {
		fmt.Printf("Error: %v \n", err)
		return
	}

	columns, _ := rows.Columns()

	PrintCloumns(columns)

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		_ = rows.Scan(scanArgs...)
		ncols := len(values) - 1
		for ncol, col := range values {
			if col != nil {
				v := GetString(col)
				if ncol == ncols {
					fmt.Printf("%s\n", v)
				} else {
					fmt.Printf("%s,", v)
				}
			}
		}
	}
	fmt.Println("============ End.. Print Query =================")
}

// PrintQueryVertical prints query vertical.
func (s *Sql) PrintQueryVertical(sql string, args ...interface{}) {
	fmt.Println("============ Begin Print Query =================")

	rows, err := s.Query(sql, args...)
	defer rows.Close()

	// check error for wrong query sql.
	if nil != err {
		fmt.Printf("Error: %v \n", err)
		return
	}

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	rowIndex := 0
	for rows.Next() {
		rowIndex++

		_ = rows.Scan(scanArgs...)

		fmt.Printf("-----------Row: %d ------------------\n", rowIndex)

		for colIndex, col := range values {
			if col != nil {
				colName := columns[colIndex]
				fmt.Printf("%s:", colName)
				colValue := GetString(col)
				fmt.Printf("%v\n", colValue)
			}
		}

		fmt.Println()
	}

	fmt.Println("============ End.. Print Query =================")
}
