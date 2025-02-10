package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	redis "github.com/redis/go-redis/v9"
	"github.com/yunya101/ozon-task/cmd/graph"
	"github.com/yunya101/ozon-task/cmd/route"
	"github.com/yunya101/ozon-task/internal/config"
	data "github.com/yunya101/ozon-task/internal/data/postgres"
	redisRepository "github.com/yunya101/ozon-task/internal/data/redis"
	"github.com/yunya101/ozon-task/internal/service"
	"github.com/yunya101/ozon-task/pkg/cheker"
)

type Application struct {
	checker *cheker.Cheker
	router  *route.Router
}

func main() {

	config.Port = *flag.Int("port", 8080, "Application port")
	config.PostgressAddr = *flag.String("postgres", "host=localhost port=5432 user=postgres password=admin dbname=ozontask sslmode=disable", "Postgres database address")
	config.RedisAddr = *flag.String("redis", config.RedisAddr, "Redis address")
	flag.Parse()

	app := NewApp()
	app.router.SetRoutes()

	strPort := fmt.Sprintf(":%v", config.Port)

	config.InfoLog(fmt.Sprintf("starting server on %v port", config.Port))

	go app.checker.CheckPopularity()

	http.ListenAndServe(strPort, app.router.GetMux())

}

func NewApp() *Application {

	db := ConnectToPostgres()
	redis := ConnectToRedis()

	postRepo := data.NewPostRepo(db)
	commentRepo := data.NewCommentRepo(db)
	userRepo := data.NewUserRepo(db)
	redisRepo := redisRepository.NewRedisRepo(redis, postRepo)

	postService := service.NewPostService(postRepo, redisRepo)
	userService := service.NewUserService(userRepo)
	commentService := service.NewCommService(commentRepo, redisRepo, postRepo)

	resolver := graph.NewResolver(postService, userService, commentService)
	checker := cheker.Cheker{}
	checker.SetRedis(redisRepo)

	router := route.NewRouter(resolver, userService, postService)

	return &Application{
		router:  router,
		checker: &checker,
	}

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
