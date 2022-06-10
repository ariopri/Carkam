package repository

import (
	"errors"
	"strconv"

	"github.com/github.com/rg-km/final-project-engineering-66/db"
)

type UserRepository struct {
	db db.DB
}

const userDbName = "users"

var userColumns = []string{"username", "password", "loggedin", "role"}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	records, err := u.db.Load(userDbName)
	if err != nil {
		records = [][]string{userColumns}
		if err := u.db.Save(userDbName, records); err != nil {
			return nil, err
		}
	}

	result := make([]User, 0)
	for i := 1; i < len(records); i++ {
		loggedin, err := strconv.ParseBool(records[i][2])
		if err != nil {
			return nil, err
		}

		user := User{
			Username: records[i][0],
			Password: records[i][1],
			Role:     records[i][3],
			Loggedin: loggedin,
		}
		result = append(result, user)
	}

	return result, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	return u.LoadOrCreate()
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	var loggedInUser *string

	users, err := u.LoadOrCreate()
	if err != nil {
		return loggedInUser, err
	}

	_, user := searchUserByUsername(users, username)
	if user == nil || user.Password != password {
		return loggedInUser, errors.New("Login Failed")
	}

	err = u.changeStatus(users, username, true, "Login Failed")
	if err != nil {
		return loggedInUser, err
	}

	loggedInUser = &username
	return loggedInUser, nil
}

func (u *UserRepository) Save(users []User) error {
	records := [][]string{userColumns}
	for i := 0; i < len(users); i++ {
		records = append(records, []string{
			users[i].Username,
			users[i].Password,
			strconv.FormatBool(users[i].Loggedin),
			users[i].Role,
		})
	}
	return u.db.Save(userDbName, records)
}

func (u *UserRepository) GetUserRole(username string) (*string, error) {
	var role *string

	// load users
	users, err := u.LoadOrCreate()
	if err != nil {
		return role, err
	}

	// search by username
	_, user := searchUserByUsername(users, username)
	if user == nil {
		return role, errors.New("user not found")
	}

	return &user.Role, nil
}

func searchUserByUsername(users []User, username string) (int, *User) {
	for i, user := range users {
		if user.Username == username {
			return i, &user
		}
	}
	return -1, nil
}

func (u *UserRepository) changeStatus(users []User, username string, status bool, errMsg string) error {
	i, user := searchUserByUsername(users, username)
	if user == nil {
		return errors.New(errMsg)
	}

	user.Loggedin = status
	users[i] = *user

	return u.Save(users)
}
