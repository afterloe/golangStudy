package authorize

import (
	"onenet/internal"
	"onenet/internal/service"
)

type (

	// LoginRequest
	// 登陆时所需要的参数
	LoginRequest struct {
		LoginName string `json:"login_name" binding:"required"` // 用户名
		Scrip     string `json:"scrip" binding:"required"`      // 编码
	}

	// LoginResponse
	// 登陆成功后的响应参数
	LoginResponse struct {
		SessionID    string `json:"session_id"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		User         string `json:"user"`
	}
)

// Login
// @BasePath /authorize/login
// @Summary 登陆
// @Description 登陆
// @Accept json
// @Produce json
// @Param json body authorize.LoginRequest true "登陆时所需要的信息"
// @Success 200 {object} internal.CommonResponse[authorize.LoginResponse] "token信息"
// @Router /authorize/login [POST]
func Login(ctx *internal.Context) {
	var loginRequest LoginRequest
	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		ctx.Fail(err)
		return
	}
	accessToken, refreshToken, err := service.AuthorizeService().Login(loginRequest.LoginName, loginRequest.Scrip)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(&LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         loginRequest.LoginName,
	})
}
