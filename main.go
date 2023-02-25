package main

import (
	"fmt"
	"golang/controller"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/students", controller.HandlingHttpReq)
	mux.HandleFunc("/todo", controller.HandlingTodos)
	fmt.Println("started server in 8080")
	http.ListenAndServe("localhost:8080", mux)
}
