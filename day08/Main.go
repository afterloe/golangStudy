package main

import (
	"runtime"
	"os"
	"fmt"
	cli "./server"
	"./integrate/logger"
	"./config"
)

var (
	defAddr string
	cpuNumber int
	pid int
)

func init() {
	defAddr = "127.0.0.1:8080"
	cpuNumber = runtime.NumCPU()
	pid = os.Getpid()
}

func startUpService(serverCfg map[string]interface{}) {
	addr, port := serverCfg["addr"], serverCfg["port"]
	if nil == addr {
		addr = "127.0.0.1"
	}
	if nil == port {
		port = "8080"
	}
	addrStr := fmt.Sprintf("%s:%s", addr, port)
	logger.Info(fmt.Sprintf("will start server in %s ", addrStr))
	cli.StartUpTCPServer(&addrStr)
}

func startUpMultiService(multiCfg map[string]interface{}) {
	flg := multiCfg["enable"]
	if nil == flg {
		logger.Info("server not allow to use multi cpu")
		return
	}
	if flg.(bool) {
		coreNumber := multiCfg["num"]
		logger.Info("server will to use multi cpu")
		if nil == coreNumber {
			coreNumber = cpuNumber
		} else if 0 >= coreNumber.(float64) {
			coreNumber = cpuNumber
		}
		logger.Info(fmt.Sprintf("server will use %v cpu", coreNumber))
		runtime.GOMAXPROCS(int(coreNumber.(float64)))
	}
}

func startDefault() {
	logger.Info(fmt.Sprintf("listen parameter is null, will start server in %s default", defAddr))
	logger.Info(fmt.Sprintf("server is init success... started pid is %d", pid))
	logger.Info("started server success.")
	cli.StartUpTCPServer(&defAddr)
}

func main() {
	logger.Info("server init ...")
	logger.Info(fmt.Sprintf("machine is %d cpus.", cpuNumber))
	serverCfg := config.Get("server").(map[string]interface{})
	if nil == serverCfg {
		startDefault()
		return
	}
	multiCfg := serverCfg["multiCore"].(map[string]interface{})
	if nil == multiCfg {
		startUpService(serverCfg)
	} else {
		startUpMultiService(multiCfg)
		startUpService(serverCfg)
	}
}