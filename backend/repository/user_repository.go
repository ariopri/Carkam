package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rg-km/final-project-engineering-66/models/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.EntityUser) (domain.EntityUser, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.EntityUser) (domain.EntityUser, error)
	Delete(ctx context.Context, tx *sql.Tx, user domain.EntityUser)
	FindByID(ctx context.Context, tx *sql.Tx, userId string) (domain.EntityUser, error)
	FindByUserName(ctx context.Context, tx *sql.Tx, username string) (domain.EntityUser, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.EntityUser, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.EntityUser, error)
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUsersRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.EntityUser) (domain.EntityUser, error) {
	//TODO implement method
	var username, email, phone string
	var usernameArr, emailArr, phoneArr []string

	rowsCheck, err := tx.QueryContext(ctx, "SELECT username, email, phone FROM users")
	if err != nil {
		return domain.EntityUser{}, err
	}

	for rowsCheck.Next() {
		err := rowsCheck.Scan(&username, &email, &phone)
		if err != nil {
			return domain.EntityUser{}, err
		}

		usernameArr = append(usernameArr, username)
		emailArr = append(emailArr, email)
		phoneArr = append(phoneArr, phone)
	}

	for _, v := range usernameArr {
		if v == user.UserName {
			return domain.EntityUser{}, errors.New("username already exist")
		}
	}

	for _, v := range emailArr {
		if v == user.Email {
			return domain.EntityUser{}, errors.New("email already exist")
		}
	}

	for _, v := range phoneArr {
		if v == user.Phone {
			return domain.EntityUser{}, errors.New("phone already exist")
		}
	}

	SQL := "INSERT INTO users (id, firstname, lastname, username, email, phone, password, kota, avatar) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err = tx.ExecContext(
		ctx,
		SQL,
		user.ID,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Email,
		user.Phone,
		user.Password,
		user.Kota,
		user.Avatar,
	)

	if err != nil {
		return domain.EntityUser{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.EntityUser) (domain.EntityUser, error) {
	//TODO implement me
	SQl := "UPDATE users SET firstname = $1, lastname = $2, username = $3, email = $4, phone = $5, update_at = $6 WHERE id = $7"
	_, err := tx.ExecContext(
		ctx,
		SQl,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Email,
		user.Phone,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		return domain.EntityUser{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.EntityUser) {
	//TODO implement me
	SQL := "DELETE FROM users WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, user.ID)
	if err != nil {
		panic(err)
	}
}

func (repository *UserRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, userId string) (domain.EntityUser, error) {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username, email, phone, password, role, kota, avatar, created_at, update_at FROM users WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		return domain.EntityUser{}, err
	}
	defer rows.Close()

	var user domain.EntityUser
	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Role,
			&user.Kota,
			&user.Avatar,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return domain.EntityUser{}, err
		}
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByUserName(ctx context.Context, tx *sql.Tx, username string) (domain.EntityUser, error) {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username, email, role, created_at, updated_at FROM users WHERE username = $1"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
		username,
	)

	if err != nil {
		return domain.EntityUser{}, err
	}

	defer rows.Close()

	var user domain.EntityUser

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Email,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return domain.EntityUser{}, err
		}
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.EntityUser, error) {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username, email, phone, password, role, kota, avatar, created_at, update_at FROM users WHERE email = $1"
	rows, err := tx.QueryContext(ctx, SQL, email)
	if err != nil {
		return domain.EntityUser{}, err
	}

	defer rows.Close()

	var user domain.EntityUser
	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Role,
			&user.Kota,
			&user.Avatar,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return domain.EntityUser{}, err
		}
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.EntityUser, error) {
	//TODO implement me
	SQL := "SELECT id, firstname, lastname, username, email, phone, password, role, kota, avatar, created_at, update_at FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return []domain.EntityUser{}, err
	}

	defer rows.Close()

	var users []domain.EntityUser

	for rows.Next() {
		var user domain.EntityUser
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UserName,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Role,
			&user.Kota,
			&user.Avatar,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return []domain.EntityUser{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
