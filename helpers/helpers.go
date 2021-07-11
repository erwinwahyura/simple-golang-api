package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"simple-golang-api/models"
)

type CustomeRes struct {
	Msg  string      `json:"message"`
	Err  error       `json:"err"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func Parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func SendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

func ValidationBrand(w http.ResponseWriter, r *http.Request, name string) error {
	w.Header().Add("Content-Type", "application/json")
	fmt.Println("data: ", name)

	if len(name) == 0 {
		return errors.New("validation error: name cannot be empty")
	} else if len(name) < 2 {
		return errors.New("validation error: name mininal 3 characters")
	}
	return nil
}

func ValidationProduct(w http.ResponseWriter, r *http.Request, p models.JSONProduct) error {
	w.Header().Add("Content-Type", "application/json")

	if len(p.Name) == 0 {
		return errors.New("validation error: name cannot be empty")
	} else if p.Price == 0 || p.Price < 0 {
		return errors.New("validation error: price cannot be 0 or minus")
	} else if reflect.ValueOf(p.Price).Kind() != reflect.Int {
		return errors.New("validation error: price must be a number")
	} else if reflect.ValueOf(p.BrandId).Kind() != reflect.Int {
		return errors.New("validation error: brand_id must be a number")
	} else if p.BrandId == 0 || p.BrandId < 0 {
		return errors.New("validation error: brand_id must be a higher than 0")
	}

	return nil
}
