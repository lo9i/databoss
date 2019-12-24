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
