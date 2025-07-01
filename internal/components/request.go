package components

import (
	"encoding/json"
	"net/http"
)

func NewDec(r *http.Request, variable any) error {
	return json.NewDecoder(r.Body).Decode(variable)
}

func NewMarshall(variable any) ([]byte, error) {
	return json.MarshalIndent(variable, "", "    ")
}
