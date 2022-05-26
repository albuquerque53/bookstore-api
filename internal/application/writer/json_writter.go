package writer

import "encoding/json"

type JsonData struct {
	Data JsonMessage `json:"data"`
}

type JsonMessage struct {
	Message interface{} `json:"message"`
}

func ToJSON(message interface{}) (string, error) {
	messageStruct := JsonMessage{Message: message}

	return buildJsonData(messageStruct)
}

func buildJsonData(message JsonMessage) (string, error) {
	structData := JsonData{Data: message}

	jsonData, _ := json.Marshal(structData)

	return string(jsonData), nil
}
