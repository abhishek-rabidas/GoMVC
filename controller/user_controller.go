package controller

import (
	"github.com/labstack/echo/v4"
	"gomvc/model"
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
	u.userService.AddUser(model.UnMarshalUser(e.Request().Body))
	return e.JSON(http.StatusOK, "User added")
}

func (u *UserController) getUsers(e echo.Context) error {
	users := u.userService.GetUsers()
	return e.JSON(http.StatusOK, users)
}

func (u *UserController) getUserById(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	user, err := u.userService.GetUserById(id)

	if err != nil {
		return echo.NewHTTPError(500, "User not found")
	}

	return e.JSON(http.StatusOK, user)
}
