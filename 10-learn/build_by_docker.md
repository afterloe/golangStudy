使用docker构建通用go打包镜像
===
使用go进行打包的步骤如下
## 使用docker镜像构建可运行对象
```shell
docker run -it -v .:/go -v .:/app --rm golang:1.23.2 go build -v
docker run -it -v .:/ --rm alpine
```
或使用如下dockerfile进行构建
```dockerfile
FROM golang:1.13 as builder

WORKDIR /app
COPY . .
RUN go build -o timelocation .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/timelocation /app

ENTRYPOINT ["./timelocation"]
```

```makefile
# create by afterloe <lm6289511@gmail.com>

.PHONY: build,compile,structure
SHELL := /bin/bash
WORKDIR = $(shell pwd)
BUILD_IMG = awpaas/builder
VERSION = $(shell more package.json | grep version | awk -F '"' 'NR==1{print$$4}')
NAME = $(shell more package.json | grep name | awk -F '"' 'NR==1{print$$4}')
BUILD_ENV = $(shell docker image ls | grep ${BUILD_IMG} | wc -l)

all: compile structure

.ONESHELL:
compile: $(src) package.json
	docker run -it -v $(src):/go/src -v $(WORKDIR):/app --rm ${BUILD_IMG}:1.0.0 go build -v
structure: app
	docker build -t awpaas/$(NAME):$(VERSION) .
```

```dockerfile
FROM alpine:3.8
MAINTAINER afterloe <lm6289511@gmail.com>

ENV \
    PROJECT_DIR="/app"
WORKDIR ${PROJECT_DIR}
COPY app ${PROJECT_DIR}
COPY package.json ${PROJECT_DIR}
EXPOSE 8080 8081
CMD ./app
```

### 构建go相关项目
使用如下脚本进行编译
```shell
#!/bin/bash
echo "ready to build image"
src=$1
out=$2
echo "use ${src} to build"
echo "bin export to ${out}"
docker run -it \
-v ${src}:/go/src \
-v ${out}:/app \
--rm \
awpaas/builder:1.0.0 \
go build
```
> 使用方法`sh build.sh ./src ./dist`

### awpaas系列组件中使用make来构建
```sbtshell
make -m src=/data/data-2/go/src
docker run -it awpaas/awpaas-route:1.0.0
```
### build
```sbtshell
docker build -t awpaas/builder:1.0.0 .
```