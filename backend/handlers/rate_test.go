package handlers

import (
	"errors"
	"log"
	"os"
	"testing"
)

type Tests struct {
	name          string
	logger        *log.Logger
	response      float32
	expectedError error
}

func TestGetRate(t *testing.T) {

	l := log.New(os.Stdout, "currency-api", log.LstdFlags)
	tests := []Tests{
		{
			name:          "get-usd-rate-request",
			logger:        l,
			response:      4.12,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			_, err := getRate(test.logger, "USD", "A")

			if !errors.Is(err, test.expectedError) {
				t.Errorf("Expected error FAILED: expected %v got %v\n", test.expectedError, err)
			}
		})
	}
}
