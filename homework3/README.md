# 模块3作业
***
## 作业要求
- 构建本地镜像
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
- 将镜像推送至 docker 官方镜像仓库
- 通过 docker 命令本地启动 httpserver
- 通过 nsenter 进入容器查看 IP 配置

## 作业完成详情
1. 编写dockerfile文件，即：[Dockerfile](../homework2/Dockerfile)

内容如下：先使用镜像作为构建工具，生成可执行文件，再生成运行部署的镜像

```dockerfile
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/main homework2/main.go


FROM alpine

RUN echo -e 'https://mirrors.aliyun.com/alpine/v3.6/main/\nhttps://mirrors.aliyun.com/alpine/v3.6/community/' > /etc/apk/repositories && \
        apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/main /app/main

CMD ["./main"]
```

2. 构建镜像：运行下面的命令即可

```shell
# 在工程根目录下
docker build -t lw1243925457/http_example:v1 -f homework2/Dockerfile .
```

可以看到自己本地的镜像大小,不到20M，比起Java确实是小很多

```text
PS D:\Code\go\self\GoTraining> docker images
REPOSITORY                  TAG       IMAGE ID       CREATED         SIZE
lw1243925457/http_example   v1        739f1bb280dc   6 minutes ago   19.2MB
```


4. 推送镜像到自己的docker hub仓库中，运行下面的命令进行推送

```shell
docker push lw1243925457/http_example:v1
```

可以看到下面的输出，则推送成功，可以登录网页账号进行查看到

```text
PS D:\Code\go\self\GoTraining> docker push lw1243925457/http_example:v1
The push refers to repository [docker.io/lw1243925457/http_example]
11dc4c8a94cf: Pushed
3789d1e7303c: Pushed
acf4cb94037d: Pushed
8d3ac3489996: Mounted from library/golang
v1: digest: sha256:14d9db05e626089c2fa667291fd9310787b15eedaf7623e54469bea1badeda26 size: 1157
```

5. 启动镜像运行：上面构建推送是window本地，下面是在服务器上面进行操作

```shell
➜  ~ docker pull  lw1243925457/http_example:v1
v1: Pulling from lw1243925457/http_example
59bf1c3509f3: Pull complete
840a0152acef: Pull complete
dc0254ddd5b3: Pull complete
331fc01c2245: Pull complete
Digest: sha256:14d9db05e626089c2fa667291fd9310787b15eedaf7623e54469bea1badeda26
Status: Downloaded newer image for lw1243925457/http_example:v1
docker.io/lw1243925457/http_example:v1

# 简单使用命令进行启动，也可以不用上面命令，直接运行，自动拉取
➜  ~ docker run -d --name http_example -p 8070:8080 lw1243925457/http_example:v1
199bb00ec82730b7a8778fcbb6b5a866c97de872c6ed973056d7dc4b64d38747

# 访问接口看是否启动成功
➜  ~ curl http://localhost:8070/v1/hello
"Hello"#     
```

7. 查看镜像

```shell
➜  ~ docker inspect -f {{.State.Pid}} http_example
1031859
➜  ~ nsenter -n -t 1031859
➜  ~ ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
75: eth0@if76: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:05 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.5/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```