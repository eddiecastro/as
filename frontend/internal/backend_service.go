/*** Arnulfo Jose Suarez Gaekel - 2020 - All rights reserved - Contact Email: ajoses@gmail.com */

package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ajoses/salesloft-test/backend/pkg/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	httpContentType = "application/json"

	backendAPIGetPeople     = "people"
	backendAPIGetFrequency  = "frequency"
	backendAPIGetDuplicates = "duplicate"
)

type BackendService struct {
	backendHost string
}

func NewBackendService(host string) *BackendService {
	ret := &BackendService{
		backendHost: host,
	}

	return ret
}

func (b *BackendService) GetPeople() ([]models.People, error) {
	var ret []models.People

	response, err := b.callRest(backendAPIGetPeople, nil, http.MethodGet)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil

}

func (b *BackendService) GetFrequency() ([]models.People, error) {
	var ret []models.People

	response, err := b.callRest(backendAPIGetFrequency, nil, http.MethodGet)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(response, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil

}

func (b *BackendService) GetDuplicate() ([]models.PossibleDuplicates, error) {
	var ret []models.PossibleDuplicates

	response, err := b.callRest(backendAPIGetDuplicates, nil, http.MethodGet)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (b *BackendService) callRest(endpoint string, data []byte, httpMethod string) ([]byte, error) {
	var response *http.Response
	var err error
	url := fmt.Sprintf("%s/%s", b.backendHost, endpoint)

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
