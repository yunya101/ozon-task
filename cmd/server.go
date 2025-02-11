package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/yunya101/ozon-task/cmd/graph"
	"github.com/yunya101/ozon-task/cmd/route"
	"github.com/yunya101/ozon-task/internal/config"
	data "github.com/yunya101/ozon-task/internal/data"
	inmem "github.com/yunya101/ozon-task/internal/data/inmemory"
	pg "github.com/yunya101/ozon-task/internal/data/postgres"
	"github.com/yunya101/ozon-task/internal/service"
)

type Application struct {
	router *route.Router
}

func main() {

	usePostgres := os.Getenv("USE_POSTGRES")
	postgresDSN := os.Getenv("POSTGRES_DSN")

	if usePostgres == "true" {
		config.UsePostgres = true
		config.PostgressAddr = postgresDSN
	}

	app := NewApp()
	app.router.SetRoutes()

	strPort := fmt.Sprintf(":%v", config.Port)

	config.InfoLog(fmt.Sprintf("starting server on %v port", config.Port))

	http.ListenAndServe(strPort, app.router.GetMux())

}

func NewApp() *Application {

	var postRepo data.PostRepository
	var commentRepo data.CommentRepository
	var userRepo data.UserRepository

	if config.UsePostgres {
		db := ConnectToPostgres()
		postRepo = pg.NewPostRepo(db)
		commentRepo = pg.NewCommentRepo(db)
		userRepo = pg.NewUserRepo(db)
	} else {
		postRepo = inmem.NewPostRepoInMem()
		commentRepo = inmem.NewCommRepoInMem()
		userRepo = inmem.NewUserRepoInMem()
	}

	postService := service.NewPostService(postRepo)
	userService := service.NewUserService(userRepo)
	commentService := service.NewCommService(commentRepo, postRepo)

	resolver := graph.NewResolver(postService, userService, commentService)

	router := route.NewRouter(resolver, userService, postService)

	return &Application{
		router: router,
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

	config.InfoLog(fmt.Sprintf("connecting to postgres db: %s", config.PostgressAddr))

	return db

}
