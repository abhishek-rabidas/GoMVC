package service

import (
	"gomvc/config"
	"gomvc/model"
)

type UserService struct {
	users []string
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) AddUser(user *model.User) {
	config.DatabaseContext.Save(user)
}

func (us *UserService) GetUsers() []string {
	return us.users
}

func (us *UserService) GetUserById(id int) string {
	return us.users[id]

}
