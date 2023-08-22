package models

import "time"

type EntityStudent struct {
	ID        string
	FirstName string
	LastName  string
	Username  string
	Phone     string
	Email     string
	Password  string
	Role      string
	Kota      string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
