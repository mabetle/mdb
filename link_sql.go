package mdb

import (
	"database/sql"
	"fmt"
)

// Link program.
type LinkSql struct{
	*Sql
	table	string
	columns string
	where	string
	order   string
	args	[]interface{}
}

func NewLinkSql(sql *Sql)*LinkSql{
	e:=new(LinkSql)
	e.Sql = sql
	return e
}

// use Sql for entries
func (s *Sql)LinkSql()*LinkSql{
	return NewLinkSql(s)
}

func (s *LinkSql)Table(table string)*LinkSql{
	s.table=table
	return s
}

func (s *LinkSql)From(table string)*LinkSql{
	return s.Table(table)
}

func (s *LinkSql)Columns(columns string)(*LinkSql){
	s.columns = columns
	return s
}

func (s *LinkSql)Select(columns string)*LinkSql{
	return s.Columns(columns)
}

func (s *LinkSql)Where(where string)(*LinkSql){
	s.where = where
	return s
}

func (s *LinkSql)Order(order string)(*LinkSql){
	s.order = order
	return s
}

func (s *LinkSql)Asc(asc string)(*LinkSql){
	s.order = asc
	return s
}

func (s *LinkSql)Desc(desc string)(*LinkSql){
	s.order = desc
	return s
}

func (s *LinkSql)Args(args ... interface{})*LinkSql{
	s.args = args
	return s
}

// =======================================
func (s *LinkSql)DoQuery()(*sql.Rows, error){
	q:= s.BuildQuery()
	return s.Query(q, s.args ... )
}

func (s *LinkSql)DoDelete()(error){
	where := QulifyWhere(s.where)
	q:=fmt.Sprintf("delete from %s %s", s.table, where)
	_, err:=s.Exec(q, s.args ...)
	return err
}

// TODO
func (s *LinkSql)DoExec()(r sql.Result, err error) {
	return
}

func (s *LinkSql)Print(){
	q:=s.BuildQuery()
	s.PrintQuery(q, s.args ...)
}

func (s *LinkSql)BuildQuery()(string){
	sql := "select "
	if (s.columns==""){
		s.columns = "*"
	}
	sql = sql + " " + s.columns + " "
	sql = sql + " from " + s.table
	sql = BuildWhereQuery(sql, s.where)
	if (s.order != ""){
		sql = sql + " " +s.order
	}
	return sql
}




