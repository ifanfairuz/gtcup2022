package services

import (
	"net/http"
	"time"

	svg "github.com/ajstarks/svgo"
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

func (service *ShareService) GenImageOnDate(d time.Time)  {
	s := svg.New(service.W)
	s.Start(1000, 1415)
	s.Style("text/css", "@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;800&display=swap');")
	s.Image(0, 0, 1000, 1415, images.GetImageBgUri())
	s.Text(60, 200, "TODAY'S", "font-size:100px;font-family:'Montserrat', sans-serif;font-weight:800")
	s.Text(60, 300, "MATCH", "font-size:120px;font-family:'Montserrat', sans-serif;font-weight:800")
	s.End()
}

func NewShareService(dbm *repositories.DatabaseManager, w http.ResponseWriter) *ShareService {
	service := &ShareService{DBM: dbm, W: w}
	service.init()
	return service
}