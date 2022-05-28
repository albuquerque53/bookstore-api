package writer

import (
	"encoding/json"
	"fmt"
	"testing"
)

type sceneries struct {
	messageSent JsonResponse
}

func TestJSONString(t *testing.T) {
	testSceneries := []sceneries{
		{JsonResponse{Message: "ok", Data: "ok"}},
		{JsonResponse{Message: "ok", Data: 1}},
		{JsonResponse{Message: "ok", Data: map[string]string{"hello": "world"}}},
		{JsonResponse{Message: "olááá", Data: "@!#!@#1"}},
		{JsonResponse{Message: "''", Data: ""}},
	}

	for _, scenery := range testSceneries {
		jsonString, err := ToJSON(scenery.messageSent)

		if err != nil {
			t.Errorf("could not generate JSON by JSONString method by the following error: %v", err)
		}

		expectedJsonString := buildExpectedJson(scenery.messageSent.Message, scenery.messageSent.Data)

		if expectedJsonString != jsonString {
			t.Errorf("the expected json (%v) does not match with returned (%v)", expectedJsonString, jsonString)
		}
	}
}

func buildExpectedJson(message string, data interface{}) string {
	jsonData, _ := json.Marshal(data)

	return fmt.Sprintf(`{"message":"%s","data":%s}`, message, jsonData)
}
