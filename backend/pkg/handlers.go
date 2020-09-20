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

func (a *api) GetTeamsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("GetTeams endpoint called...")
	teams, err := a.backendHandler.GetTeams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(teams)
}

func (a *api) GetTeamHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("GetTeam endpoint called...")
	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["id"])

	teams, err := a.backendHandler.GetTeam(int64(teamID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(teams)

}

func (a *api) AddTeamHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("AddTeam endpoint called...")
	var team models.Team

	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.backendHandler.AddTeam(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (a *api) UpdateTeamHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("UpdateTeam endpoint called...")
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

func (a *api) DeleteTeamHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("DeleteTeam endpoint called...")
	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["id"])

	err = a.backendHandler.DeleteTeam(int64(teamID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
