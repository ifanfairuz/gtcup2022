package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

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

	return e.Redirect(http.StatusFound, "/bla")
}

func AdminMatch(e echo.Context) error {
	c := e.(*server.AppContext)
	service := services.NewMatchService(c.Server.DBM())
	data := service.GetAdminMatch()

	s, err := json.Marshal(&data)
	if err != nil {
		return c.Render(http.StatusOK, "admin_match.html", "{}")
	}
	return c.Render(http.StatusOK, "admin_match.html", string(s))
}

func Generate(e echo.Context) error {
	c := e.(*server.AppContext)
	// d := e.QueryParam("date")
	// start := time.Date(2022, 9, 14, 19, 0, 0, 0, time.FixedZone("Asia/Jakarta", 25200))
	// if d != "" {
	// 	t := strings.Split(d, "-")
	// 	y,_:= strconv.Atoi(t[2])
	// 	m,_:= strconv.Atoi(t[1])
	// 	d,_:= strconv.Atoi(t[0])
	// 	start = time.Date(y, time.Month(m), d, 19, 0, 0, 0, time.FixedZone("Asia/Jakarta", 25200))
	// }
	matchService := services.NewMatchService(c.Server.DBM())
	// matchService.Regenerate(start, 2)
	return c.JSON(http.StatusOK, matchService.MatchRepo.All())
}
