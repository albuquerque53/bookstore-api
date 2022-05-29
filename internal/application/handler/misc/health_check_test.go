package misc

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()

	HealthCheck(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	expectBody := `{"message":"Hello, World","data":null}`

	assert.Equal(t, string(body), expectBody)
	assert.Equal(t, 200, resp.StatusCode)
}
