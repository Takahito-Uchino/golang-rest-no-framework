package main

import (
	"net/http"

	"github.com/Takahito-Uchino/golang-rest-no-framework/controller"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todos", controller.TodoHandler)
	server.ListenAndServe()
}
