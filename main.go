package main

import "github.com/ifanfairuz/gtcup2022/server"

var app server.Server

func init() {
	app = server.CreateServer()
}

func main() {
	app.Run()
}
