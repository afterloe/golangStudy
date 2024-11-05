package token

import (
	"github.com/o1egl/paseto"
	"time"
)

type (
	pasetoImpl struct {
		paseto               *paseto.V2
		symmetricKey         []byte
		accessTokenDuration  time.Duration
		refreshTokenDuration time.Duration
	}
)

func newPase2Inspect(symmetricKey []byte) iTokenHelper {
	return &pasetoImpl{
		paseto:       paseto.NewV2(),
		symmetricKey: symmetricKey,
	}
}

func (that *pasetoImpl) SetDurationTime(accessTokenDuration, refreshTokenDuration time.Duration) {
	that.accessTokenDuration = accessTokenDuration
	that.refreshTokenDuration = refreshTokenDuration
}

func (that *pasetoImpl) CreateToken(id, username string, role []string, tokenType string) (string, *Payload, error) {
	var duration time.Duration
	if tokenType == ACCESS {
		duration = that.accessTokenDuration
	} else if tokenType == REFRESH {
		duration = that.refreshTokenDuration
	}
	payload := generatorPayload(id, username, role, duration)
	token, err := that.paseto.Encrypt(that.symmetricKey, payload, nil)
	return token, payload, err
}

func (that *pasetoImpl) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := that.paseto.Decrypt(token, that.symmetricKey, payload, nil)
	return payload, err
}
