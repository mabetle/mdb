package mdb

type Query struct{
	Sql *Sql
	sql string
	table string
	where string
	columns string
	order string
}


// NewQuery
func (s *Sql)NewQuery()*Query{
	return NewQuery(s)
}

// NewQuery
func NewQuery(sql *Sql)*Query{
	return &Query{Sql:sql}
}


// Query Table
func (q *Query)Table(table string)(*Query){
	q.table = table
	return q
}

// Columns
func (q *Query)Columns(columns string) (*Query){
	q.columns = columns
	return q
}

// Where
func (q *Query)Where(where string)(*Query){
	q.where = where
	return q
}

// Order
func (q *Query)Order(order string)*Query{
	q.order = order
	return q
}

// SqlStr
func (q *Query)SqlStr(sql string)(*Query){
	q.sql = sql
	return q
}

// Query Run
// TODO
// what to return?
// execute query and return result.
func (q *Query)Run(args ... interface{})(){
	logger.Debug("Query Run:")
	logger.Debug("Sql:",q.sql)
	sql:=q.GetQuerySql()
	logger.Debug("build Sql:", sql)
}

// Query Print
func (q *Query)Print(args ... interface{}){
	q.Sql.PrintQuery(q.GetQuerySql(), args ...)
}

func (q *Query)GetQuerySql()string{
	if q.sql!=""{
		return q.sql
	}

	sql:=""
	sql=sql + "select "

	if "" != q.columns{
		sql = sql + q.columns
	}else{
		sql = sql + " * "
	}

	// must have table name
	sql = sql + " from " + q.table

	if "" !=q.where{
		sql = sql + " where " + q.where
	}

	if "" != q.order {
		sql = sql + " order by " + q.order
	}

	return sql
}


