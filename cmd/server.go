package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

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

	port := flag.Int("port", 8080, "Application port")
	postgressAddr := flag.String("postgres", "host=localhost port=5432 user=postgres password=admin dbname=ozontask sslmode=disable", "Postgres database address")
	usePostgres := flag.Bool("use-postgres", false, "Use postgres as storage")
	flag.Parse()

	config.Port = *port
	config.PostgressAddr = *postgressAddr
	config.UsePostgres = *usePostgres

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
