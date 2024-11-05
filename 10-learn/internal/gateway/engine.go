package gateway

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"onenet/docs"
	"onenet/internal"
	"onenet/internal/gateway/album"
	"onenet/internal/gateway/authorize"
	"onenet/internal/util"
	"time"
)

type RouterGroup struct {
	gin.RouterGroup
}

type Engine struct {
	Config     util.Config
	Background *http.Server
	*gin.Engine
}

func (that *Engine) BindRoute() {
	if that.Config.Environment == "development" {
		that._bindByDev()
	} else if that.Config.Environment == "production" {
		that._bindByPro()
	}
	that._bind()
}

func (that *Engine) _bindByDev() {
	gin.SetMode(gin.DebugMode)
	// 允许跨域
	that.Use(cors.New(cors.Config{
		AllowOrigins: that.Config.AllowedOrigins,
		AllowMethods: []string{
			http.MethodPost,
		},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
	}))

	// 加载swagger
	docs.SwaggerInfo.BasePath = "/"
	err := that.SetTrustedProxies([]string{"::1"})
	if err != nil {
		return
	}
	that.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (that *Engine) _bindByPro() {
	gin.SetMode(gin.ReleaseMode)
}

func (that RouterGroup) POST(relativePath string, handlers ...internal.HandlerFunc) RouterGroup {
	that.RouterGroup.POST(relativePath, internal.EnhanceContext(handlers[0]))
	return that
}

func (that RouterGroup) Group(relativePath string, handlers ...internal.HandlerFunc) RouterGroup {
	if len(handlers) != 0 {
		that.RouterGroup.Group(relativePath, internal.EnhanceContext(handlers[0]))
	} else {
		that.RouterGroup.Group(relativePath)
	}
	return that
}

func (that RouterGroup) Use(handlers ...gin.HandlerFunc) RouterGroup {
	that.RouterGroup.Use(handlers...)
	return that
}

func (that *Engine) _bind() {
	that.Use(gin.Logger())
	that.Use(gin.Recovery())

	routes := &RouterGroup{
		that.RouterGroup,
	}
	routes.POST("/authorize/register", authorize.Register)
	routes.POST("/authorize/login", authorize.Login)

	authorizeRoute := routes.Group("/").Use(AuthorizeValid())
	authorizeRoute.POST("/album/list", album.List)
	authorizeRoute.POST("/album/get_by_id", album.GetById)
	authorizeRoute.POST("/album/create_one", album.CreateOne)
	authorizeRoute.POST("/album/delete_by_id", album.DeleteById)
}

func Init(ctx context.Context, config util.Config) *Engine {
	engine := &Engine{
		Engine: gin.New(),
		Config: config,
	}
	engine.BindRoute()
	return engine
}

func Run(engine *Engine) error {
	server := &http.Server{
		Addr:           engine.Config.HTTPServerAddress,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	engine.Background = server
	return engine.Background.ListenAndServe()
}

func Shutdown(ctx context.Context, engine *Engine) error {
	return engine.Background.Shutdown(ctx)
}
