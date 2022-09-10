package actions

import (
	"encoding/json"
	"net/http"

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