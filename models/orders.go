package models

type Orders struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id"`
	BuyerName string `json:"buyer_name"`
	CreatedAt int    `json:"created_at"`
}

type JSONOrders struct {
	Id        int    `json:"id"`
	ProductId int    `json:"product_id"`
	BuyerName string `json:"buyer_name"`
	CreatedAt int    `json:"created_at"`
}

type JSONOrdersProduct struct {
	Id          int    `json:"id"`
	ProductId   int    `json:"product_id"`
	ProductName int    `json:"product_name"`
	BuyerName   string `json:"buyer_name"`
	BrandId     int    `json:"brand_id"`
	BrandName   string `json:"brand_name"`
	CreatedAt   string `json:"created_at"`
	Price       int    `json:"price"`
}
