package server

import "github.com/labstack/echo/v4"

type Server struct {
	e *echo.Echo
}

func (server *Server) init() {
	server.initRoute()
	server.initTemplate()
}

func (server *Server) start() {
	server.e.Logger.Fatal(server.e.Start(":1323"))
}

func (server *Server) Run() {
	server.init()
	server.start()
}

func CreateServer() Server {
	return Server{e: echo.New()}
}
