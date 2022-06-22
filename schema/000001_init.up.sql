CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null,
    password varchar(255) not null
);

CREATE TABLE books
(
    id serial not null unique,
    name varchar(255) not null,
    author varchar(255) not null
);