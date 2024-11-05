package authorize

import (
	"fmt"
	gmssl "github.com/GmSSL/GmSSL-Go"
	"onenet/internal/consts"
	"onenet/internal/dao"
	"onenet/internal/errors"
)

// _readyRegister
// 判断用户是否已经注册
func _readyRegister(username string) error {
	query := fmt.Sprintf(`{"selector":{"username":{"$eq":"%s"}},"fields":["_id"]}`, username)
	result := dao.Couchdb.Find(table, query)
	if result.Next() {
		return &errors.BizError{
			ErrorCode: 400,
			Msg:       fmt.Sprintf(consts.ReadyRegister, username),
		}
	}
	return nil
}

// _str2Ciphertext
// 国密sm3 hash
func _str2Ciphertext(str string) string {
	sm3 := gmssl.NewSm3()
	sm3.Update([]byte(str))
	cipher := sm3.Digest()
	return fmt.Sprintf("%x", cipher)
}
