package models

import "time"

type EntityUser struct {
	ID        string
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
