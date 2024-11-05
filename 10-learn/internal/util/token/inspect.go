package token

var Inspect iTokenHelper // tokenHelper 实例

// NewTokenMaker 构建TokenHelper实例
func NewTokenMaker(symmetricKey, tokenType string) {
	if "paseto" == tokenType {
		Inspect = newPase2Inspect([]byte(symmetricKey))
	} else {
		Inspect = newJWTInspect([]byte(symmetricKey))
	}
}
