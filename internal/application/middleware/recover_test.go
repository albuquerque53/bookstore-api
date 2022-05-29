package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sceneries struct {
	handler    http.HandlerFunc
	expectCode int
	expectResp string
}

func TestPanicHandler(t *testing.T) {
	successHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body := `{"message":"ok","data":null}`
		w.WriteHeader(200)
		w.Write([]byte(body))
	})
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		panic("execution error")
	})

	testSceneries := []sceneries{
		{handler: successHandler, expectCode: 200, expectResp: `{"message":"ok","data":null}`},
		{handler: panicHandler, expectCode: 500, expectResp: `{"message":"unexpected error","data":null}`},
	}

	for _, scenery := range testSceneries {
		req := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()

		PanicHandler(scenery.handler).ServeHTTP(w, req)

		result := w.Result()
		resp, _ := io.ReadAll(result.Body)

		assert.Equal(t, scenery.expectCode, result.StatusCode)
		assert.Equal(t, scenery.expectResp, string(resp))
	}
}
