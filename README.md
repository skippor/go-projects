# go-projects

## env

```shell
cat ~/.config/go/env
GO111MODULE=on
GOPATH=/mnt/dtt/go
GOPROXY=http://mirrors.xxx.org/nexus/repository/go-proxy-group,http://goproxy.xxx.com:9080
GOSUMDB=off
```


## example

```shell
$ mkdir helloworld
$ cd helloworld/
$ touch main.go
$ go mod init helloworld
$ go run main.go
$ go mod tidy
$ go build
$ ./helloworld
$ go clean
$ go build -o hello
$ ./hello
```


## todo

* convey
* grpc
* redis
* kafka
* etcd
* mongodb
* mysql
* GraphQL - gqlgen
* cgo
* restful
