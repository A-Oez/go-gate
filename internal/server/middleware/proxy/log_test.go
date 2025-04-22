package proxy

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusCapturingResponseWriter_WriteHeader(t *testing.T) {
	recorder := httptest.NewRecorder()

	writer := &LoggingResponseWriter{
		ResponseWriter: recorder,
		StatusCode:     http.StatusOK, // default
	}

	writer.WriteHeader(http.StatusTeapot)

	if writer.StatusCode != http.StatusTeapot {
		t.Errorf("Expected recorder to have status code %d, got %d", http.StatusTeapot, recorder.Result().StatusCode)
	}
}
