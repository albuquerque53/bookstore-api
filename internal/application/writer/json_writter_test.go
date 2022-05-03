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

func TestJSONString(t *testing.T) {
	expectedMessage := "Hello, World!"

	jsonString, err := ToJSON(expectedMessage)

	if err != nil {
		t.Errorf("could not generate JSON by JSONString method by the following error: %v", err)
	}

	expectedJsonString := buildExpectedJson("Hello, World!")

	if expectedJsonString != jsonString {
		t.Errorf("the expected json (%v) does not match with returned (%v)", expectedJsonString, jsonString)
	}
}

func buildExpectedJson(message string) string {
	return fmt.Sprintf(`{"data":{"message":"%s"}}`, message)
}
