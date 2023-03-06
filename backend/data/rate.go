package data

import (
	"encoding/json"
	"io"
)

type Rate struct { // used for deserialize JSON from rate call
	Rates []struct {
		Rate float32 `json:"mid"`
	} `json:"rates"`
}

func UnmarshalRate(r io.Reader, j *Rate) error {
	d := json.NewDecoder(r)
	return d.Decode(j)
}
