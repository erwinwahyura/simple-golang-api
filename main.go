package main

import (
	"fmt"
	"log"
	"net/http"
	api "simple-golang-api/controllers"
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

func brandHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		api.GetBrands(w, r)
		return
	case "POST":
		var b models.Brand
		api.CreateBrand(w, r, b)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
		return
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		var p models.Product
		api.CreateProduct(w, r, p)
		return
	case "GET":
		api.GetProductById(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
		return
	}
}

func getProductByBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		api.GetProductByBrand(w, r)
		return
	default:

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
		return
	}
}

func orderHundler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var to models.TransactionOrder
		api.CreateTrx(w, r, to)
	}
}

func getTrxDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		api.GetTrxById(w, r)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/brand", brandHandler)
	mux.HandleFunc("/product/brand", getProductByBrand)
	mux.HandleFunc("/product", productHandler)
	mux.HandleFunc("/order/trx", getTrxDetail)
	mux.HandleFunc("/order", orderHundler)

	server := &http.Server{Addr: ":4000", Handler: mux}
	log.Fatal(server.ListenAndServe())
}
