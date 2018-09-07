# 安装

- 从https://sciter.com/download/下载sciter-sdk.zip，解压，找到sciter-sdk\bin\64\sciter.dll复制到c:\windows\system32

- 下载安装tdm-gcc：http://tdm-gcc.tdragon.net/download，安装后将tdm-gcc\bin加入到环境变量Path中

- 下载安装zip：http://gnuwin32.sourceforge.net/packages/zip.htm，安装后将GnuWin32\bin加入到环境变量Path中

- 下载安装包管理工具dep：https://github.com/golang/dep/releases，下载最新的dep-windows-amd64.exe，将其放入GOPATH/bin下，修改名称为dep

# 安装依赖包

- 切换到工程目录下：cd %GOPATH%/src/github.com/zxzhang5/godesk

- 初始化：
```
dep init
```
- 安装依赖(类似npm install)
```
dep ensure
```
- 更新依赖(类似npm update)
```
dep ensure -update
```
- 安装新的依赖包(类似npm install --save)
```
dep ensure -add github.com/go-ini/ini
```

# 依赖包说明
- 静态资源嵌入exe：[go.rice](https://github.com/GeertJohan/go.rice)
- windows API封装：[lxn/win](https://github.com/lxn/win)
- windows GUI：[lxn/walk](https://github.com/lxn/walk)
- sciter的golang实现：[go-sciter](https://github.com/sciter-sdk/go-sciter)
- 优雅的http服务关闭：[manners](https://github.com/braintree/manners)
- 人性化HTTP请求库：[req](https://github.com/imroc/req)
- sqlite操作库：[go-sqlite3](https://github.com/mattn/go-sqlite3)
- ini配置文件解析：[go-ini](https://github.com/go-ini/ini) [中文文档](https://ini.unknwon.io/)

# 编译

go generate && go build -i -ldflags="-H windowsgui -w" -o dist/godesk.exe