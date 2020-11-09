package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", process)
	http.ListenAndServe(":9093", nil)
}

func process(w http.ResponseWriter, r *http.Request) {

	result := Result{Status: "E-mail to admin@admin.com sent"}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))

}
