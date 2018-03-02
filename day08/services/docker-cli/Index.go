package docker_cli

import (
	"../../config"
	"golang.org/x/net/context"
	"github.com/docker/docker/client"
)

var (
	flag bool
	addr, version string
)

func init() {
	servicesCfg := config.Get("services").(map[string]interface{})
	dockerCfg := servicesCfg["docker-cli"].(map[string]interface{})
	version = dockerCfg["version"].(string)
	socket := dockerCfg["socket"].(map[string]interface{})
	if nil == socket {
		flag = false
		return
	}
	flag = socket["enable"].(bool)
	if flag {
		addr = socket["addr"].(string)
	}
}

func getContext() context.Context{
	return context.Background()
}

func getCli() (*client.Client, error) {
	if flag {
		return client.NewClientWithOpts(client.WithHost("http://" + addr), client.WithVersion(version))
	}
	return client.NewClientWithOpts(client.WithVersion(version))
}
