package repositories

import (
	"github.com/lo9i/databoss/server/data"
	"github.com/lo9i/databoss/server/domain"
)

type UserRepository interface {
	GetById(id uint64) *domain.User
	FindOneUser(where ...interface{}) *domain.User
	Insert(user *domain.User)
}

type UserRepositoryImpl struct {
	UnitOfWork data.UnitOfWork
}

func NewUserRepository(uow data.UnitOfWork) UserRepository {
	return UserRepositoryImpl{UnitOfWork: uow}
}

func (r UserRepositoryImpl) GetById(id uint64) *domain.User {
	var user domain.User
	r.UnitOfWork.Database().Find(&user, id)
	return &user
}

func (r UserRepositoryImpl) FindOneUser(where ...interface{}) *domain.User {
	var user domain.User
	r.UnitOfWork.Database().Find(&user, where)
	return &user
}

func (r UserRepositoryImpl) Insert(user *domain.User) {
	r.UnitOfWork.Database().Create(user)
}
