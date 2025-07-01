package components

import "net/http"

func (s *Settings) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}