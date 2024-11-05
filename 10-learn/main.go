package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"onenet/internal/dao"
	"onenet/internal/gateway"
	_ "onenet/internal/logic"
	"onenet/internal/util"
	"onenet/internal/util/token"
	"os"
	"os/signal"
	"syscall"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	log.Println("ready to start server")
	config, err := util.LoadConfig("./manifest")
	if nil != err {
		log.Printf("cannot load config")
		log.Panic(err)
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	// TODO 数据库、redis、mq等中间件初始化
	runTokenKeyInit(config)
	runDBInit(config)

	waitGroup, ctx := errgroup.WithContext(ctx)
	// 启动对应的服务
	runGatewayServer(ctx, waitGroup, config)
	err = waitGroup.Wait()
	if err != nil {
		log.Printf("error from wait group")
		log.Panic(err)
	}
}

// 启动网关服务
func runGatewayServer(ctx context.Context, group *errgroup.Group, config util.Config) {
	engine := gateway.Init(ctx, config)
	group.Go(func() error {
		log.Printf("start gateway server at %s", config.HTTPServerAddress)
		err := gateway.Run(engine)
		if err != nil {
			log.Println("failed to start gateway server")
		}
		return nil
	})
	group.Go(func() error {
		<-ctx.Done()
		log.Println("graceful shutdown gateway server")
		err := gateway.Shutdown(ctx, engine)
		if err != nil {
			log.Println("failed to shutdown gateway server")
			return err
		}
		log.Println("gateway server is stopped")
		return nil
	})
}

func runDBInit(config util.Config) {
	dao.Register(config)
	log.Println("register db success")
}

func runTokenKeyInit(config util.Config) {
	token.NewTokenMaker(config.TokenSymmetricKey, config.TokenType)
	token.Inspect.SetDurationTime(config.AccessTokenDuration, config.RefreshTokenDuration)
}
