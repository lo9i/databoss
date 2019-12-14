package data


import (
	"github.com/jinzhu/gorm"
	"github.com/lo9i/databoss/server/api/domain"
)

type Database interface {
	AutoMigrate()
	Begin()
	Commit()
	Find(out interface{}, where ...interface{}) Database
	Create(value interface{}) Database
}

type DatabaseImpl struct {
	Db *gorm.DB
}

func (d *DatabaseImpl) AutoMigrate()  {
	d.Db.AutoMigrate(&domain.Candidate{})
}
func (d *DatabaseImpl) Begin()  {
	d.Db.Begin()
}

func (d *DatabaseImpl) Commit()  {
	d.Db.Commit()
}

func (d *DatabaseImpl) Find(out interface{}, where ...interface{}) Database {
	d.Db.Find(out, where)
	return d
}

func (d *DatabaseImpl) Create(value interface{}) Database {
	d.Db.Create(value)
	return d
}
