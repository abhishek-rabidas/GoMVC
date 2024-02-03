package service

import (
	"go/types"
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

func (us *UserService) GetUserById(id int) (model.User, error) {
	user := model.User{}
	var count int64
	config.DatabaseContext.First(&user, id).Count(&count)

	var err error

	if count == 0 {
		err = types.Error{Msg: "User not found"}
		return user, err
	}

	return user, nil
}
