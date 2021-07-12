package api

import (
	"fmt"
	"net/http"
	"simple-golang-api/config"
	"simple-golang-api/helpers"
	"simple-golang-api/models"
)

func mapResToJSON(p *models.Product) models.JSONProduct {
	return models.JSONProduct{
		Id:      p.Id,
		BrandId: p.BrandId,
		Name:    p.Name,
		Price:   p.Price,
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	if len(queryId) == 0 {
		m := "Error: Url Param 'id' is missing"
		fmt.Println(m)
		msg := helpers.CustomeRes{
			Msg:  m,
			Err:  nil,
			Code: http.StatusBadRequest,
		}
		helpers.SendResponse(w, r, msg, http.StatusBadRequest)
		return
	}

	req := models.Product{}

	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
		msg := helpers.CustomeRes{
			Msg:  "",
			Err:  err,
			Code: http.StatusBadRequest,
		}
		helpers.SendResponse(w, r, msg, http.StatusBadRequest)
	}

	// save data to db
	p := models.Product{
		Id:      req.Id,
		BrandId: req.BrandId,
		Name:    req.Name,
		Price:   req.Price,
	}

	sqlStatement := `select * from product where id=$1;`
	err = db.QueryRow(sqlStatement, queryId).Scan(&p.Id, &p.Price, &p.BrandId, &p.Name)
	if err != nil {
		fmt.Println(err)
		msg := helpers.CustomeRes{
			Msg:  "Error: cannot get product by id ",
			Err:  err,
			Code: http.StatusBadRequest,
		}
		helpers.SendResponse(w, r, msg, http.StatusInternalServerError)
	}

	response := mapResToJSON(&p)
	resp := helpers.CustomeRes{
		Msg:  "successfully get product by id",
		Err:  nil,
		Code: http.StatusOK,
		Data: response,
	}
	helpers.SendResponse(w, r, resp, http.StatusOK)
}

func GetProductByBrand(w http.ResponseWriter, r *http.Request) {
	queryId := r.URL.Query().Get("id")
	if len(queryId) == 0 {
		m := "Error: Url Param 'id' is missing"
		fmt.Println(m)
		msg := helpers.CustomeRes{
			Msg:  m,
			Err:  nil,
			Code: http.StatusBadRequest,
			Data: nil,
		}
		helpers.SendResponse(w, r, msg, http.StatusBadRequest)
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("select p.name as product_name, p.id, p.price, p.brand_id, b.name as brand_name from product p inner join Brand b on b.id = p.brand_id where p.brand_id = $1 ", queryId)
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

	var result []models.JSONProductBrand
	for rows.Next() {
		var pb models.JSONProductBrand

		if err := rows.Scan(&pb.Name, &pb.Id, &pb.Price, &pb.BrandId, &pb.BrandName); err != nil {
			fmt.Println(err)
			resp := helpers.CustomeRes{
				Msg:  "data not found",
				Err:  err,
				Code: http.StatusOK,
				Data: result,
			}
			helpers.SendResponse(w, r, resp, http.StatusOK)
			return
		}

		result = append(result, pb)
	}
	resp := helpers.CustomeRes{
		Msg:  "successfully get product by brand",
		Err:  nil,
		Code: http.StatusOK,
		Data: result,
	}
	helpers.SendResponse(w, r, resp, http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request, data models.Product) {
	fmt.Println(data)
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

	// save data to db
	p := models.Product{
		BrandId: req.BrandId,
		Name:    req.Name,
		Price:   req.Price,
	}

	err = helpers.ValidationProduct(w, r, models.JSONProduct{
		BrandId: req.BrandId,
		Name:    req.Name,
		Price:   req.Price,
	})
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

	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}

	sqlStatement := `insert into product (brand_id, name, price) values ($1,$2,$3)`
	res, err := db.Exec(sqlStatement, p.BrandId, p.Name, p.Price)
	if err != nil {
		fmt.Println("Error: cannot save new product in DB ", err)
		helpers.SendResponse(w, r, nil, http.StatusInternalServerError)
	}
	res.LastInsertId()

	response := mapResToJSON(&p)
	resp := helpers.CustomeRes{
		Msg:  "successfully get product by brand",
		Err:  nil,
		Code: http.StatusOK,
		Data: response,
	}
	helpers.SendResponse(w, r, resp, http.StatusOK)
}
