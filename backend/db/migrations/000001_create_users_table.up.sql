create table users(
    id uuid primary key,
    firstname varchar(50) not null,
    lastname varchar(50) not null,
    username varchar(50) not null unique,
    phone varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(255) not null,
    role varchar(50) default 'student',
    kota varchar(50),
    avatar varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

