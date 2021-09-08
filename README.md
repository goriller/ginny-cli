# ginny-cli
Ginny command line tool.


## Install

系统需要安装 Go，在 $GOPATH 目录下执行以下命令安装。成功后一个名为 ginny 的二进制可执行文件会被安装至 $GOPATH/bin/ 文件夹中。

*注意： 该命令行工具支持MacOs、Linux系统，windows系统请使用gitBash或者Cygwin*

方法一： 在GOPATH下执行

```sh
cd $GOPATH && go get github.com/gorillazer/ginny-cli/ginny

```

方法二：编译安装

```sh
git clone github.com/gorillazer/ginny-cli.git

# mac:
cd ginny-cli/ginny && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ginny
cp -f ginny $GOPATH/bin/

# linux
cd ginny-cli/ginny && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ginny
cp -f ginny $GOPATH/bin/


# windows
cd ginny-cli/ginny && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ginny
cp -f ginny $GOPATH/bin/
```
## Dependencies

protoc：

```sh
// mac
brew install protoc

// centos7
yum install protobuf-compiler
```

protoc-gen-go：

```sh
// mac
brew install protoc-gen-go	

// centos7
yum install golang-googlecode-goprotobuf
```

goimports：

```sh
go get golang.org/x/tools/cmd/goimports
```

mockgen：

```sh
# 参考 https://github.com/golang/mock 进行安装
```

## Usage
```sh

Command line tool for Ginny project bestpractice

Usage:
  ginny [flags]
  ginny [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  grpc        Create grpc server/client file
  handle      Create handle file
  help        Help about any command
  new         Create a new Ginny project
  proto       Create proto file
  repo        Create repository file
  service     Create service file

Flags:
  -h, --help   help for ginny

Use "ginny [command] --help" for more information about a command.
```
### 创建项目

根据 [ginny-template](https://github.com/gorillazer/ginny-template) 模板创建新项目

```sh
Create a new Ginny project from template

Usage:
  ginny new [flags]

Flags:
      --grpc            Create a grpc service project
  -h, --help            help for new
      --http            Create a http service project (default true)
  -m, --module string   Define the project module, ex: github.com/demo
```

可根据参数创建http 或者 grpc服务的项目，默认创建http 项目

```sh
$ ginny new hellodemo --grpc

```
还可以定制项目 module地址:
```sh

$ ginny new hellodemo -m github.com/xxx/hellodemo
```

### 创建handler

http服务项目，可以通过命令行工具创建handler：

```sh
$ ginny handle user 

```
创建的handler文件在项目 internal/handlers目录


### 创建业务层service

```sh
$ ginny service user 

```
创建的service文件在项目 internal/services目录

### 创建数据层repository

```sh
// support mysql、mongo、redis 
$ ginny repo user -d mysql

```
创建的repository文件在项目 internal/repository目录

### 定义proto协议

```sh
Create proto file from template

Usage:
  ginny proto [flags]

Flags:
  -h, --help       help for proto
  -v, --validate   Added support for parameter verification
```
根据输入名称，自动创建.proto文件

```sh
ginny proto hello
```
hello.proto文件自动保存在项目api/proto目录。

根据proto文件生成.pb.go以及serve代码:
```sh
make proto
```

创建grpc服务项目，会自动创建proto文件，无需单独创建

