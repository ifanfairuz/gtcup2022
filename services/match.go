package services

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/ifanfairuz/gtcup2022/repositories"
	"github.com/ifanfairuz/gtcup2022/repositories/match"
	"github.com/ifanfairuz/gtcup2022/repositories/set"
	"github.com/ifanfairuz/gtcup2022/repositories/team"
)

const TYPE_GROUP = "G"
const TYPE_BRACKET = "B"

type SetUpdate struct {
	ID uint
	Key int
	Home int
	Away int
	Desc string
}

type MatchService struct {
	DBM *repositories.DatabaseManager
	MatchRepo *match.MatchRepo
	TeamRepo *team.TeamRepo
	SetRepo *set.SetRepo
}

func (service *MatchService) init() {
	service.MatchRepo = service.DBM.GetRepo(&match.MatchRepo{}).(*match.MatchRepo);
	service.TeamRepo = service.DBM.GetRepo(&team.TeamRepo{}).(*team.TeamRepo);
	service.SetRepo = service.DBM.GetRepo(&set.SetRepo{}).(*set.SetRepo);
}

func (service *MatchService) generateMatchGroup(start time.Time, matchPerDay int) time.Time {
	teamsGroup := service.TeamRepo.GetAllTeamPerGroup()
	random := [][][2]int{
		{
			{0, 3},
			{1, 2},
		},
		{
			{3, 1},
			{2, 0},
		},
		{
			{1, 0},
			{3, 2},
		},
	}
	groups := []string {"A","B","C","D"}

	curentDate := start
	interval := 0
	num := 1
	var datas []match.Match = []match.Match{}
	for iii, rr := range random {
		for i, group := range groups {
			for ii, r := range rr {
				datas = append(datas, match.Match{
					Type: TYPE_GROUP,
					Group: group,
					Round: (ii+1) * (iii+1) * (i+1),
					Date: curentDate,
					Label: "Pertandingan #"+ strconv.Itoa(num),
					TeamHomeId: teamsGroup[group][r[0]].ID,
					TeamAwayId: teamsGroup[group][r[1]].ID,
				})
				num++
				interval++
				if interval % matchPerDay == 0 {
					curentDate = curentDate.AddDate(0,0,1)
				}
			}
		}
	}

	service.MatchRepo.InsertAll(&datas)
	return curentDate
}

func (service *MatchService) generateMatchBracket(start time.Time)  {
	curentDate := start
	var datas []match.Match = []match.Match{}
	datas = append(datas, match.Match{
			Type: TYPE_BRACKET,
			Group: "Perempat Final",
			Round: 1,
			Date: curentDate,
			Label: "Perempat Final 1|Peringkat 1 Grup C|Peringkat 2 Grup A",
		}, match.Match{
			Type: TYPE_BRACKET,
			Group: "Perempat Final",
			Round: 1,
			Date: curentDate.AddDate(0,0,1),
			Label: "Perempat Final 2|Peringkat 1 Grup B|Peringkat 2 Grup C",
		}, match.Match{
			Type: TYPE_BRACKET,
			Group: "Perempat Final",
			Round: 1,
			Date: curentDate.AddDate(0,0,2),
			Label: "Perempat Final 3|Peringkat 1 Grup A|Peringkat 2 Grup D",
		}, match.Match{
			Type: TYPE_BRACKET,
			Group: "Perempat Final",
			Round: 1,
			Date: curentDate.AddDate(0,0,3),
			Label: "Perempat Final 4|Peringkat 1 Grup D|Peringkat 2 Grup B",
		})
	curentDate = curentDate.AddDate(0,0,4)
	datas = append(datas, match.Match{
			Type: TYPE_BRACKET,
			Group: "Semi Final",
			Round: 2,
			Date: curentDate.AddDate(0,0,1),
			Label: "Semi Final 1|Pemenang Perempat Final 1|Pemenang Perempat Final 2",
		}, match.Match{
			Type: TYPE_BRACKET,
			Group: "Semi Final",
			Round: 2,
			Date: curentDate.AddDate(0,0,2),
			Label: "Semi Final 2|Pemenang Perempat Final 3|Pemenang Perempat Final 4",
		})
	curentDate = curentDate.AddDate(0,0,3)
	datas = append(datas, match.Match{
			Type: TYPE_BRACKET,
			Group: "Perebutan Juara 3",
			Round: 3,
			Date: curentDate.AddDate(0,0,1),
			Label: "Perebutan Juara 3|Runner Up Semifinal 1|Runner Up Semifinal 2",
		}, match.Match{
			Type: TYPE_BRACKET,
			Group: "Grand Final",
			Round: 4,
			Date: curentDate.AddDate(0,0,2),
			Label: "Grand Final|Pemenang Semifinal 1|Pemenang Semifinal 2",
		})

	service.MatchRepo.InsertAll(&datas)
}

