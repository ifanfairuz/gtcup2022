package server

import (
	"net/http"

	"github.com/ifanfairuz/gtcup2022/actions"
	"github.com/labstack/echo/v4"
)

type Route struct {
	method      string
	path        string
	handler     echo.HandlerFunc
	middlewares []echo.MiddlewareFunc
}

var routes []Route = []Route{
	{http.MethodGet, "/", actions.Index, nil},
}

func (server *Server) initRoute() {
	server.e.Static("/assets", "public/assets")
	for _, r := range routes {
		server.e.Add(r.method, r.path, r.handler, r.middlewares...)
	}
}
