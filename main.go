package main

import (
	"encoding/json"
	_ "golang-auth-database/database/migrations"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["result"] = "OK"
	response["results"] = "OK2"
	json.NewEncoder(w).Encode(response)
}

func main() {

	// Step1 Migrate Data Base
	// migrations.MigrateDB()

	http.HandleFunc("/", ResponseJson)
	http.ListenAndServe(":8080", nil)
}
