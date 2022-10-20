package requests

import (
	"fmt"

	"github.com/fahmyabdul/golibs"

	"github.com/fahmyabdul/efishery-task/fetch-app/configs"
)

type EfisheryApi struct{}

func (p *EfisheryApi) Endpoint() string {
	conf, ok := configs.Properties.Etc["endpoints"].(map[string]interface{})
	if !ok {
		return ""
	}

	return conf["efishery_api"].(string)
}

func (p *EfisheryApi) GetKomoditas() ([]byte, error) {
	endpoint := fmt.Sprintf("%s", p.Endpoint())
	response, responseBody, err := golibs.GetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("server return %d status code", response.StatusCode)
	}

	return responseBody, nil
}
