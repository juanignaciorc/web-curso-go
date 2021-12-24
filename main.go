package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
	msg := mensaje{msg: "Hola mundo again"}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)
	mux.HandleFunc("/hola-mundo", holaMundo)

	mux.Handle("/hola", msg)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}
