package writer

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(w http.ResponseWriter, status uint, resp JSONResponse) {
	jsonResp, err := toJSON(resp)

	if err != nil {
		log.Fatalf("error on response build %s", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(int(status))
	w.Write([]byte(jsonResp))
}

func toJSON(resp JSONResponse) (string, error) {
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		return "", err
	}

	return string(jsonResp), nil
}
