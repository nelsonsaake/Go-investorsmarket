package service

import (
	handler "projects/investorsmarket/posts/handlers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(
		render.Options{
			IndentJSON: true,
		},
	)

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	
	mx.HandleFunc("/posts", handler.CreatePost(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/posts/opportunity/{id}", handler.GetOpportunityPosts(formatter)).Methods(http.MethodGet)
}
