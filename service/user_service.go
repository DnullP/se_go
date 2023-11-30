package service

type UserService interface {
	Login(username string, password string) (bool, error)
	Logout(username string) (bool, error)
}

