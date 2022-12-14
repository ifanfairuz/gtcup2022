package admin

import (
	"net/http"
	"time"

	"github.com/ifanfairuz/gtcup2022/server"
	"github.com/ifanfairuz/gtcup2022/services"
	"github.com/ifanfairuz/gtcup2022/support"
	"github.com/labstack/echo/v4"
)

func GenImage(e echo.Context) error {
	c := e.(*server.AppContext)
	date, err := time.ParseInLocation("2006-01-02", c.QueryParam("date"), support.JAKARTA_TZ)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error date");
	}

	response := c.Response()
	shareService := services.NewShareService(c.Server.DBM(), response.Writer)
	shareService.GenImageOnDate(date);
	return c.Redirect(http.StatusFound, "/bla/match")
}