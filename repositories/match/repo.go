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
func (repo *MatchRepo) Update(match *Match) {
	repo.db.Save(match)
}
func (repo *MatchRepo) SetImage(matches []Match, image interface{}) {
	var ids []uint
	for _, m := range matches {
		ids = append(ids, m.ID)
	}
	repo.db.Model(&Match{}).Where("id IN ?", ids).Update("image", image)
}
func (repo *MatchRepo) DeleteSetsNotIn(match *Match, notin []uint) error {
	q := repo.db.Where("match_id = ?", match.ID)
	if len(notin) > 0 {
		q = q.Where("id NOT IN ?", notin)
	}
	del := q.Delete(&set.Set{})
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
	db := repo.db.Table("matches").Select("date").Where("done = ?", true).Order("date DESC").Group("date").First(&res)
	if (db.RowsAffected > 0) {
		return res.Date, nil
	}
	return time.Now(), errors.New("no data")
}
func (repo *MatchRepo) GetDateNextMatch(minDate ...time.Time) (time.Time, error) {
	var res struct{ Date time.Time }
	q := repo.db.Table("matches").Select("date").Where("done = ?", false).Order("date ASC").Group("date")
	if len(minDate) > 0 {
		q = q.Where("date > ?", minDate[0])
	}
	db := q.First(&res)
	if (db.RowsAffected > 0) {
		return res.Date, nil
	}
	return time.Now(), errors.New("no data")
}
func (repo *MatchRepo) GetLastMatches() *[]Match {
	var res []Match
	date, err := repo.GetDateLastMatch()
	if err != nil {
		return &[]Match{};
	}
	repo.QueryAll().Where("date = ?", date).Order("round ASC").Find(&res)
	return &res;
}
func (repo *MatchRepo) GetNextMatches() *[]Match {
	var res []Match
	lastdate := []time.Time{}
	d, e := repo.GetDateLastMatch()
	if e == nil {
		lastdate = append(lastdate, d)
	}
	date, err := repo.GetDateNextMatch(lastdate...)
	if err != nil {
		return &[]Match{};
	}
	repo.QueryAll().Where("date = ?", date).Order("round ASC").Find(&res)
	return &res;
}

func (repo *MatchRepo) GetGroupDoneMatchesByTeam(team_id uint) *[]Match {
	var res *[]Match
	repo.QueryAll().Where(
		repo.db.Where("team_home_id = ? or team_away_id = ?", team_id, team_id),
	).Where(
		repo.db.Where("done = ? and type = ?", true, "G"),
	).Find(&res)
	return res;
}
