package repositories

import (
	"github.com/lo9i/databoss/server/data"
	"github.com/lo9i/databoss/server/domain"
)

type JobRepository interface {
	Get(id uint64) *domain.Job
	Insert(job *domain.Job)
	Find(where ...interface{}) *[]domain.Job
}

type JobRepositoryImpl struct {
	UnitOfWork data.UnitOfWork
}

func NewJobRepository(uow data.UnitOfWork) JobRepository {
	return &JobRepositoryImpl{UnitOfWork: uow}
}

func (r *JobRepositoryImpl) Get(id uint64) *domain.Job {
	var Job domain.Job
	r.UnitOfWork.Database().Find(&Job, id)
	return &Job
}

func (r *JobRepositoryImpl) Insert(Job *domain.Job) {
	r.UnitOfWork.Database().Create(Job)
}

func (r *JobRepositoryImpl) Find(where ...interface{}) *[]domain.Job {
	var Jobs []domain.Job
	r.UnitOfWork.Database().Find(&Jobs, where...)
	return &Jobs
}
