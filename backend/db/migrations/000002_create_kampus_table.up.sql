create table kampus(
                       id uuid primary key,
                       nama_kampus varchar(50) not null,
                       alamat varchar(255) not null,
                       telepon varchar(50) not null,
                       email varchar(50) not null,
                       website varchar(255) not null,
                       deskripsi text,
                       logo_kampus varchar(255),
                       created_at timestamp default current_timestamp,
                       updated_at timestamp default current_timestamp
);

create table fakultas(
                         id uuid primary key,
                         nama_fakultas varchar(50) not null,
                         deskripsi text,
                         id_kampus uuid references kampus(id),
                         foreign key (id_kampus) references kampus(id),
                         created_at timestamp default current_timestamp,
                         updated_at timestamp default current_timestamp
);

create table jurusan(
                        id uuid primary key,
                        nama_jurusan varchar(50) not null,
                        deskripsi text,
                        akreditasi varchar(50),
                        id_fakultas uuid references fakultas(id),
                        foreign key (id_fakultas) references fakultas(id),
                        created_at timestamp default current_timestamp,
                        updated_at timestamp default current_timestamp
);

create table biaya(
                      id uuid primary key,
                      tahun_akademik varchar(50),
                      nominal decimal(10,2),
                      id_jurusan uuid references jurusan(id),
                      foreign key (id_jurusan) references jurusan(id),
                      created_at timestamp default current_timestamp,
                      updated_at timestamp default current_timestamp
);