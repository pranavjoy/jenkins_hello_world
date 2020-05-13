package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
        "net/http"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

type Simple struct {
	Name        string
	Description string
	Url         string
}

func SimpleFactory (host string) Simple {
	return Simple{"Hello", "World", host}
}

func handler(w http.ResponseWriter, r *http.Request) {
	simple := Simple{"Pranav", "Joy", r.Host}

	jsonOutput, _ := json.Marshal(simple)

	fmt.Fprintln(w, string(jsonOutput))
}

func main() {
	fmt.Println("Server started")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
	fmt.Println("Starting monitoring")
	http.Handle("/metrics", promhttp.Handler())
}
