package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello route was called\n")
	v := response{0, "Ok"}
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	godotenv.Load()
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
