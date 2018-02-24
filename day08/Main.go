package main

import (
	"runtime"
	"os"
	"fmt"
	cli "./server"
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

func main() {
	fmt.Printf("server init ... \n")
	fmt.Printf("machine is %d cpus. \n", cpuNumber)
	fmt.Printf("config thread number is null, will start %d thread to wrok default \n", cpuNumber)
	runtime.GOMAXPROCS(cpuNumber)
	fmt.Printf("listen parameter is null, will start server in %s default \n", defAddr)
	fmt.Printf("server is init success... started pid is %d \n", pid)
	cli.StartUpTCPServer(&defAddr)
}