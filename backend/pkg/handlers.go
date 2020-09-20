/*** Arnulfo Jose Suarez Gaekel - 2020 - All rights reserved - Contact Email: ajoses@gmail.com */

package pkg

import (
	"encoding/json"
	"github.com/ajoses/salesloft-test/backend/pkg/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
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

	startTime := time.Now()
	people, err := a.salesloftService.GetPeopleList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Info(time.Now().Sub(startTime))

	for i, _ := range people {
		helper := a.operationService.CharFrequency(people[i].EmailAddress)
		people[i].EmailCharFreq = helper.Sort()
	}

	log.Info(time.Now().Sub(startTime))
	json.NewEncoder(w).Encode(people)

}

func (a *api) ListPossibleDuplicates(w http.ResponseWriter, r *http.Request) {
	log.Info("ListPossibleDupliates endpoint called...")

	listPossibleDuplicate := []models.PossibleDuplicates{}
	peopleList, err := a.salesloftService.GetPeopleList()
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, people := range peopleList {
		for _, comparePeople := range peopleList {

			if a.operationService.PossibleDuplicate(people, comparePeople) {
				possDup := models.PossibleDuplicates{
					PeopleOne: people,
					PeopleTwo: comparePeople,
				}
				listPossibleDuplicate = append(listPossibleDuplicate, possDup)
			}
		}
	}

	json.NewEncoder(w).Encode(listPossibleDuplicate)
}
