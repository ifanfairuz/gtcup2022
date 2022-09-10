package admin

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ifanfairuz/gtcup2022/server"
	"github.com/ifanfairuz/gtcup2022/services"
	"github.com/labstack/echo/v4"
)

func Admin(e echo.Context) error {
	c := e.(*server.AppContext)
	teamService := services.NewTeamService(c.Server.DBM())
	data := teamService.GetTeams()

	s, err := json.Marshal(&data)
	if err != nil {
		return c.Render(http.StatusOK, "admin.html", "{}")
	}
	return c.Render(http.StatusOK, "admin.html", string(s))
}
func UpdateTeam(e echo.Context) error {
	c := e.(*server.AppContext)
	teamService := services.NewTeamService(c.Server.DBM())
	data := *teamService.TeamRepo.All()

	for _, t := range data {
		t.Name = c.FormValue("Name-"+strconv.Itoa(int(t.ID)))
		t.Alamat = c.FormValue("Alamat-"+strconv.Itoa(int(t.ID)))
		t.Group = c.FormValue("Group-"+strconv.Itoa(int(t.ID)))
		teamService.TeamRepo.Update(&t)
	}

	return e.Redirect(http.StatusTemporaryRedirect, "/bla")
}

func Generate(e echo.Context) error {
	c := e.(*server.AppContext)
	d := e.QueryParam("date")
	start := time.Date(2022, 9, 14, 12, 0, 0, 0, time.FixedZone("Asia/Jakarta", int(time.Second * 60 * 60 * 7)))
	if d != "" {
		t := strings.Split(d, "-")
		y,_:= strconv.Atoi(t[2])
		m,_:= strconv.Atoi(t[1])
		d,_:= strconv.Atoi(t[0])
		start = time.Date(y, time.Month(m), d, 12, 0, 0, 0, time.FixedZone("Asia/Jakarta", int(time.Second * 60 * 60 * 7)))
	}
	matchService := services.NewMatchService(c.Server.DBM())
	matchService.Regenerate(start, 2)
	return c.JSON(http.StatusOK, matchService.MatchRepo.All())
}
