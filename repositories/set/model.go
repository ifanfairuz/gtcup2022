package set

import (
	"time"

	"gorm.io/gorm"
)

type Set struct {
  gorm.Model
  ID  uint                  `gorm:"primaryKey"`
  MatchId uint              `gorm:"index"`
  Key int
  Home int
  Away int
  Winner uint               `gorm:"index"`
  Desc string
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt  `gorm:"index"`
}