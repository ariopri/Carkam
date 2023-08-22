package models

import "time"

type EntityJurusan struct {
	ID          string
	NamaJurusan string
	Deskripsi   string
	Akreditasi  string
	FakultasID  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
