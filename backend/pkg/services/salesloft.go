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

func (s *SalesLoftAPI) GetPeopleList() ([]models.People, error) {
	var ret []models.People
	response, err := s.callRest(salesloftAPIPeople, nil, http.MethodGet)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	err = json.Unmarshal(response, &ret)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return ret, nil

}

func (s *SalesLoftAPI) callRest(endpoint string, data []byte, httpMethod string) ([]byte, error) {
	var response *http.Response
	var err error
	url := fmt.Sprintf("%s/%s", s.apiURLBase, endpoint)

	switch httpMethod {
	case http.MethodGet:
		response, err = http.Get(url)
		break
	case http.MethodPost:
		body := bytes.NewReader(data)
		response, err = http.Post(url, httpContentType, body)
		break
	case http.MethodPut:
		body := bytes.NewReader(data)
		req, err := http.NewRequest(http.MethodPut, url, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", httpContentType)
		response, err = http.DefaultClient.Do(req)
		break
	case http.MethodDelete:
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, err
		}
		response, err = http.DefaultClient.Do(req)
		break
	}

	if err != nil {
		log.Error(fmt.Errorf("The HTTP request failed with error %s\n", err))
		return nil, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	if response.StatusCode == http.StatusOK {
		ret, _ := ioutil.ReadAll(response.Body)

		return ret, nil
	}

	return nil, errors.New(response.Status)
}
