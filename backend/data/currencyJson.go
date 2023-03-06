package data

import (
	"encoding/json"
	"io"
)

type JsonCurrencyTable struct { // use to deserialize JSON response
	Table string `json:"table"`
	Rates []struct {
		Currency string `json:"currency"`
		Code     string `json:"code"`
	} `json:"rates"`
}
type ListOfJsonCurrencyTable = []JsonCurrencyTable // because API returns currencies as an array

func UnmarshalJsonCurrencyTable(r io.Reader, j *ListOfJsonCurrencyTable) error {
	d := json.NewDecoder(r)
	return d.Decode(j)
}
