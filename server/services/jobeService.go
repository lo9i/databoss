package services

import (
	"github.com/lo9i/databoss/server/data/repositories"
	"github.com/lo9i/databoss/server/domain"
)

type JobServiceImpl struct {
	Repository repositories.JobRepository
}

func NewJobService(repository repositories.JobRepository) domain.JobService {
	return &JobServiceImpl{Repository: repository}
}

func (s *JobServiceImpl) Save(j *domain.Job) *domain.Job {
	s.Repository.Insert(j)
	return j
}

func (s *JobServiceImpl) Get(id uint64) *domain.Job {
	return s.Repository.Get(id)
}

func (s *JobServiceImpl) FindByUserId(id uint64) *[]domain.Job {
	return s.Repository.Find(domain.Job{CandidateId: id})
}

func (s *JobServiceImpl) Find(where ...interface{}) *[]domain.Job {
	return s.Repository.Find(where...)
}

//func (u *Job) UpdateAJob(repository repositories.JobRepository, uid uint32) (*Job, error) {
//
//	// To hash the password
//	err := u.BeforeSave()
//	if err != nil {
//		log.Fatal(err)
//	}
//	db = repository.Update
//	//.Debug().Model(&Job{}).Where("id = ?", uid).Take(&Job{}).UpdateColumns(
//	//	map[string]interface{}{
//	//		"password":  u.Password,
//	//		"nickname":  u.Nickname,
//	//		"email":     u.Email,
//	//		"update_at": time.Now(),
//	//	},
//	//)
//	if db.Error != nil {
//		return &Job{}, db.Error
//	}
//	// This is the display the updated user
//	err = db.Debug().Model(&Job{}).Where("id = ?", uid).Take(&u).Error
//	if err != nil {
//		return &Job{}, err
//	}
//	return u, nil
//}
//
//func (u *Job) DeleteAJob(db *gorm.DB, uid uint32) (int64, error) {
//
//	db = db.Debug().Model(&Job{}).Where("id = ?", uid).Take(&Job{}).Delete(&Job{})
//
//	if db.Error != nil {
//		return 0, db.Error
//	}
//	return db.RowsAffected, nil
//}
