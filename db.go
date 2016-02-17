package mdb

// DBService holds Sql and Xorm
// db base application extends DBService.
type DBService struct {
	Sql  *Sql
	Xorm *Xorm
}

// NewDBService creates DBService instance.
func NewDBService(sql *Sql, xorm *Xorm) *DBService {
	return &DBService{
		Sql:  sql,
		Xorm: xorm,
	}
}

// GetSql returns Sql
func (s *DBService) GetSql() *Sql {
	return s.Sql
}

// GetXorm returns Xorm
func (s *DBService) GetXorm() *Xorm {
	return s.Xorm
}
