package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/harrysaini/try-go-web-dev/10-mongo/models"
)

// SendErrorResponse to client
func SendErrorResponse(w http.ResponseWriter, err error, code int) {
	response := models.Response{
		Status: models.Status{
			Code:    code,
			Message: err.Error(),
		},
		Error: err,
		Data:  nil,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Panicln(err)
	}
}

// SendResponse to client
func SendResponse(w http.ResponseWriter, data interface{}) {
	response := models.Response{
		Status: models.Status{
			Code:    http.StatusOK,
			Message: "Success",
		},
		Error: nil,
		Data:  data,
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}
