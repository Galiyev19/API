CREATE TABLE admins (
    id bigserial not null primary key,
    email varchar(255) not null unique,
    encrypted_password varchar(255) not null,
    created_at timestamp not null,
    role varchar(255) not null
);

CREATE TABLE users (
    id bigserial not null primary key,
    username varchar(255) not null,
    email varchar(255) not null unique,
    encrypted_password varchar(255) not null,
    created_at timestamp not null
);
