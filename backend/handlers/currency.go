package handlers

import (
	"log"
	"net/http"

	"github.com/sebastiankul-99/currency-converter/data"
)

type CurrencytHandler struct {
	l *log.Logger
}

func GetCurrencytHandler(l *log.Logger) *CurrencytHandler {
	return &CurrencytHandler{l}
}

func getCurrenciesTable(l *log.Logger, uri string) (data.ListOfJsonCurrencyTable, error) {
	req, err := http.Get(uri)
	if err != nil {
		l.Println(err.Error())
		return nil, err
	}
	var table = data.ListOfJsonCurrencyTable{}
	err = data.UnmarshalJsonCurrencyTable(req.Body, &table)
	if err != nil {
		l.Println(err.Error())
		return nil, err
	}
	return table, nil
}

func getCurrencies(l *log.Logger) error {
	l.Println("getting currencies") // to get each currency, we need to call api twice for 2 different tables
	firstTable, err := getCurrenciesTable(l, "http://api.nbp.pl/api/exchangerates/tables/A")
	if err != nil {
		return err
	}
	secondTable, error := getCurrenciesTable(l, "http://api.nbp.pl/api/exchangerates/tables/B")
	if err != nil {
		return error
	}
	data.AddToCurrencyList((&firstTable))
	data.AddToCurrencyList((&secondTable))
	data.AddToCurrencyMap(&firstTable)
	data.AddToCurrencyMap(&secondTable)
	return nil
}

func (c *CurrencytHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if len(*data.GetCurrencyList()) == 0 { // if currencies are not cached
		err := getCurrencies(c.l)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	err := data.GetCurrencyList().CurrenciesToJSON(rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
