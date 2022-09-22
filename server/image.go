package server

import (
	"os"

	"github.com/ifanfairuz/gtcup2022/services"
)

func (server *Server) initImage()  {
	shareService := services.NewShareService(server.dbm, os.Stdout)
	shareService.RegenAllImage()
}