package main

import (
	"net/http"

	"github.com/ifanfairuz/gtcup2022/actions"
	"github.com/ifanfairuz/gtcup2022/server"
)

var app server.Server

func init() {
	routes := server.Routes{
		server.Route{Method: http.MethodGet, Path: "/", Handler: actions.Index, Middlewares: nil},
		server.Route{Method: http.MethodGet, Path: "/klasemen", Handler: actions.Klasemen, Middlewares: nil},
		server.Route{Method: http.MethodGet, Path: "/bracket", Handler: actions.Bracket, Middlewares: nil},
		server.Route{Method: http.MethodGet, Path: "/test", Handler: actions.Test, Middlewares: nil},
	}

	app = server.CreateServer(routes)
	app.Init()
}

func main() {
	app.Run()
}
