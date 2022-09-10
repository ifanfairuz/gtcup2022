package server

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_HOST string
	DB_PORT int
	DB_NAME string
	DB_USER string
	DB_PASS string
}

func (env *Env) DB_URI() string  {
	return "postgres://"+ env.DB_USER +":"+ env.DB_PASS +"@"+ env.DB_HOST +":"+ strconv.Itoa(env.DB_PORT) +"/"+ env.DB_NAME
}
func (env *Env) DB_DSN() string  {
	return "host="+ env.DB_HOST +" user="+ env.DB_USER +" password="+ env.DB_PASS +" dbname="+ env.DB_NAME +" port="+ strconv.Itoa(env.DB_PORT) +" sslmode=disable TimeZone=Asia/Jakarta"
}

func loadEnv() Env {
	env, _ := godotenv.Read()

	port, err := strconv.Atoi(env["DB_PORT"]);
	if err != nil {
		log.Fatal("Error conver db port")
	}

	return Env{
		DB_HOST: env["DB_HOST"],
		DB_PORT: port,
		DB_NAME: env["DB_NAME"],
		DB_USER: env["DB_USER"],
		DB_PASS: env["DB_PASS"],
	}
}