package match

import "github.com/ifanfairuz/gtcup2022/repositories/team"


type GrupKlasemen struct {
	Team team.Team `json:"team"`
	Matches []Match `json:"matches"`
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
	gk.Poin = gk.M * 3
	gk.Total = int64((gk.M * 30 * 5) + (gk.AS * 30) + gk.ASC)
}