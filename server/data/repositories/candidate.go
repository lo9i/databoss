package repositories

import (
	"github.com/lo9i/databoss/server/data"
	"github.com/lo9i/databoss/server/domain"
)

type CandidateRepository interface {
	Get(id uint64) *domain.Candidate
	Insert(account *domain.Candidate)
	Find(where ...interface{}) *[]domain.Candidate
}

type CandidateRepositoryImpl struct {
	UnitOfWork data.UnitOfWork
}

func NewCandidateRepository(uow data.UnitOfWork) CandidateRepository {
	return &CandidateRepositoryImpl{UnitOfWork: uow}
}

func (r *CandidateRepositoryImpl) Get(id uint64) *domain.Candidate {
	var candidate domain.Candidate
	r.UnitOfWork.Database().Find(&candidate, id)
	return &candidate
}

func (r *CandidateRepositoryImpl) Insert(candidate *domain.Candidate) {
	r.UnitOfWork.Database().Create(candidate)
}

func (r *CandidateRepositoryImpl) Find(where ...interface{}) *[]domain.Candidate {
	var candidates []domain.Candidate
	r.UnitOfWork.Database().Find(&candidates, where)
	return &candidates
}
