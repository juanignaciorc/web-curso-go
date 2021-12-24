package main

import (
	"fmt"
	"net/http"
)

func holaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola munndo </h1>")
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func main() {
	msg := mensaje{msg: "<Hola mundo again"}
	mux := http.NewServeMux()
	mux.HandleFunc("/", holaMundo)

	mux.Handle("/hola", msg)
	http.ListenAndServe(":8080", mux)
}
