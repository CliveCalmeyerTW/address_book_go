CREATE TABLE address_book (
  id          serial NOT NULL PRIMARY KEY,
  first_name  varchar(50) NOT NULL,
  last_name   varchar(50) NOT NULL,
  address_1   varchar(100) NOT NULL,
  address_2   varchar(100) NOT NULL DEFAULT '',
  city        varchar(20) NOT NULL,
  postcode    varchar(8) NOT NULL,
  email       varchar(100) NOT NULL DEFAULT '',
  telephone   varchar(20) NOT NULL DEFAULT ''
);