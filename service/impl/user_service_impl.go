package impl

type UserServiceImpl struct {}

func (userService UserServiceImpl) Login(username string, password string) (bool, error) {
	return true, nil
}

func (userService UserServiceImpl) Logout(username string) (bool, error) {
	return true, nil
}