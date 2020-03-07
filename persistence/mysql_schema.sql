CREATE DATABASE IF NOT EXISTS urlshortener;

USE urlshortener;

CREATE TABLE short_url (
  id int auto_increment primary key,
  short_key varchar(10),
  long_url varchar(2048)
);
