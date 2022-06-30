package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../../carkam.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT(255) NOT NULL,
			email TEXT(255) NOT NULL,
			password TEXT(255) NOT NULL,
			role TEXT(255) NOT NULL,
			loggedin BOOLEAN NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL

		);

		CREATE TABLE IF NOT EXISTS kampus (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			nama TEXT(255) NOT NULL,
			alamat TEXT(255) NOT NULL,
			kota TEXT(255) NOT NULL,
			provinsi TEXT(255) NOT NULL,
			kode_pos TEXT(255) NOT NULL,
			no_telp TEXT(255) NOT NULL,
			email TEXT(255) NOT NULL,
			website TEXT(255) NOT NULL,
			logo TEXT(255) NOT NULL,
			akreditasi CHAR(1) NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS fakultas (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			kampus_id INTEGER NOT NULL,
			nama TEXT(255) NOT NULL,
			akreditasi_fakultas CHAR(1) NOT NULL,
			rating_fakultas INTEGER NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (kampus_id) REFERENCES kampus(id)
		);

		CREATE TABLE IF NOT EXISTS jurusan (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			kampus_id INTEGER NOT NULL,
			fakultas_id INTEGER NOT NULL,
			nama TEXT(255) NOT NULL,
			akreditasi_jurusan CHAR(1) NOT NULL,
			rating_jurusan INTEGER NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (kampus_id) REFERENCES kampus(id),
			FOREIGN KEY (fakultas_id) REFERENCES fakultas(id)
		);

		CREATE TABLE IF NOT EXISTS review (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			kampus_id INTEGER NOT NULL,
			fakultas_id INTEGER NOT NULL,
			jurusan_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			review TEXT(255) NOT NULL,
			rating INTEGER NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			FOREIGN KEY (kampus_id) REFERENCES kampus(id),
			FOREIGN KEY (fakultas_id) REFERENCES fakultas(id),
			FOREIGN KEY (jurusan_id) REFERENCES jurusan(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);


		INSERT INTO users (username, email, password, role, loggedin, created_at, updated_at) VALUES
		 ("admin", "admin@gmail.com", "admin", "admin", "false", "2020-01-01 00:00:00", "2020-01-01 00:00:00"),
		 ("user", "user@gmail.com", "user", "user", "false", "2020-01-01 00:00:00", "2020-01-01 00:00:00");
		
		INSERT INTO kampus (user_id, nama, alamat, kota, provinsi, kode_pos, no_telp, email, website, logo, akreditasi, created_at, updated_at) VALUES
		(1, "Universitas Brawijaya", "Jl. Veteran No. 1", "Malang", "Jawa Timur", "65141", "08123456789", "infoUb@gmail.com", "www.ub.ac.id", "ub.png", "A", "2020-01-01 00:00:00", "2020-01-01 00:00:00");


		INSERT INTO fakultas (user_id, kampus_id, nama, akreditasi_fakultas, rating_fakultas, created_at, updated_at) VALUES
		(1, 1, "Fakultas Ilmu Komputer", "A", 5, "2020-01-01 00:00:00", "2020-01-01 00:00:00"),
		(1, 1, "Fakultas TEKNIK", "A", 5, "2020-01-01 00:00:00", "2020-01-01 00:00:00");

		INSERT INTO jurusan (user_id, kampus_id, fakultas_id, nama, akreditasi_jurusan, rating_jurusan, created_at, updated_at) VALUES
		(1, 1, 1, "Teknik Informatika", "A", 5, "2020-01-01 00:00:00", "2020-01-01 00:00:00"),
		(1, 1, 1, "Sistem Informasi", "A", 5, "2020-01-01 00:00:00", "2020-01-01 00:00:00"),
		(1, 1, 2, "Teknik Elektro", "A", 5, "2020-01-01 00:00:00", "2020-01-01 00:00:00"),
		(1, 1, 2, "Teknik Mesin", "A", 5, "2020-01-01 00:00:00", "2020-01-01 00:00:00");

		INSERT INTO review (kampus_id, fakultas_id, jurusan_id, user_id, review, rating, created_at, updated_at) VALUES
		(1, 1, 1, 1, "Sangat bagus", 5, "2020-01-01 00:00:00", "2020-01-01 00:00:00"),
		(1, 1, 1, 2, "Sangat bagus", 5, "2020-01-01 00:00:00", "2020-01-01 00:00:00"),
		(1, 1, 2, 1, "Sangat bagus", 5, "2020-01-01 00:00:00", "2020-01-01 00:00:00"),
		(1, 1, 2, 2, "Sangat bagus", 5, "2020-01-01 00:00:00", "2020-01-01 00:00:00");
	`)

	if err != nil {
		panic(err)
	}
}
