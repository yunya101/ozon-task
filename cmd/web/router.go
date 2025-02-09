package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yunya101/ozon-task/internal/model"
	"github.com/yunya101/ozon-task/internal/service"
)

type Controller struct {
	router       *mux.Router
	postService  *service.PostService
	commsService *service.CommsService
	userService  *service.UserService
}

func (c *Controller) SetPostService(s *service.PostService) {
	c.postService = s
}

func (c *Controller) SetCommsService(s *service.CommsService) {
	c.commsService = s
}

func (c *Controller) SetUserService(s *service.UserService) {
	c.userService = s
}

func (c *Controller) SetRouter(r *mux.Router) {
	c.router = r
}

func (c *Controller) SetHandles() {

	c.router.HandleFunc("/", c.getLastest).Methods("GET")
	c.router.HandleFunc("/sub", c.getSubsPosts).Methods("GET")
	c.router.HandleFunc("/post", c.createPost).Methods("POST")
	c.router.HandleFunc("/post", c.getPostById).Methods("GET")
	c.router.HandleFunc("/post/sub", c.subscribe).Methods("POST")
	c.router.HandleFunc("/post/comment", c.addComment).Methods("POST")
	c.router.HandleFunc("/user", c.addUser).Methods("POST")

}

func writeJson(data interface{}, w http.ResponseWriter) {

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "Cannot write json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Conteny-Type", "application/json")
	w.Write(js)

}

func (c *Controller) getLastest(w http.ResponseWriter, r *http.Request) {

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Wrong page", http.StatusBadRequest)
		return
	}

	posts, err := c.postService.GetLastestPosts(page)

	if err != nil {
		http.Error(w, "Cannot get data", http.StatusInternalServerError)
		return
	}

	writeJson(posts, w)

}

// Получаем все посты на которые подписан пользователь по его id
func (c *Controller) getSubsPosts(w http.ResponseWriter, r *http.Request) {

	strId := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		http.Error(w, "Wrong id", http.StatusBadRequest)
		return
	}

	posts, err := c.postService.GetSubsPostsByUserId(id)

	if err != nil {
		http.Error(w, "Cannot get data", http.StatusInternalServerError)
		return
	}

	writeJson(posts, w)

}

func (c *Controller) createPost(w http.ResponseWriter, r *http.Request) {

	post := &model.Post{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(post); err != nil {
		http.Error(w, "Cannot parse json", http.StatusBadRequest)
		return
	}

	if err := c.postService.AddPost(post); err != nil {
		http.Error(w, "Cannot add data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) getPostById(w http.ResponseWriter, r *http.Request) {

	strId := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		http.Error(w, "Wrong id", http.StatusBadRequest)
		return
	}

	post, err := c.postService.GetPostById(id)

	if err != nil {
		http.Error(w, "Cannot find any data", http.StatusNotFound)
		return
	}

	writeJson(post, w)

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) subscribe(w http.ResponseWriter, r *http.Request) {

	strUserId := r.URL.Query().Get("user")
	strPostId := r.URL.Query().Get("post")

	userId, err := strconv.ParseInt(strUserId, 10, 64)

	if err != nil {
		http.Error(w, "Wrong user id", http.StatusBadRequest)
		return
	}

	postId, err := strconv.ParseInt(strPostId, 10, 64)

	if err != nil {
		http.Error(w, "Wrong user id", http.StatusBadRequest)
		return
	}

	err = c.postService.SupscribeUser(postId, userId)

	if err != nil {
		http.Error(w, "Cannot subscribe user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) addComment(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	comm := &model.Comment{}

	if err := decoder.Decode(comm); err != nil {
		http.Error(w, "Cannot parse json", http.StatusBadRequest)
		return
	}

	if err := c.commsService.AddComment(comm); err != nil {
		http.Error(w, "Cannot add data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (c *Controller) addUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	user := &model.User{}

	if err := decoder.Decode(user); err != nil {
		http.Error(w, "Cannot parse json", http.StatusBadRequest)
		return
	}

	if err := c.userService.AddUser(user); err != nil {
		http.Error(w, "Cannot add data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
