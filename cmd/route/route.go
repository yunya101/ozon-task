package route

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/yunya101/ozon-task/cmd/graph"
)

type Router struct {
	resolver *graph.Resolver
	mux      *mux.Router
}

func NewRouter(r *graph.Resolver) *Router {

	return &Router{
		resolver: r,
		mux:      mux.NewRouter(),
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
}
