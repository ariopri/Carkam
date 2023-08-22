package models

import "time"

type EntityKomentar struct {
	ID        string
	Rating    string
	Content   string
	UsersID   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
