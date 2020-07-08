package service

import (
	handler "projects/investorsmarket/users/handlers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
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

	mx.HandleFunc("/users/a/", handler.GetUserA(formatter)).Methods("POST")
	mx.HandleFunc("/users", handler.CreateUser(formatter)).Methods("POST")
	mx.HandleFunc("/users", handler.GetAllUsers(formatter)).Methods("GET")
	mx.HandleFunc("/users/{id}", handler.GetUser(formatter)).Methods("GET")
	mx.HandleFunc("/users/{id}",  handler.UpdateUser(formatter)).Methods(http.MethodPut)
}
