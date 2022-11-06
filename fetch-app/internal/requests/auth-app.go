package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/fahmyabdul/self-growth/fetch-app/configs"
	"github.com/fahmyabdul/self-growth/fetch-app/internal/models/users"
)

type AuthApp struct{}

type AuthAppValidateResponse struct {
	Message string      `json:"message"`
	Data    users.Users `json:"data"`
}

func (p *AuthApp) Endpoint() string {
	conf, ok := configs.Properties.Etc["endpoints"].(map[string]interface{})
	if !ok {
		return ""
	}

	return conf["auth_app"].(string)
}

func (p *AuthApp) Validate(jwt string) (*AuthAppValidateResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", p.Endpoint(), "validate")

	// Post Request to  Auth App
	client := &http.Client{}

	formData := url.Values{}
	formData.Add("jwt", jwt)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(formData.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal response body
	var responsePayload AuthAppValidateResponse
	err = json.Unmarshal(responseBody, &responsePayload)
	if err != nil {
		return nil, err
	}

	// If jwt is valid then return jwt the content
	return &responsePayload, nil
}
