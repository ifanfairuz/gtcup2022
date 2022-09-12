package match

import (
	"errors"
	"time"

	"github.com/ifanfairuz/gtcup2022/repositories/set"
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
func (repo *MatchRepo) Update(team *Match) {
	repo.db.Save(team)
}
func (repo *MatchRepo) DeleteSetsNotIn(match *Match, notin []uint) error {
	del := repo.db.Model(&set.Set{}).Where("match_id = ? AND id NOT IN ?", match.ID, notin).Delete(&set.Set{})
	return del.Error
}
func (repo *MatchRepo) DeleteAllSets(match *Match) {
	repo.db.Model(&set.Set{}).Where("match_id = ?", match.ID).Delete(&set.Set{})
}
func (repo *MatchRepo) FindById(id uint) *Match {
	var result *Match
	repo.QueryAll().Where("ID = ?", id).First(&result)
	return result
}
func (repo *MatchRepo) All() *[]Match {
	var result *[]Match
	repo.QueryAll().Order("date ASC").Order("Round").Find(&result)
	return result
}
func (repo *MatchRepo) Truncate() *gorm.DB {
	return repo.db.Exec("TRUNCATE matches")
}
func (repo *MatchRepo) QueryAll() *gorm.DB {
	return repo.db.Model(&Match{}).Preload("TeamHome").Preload("TeamAway").Preload("Sets", func(db *gorm.DB) *gorm.DB {
		return db.Order("sets.key ASC")
	})
}
func (repo *MatchRepo) GetDateLastMatch() (time.Time, error) {
	var res struct{ Date time.Time }
	db := repo.db.Table("matches").Select("date").Where("done = ?", true).Order("date ASC").Group("date").First(&res)
	if (db.RowsAffected > 0) {
		return res.Date, nil
	}
	return time.Now(), errors.New("no data")
}
func (repo *MatchRepo) GetDateNextMatch(minDate time.Time) (time.Time, error) {
	var res struct{ Date time.Time }
	db := repo.db.Table("matches").Select("date").Where("done = ? and date > ?", false, minDate).Order("date ASC").Group("date").First(&res)
	if (db.RowsAffected > 0) {
		return res.Date, nil
	}
	return time.Now(), errors.New("no data")
}
func (repo *MatchRepo) GetLastMatches() *[]Match {
	var res *[]Match
	date, err := repo.GetDateLastMatch()
	if err != nil {
		return res;
	}

	repo.QueryAll().Where("date = ?", date).Order("round ASC").Find(&res)
	return res;
}
func (repo *MatchRepo) GetNextMatches() *[]Match {
	var res *[]Match
	lastdate, _ := repo.GetDateLastMatch()
	date, err := repo.GetDateNextMatch(lastdate)
	if err != nil {
		return res;
	}

	repo.QueryAll().Where("date = ?", date).Order("round ASC").Find(&res)
	return res;
}

