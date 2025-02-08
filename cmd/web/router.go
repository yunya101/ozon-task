package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yunya101/ozon-task/internal/config"
)

type Controller struct {
	router *mux.Router
}

func (c *Controller) SetRouter(r *mux.Router) {
	c.router = r
}

func (c *Controller) SetHandles() {

	c.router.HandleFunc("/")
}

func (c *Controller) getLastest(w http.ResponseWriter, r *http.Request) {

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		config.ErrorLog(err)
		http.Error(w, "Wrong page", http.StatusBadRequest)
		return
	}

}
