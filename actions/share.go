package actions

import (
	"net/http"
	"time"

	"github.com/ifanfairuz/gtcup2022/server"
	"github.com/ifanfairuz/gtcup2022/services"
	"github.com/ifanfairuz/gtcup2022/support"
	"github.com/labstack/echo/v4"
)

func ShareImage(e echo.Context) error {
	c := e.(*server.AppContext)
	date, err := time.ParseInLocation("2006-01-02", c.QueryParam("date"), support.JAKARTA_TZ)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error date");
	}

	response := c.Response()
	shareService := services.NewShareService(c.Server.DBM(), response.Writer)
	err = shareService.GenImageOnDate(date);
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error());
	}
	return err;
}