package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../../carkams.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id integer not null primary key AUTOINCREMENT,
			username varchar(255) not null,
			email varchar(255) not null,
			password varchar(255) not null
		);
		CREATE TABLE IF NOT EXISTS kampus (
			id integer not null primary key AUTOINCREMENT,
			name varchar(255) not null,
			email varchar(255) not null,
			jurusan1 varchar(255) not null,
			jurusan2 varchar(255) not null
		);	
		CREATE TABLE IF NOT EXISTS jurusan (
			id integer not null primary key AUTOINCREMENT,
			name varchar(255) not null,
		);
		CREATE TABLE IF NOT EXISTS review (
			id integer not null primary key AUTOINCREMENT,
			username varchar(255) not null,
			kampus_name strnig not null,
			jurusan_name string not null,
			isian varhar(255) not null
		);
		Insert into users (username, email, password) values ("user","user","user" );
		Insert into kampus (name, email, jurusan1, jurusan2) values ("Universitas Indonesia", "humas-ui@ui.ac.id", "Ilmu Komputer", "Pendidikan Dokter"),
		("Institut Teknologi Bandung", "info-center@itb.ac.id", "Teknik Informatika", "Teknik Elektro"),
		("Universitas Gadjah" , "info@ugm.ac.id", "Ilmu Komputer", "Teknik Nuklir"), 
		("Universitas Padjajaran", "humas@unpad.ac.id", "Psikologi", "Pendidikan Dokter"),
		("Institut Teknologi Sepuluh November", "humas@its.ac.id", "Teknik Elektro", "Ilmu Komputer"),
		("Institut Pertanian Bogor", "humas@ipb.ac.id", "Kehutanan", "Peternakan"),
		("Universitas Diponegoro", "humas@live.undip.ac.id", "Psikologi", "Ilmu Hukum"),
		("Universitas Negeri Yogyakarta", "humas@uny.ac.id", "Teknik Sipil", "Teknik Elektro"),
		("Universitas Negeri Jakarta", "humas@unj.ac.id", "Matematika", "Statistika"),
		("Universitas Pendidikan Indonesia", "sekuniv_upi@upi.edu", "Teknik Elektro", "Teknik Sipil");

		Insert into review (username, kampus_name, jurusan_name, isian) values ("user", "Universitas Indonesia", "Ilmu Komputer", "Isian");
	`)
	if err != nil {
		panic(err)
	}
}
