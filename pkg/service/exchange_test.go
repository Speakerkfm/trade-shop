package service

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestExchangeService_CallRates(t *testing.T) {
	body := `{"Valute": {"AUD": {"ID": "R01010","NumCode": "036","CharCode": "AUD","Nominal": 1,"Name": "Австралийский доллар","Value": 46.4457,"Previous": 46.3074}}}`
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	exchange := NewExchangeService()
	httpmock.RegisterResponder("GET", cbrRates,
		func(request *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, body)
		})

	res := exchange.CallRates()
	assert.Equal(t, res, body)
}
