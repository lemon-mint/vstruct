package main

import (
	"embed"
	"log"
	"net/http"
	"os"
	"os/signal"
)

//go:embed dist/*
var DistFS embed.FS

//go:embed index.html
var indexHTML []byte

//go:embed compiler/index.html
var compilerHTML []byte

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Cache-Control", "no-cache")
		w.Write(indexHTML)
	})

	http.HandleFunc("/compiler", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Cache-Control", "no-cache")
		w.Write(compilerHTML)
	})

	http.Handle("/dist/", http.FileServer(http.FS(DistFS)))
	server := http.Server{}
	server.Addr = ":8080"
	go func() {
		log.Println(server.ListenAndServe())
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	log.Println("Server started on http://localhost:8080")
	<-sig
	log.Println("Stopping server...")
}
