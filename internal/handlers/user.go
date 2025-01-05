package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"uala.com/core-service/internal/domain"
	"uala.com/core-service/internal/usescase"
)

type UserHandler struct {
	userUseCase usescase.CreateUser
}

func NewUserHandler(userUseCase usescase.CreateUser) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	fmt.Println("CreateUser")
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}

	if err := u.userUseCase.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusCreated, user)
}
