package album

import (
	"fmt"
	"onenet/internal/dao"
	"onenet/internal/entity"
	"onenet/internal/service"
)

const table = "album"

type _albumServiceImpl struct {
}

func init() {
	impl := &_albumServiceImpl{}
	service.RegisterAlbumService(impl)
}

func (*_albumServiceImpl) ListByPage(content string, page int, size int) ([]*entity.Album, int, error) {
	if page >= 1 {
		page -= 1
	}
	queryJSON := fmt.Sprintf("{\"selector\":{\"$or\":[{\"title\":{\"$regex\":\""+content+"\"}},{\"summary\":{\"$regex\":\""+content+"\"}},{\"description\":{\"$regex\":\""+content+"\"}}]},\"sort\":[{\"created_at\":\"desc\"}],\"limit\":%d,\"skip\":%d}", size, page*size)
	results := dao.Couchdb.Find(table, queryJSON)
	if nil != results.Err() {
		return nil, 0, results.Err()
	}
	records := make([]*entity.Album, 0)
	for results.Next() {
		var album entity.Album
		err := results.ScanDoc(&album)
		if err != nil {
			return nil, 0, err
		}
		album.FirstPic = nil
		album.ID = album.Id
		records = append(records, &album)
	}
	return records, 0, nil
}

func (*_albumServiceImpl) CreateOne(album *entity.Album) (*entity.Album, error) {
	album.Number = 0
	album.OrderIndex = 0
	id, err := dao.Couchdb.Put(table, album)
	if err != nil {
		return nil, err
	}
	album.Id = id
	return album, nil
}
