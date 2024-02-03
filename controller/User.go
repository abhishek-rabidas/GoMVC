package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
}

func NewUserController(e *echo.Group) *UserController {
	userController := UserController{}
	e.GET("/getUsers", userController.getUsers)
	e.GET("/getUser/:id", userController.getUserById)
	return &userController
}

func (u *UserController) getUsers(e echo.Context) error {
	return e.String(http.StatusOK, "Users")
}

func (u *UserController) getUserById(e echo.Context) error {
	return e.String(http.StatusOK, "User "+e.Param("id"))
}
