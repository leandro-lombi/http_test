package main

import (
	"net/http"

	"github.com/leandro-lombi/http_test.git/api"
)

func main() {
	http.HandleFunc("/hello-world", api.HelloWorld)
	http.HandleFunc("/hello-person", api.HelloPerson)
	
	http.ListenAndServe(":8080", nil)
}
