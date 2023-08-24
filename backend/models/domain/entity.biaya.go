package domain

import "time"

type EntityBiaya struct {
	ID            string
	TahunAkademik string
	Nominal       string
	JurusanID     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
