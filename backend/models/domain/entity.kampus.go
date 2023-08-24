package models

import "time"

type EntityKampus struct {
	ID         string
	NamaKampus string
	Alamat     string
	Telepon    string
	Email      string
	Website    string
	Deskripsi  string
	Logo       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
