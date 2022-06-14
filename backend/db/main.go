package main

import (
	"database/sql"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Admin (
			id integer not null primary key AutoIncrement,
			username varchar(255) not null,
			password varchar(255) not null,
			role varchar(255) not null
			loggedin boolean not null
		);
		CREATE TABLE IF NOT EXISTS users (
			id integer not null primary key AutoIncrement,
			username varchar(255) not null,
			password varchar(255) not null,
			role varhar(255) not null,
			loggedin boolean not null
		);
		CREATE TABLE IF NOT EXISTS products (
			id integer not null primary key AutoIncrement,
			kampus_name varchar(255) not null,
			kampus_address varchar(255) not null,
			kampus_phone varchar(255) not null,
			kampus_email varchar(255) not null,
			kampus_website varchar(255) not null,
			kampus_akreditasi varchar(255) not null,
			kampus_image varchar(255) not nul
		);
		CREATE TABLE IF NOT EXISTS fakultas (
			id integer not null primary key AutoIncrement,
			fakultas_name varchar(255) not null,
			fakultas_akreditasi varchar(255) not null
		);
		CREATE TABLE IF NOT EXISTS jurusan (
			id integer not null primary key AutoIncrement,
			jurusan_name varchar(255) not null,
			jurusan_akreditasi varchar(255) not null,
		);

		INSERT INTO Admin (username, password, role, loggedin) VALUES ("admin", "admin", "admin", 1);
		INSERT INTO users (username, password, role, loggedin) VALUES ("user", "user", "user", 1);
		INSERT INTO products (kampus_name, kampus_address, kampus_phone, kampus_email, kampus_website,
			kampus_akreditasi, kampus_image) VALUES
			("Universitas Indonesia", "Jl. Veteran No.1", "081212121212", "",
			"", "", "", "");
		INSERT INTO fakultas (fakultas_name, fakultas_akreditasi) VALUES
			("Fakultas Ilmu Komputer", "A"),
			("Fakultas Teknik", "A"),
			("Fakultas Ekonomi dan Bisnis", "A"),
			("Fakultas Hukum", "A"),
			("Fakultas Sastra", "A"),
			("Fakultas Kedokteran", "A"),
			("Fakultas Kesehatan Masyarakat", "A"),
			("Fakultas Psikologi", "A"),
			("Fakultas Farmasi", "A"),
			("Fakultas Kedokteran Gigi", "A"),
			("Fakultas Ilmu Administrasi", "A"),
			("Fakultas Ilmu Sosial Ilmu Politik", "A"),
			("Fakultas Matematika dan Ilmu Pengetahuan Alam", "A");

		INSERT INTO jurusan (jurusan_name, jurusan_akreditasi) VALUES
			("S1 Teknik Informatika", "A"),
			("S1 Teknik Elektro", "A"),
			("S1 Teknik Mesin", "A"),
			("S1 Teknik Sipil", "A"),
			("S1 Teknik Lingkungan", "A"),
			("S1 Teknik Kimia", "A"),
			("S1 Teknik Industri", "A"),
			("S1 Teknik Sipil", "A"),
			("S1 Kedokteran", "A"),
			("S1 Farmasi", "A"),
			("S1 Kedokteran Gigi", "A"),
			("S1 Ilmu Administrasi", "A"),
			("S1 Ilmu Sosial Ilmu Politik", "A"),
			("S1 Matematika", "A");
			("S1 Statistika", "A"),
			("S1 Ilmu Komputer", "A"),
			("S1 Ilmu Pengetahuan Alam", "A"),
			("S1 Ilmu Sosial", "A"),
			("S1 Ilmu Politik", "A"),
			("S1 Hukum", "A"),
			("S1 Sastra Jepang", "A");
		
		`)
	if err != nil {
		panic(err)
	}
}
