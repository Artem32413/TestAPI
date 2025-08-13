package errors

import (
	"net/http"

	"go.uber.org/zap"
)

func HandleError(s *zap.Logger, w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
	s.Error(err.Error())
}
