package actions

import (
	"net/http"

	"github.com/ifanfairuz/gtcup2022/server"
	"github.com/labstack/echo/v4"
)

func ShareImage(e echo.Context) error {
	c := e.(*server.AppContext)
	url := c.QueryParam("url")
	return c.Render(http.StatusOK, "share.html", url)
}