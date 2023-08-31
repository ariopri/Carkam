package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-66/models/web"
	"github.com/rg-km/final-project-engineering-66/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Register(ctx *gin.Context) {
	//TODO implement me
	var userCreateRequest web.RegisterRequest
	err := ctx.ShouldBindJSON(&userCreateRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := controller.UserService.Register(ctx, userCreateRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusCreated, web.WebResponse{
		Code:   201,
		Status: "User Register Successfully",
		Data:   response,
	})
}

func (controller *UserController) Login(ctx *gin.Context) {
	//TODO implement me
	var userLoginRequest web.LoginRequest
	err := ctx.ShouldBindJSON(&userLoginRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := controller.UserService.Login(ctx, userLoginRequest)

	if err != nil {
		ctx.Abort()
		return
	}

	if response.ID == "" {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   "Please check your username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	if response.Token == "" {
		ctx.JSON(400, web.WebResponse{
			Code:   400,
			Status: "UserName or Password is Wrong",
		})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusCreated, web.WebResponse{
		Code:   201,
		Status: "Login Successfully",
		Data:   response,
	})
}

func (controller *UserController) UpdateUser(ctx *gin.Context) {
	//TODO implement me
	var userUpdateRequest web.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&userUpdateRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userUpdateRequest.ID = ctx.Param("id")
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, web.WebResponse{
			Code:   401,
			Status: "Unauthorized",
			Data:   "Please Login First",
		})
		return
	}

	response, err := controller.UserService.UpdateUser(ctx, userUpdateRequest)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusCreated, web.WebResponse{
		Code:   201,
		Status: "User Update Successfully",
		Data:   response,
	})
}

func (controller *UserController) DeleteUser(ctx *gin.Context) {
	//TODO implement me
	userID := ctx.Param("id")
	controller.UserService.DeleteUser(ctx, userID)

	ctx.Header("Accept", "application/json")

	ctx.IndentedJSON(http.StatusOK, web.WebResponse{
		Code:   200,
		Status: "User Deleted Successfully",
	})
}

func (controller *UserController) FindByID(ctx *gin.Context) {
	//TODO implement me
	userID := ctx.Param("id")

	response, err := controller.UserService.FindByID(ctx, userID)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
		})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusOK, web.WebResponse{
		Code:   200,
		Status: "Get User By ID Successfully",
		Data:   response,
	})
}

func (controller *UserController) FindByUserName(ctx *gin.Context) {
	//TODO implement me
	userUsername := ctx.Param("username")
	response, err := controller.UserService.FindByUserName(ctx, userUsername)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
		})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusOK, web.WebResponse{
		Code:   200,
		Status: "Get User By UserName Successfully",
		Data:   response,
	})
}

func (controller *UserController) FindByEmail(ctx *gin.Context) {
	//TODO implement me
	userEmail := ctx.Param("email")
	response, err := controller.UserService.FindByEmail(ctx, userEmail)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
		})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusOK, web.WebResponse{
		Code:   200,
		Status: "Get User By Email Successfully",
		Data:   response,
	})
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	//TODO implement me
	response, err := controller.UserService.FindAll(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, web.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
		})
		return
	}

	ctx.Header("Accept", "application/json")
	ctx.Header("Content-Type", "application/json")

	ctx.IndentedJSON(http.StatusOK, web.WebResponse{
		Code:   200,
		Status: "Get All User Successfully",
		Data:   response,
	})
}
