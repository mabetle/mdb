package mdb

func (s Xorm) Ping() error {
	logger.Debugf("mdb.Xorm.Ping()")
	err := s.Engine.Ping()
	logger.CheckError(err)
	return err
}
