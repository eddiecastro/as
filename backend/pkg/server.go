/*** Arnulfo Jose Suarez Gaekel - 2020 - All rights reserved - Contact Email: ajoses@gmail.com */

package pkg

import (
	"github.com/ajoses/voxie-engineering-test/project/backend/pkg/db"
	"github.com/gorilla/mux"
	"net/http"
)

type api struct {
	router         http.Handler
	backendHandler *ServiceHandlers
}

type Server interface {
	Router() http.Handler
}

func NewBackend(dbClient db.Contract) Server {
	a := &api{}

	a.backendHandler = NewServiceHandlers(dbClient)

	r := mux.NewRouter()
	r.HandleFunc("/import", a.AddTeamHandler).Methods(http.MethodPost)
	r.HandleFunc("/teams", a.GetTeamsHandler).Methods(http.MethodGet)
	r.HandleFunc("/teams/{id}", a.GetTeamHandler).Methods(http.MethodGet)

	r.HandleFunc("/teams", a.AddTeamHandler).Methods(http.MethodPost)
	r.HandleFunc("/teams", a.UpdateTeamHandler).Methods(http.MethodPut)
	r.HandleFunc("/teams/{id}", a.DeleteTeamHandler).Methods(http.MethodDelete)

	r.HandleFunc("/contacts/search/{search}", a.SearchContactHandler).Methods(http.MethodGet)
	r.HandleFunc("/contacts/{contactid}", a.GetContactHandler).Methods(http.MethodGet)
	r.HandleFunc("/contacts/{teamid}", a.AddContactHandler).Methods(http.MethodPost)
	r.HandleFunc("/contacts", a.UpdateContactHandler).Methods(http.MethodPut)
	r.HandleFunc("/contacts/{contactid}", a.DeleteContactHandler).Methods(http.MethodDelete)

	r.HandleFunc("/customattr/{customid}", a.GetCustomAttrHandler).Methods(http.MethodGet)
	r.HandleFunc("/customattr/{contactid}", a.AddCustomAttrHandler).Methods(http.MethodPost)
	r.HandleFunc("/customattr", a.UpdateCustomAttrHandler).Methods(http.MethodPut)
	r.HandleFunc("/customattr/{customid}", a.DeleteCustomAttrHandler).Methods(http.MethodDelete)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
