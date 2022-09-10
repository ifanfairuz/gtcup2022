package server

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	PORT string
	DB_HOST string
	DB_PORT int
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_URL string
	DB_SSL string
}

func (env *Env) DB_URI() string  {
	if env.DB_URL == "" {
		return "postgres://"+ env.DB_USER +":"+ env.DB_PASS +"@"+ env.DB_HOST +":"+ strconv.Itoa(env.DB_PORT) +"/"+ env.DB_NAME
	}
	return env.DB_URL
}
func (env *Env) DB_DSN() string  {
	return "host="+ env.DB_HOST +" user="+ env.DB_USER +" password="+ env.DB_PASS +" dbname="+ env.DB_NAME +" port="+ strconv.Itoa(env.DB_PORT) +" sslmode="+ env.DB_SSL +" TimeZone=Asia/Jakarta"
}

func loadEnv() Env {
	env, _ := godotenv.Read()

	if env["DB_HOST"] == "" {
		port, err := strconv.Atoi(os.Getenv("DB_PORT"));
		if err != nil {
			log.Fatal("Error convert db port")
		}
		return Env{
			PORT: os.Getenv("PORT"),
			DB_HOST: os.Getenv("DB_HOST"),
			DB_PORT: port,
			DB_NAME: os.Getenv("DB_NAME"),
			DB_USER: os.Getenv("DB_USER"),
			DB_PASS: os.Getenv("DB_PASS"),
			DB_SSL: os.Getenv("DB_SSL"),
			DB_URL: os.ExpandEnv("DATABASE_URL"),
		}
	}

	port, err := strconv.Atoi(env["DB_PORT"]);
	if err != nil {
		log.Fatal("Error convert db port")
	}

	return Env{
		PORT: env["PORT"],
		DB_HOST: env["DB_HOST"],
		DB_PORT: port,
		DB_NAME: env["DB_NAME"],
		DB_USER: env["DB_USER"],
		DB_PASS: env["DB_PASS"],
		DB_SSL: env["DB_SSL"],
	}
}