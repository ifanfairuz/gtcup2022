package repositories

import (
	"log"

	"github.com/ifanfairuz/gtcup2022/repositories/match"
	"github.com/ifanfairuz/gtcup2022/repositories/set"
	"github.com/ifanfairuz/gtcup2022/repositories/team"
	"github.com/ifanfairuz/gtcup2022/repositories/users"
)

func (dbm *DatabaseManager) Migrate()  {
	var err error
	
	err = dbm.db.AutoMigrate(&users.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = dbm.db.AutoMigrate(&team.Team{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = dbm.db.AutoMigrate(&match.Match{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = dbm.db.AutoMigrate(&set.Set{})
	if err != nil {
		log.Fatal(err.Error())
	}
}