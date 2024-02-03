package controller

import (
	"github.com/labstack/echo/v4"
	"gomvc/service"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(e *echo.Group) *UserController {
	userController := UserController{userService: service.NewUserService()}
	e.POST("/", userController.addUser)
	e.GET("/", userController.getUsers)
	e.GET("/:id", userController.getUserById)
	return &userController
}

func (u *UserController) addUser(e echo.Context) error {
	user := e.QueryParam("name")
	u.userService.AddUser(user)
	return e.JSON(http.StatusOK, "User added")
}

func (u *UserController) getUsers(e echo.Context) error {
	users := u.userService.GetUsers()
	return e.JSON(http.StatusOK, users)
}

func (u *UserController) getUserById(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	user := u.userService.GetUserById(id)
	return e.String(http.StatusOK, user)
}
