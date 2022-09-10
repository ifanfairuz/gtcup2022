package actions

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ifanfairuz/gtcup2022/server"
	"github.com/ifanfairuz/gtcup2022/services"
	"github.com/labstack/echo/v4"
)

func Index(e echo.Context) error {
	c := e.(*server.AppContext)
	matchService := services.NewMatchService(c.Server.DBM())
	data := matchService.GetData()

	s, err := json.Marshal(&data)
	if err != nil {
		return c.Render(http.StatusOK, "index.html", "{}")
	}
	return c.Render(http.StatusOK, "index.html", string(s))
}

func Test(e echo.Context) error {
	c := e.(*server.AppContext)
	matchService := services.NewMatchService(c.Server.DBM())
	start := time.Date(2022, 9, 14, 19, 0, 0, 0, time.Local)
	matchService.Regenerate(start, 2)
	return c.JSON(http.StatusOK, matchService.MatchRepo.All())
}
