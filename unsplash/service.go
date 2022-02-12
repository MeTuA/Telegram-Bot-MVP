package unsplash

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/go-hclog"
)

type Service interface {
	GetRandomPhoto() (*RandomPhoto, error)
}

type service struct {
	credentials Credentials
	client      *http.Client
	log         hclog.Logger
}

func NewService(credentials Credentials, log hclog.Logger) Service {
	client := &http.Client{Timeout: 30 * time.Second}

	return &service{
		credentials: credentials,
		client:      client,
		log:         log,
	}
}

func (s service) GetRandomPhoto() (*RandomPhoto, error) {
	url := "https://api.unsplash.com/photos/random?client_id=" + s.credentials.ClientID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		s.log.Error("failed making request", "error", err.Error())
		return nil, err
	}

	// making request
	resp, err := s.client.Do(req)
	if err != nil {
		s.log.Error("failed do request", "error", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	// handling response and unmarshalling
	randPhoto := &RandomPhoto{}
	response, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(response, randPhoto)
	if err != nil {
		s.log.Error("failed on unmarshalling response", "error", err.Error())
		return nil, err
	}

	return randPhoto, nil
}
