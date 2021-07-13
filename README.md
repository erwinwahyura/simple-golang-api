# Simple Golang API

## How To

- create simple api golang that can do get and create
- there is 4 tables in this case, `brand`, `product`, `transactions` and `transactionsproduct`

the schema look like

`brand`

- name string

`product`

- name string
- price int
- brand_id int

`transactions`

- buyer_name string
- grand_total int
- created_at string

`transactionsproduct`

- product_id int
- transactions_id int
- qty int

---

## 1. setup database migration

- create your database first, sample: `sample_db`, run migration file in this case we use postgresql

```go
  $ migrate --path ./migrations --database "postgresql://erwinwahyuramadhan:@localhost:5432/sample_db?sslmode=disable" --verbose up
```

- details

```json
  user: erwinwahyuramadhan
  password: 
  host: localhost
  port: 5432
  dbname: sample
```

## 2. run unit testing

- run all testing file in the directory

```go
  $ go test -v ./...
```

## 3. api docs

- api using `localhost:4000`
- detail api `brand`

| EndPoint      | Path          | Method| Desc
| ------------- |:-------------:|:-----:| --------  
| Get Brand     | /brand        | GET   | get all brands
| Create Brand  | /brand        |  POST | create a new brand

- detail api `product`

| EndPoint      | Path          | Method| Desc
| ------------- |:-------------|:-----:| --------  
| Get Product By Id     | /product?id=1        | GET   | get single product by query product `id`
| Create Product  | /product        |  POST | create a new product
| Get All Product by Brand  | /product/brand?id=1        |  POST | get all data product by brand by query brand `id`

- detail api `transaction`

| EndPoint        | Path            | Method| Desc
| -------------   |:-------------   |:-----:| --------  
| Create Trx      | /order          | POST  | create trx order with multiple products in the cart
| Get Detail Trx  | /order/trx?id=1 | GET   | get detail transaction by `transaction id` including detail products

