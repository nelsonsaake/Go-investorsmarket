package service

import (
	handler "projects/investorsmarket/chats/handlers"

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

	mx.HandleFunc("/chats", handler.CreateChat(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/chats/counterpartsof/{id}", handler.GetChatCounterPartsOf(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/chats/chats/{u1}/{u2}", handler.GetChats(formatter)).Methods(http.MethodGet)
}