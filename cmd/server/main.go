package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/yunya101/ozon-task/cmd/web"
	"github.com/yunya101/ozon-task/internal/config"
	data "github.com/yunya101/ozon-task/internal/data/postgres"
	rdb "github.com/yunya101/ozon-task/internal/data/redis"
	graphstruct "github.com/yunya101/ozon-task/internal/graphQL"
	"github.com/yunya101/ozon-task/internal/service"
	"github.com/yunya101/ozon-task/pkg/cheker"
)

type Application struct {
	controller *web.Controller
	router     *mux.Router
	checker    *cheker.Cheker
}

func main() {

	config.Port = *flag.Int("port", 8080, "Application port")
	config.PostgressAddr = *flag.String("postgres", "host=localhost port=5432 user=postgres password=admin dbname=ozontask sslmode=disable", "Postgres database address")
	config.RedisAddr = *flag.String("redis", config.RedisAddr, "Redis address")

	app := NewApp()
	config.InfoLog("starting application")
	go app.checker.CheckPopularity()
	panic(app.Start())
}

func NewApp() *Application {

	controller := &web.Controller{}
	r := mux.NewRouter()
	controller.SetRouter(r)

	graph := &graphstruct.GraphQlQueries{}

	db := ConnectToPostgres()
	rDB := ConnectToRedis()

	userRepo := &data.UserRepo{}
	userRepo.SetDB(db)

	postRepo := &data.PostRepo{}
	postRepo.SetDB(db)

	commsRepo := &data.CommentRepo{}
	commsRepo.SetDB(db)

	userService := &service.UserService{}
	userService.SetRepo(userRepo)

	postService := &service.PostService{}
	postService.SetRepo(postRepo)
	graph.SetService(postService)

	commsService := &service.CommsService{}
	commsRepo.SetDB(db)
	commsService.SetRepo(commsRepo)
	commsService.SetPostRepo(postRepo)

	repoRedis := &rdb.RedisRepo{}
	repoRedis.SetPostgres(postRepo)
	repoRedis.SetRedis(rDB)
	commsService.SetRedis(repoRedis)
	postService.SetRedis(repoRedis)

	controller.SetUserService(userService)
	controller.SetCommsService(commsService)
	controller.SetPostService(postService)
	controller.SetGraph(graph)

	ch := &cheker.Cheker{}
	ch.SetRedis(repoRedis)

	controller.SetHandles()

	app := &Application{
		controller: controller,
		router:     r,
		checker:    ch,
	}

	return app
}

func (app *Application) Start() error {

	strPort := fmt.Sprintf(":%v", config.Port)
	config.InfoLog(fmt.Sprintf("starting server at %v", config.Port))
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

	config.InfoLog("connecting to postgres")

	return db

}

func ConnectToRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: "",
		DB:       0,
		Protocol: 3,
	})

	config.InfoLog("connecting to redis database")

	return client
}
