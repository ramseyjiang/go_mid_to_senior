package restymock

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type XML struct {
	Name string `xml:"Name"`
}

type JSON struct {
	Name string `json:"name"`
}

type Service struct {
	Client *resty.Client
}

func (s *Service) GetXML(url string) (*XML, error) {
	r, err := s.Client.R().
		SetResult(&XML{}).
		Get(url)

	if err != nil {
		return nil, err
	}

	if !r.IsSuccess() {
		return nil, fmt.Errorf("request faild with code %d", r.StatusCode())
	}

	return r.Result().(*XML), nil
}

func (s *Service) GetJSON(url string) (*JSON, error) {
	r, err := s.Client.R().
		SetResult(&JSON{}).
		Get(url)

	if err != nil {
		return nil, err
	}

	if !r.IsSuccess() {
		return nil, fmt.Errorf("request faild with code %d", r.StatusCode())
	}

	return r.Result().(*JSON), nil
}
