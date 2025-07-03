package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name string
}

func HelloPerson(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var person Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello %s\n", person.Name)
}
