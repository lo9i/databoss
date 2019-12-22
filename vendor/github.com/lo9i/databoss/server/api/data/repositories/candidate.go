package repositories

import (
	data2 "github.com/lo9i/databoss/server/data"
	domain2 "github.com/lo9i/databoss/server/domain"
)

type CandidateRepository interface {
	GetById(id uint64) *domain2.Candidate
	Insert(account *domain2.Candidate)
}

type CandidateRepositoryImpl struct {
	UnitOfWork data2.UnitOfWork
}

func NewCandidateRepository(uow data2.UnitOfWork) CandidateRepository {
	return &CandidateRepositoryImpl{ UnitOfWork: uow }
}

func (r CandidateRepositoryImpl) GetById(id uint64) *domain2.Candidate {
	var candidate domain2.Candidate
	r.UnitOfWork.Database().Find(&candidate, id)
	return &candidate
}

func (r CandidateRepositoryImpl) Insert(candidate *domain2.Candidate)  {
	r.UnitOfWork.Database().Create(candidate)
}

