package server

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/ifanfairuz/gtcup2022/repositories/users"
	"github.com/ifanfairuz/gtcup2022/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (server *Server) initSesion()  {
	server.e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
}

func (c *AppContext) GetSession(key string) interface{} {
	sess, _ := session.Get("session", c)
	return sess.Values[key]
}
func (c *AppContext) SetSession(key string, value interface{})  {
	sess, _ := session.Get("session", c)
	sess.Values[key] = value
	sess.Save(c.Request(), c.Response())
}

func (c *AppContext) GetAuth() *users.User {
	sess := c.GetSession("user")
	if sess == nil || sess == 0 {
		return nil
	}

	service := services.NewUserService(c.Server.DBM())
	return service.GetUser(sess.(int))
}
func (c *AppContext) SetAuth(user users.User) {
	c.SetSession("user", user.ID)
}

func AuthMiddleware(redirect string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			c := e.(*AppContext)
			if c.GetAuth() == nil {
				return c.Redirect(http.StatusMovedPermanently, redirect)
			}
			return next(c)
		}
	}
}

func UnauthMiddleware(redirect string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			c := e.(*AppContext)
			if c.GetAuth() == nil {
				return next(c)
			}
			return c.Redirect(http.StatusMovedPermanently, redirect)
		}
	}
}