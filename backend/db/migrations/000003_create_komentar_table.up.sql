create table komentar(
    id uuid primary key,
    rating int not null,
    content text not null,
    users_id uuid not null,
    foreign key (users_id) references users(id),
    creted_at timestamp default now(),
    updated_at timestamp default now()
);