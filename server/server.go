package server

import (
	"log"

	"github.com/ifanfairuz/gtcup2022/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	e *echo.Echo
	env Env
	routes Routes
	dbm *repositories.DatabaseManager
}

func (server *Server) initDB()  {
	dsn := server.env.DB_DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Connection Database", err.Error())
	}
	server.dbm = &repositories.DatabaseManager{}
	server.dbm.SetDb(db)
}

func (server *Server) DBM() *repositories.DatabaseManager {
	return server.dbm
}

func (server *Server) Init() {
	server.env = loadEnv()
	server.initDB()
	server.initRoute()
	server.initTemplate()
}

func (server *Server) start() {
	server.e.Logger.Fatal(server.e.Start(":1323"))
}

func (server *Server) GetEnv() Env {
	return server.env
}

func (server *Server) Run() {
	server.dbm.Migrate()
	server.dbm.Seed()
	server.start()
}

func CreateServer(routes Routes) Server {
	return Server{e: echo.New(), routes: routes}
}
