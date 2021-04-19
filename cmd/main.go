package main

import (
	"log"
	"net/http"
	"sync"
	"troyan/pkg/handlers"

	"github.com/gorilla/mux"
)

const ngrok = ""

func main() {
	wg := sync.WaitGroup{}
	r := mux.NewRouter()

	r.HandleFunc("/input", handlers.InputHandler).Methods("POST")
	r.HandleFunc("/normalize", handlers.NormalizeHandler).Methods("POST")
	log.Println("Listening on 8080")
	wg.Add(1)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Panic("Cannot listen on :8080")
	}
	wg.Wait()
}
