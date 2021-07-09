package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"simple-golang-api/config"
	"simple-golang-api/models"

	_ "github.com/lib/pq"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
	fmt.Println("handled /api/")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handled /")
	w.Write([]byte("ok"))
}

func getBrands(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handled /")
	db, err := config.ConnectDB()

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from brand")

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var result []models.Brand

	for rows.Next() {
		var brand models.Brand
	
		if err := rows.Scan(&brand.Id, &brand.Name); err != nil {
			fmt.Println(err)
		}

		result = append(result, brand)

	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/brands", getBrands)

	server := &http.Server{Addr: ":4000", Handler: mux}
	log.Fatal(server.ListenAndServe())
}
