package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gomvc/model"
	"gomvc/service"
	"io"
	"log"
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
	user := model.User{}
	body, err := io.ReadAll(e.Request().Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	u.userService.AddUser(&user)
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
