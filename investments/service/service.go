package service

import (
	handler "projects/investorsmarket/investments/handlers"

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

	mx.HandleFunc("/investments/investors/{id}", handler.GetInvestmentsByInvestor(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/investments/{id}", handler.GetInvestment(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/investments", handler.CreateInvestment(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/investments", handler.GetAllInvestments(formatter)).Methods(http.MethodGet)
}
