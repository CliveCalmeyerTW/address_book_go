CREATE TABLE address_book (
  id          serial,
  first_name  varchar(50),
  last_name   varchar(50),
  address_1   varchar(100),
  address_2   varchar(100),
  city        varchar(20),
  postcode    varchar(8),
  email       varchar(100),
  telephone   varchar(20)
);