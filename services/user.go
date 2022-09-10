package services

import (
	"github.com/ifanfairuz/gtcup2022/repositories"
	"github.com/ifanfairuz/gtcup2022/repositories/users"
	"github.com/ifanfairuz/gtcup2022/support"
)

type UserService struct {
	DBM *repositories.DatabaseManager
	UserRepo *users.UserRepo
}

func (service *UserService) init() {
	service.UserRepo = service.DBM.GetRepo(&users.UserRepo{}).(*users.UserRepo)
}

func (service *UserService) GetUser(id uint) *users.User {
	return service.UserRepo.FindById(id)
}

func (service *UserService) Login(username string, password string) *users.User {
	user := service.UserRepo.FindByUsername(username)
	if user != nil && support.CheckPasswordHash(password, user.Password) {
		return user
	}
	
	return nil
}

func NewUserService(dbm *repositories.DatabaseManager) *UserService {
	service := &UserService{DBM: dbm}
	service.init()
	return service
}
