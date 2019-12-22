package services

import (
	"github.com/lo9i/databoss/server/data/repositories"
	"github.com/lo9i/databoss/server/domain"
)

type UserServiceImpl struct {
	Repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) domain.UserService {
	return &UserServiceImpl{Repository: repository}
}

func (s *UserServiceImpl) Save(u *domain.User) (*domain.User, error) {

	//var err error
	//err = s.Repository.Insert(u)
	//if err != nil {
	//	return &domain.User{}, err
	//}
	return u, nil
}

func (s UserServiceImpl) Get(id uint64) *domain.User {
	return s.Repository.GetById(id)
}

func (s UserServiceImpl) FindUserByEmail(email string) *domain.User {
	return s.Repository.FindOneUser("email = ?", email)
}

//func (u *User) Find() (*[]User, error) {
//	var err error
//	var users []User
//	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
//	if err != nil {
//		return &[]User{}, err
//	}
//	return &users, err
//}
//
//func (u *User) UpdateAUser(repository repositories.UserRepository, uid uint32) (*User, error) {
//
//	// To hash the password
//	err := u.BeforeSave()
//	if err != nil {
//		log.Fatal(err)
//	}
//	db = repository.Update
//	//.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
//	//	map[string]interface{}{
//	//		"password":  u.Password,
//	//		"nickname":  u.Nickname,
//	//		"email":     u.Email,
//	//		"update_at": time.Now(),
//	//	},
//	//)
//	if db.Error != nil {
//		return &User{}, db.Error
//	}
//	// This is the display the updated user
//	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
//	if err != nil {
//		return &User{}, err
//	}
//	return u, nil
//}
//
//func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {
//
//	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
//
//	if db.Error != nil {
//		return 0, db.Error
//	}
//	return db.RowsAffected, nil
//}
