package mdb

import (
	"fmt"
	"github.com/mabetle/mcore"
)

// SqlLogger Log
func (s Sql) Log(query string, args ...interface{}) {
	if !s.ShowSql {
		return
	}
	msg := fmt.Sprintf("\nQuery:%s", query)
	if len(args) > 0 {
		argStr := mcore.Join(",", args...)
		msg = fmt.Sprintf("%s\n Args:%s", msg, argStr)
	}
	fmt.Println(msg)
}

func (s *Sql) SetShowSql(b bool) {
	s.ShowSql = b
}
