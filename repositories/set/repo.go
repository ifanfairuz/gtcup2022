package set

import (
	"gorm.io/gorm"
)

type SetRepo struct {
	db *gorm.DB
}

func (repo *SetRepo) SetDb(db *gorm.DB)  {
	repo.db = db
}
func (repo *SetRepo) InsertAll(datas *[]Set) *gorm.DB {
	return repo.db.Create(datas)
}
func (repo *SetRepo) Update(data *Set) *gorm.DB {
	return repo.db.Save(data)
}