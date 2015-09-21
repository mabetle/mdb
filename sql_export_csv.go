package mdb

import (
	"github.com/mabetle/mcsv"
)

// ExportCsv
func (s Sql) ExportCsv(location string, sql string, args ...interface{}) (err error) {
	data := s.QueryForArray(true, sql, args...)
	err = mcsv.WriteData(location, data)
	return
}
