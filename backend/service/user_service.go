package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/rg-km/final-project-engineering-66/config"
	"github.com/rg-km/final-project-engineering-66/helper"
	"github.com/rg-km/final-project-engineering-66/models/domain"
	"github.com/rg-km/final-project-engineering-66/models/web"
	"github.com/rg-km/final-project-engineering-66/repository"
	"github.com/rg-km/final-project-engineering-66/utils"
	"time"
)

type UserService interface {
	Register(ctx context.Context, request web.RegisterRequest) (web.RegisterResponse, error)
	Login(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error)
	UpdateUser(ctx context.Context, request web.UpdateUserRequest) (web.UserDetailResponse, error)
	DeleteUser(ctx context.Context, userId string)
	FindByID(ctx context.Context, userId string) (web.UserDetailResponse, error)
	FindByUserName(ctx context.Context, username string) (web.UserDetailResponse, error)
	FindByEmail(ctx context.Context, email string) (web.UserDetailResponse, error)
	FindAll(ctx context.Context) ([]web.UserDetailResponse, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.RegisterRequest) (web.RegisterResponse, error) {
	//TODO implement me
	var response web.RegisterResponse

	err := service.Validate.Struct(request)
	if err != nil {
		return web.RegisterResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.RegisterResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	request.CreatedAt = time.Now()
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return web.RegisterResponse{}, err
	}

	user := domain.EntityUser{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		UserName:  request.UserName,
		Email:     request.Email,
		Phone:     request.Phone,
		Password:  hashedPassword,
		CreatedAt: request.CreatedAt,
	}
	user, err = service.UserRepository.Save(ctx, tx, user)
	if err != nil {
		return web.RegisterResponse{}, err
	}

	response = web.RegisterResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) (web.LoginResponse, error) {
	//TODO implement me
	var response web.LoginResponse

	err := service.Validate.Struct(request)
	if err != nil {
		return web.LoginResponse{}, err
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return web.LoginResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUserName(ctx, tx, request.UserName)
	if err != nil {
		return web.LoginResponse{}, err
	}

	loadConfig, _ := config.LoadConfig()
	verifyError := utils.CheckPasswordHash(user.Password, request.Password)
	if verifyError != false {
		return web.LoginResponse{}, err
	}

	token, errToken := utils.GenerateToken(loadConfig.TokenExpiresIn, user.ID, loadConfig.TokenSecret)
	if errToken != nil {
		return web.LoginResponse{}, errToken
	}

	response = web.LoginResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		TokenType: "Bearer",
		Token:     token,
	}

	return response, nil
}

func (service *UserServiceImpl) UpdateUser(ctx context.Context, request web.UpdateUserRequest) (web.UserDetailResponse, error) {
	//TODO implement me
	var response web.UserDetailResponse
	err := service.Validate.Struct(request)

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserDetailResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByID(ctx, tx, request.ID)
	if err != nil {
		return web.UserDetailResponse{}, err
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.UserName = request.UserName
	user.Email = request.Email
	user.Phone = request.Phone
	user.Kota = &request.Kota
	user.Avatar = &request.Avatar
	user.UpdatedAt = request.UpdateAt

	user, err = service.UserRepository.Update(ctx, tx, user)
	if err != nil {
		return web.UserDetailResponse{}, err
	}

	response = web.UserDetailResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
		Kota:      *user.Kota,
		Avatar:    *user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (service *UserServiceImpl) DeleteUser(ctx context.Context, userId string) {
	//TODO implement me
	tx, err := service.DB.Begin()
	if err != nil {
		return
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByID(ctx, tx, userId)
	if err != nil {
		return
	}

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindByID(ctx context.Context, userId string) (web.UserDetailResponse, error) {
	//TODO implement me
	var response web.UserDetailResponse

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserDetailResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByID(ctx, tx, userId)
	if err != nil {
		return web.UserDetailResponse{}, err
	}

	response = web.UserDetailResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
		Kota:      *user.Kota,
		Avatar:    *user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (service *UserServiceImpl) FindByUserName(ctx context.Context, username string) (web.UserDetailResponse, error) {
	//TODO implement me
	var response web.UserDetailResponse

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserDetailResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUserName(ctx, tx, username)
	if err != nil {
		return web.UserDetailResponse{}, err
	}

	response = web.UserDetailResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
		Kota:      *user.Kota,
		Avatar:    *user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (service *UserServiceImpl) FindByEmail(ctx context.Context, email string) (web.UserDetailResponse, error) {
	//TODO implement me
	var response web.UserDetailResponse

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserDetailResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, email)
	if err != nil {
		return web.UserDetailResponse{}, err
	}

	response = web.UserDetailResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
		Kota:      *user.Kota,
		Avatar:    *user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserDetailResponse, error) {
	//TODO implement me
	var response []web.UserDetailResponse

	tx, err := service.DB.Begin()
	if err != nil {
		return []web.UserDetailResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	users, err := service.UserRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.UserDetailResponse{}, err
	}

	for _, user := range users {
		response = append(response, web.UserDetailResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			UserName:  user.UserName,
			Email:     user.Email,
			Phone:     user.Phone,
			Kota:      *user.Kota,
			Avatar:    *user.Avatar,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return response, nil
}
