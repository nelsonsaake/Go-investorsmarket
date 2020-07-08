package service

import (
	handler "projects/investorsmarket/opportunities/handlers"

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

	mx.HandleFunc("/opportunities", handler.GetAllOpportunities(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/opportunities", handler.CreateOpportunity(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/opportunities/creators", handler.GetOpportunityCreators(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/opportunities/creators/{id}/history", handler.GetCreatorHistory(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/opportunities/{id}", handler.GetOpportunity(formatter)).Methods(http.MethodGet)
}
