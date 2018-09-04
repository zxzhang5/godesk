# 安装

- 从https://sciter.com/download/下载sciter-sdk.zip，解压，找到sciter-sdk\bin\64\sciter.dll复制到c:\windows\system32

- 下载安装tdm-gcc：http://tdm-gcc.tdragon.net/download，安装后将tdm-gcc\bin加入到环境变量Path中

- 下载安装zip：http://gnuwin32.sourceforge.net/packages/zip.htm，安装后将GnuWin32\bin加入到环境变量Path中

- cmd进入gopath目录src下运行：
```
go get github.com/GeertJohan/go.rice
go get github.com/GeertJohan/go.rice/rice
go get github.com/lxn/win
go get github.com/lxn/walk
go get github.com/sciter-sdk/go-sciter
go get github.com/mattn/go-sqlite3
go get github.com/imroc/req
go get github.com/braintree/manners
```
# 编译

go generate && go build -i -ldflags="-H windowsgui -w" -o dist/godesk.exe