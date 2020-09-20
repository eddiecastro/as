/*** Arnulfo Jose Suarez Gaekel - 2020 - All rights reserved - Contact Email: ajoses@gmail.com */

package pkg

import (
	"encoding/json"
	"github.com/ajoses/voxie-engineering-test/project/backend/pkg/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (a *api) ListPeopleHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("ListPeopleHandler endpoint called...")
	people, err := a.salesloftService.GetPeopleList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(people)
}

func (a *api) GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("GetPeopleHandler endpoint called...")

	peopleID := mux.Vars(r)["id"]

	people, err := a.salesloftService.GetPeople(peopleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(people)

}

func (a *api) GetFrequencyHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("GetFrequencyHandler endpoint called...")

	people, err := a.salesloftService.GetPeopleList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, _ := range people {
		people[i].EmailCharFreq = a.operationService.CharFrequency(people[i].EmailAddress)
	}

	json.NewEncoder(w).Encode(people)

}

func (a *api) ListPossibleDuplicates(w http.ResponseWriter, r *http.Request) {
	log.Info("ListPossibleDupliates endpoint called...")
	var team models.Team

	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.backendHandler.UpdateTeam(team)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
