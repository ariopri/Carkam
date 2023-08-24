package models

import "time"

type EntityFakultas struct {
	ID           string
	NamaFakultas string
	Deskripsi    string
	KampusID     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
