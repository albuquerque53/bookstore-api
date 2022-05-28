package helper

import (
	"bookstoreapi/internal/application/writer"
	"log"
	"net/http"
)

func HandleHttpError(w http.ResponseWriter, status uint, message string, data interface{}) {
	if data != nil {
		log.Print(data)
	}

	writer.SendResponse(w, status, writer.JSONResponse{
		Message: message,
		Data:    data,
	})
}
