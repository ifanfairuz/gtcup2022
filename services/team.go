package services

import (
	"sort"
	"sync"

	"github.com/ifanfairuz/gtcup2022/repositories"
	"github.com/ifanfairuz/gtcup2022/repositories/match"
	"github.com/ifanfairuz/gtcup2022/repositories/team"
)

type GrupKlasemen struct {
	Team team.Team `json:"team"`
	Matches []match.Match `json:"matches"`
	Pos int `json:"pos"`
	P int `json:"P"`
	M int `json:"M"`
	K int `json:"K"`
	SM int `json:"SM"`
	SK int `json:"SK"`
	AS int `json:"AS"`
	SCM int `json:"SCM"`
	SCK int `json:"SCK"`
	ASC int `json:"ASC"`
	Poin int `json:"poin"`
	Total int64 `json:"total"`
}

func (gk *GrupKlasemen) Count() {
	gk.P = len(gk.Matches)
	for _, m := range gk.Matches {
		if m.Winner == gk.Team.ID {
			gk.M++
		} else {
			gk.K++
		}

		for _, s := range m.Sets {
			if m.TeamHomeId == gk.Team.ID {
				gk.SCM += s.Home
				gk.SCK += s.Away
			} else if m.TeamAwayId == gk.Team.ID {
				gk.SCM += s.Away
				gk.SCK += s.Home
			}
			if s.Winner == gk.Team.ID {
				gk.SM++
			} else {
				gk.SK++
			}
		}
	}
	gk.AS = gk.SM - gk.SK
	gk.ASC = gk.SCM - gk.SCK
	gk.Poin = gk.M * 1
	gk.Total = int64((gk.Poin * 4) + (gk.AS * 3) + (gk.ASC * 2))
}

type TeamService struct {
	DBM *repositories.DatabaseManager
	TeamRepo *team.TeamRepo
	MatchRepo *match.MatchRepo
}

func (service *TeamService) init() {
	service.TeamRepo = service.DBM.GetRepo(&team.TeamRepo{}).(*team.TeamRepo)
	service.MatchRepo = service.DBM.GetRepo(&match.MatchRepo{}).(*match.MatchRepo)
}

func (service *TeamService) GetKlasemenGroup(group string) []GrupKlasemen {
	var (
		res []GrupKlasemen
		teams []team.Team
		wg sync.WaitGroup
		m sync.RWMutex
	)
	service.TeamRepo.AllByGroupQuery(group).Find(&teams)
	for _, t := range teams {
		wg.Add(1)
		go func (t team.Team)  {
			defer wg.Done()
			defer m.Unlock()
			m.Lock()
			matches := service.MatchRepo.GetGroupDoneMatchesByTeam(t.ID)
			klasemen := GrupKlasemen{Team: t, Matches: *matches}
			klasemen.Count()
			res = append(res, klasemen)
		}(t)
	}
	wg.Wait()
	sort.Slice(res, func(i, j int) bool {
		return res[i].Total > res[j].Total
	});
	pos := 1
	for i, gk := range res {
		if i > 0 {
			if res[i-1].Total != gk.Total {
				pos = i+1
			}
		}
		gk.Pos = pos
		res[i] = gk
	}
	return res
}

func (service *TeamService) getKlasemen() map[string][]GrupKlasemen {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	
	groups := []string{"A", "B", "C", "D"}
	klasemen := make(map[string][]GrupKlasemen)
	for _, group := range groups {
		wg.Add(1)
		go func (group string)  {
			defer wg.Done()
			defer mutex.Unlock()
			mutex.Lock()
			klasemen[group] = service.GetKlasemenGroup(group)
		}(group)
	}
	wg.Wait()
	return klasemen
}

func (service *TeamService) GetKlasemen() interface{} {
	return struct {
		Klasemen map[string][]GrupKlasemen `json:"klasemen"`
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
