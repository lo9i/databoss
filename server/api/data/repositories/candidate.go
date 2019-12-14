package repositories

import (
	"github.com/lo9i/databoss/server/api/data"
	"github.com/lo9i/databoss/server/api/domain"
)

type CandidateRepository interface {
	GetById(id uint64) *domain.Candidate
	Insert(account *domain.Candidate)
}

type CandidateRepositoryImpl struct {
	UnitOfWork data.UnitOfWork
}

func NewCandidateRepository(uow data.UnitOfWork) CandidateRepository {
	return &CandidateRepositoryImpl{ UnitOfWork: uow }
}

func (r CandidateRepositoryImpl) GetById(id uint64) *domain.Candidate {
	var candidate domain.Candidate
	r.UnitOfWork.Database().Find(&candidate, id)
	return &candidate
}

func (r CandidateRepositoryImpl) Insert(candidate *domain.Candidate)  {
	r.UnitOfWork.Database().Create(candidate)
}

