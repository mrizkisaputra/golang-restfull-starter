-- migrate -path=db/migrations -database="mysql://root:root@tcp(localhost:3306)/belajar" up

CREATE TABLE products(
                         id char(100) not null,
                         item varchar(200) not null,
                         price int not null,
                         quantity int not null
);

ALTER TABLE products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);