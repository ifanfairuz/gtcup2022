package actions

import (
	"encoding/json"
	"net/http"

	"github.com/ifanfairuz/gtcup2022/server"
	"github.com/ifanfairuz/gtcup2022/services"
	"github.com/labstack/echo/v4"
)

func Klasemen(e echo.Context) error {
	c := e.(*server.AppContext)
	matchService := services.NewTeamService(c.Server.DBM())
	data := matchService.GetKlasemen()

	s, err := json.Marshal(&data)
	if err != nil {
		return c.Render(http.StatusOK, "klasemen.html", "{}")
	}
	return c.Render(http.StatusOK, "klasemen.html", string(s))
}