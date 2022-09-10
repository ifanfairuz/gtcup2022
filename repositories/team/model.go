package team

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

const COUNT_TEAM = 16

type Team struct {
  gorm.Model
  ID  uint `gorm:"primaryKey"`
  Name string
  Alamat string
  Group string
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (team *Team) Seed(db *gorm.DB) {
  var teams [COUNT_TEAM]Team = [COUNT_TEAM]Team{}
  for i := 0; i < COUNT_TEAM; i++ {
    var group string
    if i < 4 {
      group = "A"
    } else if i < 8 {
        group = "B"
    } else if i < 12 {
        group = "C"
    } else {
        group = "D"
    }
    teams[i] = Team{Name: "Team #"+ strconv.Itoa(i+1), Alamat: "", Group: group}
  }
  db.Create(&teams)
}