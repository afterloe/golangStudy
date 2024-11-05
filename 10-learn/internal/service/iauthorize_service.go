package service

import "onenet/internal/entity"

// IAuthorizeService
// 授权相关接口
type IAuthorizeService interface {

	// Register
	// 注册
	// @Param *entity.User 用户
	// @Param string 密码
	// @Return string 用户登陆凭证
	// @Return error  异常
	Register(*entity.User, string) (string, error)

	// Login
	// 登陆
	// @Param string 登陆名
	// @Param string 登陆凭证
	// @Return accessToken 访问token
	// @Return refreshToken 刷新token
	// @Return error  异常
	Login(string, string) (string, string, error)
}

var (
	_authorizeService IAuthorizeService
)

func RegisterAuthorizeService(service IAuthorizeService) {
	_authorizeService = service
}

func AuthorizeService() IAuthorizeService {
	if _authorizeService == nil {
		panic("authorizeService is nil")
		return nil
	}
	return _authorizeService
}
