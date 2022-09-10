package users

import (
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (repo *UserRepo) SetDb(db *gorm.DB)  {
	repo.db = db
}
func (repo *UserRepo) FindById(id uint) *User {
	var result *User
	repo.db.Model(&User{}).Where("id = ?", id).First(&result)
	return result
}
func (repo *UserRepo) FindByUsername(username string) *User {
	var result *User
	repo.db.Model(&User{}).Where("username = ?", username).First(&result)
	return result
}