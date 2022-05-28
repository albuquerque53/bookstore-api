package writer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type sendResponseSceneries struct {
	sentStatus       uint
	sentResponseJSON JSONResponse
}

func TestSendResponse(t *testing.T) {
	testSceneries := []sendResponseSceneries{
		{
			sentStatus:       200,
			sentResponseJSON: JSONResponse{Message: "ok", Data: "ok"},
		},
		{
			sentStatus:       200,
			sentResponseJSON: JSONResponse{Message: "ok", Data: map[string]string{"hello": "world"}},
		},
		{
			sentStatus:       400,
			sentResponseJSON: JSONResponse{Message: "ok", Data: 1},
		},
		{
			sentStatus:       200,
			sentResponseJSON: JSONResponse{Message: "ok", Data: nil},
		},
		{
			sentStatus:       500,
			sentResponseJSON: JSONResponse{Message: "Olá Mundo!", Data: "!@#$%*()"},
		},
		{
			sentStatus:       200,
			sentResponseJSON: JSONResponse{Message: "''", Data: ""},
		},
	}

	for _, scenery := range testSceneries {
		sentStatus := scenery.sentStatus
		sentResp := scenery.sentResponseJSON

		handler := func(w http.ResponseWriter, req *http.Request) {
			SendResponse(w, sentStatus, sentResp)
		}

		req := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		handler(w, req)

		resp := w.Result()
		respJSON, _ := io.ReadAll(resp.Body)

		expectJSON := buildExpectedJson(sentResp.Message, sentResp.Data)

		if expectJSON != string(respJSON) {
			t.Errorf("expected JSON (%s) does not match with returned in test: %s", expectJSON, string(respJSON))
		}

		expectContType := "application/json"
		contType := resp.Header.Get("Content-Type")

		if expectContType != contType {
			t.Errorf("expected Content-Type (%s) does not match with returned in test: %s", expectContType, contType)
		}

		status := resp.StatusCode

		if int(sentStatus) != status {
			t.Errorf("expected status code (%d) does not match with returned in test: %d", sentStatus, status)
		}
	}
}

func buildExpectedJson(message string, data interface{}) string {
	jsonData, _ := json.Marshal(data)

	return fmt.Sprintf(`{"message":"%s","data":%s}`, message, jsonData)
}
