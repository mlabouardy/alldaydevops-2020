package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
)

type Response struct {
	Environment string `json:"environment"`
}

func EnvironmentHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Environment: os.Getenv("ENVIRONMENT"),
	}
	json.NewEncoder(w).Encode(response)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	http.ServeFile(w, r, path.Join(".", "static/index.html"))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/environment", EnvironmentHandler)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
