package services

import (
	"strconv"
	"time"

	"github.com/ifanfairuz/gtcup2022/repositories"
	"github.com/ifanfairuz/gtcup2022/repositories/match"
	"github.com/ifanfairuz/gtcup2022/repositories/team"
)

const TYPE_GROUP = "G"
const TYPE_BRACKET = "B"

type MatchService struct {
	DBM *repositories.DatabaseManager
	MatchRepo *match.MatchRepo
	TeamRepo *team.TeamRepo
}

func (service *MatchService) init() {
	service.MatchRepo = service.DBM.GetRepo(&match.MatchRepo{}).(*match.MatchRepo);
	service.TeamRepo = service.DBM.GetRepo(&team.TeamRepo{}).(*team.TeamRepo);
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
	var lastMatch, nextMatch match.Match
	var lastMatches, nextMatches []match.Match
	service.MatchRepo.QueryAll().Order("date asc").Order("type desc").Order("round asc").Order("\"group\" asc").Find(&matches)
	service.MatchRepo.QueryAll().Where("done = ?", true).Order("date desc").First(&lastMatch)
	service.MatchRepo.QueryAll().Where("done = ?", false).Order("date asc").First(&nextMatch)

	if lastMatch.ID > 0 {
		service.MatchRepo.QueryAll().Where("date = ?", lastMatch.Date).Find(&lastMatches)
	}
	if nextMatch.ID > 0 {
		service.MatchRepo.QueryAll().Where("date = ?", nextMatch.Date).Find(&nextMatches)
	}

	result := struct {
		Matches []match.Match `json:"matches"`
		NextMatches []match.Match `json:"nextMatches"`
		LastMatches []match.Match `json:"lastMatches"`
	}{
		Matches: matches,
		NextMatches: nextMatches,
		LastMatches: lastMatches,
	}

	return result;
}

func (service *MatchService) GetBracket() interface{} {
	var matches []match.Match
	service.MatchRepo.QueryAll().Where("type = ?", TYPE_BRACKET).Order("date asc").Order("type desc").Order("round asc").Order("\"group\" asc").Find(&matches)

	result := struct {
		Matches []match.Match `json:"matches"`
	}{
		Matches: matches,
	}

	return result;
}

func NewMatchService(dbm *repositories.DatabaseManager) *MatchService {
	service := &MatchService{DBM: dbm}
	service.init()
	return service
}