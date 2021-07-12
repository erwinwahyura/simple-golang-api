package api

import (
	"fmt"
	"net/http"
	"simple-golang-api/config"
	"simple-golang-api/helpers"
	"simple-golang-api/models"
	"time"
)

func mapTrxDetailToJSON(t *models.Transaction) models.JSONTransaction {
	return models.JSONTransaction{
		Id:         t.Id,
		GrandTotal: t.GrandTotal,
		BuyerName:  t.BuyerName,
		CreatedAt:  t.CreatedAt,
	}
}

func GetTrxById(w http.ResponseWriter, r *http.Request) {
	fmt.Print("DISNI KAN??")
	queryId := r.URL.Query().Get("id")

	if len(queryId) == 0 {
		resp := helpers.CustomeRes{
			Msg:  "Error: url param 'id' is missing",
			Err:  nil,
			Code: http.StatusBadRequest,
			Data: nil,
		}

		helpers.SendResponse(w, r, resp, http.StatusBadRequest)
	}

	req := models.Transaction{}

	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
		msg := helpers.CustomeRes{
			Msg:  "Error connecting to db",
			Err:  err,
			Code: http.StatusBadRequest,
		}
		helpers.SendResponse(w, r, msg, http.StatusBadRequest)
	}
	defer db.Close()

	trx := models.Transaction{
		Id:         req.Id,
		GrandTotal: req.GrandTotal,
		BuyerName:  req.BuyerName,
		CreatedAt:  req.CreatedAt,
	}

	fmt.Println(queryId)
	sqlStatementTrx := `select * from transactions where id = $1`
	err = db.QueryRow(sqlStatementTrx, queryId).Scan(&trx.Id, &trx.GrandTotal, &trx.BuyerName, &trx.CreatedAt)
	if err != nil {
		fmt.Println(err)
		msg := helpers.CustomeRes{
			Msg:  "Error: cannot get detail trx by id ",
			Err:  err,
			Code: http.StatusBadRequest,
		}
		helpers.SendResponse(w, r, msg, http.StatusInternalServerError)
	}

	responseDetailTrx := mapTrxDetailToJSON(&trx)

	sqlStatementProducts := `
		select p.name as product_name, p.id as product_id, p.price, p.brand_id, b.name as brand_name, tp.qty 
		from transactionsproduct tp
		inner join product p on p.id = tp.product_id
		inner join transactions t on t.id = tp.transaction_id
		inner join brand b on b.id = p.brand_id
		where tp.transaction_id = $1
		`
	rows, err := db.Query(sqlStatementProducts, queryId)
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
	var result []models.JSONProductTransaction

	for rows.Next() {
		var pb models.JSONProductTransaction

		if err := rows.Scan(&pb.Name, &pb.Id, &pb.Price, &pb.BrandId, &pb.BrandName, &pb.Quantity); err != nil {
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

	tp := models.JSONTransactionProductDetail{
		Id:         responseDetailTrx.Id,
		CreatedAt:  responseDetailTrx.CreatedAt,
		BuyerName:  responseDetailTrx.BuyerName,
		GrandTotal: responseDetailTrx.GrandTotal,
		Product:    result,
	}

	resp := helpers.CustomeRes{
		Msg:  "successfully get detail trx by id",
		Err:  nil,
		Code: http.StatusOK,
		Data: tp,
	}
	helpers.SendResponse(w, r, resp, http.StatusOK)
}

func CreateTrx(w http.ResponseWriter, r *http.Request, data models.TransactionOrder) {
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

	p := models.TransactionOrder{
		BuyerName: req.BuyerName,
		Product:   req.Product,
	}

	var grandTotal int
	for _, v := range p.Product {
		grandTotal += v.Price * v.Quantity
	}

	// create trx
	trx := models.Transaction{
		GrandTotal: grandTotal,
		BuyerName:  p.BuyerName,
		CreatedAt:  time.Now(),
	}

	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	var id int
	// trx sql statement
	sqlTrxStatement := `insert into transactions (grand_total, buyer_name, created_at) values ($1, $2, $3) RETURNING id`
	err = db.QueryRow(sqlTrxStatement, trx.GrandTotal, trx.BuyerName, trx.CreatedAt).Scan(&id)

	if err != nil {
		fmt.Println("Error: cannot save new product in DB ", err)
		resp := helpers.CustomeRes{
			Msg:  "Error: cannot save transaction",
			Err:  err,
			Code: http.StatusInternalServerError,
			Data: p,
		}
		helpers.SendResponse(w, r, resp, http.StatusInternalServerError)
		return
	}

	resWrapper := mapTrxDetailToJSON(&trx)

	// trx sql product statement
	sqlTrxProduct := `insert into transactionsproduct (product_id, transaction_id, qty) values ($1, $2, $3)`

	for _, v := range p.Product {
		_, err := db.Exec(sqlTrxProduct, v.Id, id, v.Quantity)

		if err != nil {
			panic(err)
		}

		fmt.Println("done")
	}

	finalRes := models.JSONTransactionProductDetail{
		Id:         resWrapper.Id,
		GrandTotal: resWrapper.GrandTotal,
		BuyerName:  resWrapper.BuyerName,
		CreatedAt:  resWrapper.CreatedAt,
		Product:    p.Product,
	}

	resp := helpers.CustomeRes{
		Msg:  "Successfully",
		Err:  nil,
		Code: http.StatusOK,
		Data: finalRes,
	}
	helpers.SendResponse(w, r, resp, http.StatusOK)
}
