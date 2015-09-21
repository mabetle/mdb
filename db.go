package mdb

// db base application extends DBService.
type DBService struct {
	*Sql
	*Xorm
}

func NewDBService(sql *Sql, xorm *Xorm) *DBService {
	return &DBService{
		Sql:  sql,
		Xorm: xorm,
	}
}

func (s *DBService) GetSql() *Sql {
	return s.Sql
}

func (s *DBService) GetXorm() *Xorm {
	return s.Xorm
}
