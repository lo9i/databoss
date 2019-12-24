package repositories

import (
	"github.com/lo9i/databoss/server/data"
	"github.com/lo9i/databoss/server/domain"
)

type NosisUserRepository interface {
	GetById(id uint64) *domain.NosisUser
	Insert(NosisUser *domain.NosisUser)
}

type NosisUserRepositoryImpl struct {
	UnitOfWork data.UnitOfWork
}

func NewNosisUserRepository(uow data.UnitOfWork) NosisUserRepository {
	return NosisUserRepositoryImpl{UnitOfWork: uow}
}

func (r NosisUserRepositoryImpl) GetById(id uint64) *domain.NosisUser {
	var NosisUser domain.NosisUser
	r.UnitOfWork.Database().Find(&NosisUser, id)
	return &NosisUser
}

func (r NosisUserRepositoryImpl) Insert(NosisUser *domain.NosisUser) {
	r.UnitOfWork.Database().Create(NosisUser)
}
