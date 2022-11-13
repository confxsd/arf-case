CREATE TABLE users (
    id bigserial PRIMARY KEY,
    username VARCHAR (32) UNIQUE NOT NULL,
    password VARCHAR (32) NOT NULL
);