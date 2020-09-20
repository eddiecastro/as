package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ajoses/salesloft-test/backend/pkg/models"
	"github.com/ajoses/salesloft-test/backend/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	salesloftAPIPeople = "/v2/people.json"
	httpContentType    = "application/json"
)

type SalesLoftAPI struct {
	apiURLBase string
	apiKey     string
}

type responseHelper struct {
	Data []models.People
}

func NewSalesloftAPI(apiURLBase, apiKey string) *SalesLoftAPI {
	ret := &SalesLoftAPI{
		apiURLBase: apiURLBase,
		apiKey:     apiKey,
	}

	return ret
}

func (s *SalesLoftAPI) GetPeopleList() ([]models.People, error) {
	var ret responseHelper
	response, err := s.callRest(salesloftAPIPeople, nil, http.MethodGet)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	err = json.Unmarshal(response, &ret)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return ret.Data, nil

}

func (s *SalesLoftAPI) GetPeople(id string) (*models.People, error) {
	var ret responseHelper

	url := fmt.Sprintf("%s?ids%%5B%%5D=%s", id)
	response, err := s.callRest(url, nil, http.MethodGet)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	err = json.Unmarshal(response, &ret)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	if len(ret.Data) > 0 {
		return &ret.Data[0], nil
	}

	return nil, nil

}

func (s *SalesLoftAPI) callRest(endpoint string, data []byte, httpMethod string) ([]byte, error) {
	var response *http.Response
	var err error
	url := fmt.Sprintf("%s/%s", s.apiURLBase, endpoint)
	authHeader := fmt.Sprintf("Bearer %s", s.apiKey)

	switch httpMethod {
	case http.MethodGet:
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", authHeader)
		response, err = http.DefaultClient.Do(req)

		break
	case http.MethodPost:
		body := bytes.NewReader(data)

		req, err := http.NewRequest(http.MethodPost, url, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", httpContentType)
		req.Header.Set("Authorization", authHeader)
		response, err = http.DefaultClient.Do(req)
		break
	case http.MethodPut:
		body := bytes.NewReader(data)
		req, err := http.NewRequest(http.MethodPut, url, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", httpContentType)
		req.Header.Set("Authorization", authHeader)
		response, err = http.DefaultClient.Do(req)
		break
	case http.MethodDelete:
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", authHeader)
		response, err = http.DefaultClient.Do(req)
		break
	}

	if err != nil {
		log.Error(fmt.Errorf("The HTTP request failed with error %s\n", err))
		return nil, utils.HandleError(fmt.Errorf("The HTTP request failed with error %s\n", err))
	}

	if response.StatusCode == http.StatusOK {
		ret, _ := ioutil.ReadAll(response.Body)

		return ret, nil
	}

	return nil, utils.HandleError(errors.New(response.Status))
}
