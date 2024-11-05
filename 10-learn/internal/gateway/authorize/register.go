package authorize

import (
	"onenet/internal"
	"onenet/internal/entity"
	"onenet/internal/service"
	"onenet/internal/util"
)

type (
	// RegisterRequest
	// 注册时所需要的参数
	RegisterRequest struct {
		Username string `json:"username" binding:"required"` // 用户名
		Password string `json:"password" binding:"required"` // 密码
	}
)

// Register
// @BasePath /authorize/register
// @Summary 注册用户
// @Description 注册用户
// @Accept json
// @Produce json
// @Param json body authorize.RegisterRequest true "注册用户所需的信息"
// @Success 200 {object} internal.CommonResponse[string] "请求成功"
// @Router /authorize/register [POST]
func Register(ctx *internal.Context) {
	var req RegisterRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Fail(err)
		return
	}

	user := entity.User{}
	_ = util.Copy(&user, req)
	loginName, err := service.AuthorizeService().Register(&user, req.Password)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(loginName)
}
