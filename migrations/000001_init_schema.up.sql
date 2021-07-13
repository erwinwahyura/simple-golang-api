CREATE TABLE Brand (
	id serial NOT NULL primary key,
	name VARCHAR(255) NOT NULL
);
CREATE TABLE Product (
	id serial NOT NULL primary key,
	price int not null,
	brand_id int,
	name VARCHAR(255) NOT NULL,
	foreign key (brand_id) references Brand (id)
);


CREATE TABLE Transactions (
	id serial NOT NULL primary key,
	grand_total int,
	buyer_name VARCHAR(255) NOT NULL,
	created_at date
);

CREATE TABLE TransactionsProduct (
	id serial NOT NULL primary key,
	product_id int,
	transaction_id int,
  qty int,
	foreign key (product_id) references Product (id),
	foreign key (transaction_id) references Transactions (id)
);
