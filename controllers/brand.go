package api

import (
	"fmt"
	"net/http"
	"simple-golang-api/config"
	"simple-golang-api/helpers"
	"simple-golang-api/models"
)

func mapBrandToJSON(b *models.Brand) models.JSONBrand {
	return models.JSONBrand{
		Id:   b.Id,
		Name: b.Name,
	}
}

func GetBrands(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from brand")

	if err != nil {
		resp := helpers.CustomeRes{
			Msg:  "Error: while retrieve data from db",
			Err:  err,
			Code: http.StatusInternalServerError,
			Data: nil,
		}
		helpers.SendResponse(w, r, resp, http.StatusOK)
		return
	}
	defer rows.Close()

	var result []models.JSONBrand

	for rows.Next() {
		var brand models.JSONBrand

		if err := rows.Scan(&brand.Id, &brand.Name); err != nil {
			fmt.Println(err)
		}

		result = append(result, brand)
	}
	resp := helpers.CustomeRes{
		Msg:  "successfully get brands",
		Err:  nil,
		Code: http.StatusOK,
		Data: result,
	}
	helpers.SendResponse(w, r, resp, http.StatusOK)
}

func CreateBrand(w http.ResponseWriter, r *http.Request, data models.Brand) {
	req := data
	err := helpers.Parse(w, r, &req)
	if err != nil {
		resp := helpers.CustomeRes{
			Msg:  "Error: cannot parsing data body",
			Err:  err,
			Code: http.StatusBadRequest,
			Data: nil,
		}
		helpers.SendResponse(w, r, resp, http.StatusBadRequest)
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// save data to db
	b := models.Brand{
		Id:   req.Id,
		Name: req.Name,
	}

	err = helpers.ValidationBrand(w, r, b.Name)
	if err != nil {
		resp := helpers.CustomeRes{
			Msg:  err.Error(),
			Err:  err,
			Code: http.StatusBadRequest,
			Data: nil,
		}
		helpers.SendResponse(w, r, resp, http.StatusInternalServerError)
		return
	}

	sqlStatement := `insert into brand (name) values ($1)`
	res, err := db.Exec(sqlStatement, b.Name)

	if err != nil {
		resp := helpers.CustomeRes{
			Msg:  "Error: cannot save new brand in DB",
			Err:  err,
			Code: http.StatusInternalServerError,
			Data: nil,
		}
		helpers.SendResponse(w, r, resp, http.StatusInternalServerError)
		return
	}
	res.LastInsertId()

	response := mapBrandToJSON(&b)

	resp := helpers.CustomeRes{
		Msg:  "successfully create new brand",
		Err:  err,
		Code: http.StatusOK,
		Data: response,
	}
	helpers.SendResponse(w, r, resp, http.StatusInternalServerError)
}
