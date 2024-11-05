package album

import (
	"onenet/internal"
	"onenet/internal/entity"
	"onenet/internal/service"
	"onenet/internal/util"
)

type (
	CreateAlbumRequest struct {
		Title       string `json:"title" binding:"required"` // 相册标题
		Summary     string `json:"summary"`                  // 概况
		Description string `json:"description"`              // 描述
	}

	CreateAlbumResponse struct {
		ID string `json:"id"`
		CreateAlbumRequest
	}
)

// CreateOne
// @BasePath /album/create_one
// @Summary 创建相册
// @Description 创建相册
// @Accept json
// @Produce json
// @Param json body album.CreateAlbumRequest true "创建相册所需要的参数"
// @Success 200 {object} internal.CommonResponse[album.CreateAlbumResponse] "请求成功"
// @Router /album/create_one [POST]
func CreateOne(ctx *internal.Context) {
	var request CreateAlbumRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.FailWithStr("绑定失败")
		ctx.Fail(err)
		return
	}

	album := entity.Album{}
	err = util.Copy(&album, request)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.BindAuthorizeInfo(&album)
	info := ctx.AuthorizeInfo()
	album.Artist = info.Username
	r, err := service.AlbumService().CreateOne(&album)
	if err != nil {
		ctx.Fail(err)
		return
	}

	res := CreateAlbumResponse{}
	_ = util.Copy(&res, r)
	res.ID = r.Id
	ctx.Success(res)
}
