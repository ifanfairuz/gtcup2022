package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AppContext struct {
	echo.Context
	Server *Server
}

type Route struct {
	Method string
	Path string
	Handler echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}
type Routes []Route

func (server *Server) initRoute() {
	server.e.Pre(middleware.HTTPSRedirect())
	server.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &AppContext{c, server}
			return next(cc)
		}
	})

	server.e.Static("/assets", "public/assets")

	for _, r := range server.routes {
		server.e.Add(r.Method, r.Path, r.Handler, r.Middlewares...)
	}
}
