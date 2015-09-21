package mdb

import (
	"fmt"
	"strings"
)

func (s Sql) QueryForFloat(sql string, args ...interface{}) (r float64, err error) {
	row := s.QueryRow(sql, args...)
	err = row.Scan(&r)
	if logger.CheckError(err) {
		logger.Warn("Scan float error:", err)
	}
	return
}

func (s Sql) QueryForFloatNoError(sql string, args ...interface{}) float64 {
	r, _ := s.QueryForFloat(sql, args...)
	return r
}

func BuildWhere(where string) string {
	where = strings.TrimSpace(where)
	if where != "" && !strings.HasPrefix(where, "where") {
		where = " where " + where
	}
	return where
}

// SumNoError(stat_sum_org, income, BuFlag = ?, 'jijian')
func (s Sql) StatSumColumn(table string, column string, where string, args ...interface{}) float64 {
	q := fmt.Sprintf("select sum(%s) from %s %s", column, table, BuildWhere(where))
	return s.QueryForFloatNoError(q, args...)
}

func (s Sql) StatMaxColumn(table string, column string, where string, args ...interface{}) float64 {
	q := fmt.Sprintf("select max(%s) from %s %s", column, table, BuildWhere(where))
	return s.QueryForFloatNoError(q, args...)
}

func (s Sql) StatMinColumn(table string, column string, where string, args ...interface{}) float64 {
	q := fmt.Sprintf("select min(%s) from %s %s", column, table, BuildWhere(where))
	return s.QueryForFloatNoError(q, args...)
}

func (s Sql) StatAvgColumn(table string, column string, where string, args ...interface{}) float64 {
	q := fmt.Sprintf("select avg(%s) from %s %s", column, table, BuildWhere(where))
	return s.QueryForFloatNoError(q, args...)
}
