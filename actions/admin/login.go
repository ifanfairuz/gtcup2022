package admin

import (
	"net/http"

	"github.com/ifanfairuz/gtcup2022/server"
	"github.com/ifanfairuz/gtcup2022/services"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", "")
}

func DoLogin(e echo.Context) error {
	username := e.FormValue("username");
	password := e.FormValue("password");
	c := e.(*server.AppContext)
	service := services.NewUserService(c.Server.DBM())
	user := service.Login(username, password)
	if user == nil {
		return e.Redirect(http.StatusTemporaryRedirect, e.Path())
	}
	c.SetAuth(*user)
	return e.Redirect(http.StatusTemporaryRedirect, "/bla")
}

func DoLogout(e echo.Context) error {
	c := e.(*server.AppContext)
	c.SetAuth(nil)
	return e.Redirect(http.StatusTemporaryRedirect, "/bla/auth")

}