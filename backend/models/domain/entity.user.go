package models

import "time"

type EntityUser struct {
	ID        string
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Phone     string
	Password  string
	Role      string
	Kota      *string
	Avatar    *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
