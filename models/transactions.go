package models

import "time"

type Transaction struct {
	Id         int       `json:"id"`
	GrandTotal int       `json:"grand_total"`
	BuyerName  string    `json:"buyer_name"`
	CreatedAt  time.Time `json:"created_at"`
}

type JSONTransaction struct {
	Id         int       `json:"id"`
	GrandTotal int       `json:"grand_total"`
	BuyerName  string    `json:"buyer_name"`
	CreatedAt  time.Time `json:"created_at"`
}

type TransactionProduct struct {
	Id            int `json:"id"`
	TransactionId int `json:"transaction_id"`
	ProductId     int `json:"product_id"`
}

type TransactionOrder struct {
	BuyerName string                   `json:"buyer_name"`
	Product   []JSONProductTransaction `json:"products"`
}

type JSONTransactionProductDetail struct {
	Id         int                      `json:"id"`
	GrandTotal int                      `json:"grand_total"`
	BuyerName  string                   `json:"buyer_name"`
	CreatedAt  time.Time                `json:"created_at"`
	Product    []JSONProductTransaction `json:"products"`
}

type JSONProductTransaction struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	BrandId   int    `json:"brand_id"`
	BrandName string `json:"brand_name"`
	Quantity  int    `json:"qty"`
}
