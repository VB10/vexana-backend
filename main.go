package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello",HelloHandler)
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	io.WriteString(writer, `{"alive": true}`)

	fmt.Print("hello")
}
