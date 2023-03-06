package handlers

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/sebastiankul-99/currency-converter/data"
)

type Test struct {
	name          string
	logger        *log.Logger
	response      data.ListOfJsonCurrencyTable
	expectedError error
}

func TestGetCurrenciesTable(t *testing.T) {

	l := log.New(os.Stdout, "currency-api", log.LstdFlags)
	resp := data.ListOfJsonCurrencyTable{
		data.JsonCurrencyTable{
			Table: "A",
		},
	}
	tests := []Test{
		{
			name:          "get-currencies--request",
			logger:        l,
			response:      resp,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			resp, err := getCurrenciesTable(test.logger, "http://api.nbp.pl/api/exchangerates/tables/A")

			if !errors.Is(err, test.expectedError) {
				t.Errorf("Expected error FAILED: expected %v got %v\n", test.expectedError, err)
			}
			if resp[0].Table != test.response[0].Table {
				t.Errorf("FAILED: expected %v, got %v\n", test.response[0].Table, resp[0].Table)
			}

		})
	}
}
