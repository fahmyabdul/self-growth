package requests

import (
	"fmt"

	"github.com/fahmyabdul/golibs"

	"github.com/fahmyabdul/self-growth/fetch-app/configs"
)

type DataApp struct{}

func (p *DataApp) Endpoint() string {
	conf, ok := configs.Properties.Etc["endpoints"].(map[string]interface{})
	if !ok {
		return ""
	}

	return conf["data_app"].(string)
}

func (p *DataApp) GetAll() ([]byte, error) {
	endpoint := fmt.Sprintf("%s/transactions/get", p.Endpoint())
	fmt.Println(endpoint)
	response, responseBody, err := golibs.GetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("server return %d status code", response.StatusCode)
	}

	return responseBody, nil
}
