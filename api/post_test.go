package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloPerson(t *testing.T) {
	person := Person{
		Name: "Leandro",
	}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(person)
	if err != nil {
		t.Error(err)
	}

	nr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/hello-person", &b)

	HelloPerson(nr, r)

	if http.StatusOK != nr.Code {
		t.Errorf("status code expected: %d got %d", http.StatusOK, nr.Code)
	}

	response, err := io.ReadAll(nr.Body)
	if err != nil {
		t.Error(err)
	}

	expBody := fmt.Sprintf("Hello %s\n", person.Name)
	if expBody != string(response) {
		t.Errorf("response expected: %s got %s", expBody, response)
	}
}
