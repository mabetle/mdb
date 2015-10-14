package mdb

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/mabetle/mcore/mmodel"
)

// Xorm extends xorm.Engine and mdb.Sql.
// so we can provide more functions.
type Xorm struct {
	*xorm.Engine
}

func NewXorm(e *xorm.Engine) *Xorm {
	//logger.Trace("Create new mdb.Xorm")
	return &Xorm{Engine: e}
}

// Overide Engine.Insert
func (s Xorm) Insert(model interface{}) (int64, error) {
	return s.UuidSave(model)
}

// equal to Insert.
func (s Xorm) Save(model interface{}) (int64, error) {
	return s.Insert(model)
}

// SaveOrUpdate
func (s Xorm) SaveOrUpdate(model interface{}) (int64, error) {
	// id == ""; insert with new uuid.
	id := mmodel.GetModelId(model)
	if id == "" {
		logger.Trace("uuid save")
		return s.UuidSave(model)
	}

	// id!=null && id in db, update
	tableName := mmodel.GetModelTableName(model)
	q := fmt.Sprintf("select Id from %s where Id = ? ", tableName)
	row, _ := s.Engine.Query(q, id)
	if len(row) > 0 {
		logger.Trace("update")
		n, err := s.Engine.Id(id).Update(model)
		logger.CheckError(err)
		return n, err
	}

	// id != null && id not in db, insert with provide id
	logger.Trace("normal insert.")
	n, err := s.Engine.Insert(model)
	logger.CheckError(err)
	return n, err
}

// model must have field Id and Id not null.
func (s Xorm) Update(model interface{}) (int64, error) {
	id := mmodel.GetModelId(model)
	if id == "" {
		err := fmt.Errorf("Cannot update model without Id")
		return 0, err
	}
	n, err := s.Engine.Id(id).Update(model)
	logger.CheckError(err)
	return n, err
}

// Id is uuid value.
// model is a pointer.
func (s Xorm) UuidSave(model interface{}) (int64, error) {
	logger.Trace("UuidSave modell")
	if mmodel.GetModelId(model) == "" {
		mmodel.AddModelUuid(model)
	}
	n, err := s.Engine.Insert(model)
	logger.CheckError(err)
	return n, err
}

// GetById
func (s Xorm) GetById(model interface{}, id interface{}) (bool, error) {
	b, err := s.Engine.Id(id).Get(model)
	logger.CheckError(err)
	return b, err
}

// Migrate models
func (s Xorm) Migrate(models ...interface{}) error {
	err := s.Engine.Sync(models...)
	logger.CheckError(err)
	return err
}

// Drop
func (s Xorm) DropTables(models ...interface{}) error {
	var err error
	for _, model := range models {
		q := fmt.Sprintf("drop %s", GetTableName(model))
		_, e := s.Engine.Exec(q)
		if logger.CheckError(e) {
			err = e
		}
	}
	return err
}

func (s Xorm) DropMigrate(models ...interface{}) error {
	var err error
	e := s.DropTables(models...)
	if logger.CheckError(e) {
		err = e
	}
	e2 := s.Migrate(models...)
	if logger.CheckError(e2) {
		err = e2
	}
	return err
}

// Find wraps xorm Find.
func (s Xorm) Find(beans interface{}, condiBeans ...interface{}) error {
	err := s.Engine.Find(beans, condiBeans...)
	logger.CheckError(err)
	return err
}
