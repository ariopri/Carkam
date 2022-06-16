package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id integer not null primary key AUTOINCREMENT,
			name varchar(255) not null,
			email varchar(255) not null,
			password varchar(255) not null
		)
		CREATE TABLE IF NOT EXISTS kampus (
			id integer not null primary key AUTOINCREMENT,
			name varchar(255) not null,
			email varchar(255) not null,
			info varchar(255) not null,
			jurusan1 varchar(255) not null,
			jurusan2 varchar(255) not null
		)	

		CREATE TABLE IF NOT EXISTS review (
			id integer not null primary key AUTOINCREMENT,
			username varchar(255) not null,
			email varchar(255) not null,
			nama_kampus varchar(255) not null,
			jurusan varchar(255) not null,
			isian longtext not null
		)
		Insert into users (username, email, password) values (?, ?, ?)
		Insert into kampus (name, email, info, jurusan1, jurusan2) values (?, ?, ?, ?, ?)
		Insert into review (username, email, nama_kampus, jurusan, isian) values (?, ?, ?, ?, ?)

	
		select * from review where nama_kampus = 'Universitas Indonesia' and jurusan = 'Teknik'
	`)
	if err != nil {
		panic(err)
	}

}
