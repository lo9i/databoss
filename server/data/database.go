package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
	"github.com/lo9i/databoss/server/domain"
)

type DatabaseConfiguration struct {
	Driver   string
	User     string
	Password string
	Port     string
	Host     string
	Name     string
}

type Database interface {
	AutoMigrate()
	Begin()
	Commit()
	Find(out interface{}, where ...interface{}) Database
	Get(out interface{}, where ...interface{}) Database
	Create(value interface{}) Database
}

func NewDatabase(configuration DatabaseConfiguration) Database {
	db, err := gorm.Open(configuration.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configuration.User, configuration.Password,
		configuration.Host, configuration.Port,
		configuration.Name))
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	database := &DatabaseImpl{Db: db}
	database.AutoMigrate()
	return database
}

type DatabaseImpl struct {
	Db *gorm.DB
}

func (d *DatabaseImpl) AutoMigrate() {
	d.Db.AutoMigrate(&domain.Candidate{})
	d.Db.AutoMigrate(&domain.NosisUser{})
	d.Db.AutoMigrate(&domain.User{})
	d.Db.AutoMigrate(&domain.Job{})
}

func (d *DatabaseImpl) Begin() {
	d.Db.Begin()
}

func (d *DatabaseImpl) Commit() {
	d.Db.Commit()
}

func (d *DatabaseImpl) Find(out interface{}, where ...interface{}) Database {
	if len(where) > 0 && where[0] != nil {
		d.Db.Find(out, where...)
	} else {
		d.Db.Find(out)
	}
	return d
}

func (d *DatabaseImpl) Get(out interface{}, where ...interface{}) Database {
	d.Db.Take(out, where)
	return d
}

func (d *DatabaseImpl) Create(value interface{}) Database {
	d.Db.Create(value)
	return d
}
