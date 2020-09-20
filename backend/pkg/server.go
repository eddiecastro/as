/*** Arnulfo Jose Suarez Gaekel - 2020 - All rights reserved - Contact Email: ajoses@gmail.com */

package pkg

import (
	"github.com/ajoses/salesloft-test/backend/pkg/services"
	"github.com/gorilla/mux"
	"net/http"
)

type api struct {
	router           http.Handler
	salesloftService *services.SalesLoftAPI
	operationService *services.OperationsService
}

type Server interface {
	Router() http.Handler
}

func NewBackend(apiSalesloftURLBase, apiSalesloftURLKey string) Server {
	a := &api{}

	a.salesloftService = services.NewSalesloftAPI(apiSalesloftURLBase, apiSalesloftURLKey)
	a.operationService = services.NewOperationsService()

	r := mux.NewRouter()
	r.HandleFunc("/", a.ListPeopleHandler).Methods(http.MethodGet)
	r.HandleFunc("/people", a.ListPeopleHandler).Methods(http.MethodGet)
	r.HandleFunc("/people/{id}", a.GetPeopleHandler).Methods(http.MethodGet)

	r.HandleFunc("/frequency", a.GetFrequencyHandler).Methods(http.MethodGet)
	r.HandleFunc("/duplicate", a.ListPossibleDuplicates).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
