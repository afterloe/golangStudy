package service

import (
	"onenet/internal/entity"
)

// IAlbumService
// 相册相关接口
type IAlbumService interface {

	// CreateOne
	// 创建相册
	// @Param *entity.Album 相册创建对象
	// @Return entity.Album 保存后的相册对象
	// @Return error bizError.BizError
	CreateOne(*entity.Album) (*entity.Album, error)

	// ListByPage
	// 查询相册
	// @Param string 相册名称/描述/简介
	// @Param int 第几页，从0开始
	// @Param int 每页多少条
	// @Return []*entity.Album 查询结果
	// @Return int 总数
	// @Return error bizError.BizError
	ListByPage(string, int, int) ([]*entity.Album, int, error)
}

var (
	_albumService IAlbumService
)

func RegisterAlbumService(i IAlbumService) {
	_albumService = i
}

func AlbumService() IAlbumService {
	if _albumService == nil {
		panic("albumService is nil")
		return nil
	}
	return _albumService
}
