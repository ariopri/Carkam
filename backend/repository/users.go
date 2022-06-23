package repository

import (
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (User, error) {

	var user User
	err := u.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil

}

func (u *UserRepository) FetchUsers() ([]User, error) {

	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

func (u *UserRepository) Login(username string, password string) (*string, error) {

	var user User
	err := u.db.QueryRow("SELECT username FROM users WHERE username = ? AND password = ?", username, password).Scan(&user.Username)
	if err != nil {
		return nil, errors.New("Login Failed")
	}
	return &user.Username, nil

}

func (u *UserRepository) InsertUser(username string, email string, password string) error {

	_, err := u.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, password)
	if err != nil {
		return err
	}
	return nil

}
