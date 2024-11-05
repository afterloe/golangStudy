package album

import (
	"onenet/internal"
	"onenet/internal/entity"
	"onenet/internal/service"
)

// List
// @BasePath /album/list
// @Summary 列表查询所有相册
// @Description 列表查询所有相册
// @Accept json
// @Produce json
// @Param json body album.ListRequest true "分页参数"
// @Success 200 {object} internal.CommonResponse[album.ListResponse] "请求成功"
// @Router /album/list [POST]
func List(ctx *internal.Context) {
	var req ListRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.FailWithStr("绑定失败")
	}
	records, total, err := service.AlbumService().ListByPage(req.Content, req.Page, req.Size)
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Success(&ListResponse{
		Total:   total,
		Records: records,
	})
}

type (
	ListRequest struct {
		Content string `json:"content"`
		internal.PageParam
	}

	ListResponse struct {
		Total   int             `json:"total"`
		Records []*entity.Album `json:"records"`
	}
)
