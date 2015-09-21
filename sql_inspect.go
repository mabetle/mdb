package mdb

func (s Sql) Ping() error {
	err := s.DB.Ping()
	logger.CheckError(err)
	return err
}
