package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type (
	jwtImpl struct {
		symmetricKey         []byte
		accessTokenDuration  time.Duration
		refreshTokenDuration time.Duration
	}
)

const (
	timeLayout = "2006-01-02T15:04:05.999999-07:00"
)

func newJWTInspect(symmetricKey []byte) iTokenHelper {
	return &jwtImpl{symmetricKey: symmetricKey}
}

func (that *jwtImpl) SetDurationTime(accessTokenDuration, refreshTokenDuration time.Duration) {
	that.accessTokenDuration = accessTokenDuration
	that.refreshTokenDuration = refreshTokenDuration
}

func (that *jwtImpl) CreateToken(id, username string, role []string, tokenType string) (string, *Payload, error) {
	var duration time.Duration
	if tokenType == ACCESS {
		duration = that.accessTokenDuration
	} else if tokenType == REFRESH {
		duration = that.refreshTokenDuration
	}
	payload := generatorPayload(id, username, role, duration)
	claims := jwt.MapClaims{}
	claims["id"] = payload.ID
	claims["username"] = payload.Username
	claims["role"] = payload.Role
	claims["issued_at"] = payload.IssuedAt
	claims["expired_at"] = payload.ExpiredAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(that.symmetricKey)
	return tokenStr, payload, err
}

func (that *jwtImpl) VerifyToken(token string) (*Payload, error) {
	_token, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return that.symmetricKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !_token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := _token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("passer token failed")
	}

	payload := &Payload{
		ID:       claims["id"].(string),
		Username: claims["username"].(string),
	}

	iRoles := claims["role"].([]interface{})
	if len(iRoles) == 0 {
		payload.Role = []string{}
	} else {
		roles := make([]string, len(iRoles))
		for i, role := range iRoles {
			roles[i] = role.(string)
		}
		payload.Role = roles
	}
	issuedAtTime, _ := time.Parse(timeLayout, claims["issued_at"].(string))
	payload.IssuedAt = issuedAtTime
	expiredAtTime, _ := time.Parse(timeLayout, claims["expired_at"].(string))
	payload.ExpiredAt = expiredAtTime

	return payload, nil
}
