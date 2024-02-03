package service

import (
	"gomvc/config"
	"gomvc/model"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) AddUser(user *model.User) {
	config.DatabaseContext.Save(user)
}

func (us *UserService) GetUsers() []model.User {
	users := make([]model.User, 0)
	config.DatabaseContext.Find(&users)
	return users
}

func (us *UserService) GetUserById(id int) model.User {
	user := model.User{}
	//config.DatabaseContext.Table("users").Find(&user)
	config.DatabaseContext.First(&user, id)
	return user
}
