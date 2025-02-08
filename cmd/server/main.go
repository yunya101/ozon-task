package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/yunya101/ozon-task/cmd/web"
	"github.com/yunya101/ozon-task/internal/config"
)

type Application struct {
	controller *web.Controller
	router     *mux.Router
}

func main() {

	config.Port = *flag.Int("port", 8080, "Application port")
	config.PostgressAddr = *flag.String("postgres", "host=localhost port=5432 user=postgres password=admin dbname=ozon-task sslmode=disable", "Postgres database address")
	config.RedisAddr = *flag.String("redis", "localhost:6321", "Redis address")

	app := NewApp()
	app.controller.SetHandles()
	panic(app.Start())
}

func NewApp() *Application {

	controller := &web.Controller{}
	r := mux.NewRouter()
	controller.SetRouter(r)

	app := &Application{
		controller: controller,
		router:     r,
	}

	return app
}

func (app *Application) Start() error {

	strPort := fmt.Sprintf(":%v", config.Port)
	return http.ListenAndServe(strPort, app.router)
}

func ConnectToPostgres() *sql.DB {
	db, err := sql.Open("postgres", config.PostgressAddr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db

}

func ConnectToRedis() *redis.Client {
	opt, err := redis.ParseURL(config.RedisAddr)

	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	return client
}
