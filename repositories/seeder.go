package repositories

import (
	"github.com/ifanfairuz/gtcup2022/repositories/team"
	"github.com/ifanfairuz/gtcup2022/repositories/users"
	"gorm.io/gorm"
)

type ModelWithSeeder interface {
	Seed(db *gorm.DB)
}

func (dbm *DatabaseManager) Seed()  {
	models := []ModelWithSeeder{
		&users.User{},
		&team.Team{},
	}

	for _, model := range(models) {
		if !dbm.recordExist(model) {
			model.Seed(dbm.db)
		}
	}

}

func (dbm *DatabaseManager) recordExist(model interface{}) bool  {
	var count int64
	dbm.db.Model(model).Count(&count)
	return count > 0
}