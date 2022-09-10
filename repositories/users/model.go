package users

import (
	"time"

	"github.com/ifanfairuz/gtcup2022/support"
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  ID  uint `gorm:"primaryKey"`
  Username string
  Password string
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *User) Seed(db *gorm.DB) {
  var datas []User = []User{}
  password, _ := support.HashPassword("Iwanfalsmania.123#")
  datas = append(datas, User{Username: "ifanfairuz", Password: password})
  db.Create(&datas)
}