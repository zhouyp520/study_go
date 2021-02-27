# studygo
日常学习go

win7_x64+go1.15.8+vscode1.53.2

# 环境配置 for windows

## 1.go配置

### 安装go

https://golang.google.cn/dl/

### go程序路径

set GOROOT=C:\Program Files\Go

### gopath路径

set GOPATH=C:\go

## 2.VSCODE配置

### 安装vscode

https://code.visualstudio.com/Download

### 安装go插件

vscode->扩展->go 

#离线安装

#golang.Go-0.22.1.vsix

### 安装中文语言包插件

vscode->扩展->chinese

#离线安装vsix

### 安装go工具 

set GOPROXY=https://goproxy.io,direct

vscode->shift+ctrl+p->go:Install/Update Tools->选择全部安装

#离线安装gotools，在线机器GOPATH/bin下文件打包放到离线机器相同目录

# 编译

## go直接运行

go run main.go

## go编译成可执行文件

go build main.go

## go编译成linux可执行文件

Mac 下编译 Linux 和 Windows 64位可执行程序

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

Linux 下编译 Mac 和 Windows 64位可执行程序

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

Windows 下编译 Mac 和 Linux 64位可执行程序,一定要在cmd下运行

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go