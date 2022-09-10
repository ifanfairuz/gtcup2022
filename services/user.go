package services

import (
	"github.com/ifanfairuz/gtcup2022/repositories"
	"github.com/ifanfairuz/gtcup2022/repositories/users"
)

type UserService struct {
	DBM *repositories.DatabaseManager
	UserRepo *users.UserRepo
}

func (service *UserService) init() {
	service.UserRepo = service.DBM.GetRepo(&users.UserRepo{}).(*users.UserRepo)
}

func (service *UserService) GetUser(id int) *users.User {
	return service.UserRepo.FindById(id)
}

func NewUserService(dbm *repositories.DatabaseManager) *UserService {
	service := &UserService{DBM: dbm}
	service.init()
	return service
}
