package requests

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fahmyabdul/golibs"

	"github.com/fahmyabdul/efishery-task/fetch-app/configs"
)

type CurrencyConverterApi struct{}

type ExchangeRatesData struct {
	Base       string             `json:"base"`
	Date       string             `json:"date"`
	Rates      map[string]float64 `json:"rates"`
	Success    bool               `json:"success"`
	Timestamp  float64            `json:"timestamp"`
	LimitMonth int                `json:"limit_month"`
}

func (p *CurrencyConverterApi) Endpoint() string {
	conf, ok := configs.Properties.Etc["endpoints"].(map[string]interface{})
	if !ok {
		return ""
	}

	return conf["currencyconverter_api"].(string)
}

func (p *CurrencyConverterApi) ApiKey() string {
	conf, ok := configs.Properties.Etc["endpoints"].(map[string]interface{})
	if !ok {
		return ""
	}

	return conf["currencyconverter_api_key"].(string)
}

func (p *CurrencyConverterApi) GetExchangeRates(from, to string) (*ExchangeRatesData, error) {
	endpoint := fmt.Sprintf("%s/latest?symbols=%s&base=%s", p.Endpoint(), to, from)

	// Set Apikey in header
	header := map[string]string{
		"apikey": p.ApiKey(),
	}

	response, responseBody, err := golibs.SendRequestWithHeader("GET", endpoint, header, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("request failed")
	}

	var exchangeRatesData ExchangeRatesData
	err = json.Unmarshal(responseBody, &exchangeRatesData)
	if err != nil {
		return nil, err
	}

	// Set Current Month Limit
	exchangeRatesData.LimitMonth, err = strconv.Atoi(response.Header.Get("x-ratelimit-remaining-month"))
	if err != nil {
		return nil, err
	}

	return &exchangeRatesData, nil
}
