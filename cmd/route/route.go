package route

import (
	"encoding/json"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/yunya101/ozon-task/cmd/graph"
	"github.com/yunya101/ozon-task/internal/model"
	"github.com/yunya101/ozon-task/internal/service"
)

type Router struct {
	resolver    *graph.Resolver
	mux         *mux.Router
	userService *service.UserService
	postService *service.PostService
}

func NewRouter(r *graph.Resolver, uServ *service.UserService, pServ *service.PostService) *Router {

	return &Router{
		resolver:    r,
		mux:         mux.NewRouter(),
		userService: uServ,
		postService: pServ,
	}

}

func (r *Router) GetMux() *mux.Router {
	return r.mux
}

func (r *Router) SetRoutes() {

	srv := handler.New(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: r.resolver,
		},
	))

	srv.AddTransport(transport.POST{}) // HTTP POST
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	srv.Use(extension.Introspection{})

	r.mux.Handle("/", playground.Handler("Playground", "/query"))
	r.mux.Handle("/query", srv)

	r.mux.HandleFunc("/user/add", r.addUser)
	r.mux.HandleFunc("/post/add", r.addPost)

}

func (router *Router) addUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	user := &model.User{}

	if err := decoder.Decode(user); err != nil {
		http.Error(w, "Cannot parse json", http.StatusBadRequest)
		return
	}

	if err := router.userService.AddUser(user); err != nil {
		http.Error(w, "Cannot add data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (router *Router) addPost(w http.ResponseWriter, r *http.Request) {

	post := &model.Post{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(post); err != nil {
		http.Error(w, "Cannot parse json", http.StatusBadRequest)
		return
	}

	if err := router.postService.AddPost(post); err != nil {
		http.Error(w, "Cannot add data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
