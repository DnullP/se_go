package utils

import "github.com/DnullP/se_work/service/impl"

func GetUserService() impl.UserServiceImpl {
	return impl.UserServiceImpl{}
}

func GetControlService() impl.ControlServiceImpl {
	return impl.ControlServiceImpl{}
}

func GetQueryService() impl.QueryServiceImpl {
	return impl.QueryServiceImpl{}
}
