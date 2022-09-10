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
	adminPath := "/qwerty"
	adminGroup := e.Group(adminPath+"/app", server.AuthMiddleware(adminPath+"/auth"))
	authGroup := e.Group(adminPath+"/auth", server.UnauthMiddleware(adminPath+"/app"))
	adminGroup.GET("/", admin.Admin)
	authGroup.GET("/", admin.Admin)
}

func main() {
	app.Run()
}