func (service *MatchService) Generate(start time.Time, matchPerDay int)  {
	currentDate := service.generateMatchGroup(start, matchPerDay)
	service.generateMatchBracket(currentDate)
}

func (service *MatchService) Regenerate(start time.Time, matchPerDay int)  {
	service.MatchRepo.Truncate()
	service.Generate(start, matchPerDay)
}

func (service *MatchService) GetData() interface{} {
	var matches []match.Match
	service.MatchRepo.QueryAll().Order("date ASC").Order("type DESC").Order("round ASC").Order("\"group\" ASC").Find(&matches)
	lastMatches := service.MatchRepo.GetLastMatches()
	nextMatches := service.MatchRepo.GetNextMatches()

	result := struct {
		Matches []match.Match `json:"matches"`
		NextMatches []match.Match `json:"nextMatches"`
		LastMatches []match.Match `json:"lastMatches"`
	}{
		Matches: matches,
		LastMatches: *lastMatches,
		NextMatches: *nextMatches,
	}

	return result;
}

func (service *MatchService) GetBracket() interface{} {
	var matches []match.Match
	service.MatchRepo.QueryAll().Where("type = ?", TYPE_BRACKET).Order("date ASC").Order("type DESC").Order("round ASC").Order("\"group\" ASC").Find(&matches)

	result := struct {
		Matches []match.Match `json:"matches"`
	}{
		Matches: matches,
	}

	return result;
}

func (service *MatchService) GetMatches() interface{} {
	return struct {
		Matches []match.Match `json:"matches"`
		Teams []team.Team `json:"teams"`
	}{
		Matches: *service.MatchRepo.All(),
		Teams: *service.TeamRepo.All(),
	}
}

func (service *MatchService) replaceSet(m *match.Match, setsUpdate []SetUpdate) map[string]int {
	var (
		ids []uint
		insertSets []set.Set
	)
	winner := map[string]int{
		"home": 0,
		"away": 0,
	}
	for _, s := range setsUpdate {
		ss := set.Set{
			MatchId: m.ID,
			Key: s.Key,
			Home: s.Home,
			Away: s.Away,
			Desc: s.Desc,
		}
		if m.Done {
			if (ss.Home > ss.Away) {
				ss.Winner = m.TeamHomeId
				winner["home"] += 1
			} else {
				ss.Winner = m.TeamAwayId
				winner["away"] += 1
			}
		} else {
			ss.Winner = 0
		}
		if s.ID > 0 {
			ids = append(ids, s.ID)
			ss.ID = s.ID
			service.SetRepo.Update(&ss)
		} else {
			insertSets = append(insertSets, ss);
		}
	}

	err := service.MatchRepo.DeleteSetsNotIn(m, ids)
	if err != nil {
		log.Fatal("Error delete sets", err.Error())
	}
	if len(insertSets) > 0 {
		service.SetRepo.InsertAll(&insertSets)
	}

	return winner
}

func (service *MatchService) UpdateSets(id uint, date time.Time, done bool, jsonData []byte) error {
	m := service.MatchRepo.FindById(id)
	m.Done = done
	m.Date = date
	
	var sets []SetUpdate

	err := json.Unmarshal(jsonData, &sets)
	if err != nil {
        log.Fatal("Cannot parse JSON");
		return err
    }

	winner := service.replaceSet(m, sets)
	if done {
		if winner["home"] > winner["away"] {
			m.Winner = m.TeamHomeId
		} else if winner["home"] < winner["away"] {
			m.Winner = m.TeamAwayId
		} else {
			m.Winner = 0
		}
	} else {
		m.Winner = 0
	}
	service.MatchRepo.Update(m)
	return nil;
}
func (service *MatchService) UpdateTeam(id uint, home_id uint, away_id uint) error {
	m := service.MatchRepo.FindById(id)
	m.TeamHomeId = home_id
	m.TeamAwayId = away_id
	service.MatchRepo.Update(m)
	json, _ := json.Marshal(m.Sets)
	service.UpdateSets(m.ID, m.Date, m.Done, json)
	return nil;
}

func NewMatchService(dbm *repositories.DatabaseManager) *MatchService {
	service := &MatchService{DBM: dbm}
	service.init()
	return service
}