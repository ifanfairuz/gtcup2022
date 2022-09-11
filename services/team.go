package services

import (
	"sync"

	"github.com/ifanfairuz/gtcup2022/repositories"
	"github.com/ifanfairuz/gtcup2022/repositories/team"
)

type TeamService struct {
	DBM *repositories.DatabaseManager
	TeamRepo *team.TeamRepo
}

func (service *TeamService) init() {
	service.TeamRepo = service.DBM.GetRepo(&team.TeamRepo{}).(*team.TeamRepo)
}

func (service *TeamService) GetKlasemenGroup(group string) []team.Team {
	var teams []team.Team
	service.TeamRepo.AllByGroupQuery(group).Find(&teams)
	return teams
}

func (service *TeamService) getKlasemen() map[string][]team.Team {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	
	groups := []string{"A", "B", "C", "D"}
	klasemen := make(map[string][]team.Team)
	for _, group := range groups {
		wg.Add(1)
		go func (group string)  {
			defer wg.Done()
			mutex.Lock()
			klasemen[group] = service.GetKlasemenGroup(group)
			mutex.Unlock()
		}(group)
	}
	wg.Wait()
	return klasemen
}

func (service *TeamService) GetKlasemen() interface{} {
	return struct {
		Klasemen map[string][]team.Team `json:"klasemen"`
	}{
		Klasemen: service.getKlasemen(),
	}
}

func (service *TeamService) GetTeams() interface{} {
	return struct{
		Teams *[]team.Team `json:"teams"`
	} {
		Teams: service.TeamRepo.All(),
	}
}

func NewTeamService(dbm *repositories.DatabaseManager) *TeamService {
	service := &TeamService{DBM: dbm}
	service.init()
	return service
}
