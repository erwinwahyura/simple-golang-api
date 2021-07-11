package models

type Product struct {
	Id      int    `json:"id"`
	Price   int    `json:"price"`
	Name    string `json:"name"`
	BrandId int    `json:"brand_id"`
}

type JSONProduct struct {
	Id      int    `json:"id"`
	Price   int    `json:"price"`
	Name    string `json:"name"`
	BrandId int    `json:"brand_id"`
}

type JSONProductBrand struct {
	Name      string `json:"name"`
	Id        int    `json:"id"`
	Price     int    `json:"price"`
	BrandId   int    `json:"brand_id"`
	BrandName string `json:"brand_name"`
}
