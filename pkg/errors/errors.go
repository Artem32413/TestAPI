package errors

import (
	"apiGo/internal/inventory/model/structs"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func HandleError(s *zap.Logger, w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    
    resp := structs.ErrorResponse{
        Status:  status,
        Message: err.Error(),
    }
    
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        http.Error(w, "Не удалось закодировать ответ об ошибке: ", http.StatusInternalServerError)
    }
	
	s.Error(err.Error())
}
