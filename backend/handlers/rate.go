package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/sebastiankul-99/currency-converter/data"
)

type RateHandler struct {
	l *log.Logger
}

func GetRateHandler(l *log.Logger) *RateHandler {
	return &RateHandler{l}
}

func getRate(l *log.Logger, symbol string, table string) (float32, error) {
	uri := "http://api.nbp.pl/api/exchangerates/rates/" + table + "/" + symbol
	req, err := http.Get(uri)
	if err != nil {
		l.Println(err.Error())
		return 0, err
	}
	rate := data.Rate{}
	err = data.UnmarshalRate(req.Body, &rate)
	if err != nil {
		l.Println(err.Error())
		return 0, err
	}
	if len(rate.Rates) == 0 {
		return 0, errors.New("could not get currency rate")
	}

	return rate.Rates[0].Rate, nil
}

func (c *RateHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*") // preventing from CORS' block because of running frontend on localhost
	symbol := strings.TrimPrefix(r.URL.Path, "/rate/")
	if len(*data.GetCurrencyList()) == 0 { // if data not cached
		err := getCurrencies(c.l)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	table := (*data.GetCurrencyMap())[symbol]
	if table == "" {
		http.Error(rw, "Currency Not Found", http.StatusBadRequest)
		return
	}

	rate, err := getRate(c.l, symbol, table)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	responseBody := make(map[string]float32)
	responseBody["rate"] = rate

	err = json.NewEncoder(rw).Encode(responseBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
