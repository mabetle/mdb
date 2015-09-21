package mdb

// PrintModelTable
func (s *Sql)PrintModelTable(model interface{}){
	s.PrintModelTableFriendly(model)
}

// PrintModelTableFriendly
func (s *Sql)PrintModelTableFriendly(model interface{}){
	s.PrintTableFriendly(TableName(model))
}

// PrintModelTableVertical
func (s *Sql)PrintModelTableVertical(model interface{}){
	s.PrintTableVertical(TableName(model))
}

// PrintModels
func (s *Sql)PrintModels(models ... interface{}){
	PrintModels(models ... )
}

// PrintModel
func (s *Sql)PrintModel(model interface{}){
	PrintModel(model)
}

// PrintModelTableQuery
func (s *Sql)PrintModelTableQuery(model interface{}, ql string, args ... interface{}){
	s.PrintTableQuery(TableName(model), ql, args ... )
}


// PrintModelTableQueryVertical
func (s *Sql)PrintModelTableQueryVertical(model interface{}, ql string, args ... interface{}){
	s.PrintTableQueryVertical(TableName(model), ql, args ... )
}


