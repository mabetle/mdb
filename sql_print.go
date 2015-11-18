package mdb

import (
	"fmt"
	"github.com/mabetle/mcore"
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
	datas, columns, err := s.GetTableRowsMap(table)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		PrintMapWithColumns(datas, columns)
	}
}

//PrintTableFriendly print table friendly.
func (s *Sql) PrintTableFriendly(table string) {
	s.PrintQueryFriendly("select * from " + table)
}

//PrintQuery print query.
func (s *Sql) PrintQuery(q string, args ...interface{}) {
	s.PrintQueryFriendly(q, args...)
}

// PrintTableQuery print table query.
func (s *Sql) PrintTableQuery(table string, ql string, args ...interface{}) {
	sql := "select * from " + table + " "
	sql = BuildWhereQuery(sql, ql)
	s.PrintQuery(sql, args...)
}

// PrintTableQueryVertical print table query vertical.
func (s *Sql) PrintTableQueryVertical(table string, ql string, args ...interface{}) {
	q := "select * from " + table + " "
	q = BuildWhereQuery(q, ql)
	s.PrintQueryVertical(q, args...)
}

// PrintQueryInJSONFormat prints query in JSON format.
func (s *Sql) PrintQueryInJSONFormat(sql string, args ...interface{}) {
	datas, columns, err := s.QueryForMaps(sql, args...)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		PrintMapWithColumns(datas, columns)
	}
}

func maxIndexLen(data [][]string, index int) int {
	max := 10
	for _, row := range data {
		cols := len(row)
		// index out of range
		if cols < index {
			break
		}
		rn := mcore.StringWidth(row[index])
		if rn > max {
			max = rn
		}
	}
	return max + 2
}

func maxLen(data [][]string) []int {
	if len(data) < 1 {
		return []int{}
	}
	// first line
	n := len(data[0])
	maxLen := make([]int, n)
	for index := 0; index < n; index++ {
		maxLen[index] = maxIndexLen(data, index)
	}
	return maxLen
}

func maxCol(row []string) int {
	n := 10
	for _, v := range row {
		vn := mcore.StringWidth(v)
		if vn > n {
			n = vn
		}
	}
	return n + 2
}

// PrintArrayFriendly
func PrintArrayFriendly(data [][]string) {
	// no data
	if len(data) == 0 {
		return
	}
	maxLen := maxLen(data)
	//fmt.Printf("%v\n", maxLen)
	for _, row := range data {
		for i, v := range row {
			if i < len(maxLen) {
				v = mcore.GetFixedWidthStringAlignLeft(v, maxLen[i])
			}
			fmt.Printf("%s", v)
		}
		fmt.Printf("\n")
	}
}

func PrintArrayVerticalFriendly(data [][]string) {
	if len(data) == 0 {
		fmt.Printf("no results.\n")
		return
	}
	hd := data[0]
	maxCol := maxCol(hd)
	for k, row := range data {
		if k == 0 {
			// skip head line
			continue
		}
		for i, v := range row {
			colLabel := ""
			if i < len(hd) {
				colLabel = mcore.GetFixedWidthStringAlignRight(hd[i], maxCol)
			}
			fmt.Printf("%s:%s\n", colLabel, v)
		}
		fmt.Printf("\n")
	}
}

// PrintQueryFriendly prints query friendly.
func (s *Sql) PrintQueryFriendly(q string, args ...interface{}) {
	data := s.QueryForArrayWithHeader(q, args...)
	PrintArrayFriendly(data)
}

// PrintQueryVertical prints query vertical.
func (s *Sql) PrintQueryVertical(q string, args ...interface{}) {
	data := s.QueryForArrayWithHeader(q, args...)
	PrintArrayVerticalFriendly(data)
}
