package service

import (
	handler "projects/investorsmarket/auths/handlers"

	"net/http"

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

	mx.HandleFunc("/auths", handler.CreateAuth(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/auths/ep", handler.GetAuthGivenEP(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/auths/t1", handler.GetAuthGivenToken(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/auths/ut", handler.GetAuthGivenUserAndToken(formatter)).Methods(http.MethodPost)
}
