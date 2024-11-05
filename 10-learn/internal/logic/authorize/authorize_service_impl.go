package authorize

import (
	"fmt"
	"onenet/internal/consts"
	"onenet/internal/dao"
	"onenet/internal/entity"
	"onenet/internal/errors"
	"onenet/internal/service"
	"onenet/internal/util/token"
)

const table = "user"

type _authorizeServiceImpl struct{}

func init() {
	impl := &_authorizeServiceImpl{}
	service.RegisterAuthorizeService(impl)
}

func (*_authorizeServiceImpl) Login(loginName, scrip string) (string, string, error) {
	ciphertext := _str2Ciphertext(scrip)
	query := fmt.Sprintf(`{"selector":{"$and":[{"username":"%s"},{"ciphertext":"%s"}]},"fields":["_id", "username"]}`, loginName, ciphertext)
	result := dao.Couchdb.Find(table, query)
	if !result.Next() {
		return "", "", &errors.BizError{
			ErrorCode: 401,
			Msg:       consts.LoginNameOrPassword,
		}
	}
	var user entity.User
	_ = result.ScanDoc(&user)
	accessToken, _, err := token.Inspect.CreateToken(user.Id, user.Username, []string{}, token.ACCESS)
	if err != nil {
		return "", "", err
	}
	refreshToken, _, err := token.Inspect.CreateToken(user.Id, user.Username, []string{}, token.REFRESH)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}

func (*_authorizeServiceImpl) Register(user *entity.User, pwd string) (string, error) {
	err := _readyRegister(user.Username)
	if err != nil {
		return "", err
	}
	user.Ciphertext = _str2Ciphertext(pwd)
	_, err = dao.Couchdb.Put(table, user)
	if err != nil {
		return "", err
	}
	return user.Username, err
}
