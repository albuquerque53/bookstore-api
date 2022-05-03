package writer

import "encoding/json"

type JsonMessageWriter struct {
	Message string `json:"message"`
}

func NewMessageWriter(message string) *JsonMessageWriter {
	return &JsonMessageWriter{
		Message: message,
	}
}

func (message *JsonMessageWriter) JSONString() (string, error) {
	messageResponse := map[string]interface{}{
		"data": map[string]string{
			"message": message.Message,
		},
	}

	bytesValue, err := json.Marshal(messageResponse)

	if err != nil {
		return "", err
	}

	return string(bytesValue), nil
}
