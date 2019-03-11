package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	cbrRates = "https://www.cbr-xml-daily.ru/daily_json.js"
)

type ExchangeService struct {
	client *http.Client
}

func NewExchangeService() *ExchangeService {
	return &ExchangeService{
		client: http.DefaultClient,
	}
}

func (e *ExchangeService) CallRates() interface{} {
	resp, err := e.client.Get(cbrRates)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(err)
	}

	var response interface{}
	bits, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bits, &response)
	if err != nil {
		panic(err)
	}

	return response
}
