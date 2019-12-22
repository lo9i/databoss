package services

import (
	"github.com/lo9i/databoss/server/data/repositories"
	"github.com/lo9i/databoss/server/domain"
)

type CandidateServiceImpl struct {
	Repository repositories.CandidateRepository
}

func NewCandidateService(repository repositories.CandidateRepository) domain.CandidateService {
	return &CandidateServiceImpl{Repository: repository}
}

func (s *CandidateServiceImpl) Save(u *domain.Candidate) (*domain.Candidate, error) {

	//var err errore
	//err = s.Repository.Insert(u)
	//if err != nil {
	//	return &domain.Candidate{}, err
	//}
	return u, nil
}

func (s *CandidateServiceImpl) Get(id uint64) *domain.Candidate {
	return s.Repository.Get(id)
}

func (s *CandidateServiceImpl) Find(where ...interface{}) *[]domain.Candidate {
	return s.Repository.Find(where)
}

//func (u *Candidate) Find() (*[]Candidate, error) {
//	var err error
//	var users []Candidate
//	err = db.Debug().Model(&Candidate{}).Limit(100).Find(&users).Error
//	if err != nil {
//		return &[]Candidate{}, err
//	}
//	return &users, err
//}
//
//func (u *Candidate) UpdateACandidate(repository repositories.CandidateRepository, uid uint32) (*Candidate, error) {
//
//	// To hash the password
//	err := u.BeforeSave()
//	if err != nil {
//		log.Fatal(err)
//	}
//	db = repository.Update
//	//.Debug().Model(&Candidate{}).Where("id = ?", uid).Take(&Candidate{}).UpdateColumns(
//	//	map[string]interface{}{
//	//		"password":  u.Password,
//	//		"nickname":  u.Nickname,
//	//		"email":     u.Email,
//	//		"update_at": time.Now(),
//	//	},
//	//)
//	if db.Error != nil {
//		return &Candidate{}, db.Error
//	}
//	// This is the display the updated user
//	err = db.Debug().Model(&Candidate{}).Where("id = ?", uid).Take(&u).Error
//	if err != nil {
//		return &Candidate{}, err
//	}
//	return u, nil
//}
//
//func (u *Candidate) DeleteACandidate(db *gorm.DB, uid uint32) (int64, error) {
//
//	db = db.Debug().Model(&Candidate{}).Where("id = ?", uid).Take(&Candidate{}).Delete(&Candidate{})
//
//	if db.Error != nil {
//		return 0, db.Error
//	}
//	return db.RowsAffected, nil
//}
