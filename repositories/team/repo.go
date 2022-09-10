package team

import (
	"gorm.io/gorm"
)

type TeamRepo struct {
	db *gorm.DB
}

func (repo *TeamRepo) SetDb(db *gorm.DB)  {
	repo.db = db
}
func (repo *TeamRepo) Create(name string, alamat string) *Team {
	model := &Team{ Name: name, Alamat: alamat }
	repo.db.Create(model)
	return model
}
func (repo *TeamRepo) Update(team *Team) {
	repo.db.Save(team)
}
func (repo *TeamRepo) All() *[]Team {
	var result *[]Team
	repo.db.Model(&Team{}).Find(&result)
	return result
}
func (repo *TeamRepo) AllByGroup(group string) *[]Team {
	var result *[]Team
	repo.db.Model(&Team{}).Where("\"group\" = ?", group).Find(&result)
	return result
}
func (repo *TeamRepo) AllByGroupQuery(group string) *gorm.DB {
	return repo.db.Model(&Team{}).Where("\"group\" = ?", group)
}
func (repo *TeamRepo) GetAllTeamPerGroup() map[string][]Team {
	teams := repo.All()
	teamsGroup := make(map[string][]Team)

	for _, t := range *teams {
		if teamsGroup[t.Group] == nil {
		teamsGroup[t.Group] = []Team{}
		}
		teamsGroup[t.Group] = append(teamsGroup[t.Group], t)
	}

	return teamsGroup
}

