# 第八天

进入项目实战篇，这里打算使用go来构建一个SOA服务管理平台。需求是这样可以使用react等前后端分离的页面框架实现在页面上托拉拽来构成SOA服务。  
SOA这一部分就不在这里多扯了，这里主要使用web来控制docker来实现镜像拉取、部署、发布等一套活动

## docker go sdk

docker的go sdk可以在[这里获取](https://docs.docker.com/develop/sdk/#install-the-sdks)到,这里使用的是go的sdk。docker的封装比较彻底，一般对外的联系都是采用`/var/run/docker.sock`来进行沟通的。所以在开发的时候需要安装docker，当然使用远程的docker也是可以的，可以修改docker的启动方式，让socket的连接方式改成tcp方式即可

### 修改远程docker为tcp连接
```sbtshell
[root@localhost docker]# cd /etc/docker/
[root@localhost docker]# vim daemon.json

... 
[root@localhost docker]# more daemon.json

{
  "registry-mirrors": ["https://mq8v733y.mirror.aliyuncs.com"],
  "hosts": ["tcp://0.0.0.0:2376", "unix:///var/run/docker.sock"]
}

[root@localhost docker]# firewall-cmd --zone=public --add-port=2376/tcp --permanent
success
[root@localhost docker]# firewall-cmd --reload
success

[root@localhost docker]# systemctl restart docker.service
```

关于daemon.json 的配置文件的修改可以在[这里](https://docs.docker.com/config/daemon/#configure-the-docker-daemon)查看到。

测试一下
```sbtshell
[afterloe@localhost docker]# curl 127.0.0.1:2376/v1.32/images/json

[{"Containers":-1,"Created":1515547359,"Id":"sha256:d1fd7d86a8257f3404f92c4474fb3353076883062d64a09232d95d940627459d","Labels":null,"ParentId":"","RepoDigests":["registry@sha256:672d519d7fd7bbc7a448d17956ebeefe225d5eb27509d8dc5ce67ecb4a0bce54"],"RepoTags":["registry:latest"],"SharedSize":-1,"Size":33258091,"VirtualSize":33258091}]
```

关于docker的api可以在[这里](https://docs.docker.com/engine/api/v1.32/#tag/Image) 查看到，[这里是go Doc关于docker的源码文档](https://godoc.org/github.com/docker/docker/client)    

然后在go使用sdk进行访问.这里的事例是获取镜像列表的。[官网给的api示例](https://docs.docker.com/develop/sdk/examples/)是过时的，我这里用的是最新的api
```golang
import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func GetImageList() interface{} {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithHost("http://yyy:2376"),
		client.WithVersion("1.36"))
	if nil != err {
		panic(err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})

	if nil != err {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image)
	}

	return images
}
```
