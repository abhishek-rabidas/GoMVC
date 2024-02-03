package service

type UserService struct {
	users []string
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) AddUser(user string) {
	us.users = append(us.users, user)
}

func (us *UserService) GetUsers() []string {
	return us.users
}

func (us *UserService) GetUserById(id int) string {
	return us.users[id]

}
