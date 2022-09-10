package repositories

import (
	"gorm.io/gorm"
)

type Repo interface {
	SetDb(db *gorm.DB)
}

type DatabaseManager struct {
	db *gorm.DB
}

func (dbm *DatabaseManager) SetDb(db *gorm.DB)  {
	dbm.db = db
}

func (dbm *DatabaseManager) GetRepo(repo Repo) Repo {
	repo.SetDb(dbm.db)
	return repo;
}