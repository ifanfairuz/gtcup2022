package main

import (
	"net/http"

	"github.com/ifanfairuz/gtcup2022/actions"
	"github.com/ifanfairuz/gtcup2022/actions/admin"
	"github.com/ifanfairuz/gtcup2022/server"
)

var app server.Server

func init() {
	routes := server.Routes{
		server.Route{Method: http.MethodGet, Path: "/", Handler: actions.Index, Middlewares: nil},
		server.Route{Method: http.MethodGet, Path: "/klasemen", Handler: actions.Klasemen, Middlewares: nil},
		server.Route{Method: http.MethodGet, Path: "/bracket", Handler: actions.Bracket, Middlewares: nil},
	}

	app = server.CreateServer(routes)
	app.Init()

	e := app.E()
	mustAuth := server.AuthMiddleware("/bla/auth")
	mustUnauth := server.UnauthMiddleware("/bla")
	e.GET("/bla", admin.Admin, mustAuth)
	e.POST("/bla/update-team", admin.UpdateTeam, mustAuth)
	e.GET("/bla/generate", admin.Generate, mustAuth)
	e.GET("/bla/auth", admin.Login, mustUnauth)
	e.POST("/bla/auth", admin.DoLogin, mustUnauth)
	e.POST("/bla/logout", admin.DoLogout, mustUnauth)
}

func main() {
	app.Run()
}
