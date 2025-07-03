package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	nr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/hello-world", nil)
	HelloWorld(nr, r)

	if http.StatusOK != nr.Code {
		t.Errorf(" Status code expected: %d got %d", http.StatusOK, nr.Body)
	}
}
