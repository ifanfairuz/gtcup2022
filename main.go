package main

import (
	"net/http"

	"github.com/ifanfairuz/gtcup2022/actions"
	"github.com/ifanfairuz/gtcup2022/actions/admin"
	"github.com/ifanfairuz/gtcup2022/server"
)

var app server.Server

func initAdmin(app *server.Server)  {
	e := app.E()
	
	mustUnauth := server.UnauthMiddleware("/bla")
	e.GET("/bla/auth", admin.Login, mustUnauth)
	e.POST("/bla/auth", admin.DoLogin, mustUnauth)
	
	mustAuth := server.AuthMiddleware("/bla/auth")
	e.POST("/bla/logout", admin.DoLogout, mustAuth)
	e.GET("/bla", admin.Admin, mustAuth)
	e.POST("/bla/team/update", admin.UpdateTeam, mustAuth)
	e.GET("/bla/match", admin.AdminMatch, mustAuth)
	e.POST("/bla/match/update", admin.UpdateMatch, mustAuth)
	e.GET("/bla/bracket", admin.AdminMatch, mustAuth)
	e.GET("/bla/genimage", admin.GenImage, mustAuth)
}

func init() {
	routes := server.Routes{
		server.Route{Method: http.MethodGet, Path: "/", Handler: actions.Index, Middlewares: nil},
		server.Route{Method: http.MethodGet, Path: "/klasemen", Handler: actions.Klasemen, Middlewares: nil},
		server.Route{Method: http.MethodGet, Path: "/bracket", Handler: actions.Bracket, Middlewares: nil},
		server.Route{Method: http.MethodGet, Path: "/shareimage", Handler: actions.ShareImage, Middlewares: nil},
	}

	app = server.CreateServer(routes)
}

func main() {
	app.Init()
	initAdmin(&app)
	app.Run()
}
