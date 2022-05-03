package writer

import (
	"fmt"
	"testing"
)

type MockedJsonData struct {
	Data MockedJsonMessage `json:"data"`
}

type MockedJsonMessage struct {
	Message string `json:"message"`
}

type sceneries struct {
	messageSent string
}

func TestJSONString(t *testing.T) {
	testSceneries := []sceneries{
		{"hello world"},
		{"Hello, World!"},
		{"Olá, Mundo!"},
		{"''"},
	}

	for _, scenery := range testSceneries {
		jsonString, err := ToJSON(scenery.messageSent)

		if err != nil {
			t.Errorf("could not generate JSON by JSONString method by the following error: %v", err)
		}

		expectedJsonString := buildExpectedJson(scenery.messageSent)

		if expectedJsonString != jsonString {
			t.Errorf("the expected json (%v) does not match with returned (%v)", expectedJsonString, jsonString)
		}
	}
}

func buildExpectedJson(message string) string {
	return fmt.Sprintf(`{"data":{"message":"%s"}}`, message)
}
