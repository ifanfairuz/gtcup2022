package services

import (
	"net/http"
	"time"

	"github.com/ifanfairuz/gtcup2022/images"
	"github.com/ifanfairuz/gtcup2022/repositories"
	"github.com/ifanfairuz/gtcup2022/repositories/match"
)

type ShareService struct {
	DBM *repositories.DatabaseManager
	W http.ResponseWriter
	MatchRepo *match.MatchRepo
}

func (service *ShareService) init() {
	service.MatchRepo = service.DBM.GetRepo(&match.MatchRepo{}).(*match.MatchRepo)
}

func (service *ShareService) GenImageOnDate(d time.Time) {
	go func() {
		teamService := NewTeamService(service.DBM)
		var m []match.Match
		service.MatchRepo.QueryAll().Where("TO_CHAR(date, 'YYYY-MM-DD') = ?", d.Format("2006-01-02")).Find(&m)
		if len(m) <= 0 {
			return;
		}
		var k []match.GrupKlasemen
		if m[0].Type == "G" {
			k = teamService.GetKlasemenGroup(m[0].Group);
		}
		
		images.RemoveOldImage(m)
		service.MatchRepo.SetImage(m, nil)
		img, err := images.GenImage(m, k)
		if err == nil {
			service.MatchRepo.SetImage(m, img)
		}
	}()
}

func NewShareService(dbm *repositories.DatabaseManager, w http.ResponseWriter) *ShareService {
	service := &ShareService{DBM: dbm, W: w}
	service.init()
	return service
}