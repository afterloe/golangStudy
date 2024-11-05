go中如何实现bear token拦截和生成
===

关于jwt token，这里一块可以自行去百度，这里使用的是两个库
jwt [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)    
passto [github.com/o1egl/paseto](https://github.com/o1egl/paseto)    

## 实践过程

### 构建通用token工具
设计一个Token需要存储信息的对象，全部代码可参考[internal/util/token](internal/util/token)包内， 部分代码如下:
```go
package token

import (
	"time"
)

// Payload token中存储的信息
type Payload struct {
	ID        string    `json:"id"`         // 用户id
	Username  string    `json:"username"`   // 用户名
	Role      []string  `json:"role"`       // 用户角色列表
	IssuedAt  time.Time `json:"issued_at"`  // token创建时间
	ExpiredAt time.Time `json:"expired_at"` // token过期时间
}
```
代码参考:  [internal/util/token/payload.go](internal/util/token/payload.go)       

基于这个对象，构建TokenHelper工具和接口，以及实现方式，若只有一种则自行去掉多余的内容
```go
package token

import (
	"errors"
	"time"
)

// token类型，是access_token 还是 refresh_token
const (
	ACCESS  = "access" 
	REFRESH = "refresh"
)

// ITokenHelper tokenHelper接口
type iTokenHelper interface {
	// CreateToken
	// 为登陆用户创建token
	// @Param id string 用户id
	// @Param username string 用户名
	// @Param role []string 用户角色列表
	// @Param tokenType string token类型 access, refresh
	// @Return token， token.Payload， 异常
	CreateToken(id, username string, role []string, tokenType string) (string, *Payload, error)

	// VerifyToken
	// 验证Token字符串是否有效
	// @Param token string token字符串
	// @Return token.Payload，异常
	VerifyToken(token string) (*Payload, error)

	// SetDurationTime
	// 设置token存在时间
	// @Param accessTokenDuration 访问token存在时间
	// @Param refreshTokenDuration 刷新token存在时间
	SetDurationTime(accessTokenDuration, refreshTokenDuration time.Duration)
}

// ValidExpired
// 验证token是否超时
// Param expired 设置的时间
// Return error 未超时返回nil
func ValidExpired(expired time.Time) error {
	if time.Now().After(expired) {
		return errors.New("token 已过期")
	}
	return nil
}

// generatorPayload
// 生成token封装信息
// Param id string 用户id
// Param username string 用户名
// Param role []string 用户角色列表
// Param duration time.Duration token存在时间
// Return token.Payload
func generatorPayload(id, username string, role []string, duration time.Duration) *Payload {
	return &Payload{
		ID:        id,
		Username:  username,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}
```
代码参考:  [internal/util/token/token_helper.go](internal/util/token/token_helper.go)   
#### JWT 实现
```go
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
```
具体代码实现可参考[internal/util/token/jwt_impl.go](internal/util/token/jwt_impl.go)。

#### Passto实现
```go
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
```
具体代码实现可参考[internal/util/token/paseto_impl.go](internal/util/token/paseto_impl.go)

### tokenHelper实例化
```go
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
```
代码参考: [internal/util/token/inspect.go](internal/util/token/inspect.go)     

这样tokenHelper工具就支持jwt和passto两种方式。

### 工具使用

#### 初始化
```go
func runTokenKeyInit(config util.Config) {
	token.NewTokenMaker(config.TokenSymmetricKey, config.TokenType)
	token.Inspect.SetDurationTime(config.AccessTokenDuration, config.RefreshTokenDuration)
}
```
从配置文件中读取是实现jwt还是passto，同时配置AccessToken和RefreshToken两种token的存在时间. 初始化是在[main.go](main.go)中`78`行实现。

#### 登陆成功后生成Token
```go
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
    // 查询到用户信息后 调用 tokenHelper进行token创建
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
```
该段代码在[internal/logic/authorize/authorize_service_impl.go](internal/logic/authorize/authorize_service_impl.go)的第`22`行。

#### 在路由时对请求进行拦截
这边使用的是Gin的中间件实现
```go
package gateway

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onenet/internal"
	"onenet/internal/util/token"
	"strings"
)

func AuthorizeValid() gin.HandlerFunc {
	return func(c *gin.Context) {
        // 增强Context
		ctx := &internal.Context{Context: c}

        // 得到请求头
		bearerToken := ctx.GetHeader("Authorization")
        
        // 若不是 Bear  xx 格式则报错
		if len(strings.Split(bearerToken, " ")) != 2 {
			ctx.IndentedJSON(http.StatusUnauthorized, &internal.CommonResponse[interface{}]{
				BizCode: http.StatusUnauthorized,
				Error:   "未授权",
			})
			ctx.Abort()
			return
		}
        // 获取token字符串
		tokenStr := strings.Split(bearerToken, " ")[1]
        
        // 验证token字符串是否合法
		payload, err := token.Inspect.VerifyToken(tokenStr)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, &internal.CommonResponse[interface{}]{
				BizCode: http.StatusUnauthorized,
				Error:   "未授权",
			})
			ctx.Abort()
			return
		}
        
        // 验证token字符串是否超时
		err = token.ValidExpired(payload.ExpiredAt)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, &internal.CommonResponse[interface{}]{
				BizCode: http.StatusUnauthorized,
				Error:   "登陆超时",
			})
			ctx.Abort()
			return
		}
		ctx.Set("authorize", payload)
		ctx.Next()
	}
}
```
该段代码在[internal/gateway/authorize_valid.go](internal/gateway/authorize_valid.go)    

被用于[internal/gateway/engine.go](internal/gateway/engine.go)的`94`行
```go
authorizeRoute := routes.Group("/").Use(AuthorizeValid())
authorizeRoute.POST("/album/list", album.List)
authorizeRoute.POST("/album/get_by_id", album.GetById)
```