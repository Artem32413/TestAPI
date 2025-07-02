package components

import (
	"encoding/json"
	"io"
	"net/http"
)

func NewMarshall(variable any) ([]byte, error) {
	return json.MarshalIndent(variable, "", "    ")
}

func NewDec(r *http.Request, variable any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, variable)
}
