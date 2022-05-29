package helper

import (
	"bookstoreapi/internal/application/writer"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

func GetPathId(w http.ResponseWriter, urlPath string) (int, error) {
	re := regexp.MustCompile("[0-9]")
	pref := re.ReplaceAllString(urlPath, "")

	strId := strings.TrimPrefix(urlPath, pref)

	id, err := strconv.Atoi(strId)

	if err != nil {
		HandleHttpError(w, 400, "invalid fields", nil)

		return 0, err
	}

	return id, nil
}
