package models

import "time"

type EntityUser struct {
	ID        string
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
