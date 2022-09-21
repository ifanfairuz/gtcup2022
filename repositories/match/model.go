package match

import (
	"time"

	"github.com/ifanfairuz/gtcup2022/repositories/set"
	"github.com/ifanfairuz/gtcup2022/repositories/team"
	"gorm.io/gorm"
)

type Match struct {
  gorm.Model
  ID  uint                  `gorm:"primaryKey"`
  TeamHomeId uint           `gorm:"index"`
  TeamAwayId uint           `gorm:"index"`
  TeamHome team.Team        `gorm:"foreignKey:ID;references:TeamHomeId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
  TeamAway team.Team        `gorm:"foreignKey:ID;references:TeamAwayId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
  Sets []set.Set            `gorm:"foreignKey:MatchId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
  Type string               `gorm:"index"`
  Group string              `gorm:"index"`
  Round int                 `gorm:"index"`
  Label string
  Winner uint
  Done bool
  Date time.Time            `gorm:"index"`
  Image string
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt  `gorm:"index"`
}