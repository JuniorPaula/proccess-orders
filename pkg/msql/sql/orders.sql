CREATE DATABASE IF NOT EXISTS orders;
USE orders;

CREATE TABLE orders(
    id varchar(50) primary key,
    price float,
    tax float,
    final_price float
)