package match

import (
	"gorm.io/gorm"
)

type MatchRepo struct {
	db *gorm.DB
}

func (repo *MatchRepo) SetDb(db *gorm.DB)  {
	repo.db = db
}
func (repo *MatchRepo) InsertAll(datas *[]Match) *gorm.DB {
	return repo.db.Create(datas)
}
func (repo *MatchRepo) Create(name string, alamat string) *Match {
	model := &Match{}
	repo.db.Create(model)
	return model
}
func (repo *MatchRepo) All() *[]Match {
	var result *[]Match
	repo.db.Model(&Match{}).Preload("TeamHome").Preload("TeamAway").Find(&result)
	return result
}
func (repo *MatchRepo) AllByGroup(group string) *[]Match {
	var result *[]Match
	repo.db.Model(&Match{}).Where("").Find(&result)
	return result
}
func (repo *MatchRepo) Truncate() *gorm.DB {
	return repo.db.Exec("TRUNCATE matches")
}
func (repo *MatchRepo) QueryAll() *gorm.DB {
	return repo.db.Model(&Match{}).Preload("TeamHome").Preload("TeamAway")
}

