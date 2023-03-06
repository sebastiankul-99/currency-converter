package data

import (
	"encoding/json"
	"io"
)

type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Currencies []*Currency

var currencyList = Currencies{}               // use to cache currencies, so that there is no need to call NBP API each time
var currencyToTable = make(map[string]string) // mapping currency to table because currency can be in either A or B table
// this way we don't need to provide frontend with table's name for each currency

func GetCurrencyList() *Currencies {
	return &currencyList
}

func GetCurrencyMap() *map[string]string {
	return &currencyToTable
}

func AddToCurrencyList(list *ListOfJsonCurrencyTable) {
	for _, object := range *list {
		for _, currency := range object.Rates {
			cur := Currency{
				currency.Code,
				currency.Currency,
			}
			currencyList = append(currencyList, &cur)
		}
	}
}

func AddToCurrencyMap(list *ListOfJsonCurrencyTable) {
	for _, object := range *list {
		for _, currency := range object.Rates {
			currencyToTable[currency.Code] = object.Table
		}
	}
}

func (c *Currencies) CurrenciesToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}
