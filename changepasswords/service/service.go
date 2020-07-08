package service

import (
	handler "projects/investorsmarket/changepasswords/handlers"

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
	
	mx.HandleFunc("/changepasswords", handler.CreateChangePassword(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/changepasswords", handler.GetAllChangePasswords(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/changepasswords/g/1", handler.GetChangePassword(formatter)).Methods(http.MethodPost)
}
