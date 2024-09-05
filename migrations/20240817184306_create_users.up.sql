CREATE TABLE users (
    id bigserial not null primary key,
    name varchar not null unique,
    fullname varchar not null,
    email varchar not null unique,
    password_hash varchar not null
);